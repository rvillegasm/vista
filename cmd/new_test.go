package cmd

import (
	"os"
	"testing"
)

func TestNewCmd(t *testing.T) {
	dir := t.TempDir()

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}

	// Restore working directory when the test ends.
	t.Cleanup(func() {
		if err := os.Chdir(cwd); err != nil {
			t.Fatalf("chdir back: %v", err)
		}
	})

	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir into temp dir: %v", err)
	}

	if err := newCmd.RunE(newCmd, []string{"test"}); err != nil {
		t.Fatalf("newCmd returned error: %v", err)
	}

	if _, err := os.Stat("FEATURE_test.md"); err != nil {
		t.Fatalf("feature file not created: %v", err)
	}

	if err := newCmd.RunE(newCmd, []string{"test"}); err == nil {
		t.Fatalf("expected error when feature file already exists")
	}
}
