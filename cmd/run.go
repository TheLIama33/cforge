package cmd

import (
	"fmt"
	"os"

	"github.com/TheLIama33/cforge/internal/output"
	"github.com/TheLIama33/cforge/internal/scanner"
	"github.com/TheLIama33/cforge/internal/tokenizer"
	"github.com/TheLIama33/cforge/internal/utils"
	"github.com/spf13/cobra"
)

func runScan(cmd *cobra.Command, args []string) error {
	if appConfig == nil {
		return fmt.Errorf("configuration not initialized")
	}

	profile, err := appConfig.GetProfile(profileName)
	if err != nil {
		return err
	}

	useStdout, _ := cmd.Flags().GetBool("stdout")
	forceXML, _ := cmd.Flags().GetBool("xml")

	format := appConfig.Global.Formatting
	if forceXML {
		format = "xml"
	}

	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("cannot get workdir: %w", err)
	}

	if !useStdout {
		fmt.Printf("Scanning %s...\n", wd)
		if profileName != "" {
			fmt.Printf("Profile: %s\n", profileName)
		}
	}

	sc, err := scanner.New(wd, profile, appConfig.Global.UseGitIgnore)
	if err != nil {
		return fmt.Errorf("scanner init failed: %w", err)
	}

	results, err := sc.Scan()
	if err != nil {
		return fmt.Errorf("scan failed: %w", err)
	}

	if len(results) == 0 {
		if !useStdout {
			fmt.Println("No files matched. Check your profile configuration or .gitignore.")
		}
		return nil
	}

	finalOutput := output.Format(results, format)

	if useStdout {
		fmt.Print(finalOutput)
	} else {
		if appConfig.Global.CopyToClipboard {
			if err := utils.WriteToClipboard(finalOutput); err != nil {
				return fmt.Errorf("clipboard failure: %w", err)
			}
			fmt.Printf("Copied %d files to system clipboard!\n", len(results))
		} else {
			fmt.Println("Clipboard copy disabled. Use --stdout to see output.")
		}

		if appConfig.Global.ShowTokenCount {
			tokens := tokenizer.Estimate(finalOutput)
			lines := tokenizer.CountLines(finalOutput)
			fmt.Printf("Stats: ~%d Tokens | %d Lines\n", tokens, lines)
		}
	}

	return nil
}
