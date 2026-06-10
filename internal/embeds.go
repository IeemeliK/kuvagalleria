// Package internal
package internal

import (
	"embed"
	"fmt"
	"io/fs"
)

var (
	//go:embed templates
	templateFS embed.FS

	Templates fs.FS

	//go:embed static
	Static embed.FS
)

func InitTemplates() error {
	sub, err := fs.Sub(templateFS, "templates")
	if err != nil {
		return fmt.Errorf("create template sub fs: %w", err)
	}
	Templates = sub
	return nil
}
