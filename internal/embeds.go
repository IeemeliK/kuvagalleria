// Package internal
package internal

import (
	"embed"
	"io/fs"
	"log"
)

var (
	//go:embed templates
	templateFS embed.FS

	Templates fs.FS

	//go:embed static
	Static embed.FS
)

func init() {
	sub, err := fs.Sub(templateFS, "templates")
	if err != nil {
		log.Fatalf("failed to create template fs: %v", err)
	}
	Templates = sub
}
