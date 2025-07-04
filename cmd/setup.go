package cmd

import (
	"fmt"
	"os"

	"github.com/rvillegasm/vista/internal/templates"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Bootstrap current directory with baseline context files",
	RunE: func(cmd *cobra.Command, args []string) error {
		files := map[string]string{
			"RULES.md":        templates.RulesMD,
			"GENERATE_PRP.md": templates.GeneratePRPMD,
			"EXECUTE_PRP.md":  templates.ExecutePRPMD,
			"prp_template.md": templates.PrpTemplateMD,
		}

		for name, content := range files {
			if _, err := os.Stat(name); err == nil {
				cmd.Printf("%s already exists. Skipping.\n", name)
				continue
			} else if !os.IsNotExist(err) {
				return fmt.Errorf("checking %s: %w", name, err)
			}

			if err := os.WriteFile(name, []byte(content), 0o644); err != nil {
				return fmt.Errorf("writing %s: %w", name, err)
			}
			cmd.Printf("Created %s\n", name)
		}

		cmd.Println("Setup complete.")
		return nil
	},
}
