package scanner

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/TheLIama33/cforge/internal/config"
	"github.com/monochromegane/go-gitignore"
)

var systemExcludes = map[string]bool{
	".git": true, "node_modules": true, ".idea": true, ".vscode": true, "vendor": true,
	".next": true, "target": true, "build": true, "dist": true, "bin": true, "obj": true,
}

const MaxFileSize = 1024 * 1024

type FileResult struct {
	Path    string
	Content string
}

type Scanner struct {
	Root         string
	Profile      config.Profile
	UseGitIgnore bool
	gitMatcher   gitignore.IgnoreMatcher
}

func New(root string, profile config.Profile, useGitIgnore bool) (*Scanner, error) {
	s := &Scanner{
		Root:         root,
		Profile:      profile,
		UseGitIgnore: useGitIgnore,
	}

	if useGitIgnore {
		gitignorePath := filepath.Join(root, ".gitignore")
		if info, err := os.Stat(gitignorePath); err == nil && !info.IsDir() {
			matcher, err := gitignore.NewGitIgnore(gitignorePath)
			if err != nil {
				return nil, fmt.Errorf("failed to parse .gitignore: %w", err)
			}
			s.gitMatcher = matcher
		}
	}
	return s, nil
}

func (s *Scanner) Scan() ([]FileResult, error) {
	var results []FileResult

	hasIncludeRules := len(s.Profile.IncludeFiles) > 0 ||
		len(s.Profile.IncludePaths) > 0 ||
		len(s.Profile.IncludePatterns) > 0

	err := filepath.WalkDir(s.Root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		relPath, err := filepath.Rel(s.Root, path)
		if err != nil {
			return nil
		}
		if relPath == "." {
			return nil
		}

		slashPath := filepath.ToSlash(relPath)
		name := d.Name()

		if d.IsDir() {
			if systemExcludes[name] {
				return filepath.SkipDir
			}

			if s.UseGitIgnore && s.gitMatcher != nil {
				if s.gitMatcher.Match(path, true) {
					return filepath.SkipDir
				}
			}

			for _, exPath := range s.Profile.ExcludePaths {
				cleanEx := filepath.ToSlash(exPath)
				if slashPath == cleanEx || strings.HasPrefix(slashPath, cleanEx+"/") {
					return filepath.SkipDir
				}
			}

			if hasIncludeRules {
				shouldTraverse := false

				for _, incPath := range s.Profile.IncludePaths {
					cleanInc := filepath.ToSlash(incPath)
					if strings.HasPrefix(cleanInc, slashPath+"/") {
						shouldTraverse = true
						break
					}
					if strings.HasPrefix(slashPath, cleanInc) {
						shouldTraverse = true
						break
					}
				}

				if !shouldTraverse {
					for _, incFile := range s.Profile.IncludeFiles {
						cleanFile := filepath.ToSlash(incFile)
						if strings.HasPrefix(cleanFile, slashPath+"/") {
							shouldTraverse = true
							break
						}
					}
				}

				if !shouldTraverse && len(s.Profile.IncludePatterns) == 0 {
					return filepath.SkipDir
				}
			}
			return nil
		}

		info, err := d.Info()
		if err == nil && info.Size() > MaxFileSize {
			return nil
		}

		if s.UseGitIgnore && s.gitMatcher != nil {
			if s.gitMatcher.Match(path, false) {
				return nil
			}
		}

		for _, exFile := range s.Profile.ExcludeFiles {
			if slashPath == filepath.ToSlash(exFile) {
				return nil
			}
		}

		for _, pat := range s.Profile.ExcludePatterns {
			if matched, _ := filepath.Match(pat, name); matched {
				return nil
			}
		}

		keep := false
		if !hasIncludeRules {
			keep = true
		} else {
			for _, incPath := range s.Profile.IncludePaths {
				cleanInc := filepath.ToSlash(incPath)
				if strings.HasPrefix(slashPath, cleanInc) {
					keep = true
					break
				}
			}
			if !keep {
				for _, incFile := range s.Profile.IncludeFiles {
					if slashPath == filepath.ToSlash(incFile) {
						keep = true
						break
					}
				}
			}
			if !keep {
				for _, pat := range s.Profile.IncludePatterns {
					if matched, _ := filepath.Match(pat, name); matched {
						keep = true
						break
					}
				}
			}
		}

		if keep {
			content, err := readFileIfText(path)
			if err != nil {
				return nil
			}

			results = append(results, FileResult{
				Path:    slashPath,
				Content: content,
			})
		}

		return nil
	})

	return results, err
}

func readFileIfText(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	buf := make([]byte, 512)
	n, err := f.Read(buf)
	if err != nil && err != io.EOF {
		return "", err
	}

	if n == 0 {
		return "", nil
	}

	if bytes.IndexByte(buf[:n], 0) != -1 {
		return "", fmt.Errorf("binary file detected")
	}

	if !utf8.Valid(buf[:n]) {
		return "", fmt.Errorf("invalid utf-8")
	}

	if _, err := f.Seek(0, 0); err != nil {
		return "", err
	}

	content, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
