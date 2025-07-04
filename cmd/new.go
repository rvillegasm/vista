package cmd

import (
	"fmt"
	"os"

	"github.com/rvillegasm/vista/internal/templates"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new <feature_name>",
	Short: "Create new feature file initialized from template",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		featureName := args[0]
		if featureName == "" {
			return fmt.Errorf("feature name cannot be empty")
		}

		dest := fmt.Sprintf("FEATURE_%s.md", featureName)

		if _, err := os.Stat(dest); err == nil {
			return fmt.Errorf("%s already exists", dest)
		} else if err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("checking %s: %w", dest, err)
		}

		if err := os.WriteFile(dest, []byte(templates.FeatureTemplateMD), 0o644); err != nil {
			return fmt.Errorf("writing %s: %w", dest, err)
		}

		cmd.Printf("Created %s\n", dest)
		return nil
	},
}
