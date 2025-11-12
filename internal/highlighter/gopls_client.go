package highlighter

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
)

type SemanticToken struct {
	Offset    int
	Type      string
	Modifiers []string
	Value     string
}

func GetTokens(filePath string) (string, error) {
	// Convert to absolute, clean path
	abs, err := filepath.Abs(filePath)
	if err != nil {
		return "", fmt.Errorf("invalid path: %w", err)
	}

	// Run gopls from the same directory as the file
	cmd := exec.Command("gopls", "semtok", abs)
	cmd.Dir = filepath.Dir(abs)

	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("gopls error: %v\nstderr: %s", err, stderr.String())
	}

	return out.String(), nil
}
