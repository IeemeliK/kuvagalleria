package web

import (
	"embed"
	"io/fs"
)

//go:embed static
var Static embed.FS

//go:embed templates
var templateRoot embed.FS

func Templates() fs.FS {
	sub, err := fs.Sub(templateRoot, "templates")
	if err != nil {
		panic(err)
	}
	return sub
}
