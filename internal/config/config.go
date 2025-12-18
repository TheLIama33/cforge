package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const DefaultConfigName = ".cforge.json"

type Profile struct {
	IncludePatterns []string `mapstructure:"includePatterns" json:"includePatterns"`
	IncludeFiles    []string `mapstructure:"includeFiles" json:"includeFiles"`
	IncludePaths    []string `mapstructure:"includePaths" json:"includePaths"`
	ExcludePatterns []string `mapstructure:"excludePatterns" json:"excludePatterns"`
	ExcludeFiles    []string `mapstructure:"excludeFiles" json:"excludeFiles"`
	ExcludePaths    []string `mapstructure:"excludePaths" json:"excludePaths"`
}

type GlobalSettings struct {
	CopyToClipboard bool   `mapstructure:"copyToClipboard" json:"copyToClipboard"`
	ShowTokenCount  bool   `mapstructure:"showTokenCount" json:"showTokenCount"`
	DefaultProfile  string `mapstructure:"defaultProfile" json:"defaultProfile"`
	UseGitIgnore    bool   `mapstructure:"useGitIgnore" json:"useGitIgnore"`
	Formatting      string `mapstructure:"formatting" json:"formatting"`
}

type Config struct {
	Global   GlobalSettings     `mapstructure:"global" json:"global"`
	Profiles map[string]Profile `mapstructure:"profiles" json:"profiles"`
}

func DefaultConfig() Config {
	return Config{
		Global: GlobalSettings{
			CopyToClipboard: true,
			ShowTokenCount:  true,
			DefaultProfile:  "default",
			UseGitIgnore:    true,
			Formatting:      "markdown",
		},
		Profiles: map[string]Profile{
			"default": {
				IncludePatterns: []string{"*.md", "*.go", "*.js", "*.ts", "*.py", "*.java", "*.c", "*.cpp"},
				ExcludePatterns: []string{"*.spec.ts", "*.test.ts", "*.log", "*.tmp", "*.exe"},
				ExcludePaths:    []string{"node_modules", "dist", "vendor", ".git", "bin", "obj", "build", ".vscode"},
			},
		},
	}
}

func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigName(".cforge")
	v.SetConfigType("json")
	v.AddConfigPath(".")

	home, err := os.UserHomeDir()
	if err == nil {
		v.AddConfigPath(home)
	}

	defaults := DefaultConfig()
	v.SetDefault("global", defaults.Global)
	v.SetDefault("profiles", defaults.Profiles)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("config file found but invalid: %w", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse configuration: %w", err)
	}

	return &cfg, nil
}

func (c *Config) GetProfile(name string) (Profile, error) {
	if name == "" {
		name = c.Global.DefaultProfile
	}

	if c.Profiles == nil {
		return Profile{}, nil
	}

	if p, ok := c.Profiles[name]; ok {
		return p, nil
	}

	if name != "default" {
		return Profile{}, fmt.Errorf("profile '%s' not found", name)
	}

	return Profile{}, nil
}

func GenerateDefault() error {
	cfg := DefaultConfig()
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(DefaultConfigName, data, 0644)
}
