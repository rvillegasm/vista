package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSetupCmdCreatesFiles(t *testing.T) {
	dir := t.TempDir()
	cwd, _ := os.Getwd()
	defer os.Chdir(cwd)
	os.Chdir(dir)

	if err := setupCmd.RunE(setupCmd, []string{}); err != nil {
		t.Fatalf("setupCmd returned error: %v", err)
	}

	expected := []string{"RULES.md", "GENERATE_PRP.md", "EXECUTE_PRP.md", "prp_template.md"}
	for _, f := range expected {
		if _, err := os.Stat(filepath.Join(dir, f)); err != nil {
			t.Errorf("expected %s to be created: %v", f, err)
		}
	}
}
