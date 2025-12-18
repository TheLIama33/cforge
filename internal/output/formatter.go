package output

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/TheLIama33/cforge/internal/scanner"
)

func Format(files []scanner.FileResult, formatType string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("ContextForge Export | Total Files: %d\n\n", len(files)))

	sb.WriteString("Included Files:\n")
	for _, f := range files {
		sb.WriteString(fmt.Sprintf("- %s\n", f.Path))
	}
	sb.WriteString("\n")

	for _, f := range files {
		if strings.ToLower(formatType) == "xml" {
			sb.WriteString(formatXML(f))
		} else {
			sb.WriteString(formatMarkdown(f))
		}
	}
	return sb.String()
}

func formatMarkdown(f scanner.FileResult) string {
	ext := strings.TrimPrefix(filepath.Ext(f.Path), ".")
	if ext == "" {
		ext = "text"
	}

	delimiter := "```"
	if strings.Contains(f.Content, "```") {
		delimiter = "````"
	}

	var sb strings.Builder
	sb.WriteString("### File: " + f.Path + "\n")
	sb.WriteString(delimiter + ext + "\n")
	sb.WriteString(f.Content)
	if !strings.HasSuffix(f.Content, "\n") {
		sb.WriteString("\n")
	}
	sb.WriteString(delimiter + "\n\n")
	return sb.String()
}

func formatXML(f scanner.FileResult) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("<file path=\"%s\">\n", f.Path))
	sb.WriteString("<![CDATA[\n")

	safeContent := strings.ReplaceAll(f.Content, "]]>", "]]]]><![CDATA[>")

	sb.WriteString(safeContent)
	if !strings.HasSuffix(safeContent, "\n") {
		sb.WriteString("\n")
	}
	sb.WriteString("]]>\n")
	sb.WriteString("</file>\n\n")
	return sb.String()
}
