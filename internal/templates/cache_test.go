package templates

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/IeemeliK/kuvagalleria/internal"
)

func TestNewTemplateCache(t *testing.T) {
	tc := NewTemplateCache()
	if tc == nil {
		t.Fatal("NewTemplateCache() returned nil")
	}
	if len(tc.templates) != 0 {
		t.Errorf("expected empty cache, got %d entries", len(tc.templates))
	}
}

func TestCacheGet(t *testing.T) {
	tc := NewTemplateCache()

	t.Run("missing template", func(t *testing.T) {
		_, err := tc.Get("nonexistent.html")
		if err == nil {
			t.Error("expected error for missing template")
		}
	})
}

func TestLoadAll(t *testing.T) {
	if err := internal.InitTemplates(); err != nil {
		t.Fatalf("InitTemplates() failed: %v", err)
	}

	tc := NewTemplateCache()
	if err := tc.LoadAll(); err != nil {
		t.Fatalf("LoadAll() failed: %v", err)
	}

	tests := []struct {
		name     string
		tmplName string
		wantOk   bool
	}{
		{name: "index page", tmplName: "index.html", wantOk: true},
		{name: "login page", tmplName: "login.html", wantOk: true},
		{name: "nonexistent", tmplName: "missing.html", wantOk: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tc.Get(tt.tmplName)
			if tt.wantOk && err != nil {
				t.Errorf("Get(%q) unexpected error: %v", tt.tmplName, err)
			}
			if !tt.wantOk && err == nil {
				t.Errorf("Get(%q) expected error", tt.tmplName)
			}
		})
	}
}

func TestRender(t *testing.T) {
	if err := internal.InitTemplates(); err != nil {
		t.Fatalf("InitTemplates() failed: %v", err)
	}

	if err := Init(); err != nil {
		t.Fatalf("Init() failed: %v", err)
	}

	t.Run("successful render", func(t *testing.T) {
		w := httptest.NewRecorder()
		err := Render(w, "index.html", "", nil)
		if err != nil {
			t.Fatalf("Render() error: %v", err)
		}
		if w.Code != 200 {
			t.Errorf("expected 200, got %d", w.Code)
		}
		if w.Body.Len() == 0 {
			t.Error("expected non-empty body")
		}
	})

	t.Run("missing template", func(t *testing.T) {
		w := httptest.NewRecorder()
		err := Render(w, "missing.html", "", nil)
		if err == nil {
			t.Error("expected error for missing template")
		}
	})
}

func TestRenderNotInitialized(t *testing.T) {
	old := cache
	cache = nil
	t.Cleanup(func() { cache = old })

	w := httptest.NewRecorder()
	err := Render(w, "index.html", "", nil)
	if err == nil {
		t.Error("expected error when cache not initialized")
	}
}

func TestRenderToBuffer(t *testing.T) {
	if err := internal.InitTemplates(); err != nil {
		t.Fatalf("InitTemplates() failed: %v", err)
	}

	if err := Init(); err != nil {
		t.Fatalf("Init() failed: %v", err)
	}

	tc := cache
	tmpl, err := tc.Get("index.html")
	if err != nil {
		t.Fatalf("Get() error: %v", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, nil); err != nil {
		t.Fatalf("Execute() error: %v", err)
	}
	if buf.Len() == 0 {
		t.Error("expected non-empty output")
	}
}
