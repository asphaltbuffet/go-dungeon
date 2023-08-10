package cmd

import (
	"github.com/spf13/cobra"
)

var generateCmd *cobra.Command

func GetGenerateCmd() *cobra.Command {
	if generateCmd == nil {
		generateCmd = &cobra.Command{
			Use:     "generate [flags]",
			Aliases: []string{"play", "run"},
			Short:   "explore a dungeon",
			RunE: func(cmd *cobra.Command, args []string) error {
				return runGenerate()
			},
		}
	}

	return generateCmd
}

func runGenerate() error {
	return nil
}
