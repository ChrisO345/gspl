package examples_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestExamplesDoNotPanic(t *testing.T) {
	examplesRoot := "."

	entries, err := os.ReadDir(examplesRoot)
	if err != nil {
		t.Fatalf("Failed to read examples directory: %v", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		examplePath := filepath.Join(examplesRoot, entry.Name())
		mainGo := filepath.Join(examplePath, "main.go")

		if _, err := os.Stat(mainGo); err != nil {
			continue
		}

		t.Run(entry.Name(), func(t *testing.T) {
			cmd := exec.Command("go", "run", mainGo)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				t.Errorf("Example %q failed: %v", entry.Name(), err)
			}
		})
	}

	// panic("")
}
