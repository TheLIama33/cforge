package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var Version = "0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ContextForge",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("ContextForge v%s\n", Version)
		fmt.Printf("Go Version: %s\n", runtime.Version())
		fmt.Printf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
