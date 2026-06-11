package templates

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
	"sync"
)

type TemplateCache struct {
	templates map[string]*template.Template
	mu        sync.RWMutex
}

var cache *TemplateCache

func Init(templatesFS fs.FS) error {
	cache = NewTemplateCache()
	return cache.LoadAll(templatesFS)
}

func NewTemplateCache() *TemplateCache {
	return &TemplateCache{
		templates: make(map[string]*template.Template),
	}
}

func (tc *TemplateCache) LoadAll(templatesFS fs.FS) error {
	layouts, err := fs.Glob(templatesFS, "layouts/*.html")
	if err != nil {
		return fmt.Errorf("failed to find layouts: %w", err)
	}

	pages, err := fs.Glob(templatesFS, "pages/*.html")
	if err != nil {
		return fmt.Errorf("failed to find pages: %w", err)
	}

	for _, page := range pages {
		name := filepath.Base(page)
		files := append(layouts, page)

		tmpl := template.New(name).Funcs(template.FuncMap{})
		_, err = tmpl.ParseFS(templatesFS, files...)
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", name, err)
		}

		tc.templates[name] = tmpl
	}

	return nil
}

func (tc *TemplateCache) Get(name string) (*template.Template, error) {
	tc.mu.RLock()
	defer tc.mu.RUnlock()

	tmpl, ok := tc.templates[name]
	if !ok {
		return nil, fmt.Errorf("template %s not found", name)
	}

	return tmpl, nil
}

func (tc *TemplateCache) Render(w http.ResponseWriter, name string, layout string, data any) error {
	tmpl, err := tc.Get(name)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if layout != "" {
		err = tmpl.ExecuteTemplate(buf, layout, data)
	} else {
		err = tmpl.Execute(buf, data)
	}
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err = buf.WriteTo(w)
	return err
}

func Render(w http.ResponseWriter, name string, layout string, data any) error {
	if cache == nil {
		return fmt.Errorf("template cache not initialized")
	}
	return cache.Render(w, name, layout, data)
}
