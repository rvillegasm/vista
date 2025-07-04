package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSetupCmdCreatesFiles(t *testing.T) {
	dir := t.TempDir()

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}

	t.Cleanup(func() {
		if err := os.Chdir(cwd); err != nil {
			t.Fatalf("chdir back: %v", err)
		}
	})

	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir into temp dir: %v", err)
	}

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
