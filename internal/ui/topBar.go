package ui

import (
	"fmt"
	"strings"
)

func TopBar(file string, b strings.Builder) error {
	if file == "" {
		file = "[untitled]"
	}
	_, err := b.WriteString(fmt.Sprintf("File: %s  —  Ctrl+S save  |  Ctrl+Q quit\n\n", file))
	return err
}
