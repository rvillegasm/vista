package cmd

import (
	"os"
	"testing"
)

func TestNewCmd(t *testing.T) {
	dir := t.TempDir()
	cwd, _ := os.Getwd()
	defer os.Chdir(cwd)
	os.Chdir(dir)

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
