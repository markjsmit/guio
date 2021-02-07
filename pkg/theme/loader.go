package theme

import (
	"bytes"
	"context"
	"path/filepath"
	"text/template"
)

type Loader interface {
	LoadComponent(ctx context.Context, component string, data interface{}) ([]byte, error)
	Path(category string, file string) string
	LoadStylesheet(style string) ([]byte, error)
}

type loader struct {
	themeDirectory string
	cache          map[string]*template.Template
}

func (t *loader) Path(category string, file string) string {
	return filepath.Join(t.themeDirectory, category, file)
}

func NewLoader(themeDirectory string) Loader {
	return &loader{themeDirectory: themeDirectory, cache: map[string]*template.Template{}}
}

func (t *loader) get(category string, file string) (*template.Template, error) {
	
	tplName := file
	fileName := t.Path(category, tplName)
	
	if tmpl, ok := t.cache[fileName]; ok {
		return tmpl, nil
	}
	tmpl, err := template.New(tplName).Funcs(TemplateFuncs(t)).ParseFiles(fileName)
	
	if err == nil {
		t.cache[fileName] = tmpl
	}
	
	return tmpl, err
}

func (t *loader) LoadComponent(ctx context.Context, component string, data interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	tmpl, err := t.get("component", component+".xml")
	err = tmpl.Execute(&buffer, data)
	if err != nil {
		return []byte{}, err
	}
	out := buffer.Bytes()
	return out, err
}

func (t *loader) LoadStylesheet(style string) ([]byte, error) {
	var buffer bytes.Buffer
	
	tmpl, err := t.get("style", style+".json")
	if err != nil {
		return nil, err
	}
	err = tmpl.Execute(&buffer, map[string]string{})
	out := buffer.Bytes()
	return out, err
}
