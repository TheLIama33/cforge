package cmd

import (
	"fmt"
	"os"

	"github.com/TheLIama33/cforge/internal/config"
	"github.com/spf13/cobra"
)

var (
	cfgFile     string
	profileName string
	appConfig   *config.Config
)

var rootCmd = &cobra.Command{
	Use:   "cforge",
	Short: "ContextForge: Copy code context to clipboard",
	Long:  `Scan your project, filter files, and copy them to the system clipboard for LLM usage.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Name() == "init" || cmd.Name() == "version" || cmd.Name() == "update" {
			return nil
		}

		var err error
		appConfig, err = config.Load()
		if err != nil {
			return fmt.Errorf("failed to load configuration (run 'cforge init' to fix): %w", err)
		}
		return nil
	},
	RunE: runScan,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&profileName, "profile", "p", "", "Profile to use (defined in .cforge.json)")
	rootCmd.Flags().Bool("stdout", false, "Print to stdout instead of clipboard")
	rootCmd.Flags().Bool("xml", false, "Force XML output format")
}
