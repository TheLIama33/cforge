package cmd

import (
	"fmt"

	"github.com/TheLIama33/cforge/internal/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a default .cforge.json configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.GenerateDefault(); err != nil {
			return fmt.Errorf("failed to create config: %w", err)
		}
		fmt.Println("Created .cforge.json in current directory.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
