package cmd

import (
	"fmt"
	"os"

	"github.com/TheLIama33/cforge/internal/updater"
	"github.com/spf13/cobra"
)

var RepoSlug = "TheLIama33/cforge"

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update cforge to the latest version from GitHub",
	Long:  fmt.Sprintf("Checks %s for the latest release and updates the binary if a newer version is available.", RepoSlug),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Checking updates for %s...\n", RepoSlug)
		result, err := updater.CheckAndUpdate(Version, RepoSlug)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Update failed: %v\n", err)
			os.Exit(1)
		}

		if result.Updated {
			fmt.Printf("Successfully updated to v%s!\n", result.LatestVersion)
			fmt.Println("Release Notes:")
			fmt.Println(result.ReleaseNotes)
		} else {
			fmt.Printf("You are already using the latest version (v%s).\n", result.LatestVersion)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
