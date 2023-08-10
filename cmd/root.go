package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd  *cobra.Command
	verbose  bool
	textOnly bool
	language string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(GetRootCommand().Execute())
}

// GetRootCommand returns the root command for the CLI.
func GetRootCommand() *cobra.Command {
	if rootCmd == nil {
		rootCmd = &cobra.Command{
			Use:               "dungeon [command]",
			Version:           "1.0.0",
			Short:             "dungeon is a CLI for generating single-player dungeons",
			PersistentPreRunE: getConfig,
			RunE:              RunRootCmd,
		}
	}

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Show extra information.")
	rootCmd.PersistentFlags().BoolVarP(&textOnly, "text", "t", false, "Run without user interaction.")
	rootCmd.PersistentFlags().StringVarP(&language, "lang", "l", "", "Custom language file to use")

	rootCmd.AddCommand(GetGenerateCmd())

	return rootCmd
}

// getConfig is a pre-run function that loads the configuration and sets defaults.
func getConfig(cmd *cobra.Command, args []string) error {
	return nil
}

// RunRootCmd is the entry point for the CLI.
func RunRootCmd(cmd *cobra.Command, args []string) error {
	_ = cmd.Help()

	return nil
}
