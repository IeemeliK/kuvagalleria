// Package internal
package internal

import "embed"

var (
	//go:embed all:static
	Static embed.FS

	//go:embed templates/*.html
	Templates embed.FS
)
