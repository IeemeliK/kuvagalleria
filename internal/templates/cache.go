package templates

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/IeemeliK/kuvagalleria/internal"
)

type TemplateCache struct {
	templates map[string]*template.Template
	mu        sync.RWMutex
}

var cache *TemplateCache

func init() {
	cache = NewTemplateCache()
	if err := cache.LoadAll(); err != nil {
		log.Fatalf("failed to load templates: %v", err)
	}
}

func NewTemplateCache() *TemplateCache {
	return &TemplateCache{
		templates: make(map[string]*template.Template),
	}
}

func (tc *TemplateCache) LoadAll() error {
	layouts, err := fs.Glob(internal.Templates, "layouts/*.html")
	if err != nil {
		return fmt.Errorf("failed to find layouts: %w", err)
	}

	pages, err := fs.Glob(internal.Templates, "pages/*.html")
	if err != nil {
		return fmt.Errorf("failed to find pages: %w", err)
	}

	for _, page := range pages {
		name := filepath.Base(page)
		files := append(layouts, page)

		tmpl := template.New(name).Funcs(template.FuncMap{})
		_, err = tmpl.ParseFS(internal.Templates, files...)
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
	return cache.Render(w, name, layout, data)
}
