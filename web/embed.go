package web

import (
	"embed"
	"io/fs"
)

//go:embed static
var staticRoot embed.FS

//go:embed templates
var templateRoot embed.FS

func StaticFS() fs.FS {
	sub, err := fs.Sub(staticRoot, "static")
	if err != nil {
		panic(err)
	}
	return sub
}

func Templates() fs.FS {
	sub, err := fs.Sub(templateRoot, "templates")
	if err != nil {
		panic(err)
	}
	return sub
}
