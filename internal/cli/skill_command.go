package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tamnguyendinh/anvien/internal/aicontext"
)

const codexSkillHelpPrefix = ".agents/skills/"

func newSkillCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "skill",
		Short: "Show available AI agent skills and when to use them",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return writeSkillHelp(cmd)
		},
	}
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		if err := writeSkillHelp(cmd); err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Error: %v\n", err)
		}
	})
	return cmd
}

func writeSkillHelp(cmd *cobra.Command) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	guide, err := aicontext.RenderSkillSelectionGuideForRepo(cwd, codexSkillHelpPrefix)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(cmd.OutOrStdout(), guide)
	return err
}
