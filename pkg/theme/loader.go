package theme

import (
	"bytes"
	"context"
	"path/filepath"
	"text/template"
)

type Loader interface {
	Load(ctx context.Context, component string, data map[string]string) ([]byte, error)
	GetPathFor(file string) string
}

type loader struct {
	themeDirectory string
	cache          map[string]*template.Template
}

func (t *loader) GetPathFor(file string) string {
	return filepath.Join(t.themeDirectory, file)
}

func NewLoader(themeDirectory string) Loader {
	return &loader{themeDirectory: themeDirectory, cache: map[string]*template.Template{}}
}

func (t *loader) get(component string) (*template.Template, error) {
	
	tplName := component + ".xml"
	fileName := t.GetPathFor(tplName)
	
	if tmpl, ok := t.cache[fileName]; ok {
		return tmpl, nil
	}
	
	tmpl, err := template.New(tplName).Funcs(funcMap).ParseFiles(fileName)
	
	if err == nil {
		t.cache[fileName] = tmpl
	}
	
	return tmpl, err
}

func (t *loader) Load(ctx context.Context, component string, data map[string]string) ([]byte, error) {
	var tpl bytes.Buffer
	tmpl, err := t.get(component)
	err = tmpl.Execute(&tpl, data)
	if err != nil {
		return []byte{}, err
	}
	out := tpl.Bytes()
	return out, err
}
