package cmd

import (
	"fmt"

	"github.com/matkv/utils/internal/dotfiles"
	"github.com/matkv/utils/internal/registry"
	"github.com/spf13/cobra"
)

// dotfilesCmd represents the dotfiles command
var dotfilesCmd = &cobra.Command{
	Use:   "dotfiles",
	Short: "Tools to manage my dotfiles.",
	Long:  `Tools to manage my dotfiles. For example moving dotfiles to the correct location on Windows so I can update my dotfiles git repository more easily.`,
	Annotations: map[string]string{
		"IsWindowsOnly": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("Unknown command:", args[0])
			fmt.Println("Use 'utils dotfiles --help' for more information.")
			return
		}

		if len(args) == 0 {
			cmd.Help()
		}

		if args[0] == "sync" {
			if err := dotfiles.SyncConfigFiles(); err != nil {
				fmt.Printf("Error syncing config files: %v\n", err)
				return
			}
			fmt.Println("✓ Config files synced successfully!")
			return
		}
		if args[0] == "status" {
			if err := dotfiles.ShowStatus(); err != nil {
				fmt.Printf("Error showing status: %v\n", err)
			}
			return
		}
		if args[0] == "pull" {
			if err := dotfiles.PullLatestDotfiles(); err != nil {
				fmt.Printf("Error pulling dotfiles: %v\n", err)
				return
			}
			fmt.Println("✓ Pulled latest dotfiles successfully!")
			return
		}
	},
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync config files to dotfiles repository",
	Long:  `Copy configured files from their locations to the dotfiles repository. Ensures the repo is clean and up to date first.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := dotfiles.SyncConfigFiles(); err != nil {
			fmt.Printf("Error syncing config files: %v\n", err)
			return
		}
		fmt.Println("✓ Config files synced successfully!")
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show dotfiles repository status",
	Long:  `Display the current git status of the dotfiles repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := dotfiles.ShowStatus(); err != nil {
			fmt.Printf("Error showing status: %v\n", err)
		}
	},
}

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull latest changes from dotfiles repository",
	Long:  `Pull the latest changes from the remote dotfiles repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := dotfiles.PullLatestDotfiles(); err != nil {
			fmt.Printf("Error pulling dotfiles: %v\n", err)
			return
		}
	},
}

func init() {
	dotfilesCmd.AddCommand(syncCmd)
	dotfilesCmd.AddCommand(statusCmd)
	dotfilesCmd.AddCommand(pullCmd)
	registry.RegisterCommand(dotfilesCmd)
}
