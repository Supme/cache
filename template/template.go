package template

import (
	"cache/lruttl"
	htmlTemplate "html/template"
	"path/filepath"
	textTemplate "text/template"
	"time"
)

type templateH struct {
	cache *lruttl.LRU
	htmlTemplate.Template
}

type templateT struct {
	cache *lruttl.LRU
	textTemplate.Template
}

// NewCacheHTMLTemplate return cached wrapper over html/template with custom its own analogue function ParseFiles
func NewCacheHTMLTemplate(rootPath string, capacity int, ttl time.Duration) *templateH {
	get := func(name string) (interface{}, time.Duration, error) {
		t, err := htmlTemplate.ParseFiles(filepath.Join(rootPath, name))
		return t, 0, err
	}
	cache := lruttl.NewCache(capacity, ttl, get)
	r := new(templateH)
	r.cache = cache
	return r
}

// NewCacheTextTemplate return cached wrapper over text/template with custom its own analogue function ParseFiles
func NewCacheTextTemplate(rootPath string, capacity int, ttl time.Duration) *templateT {
	get := func(p string) (interface{}, time.Duration, error) {
		t, err := textTemplate.ParseFiles(filepath.Join(rootPath, p))
		return t, 0, err
	}
	cache := lruttl.NewCache(capacity, ttl, get)
	r := new(templateT)
	r.cache = cache
	return r
}

// ParseFiles cached wrapper function over html/template ParseFiles
func (t templateH) ParseFiles(filenames ...string) (*htmlTemplate.Template, error) {
	tmpl := htmlTemplate.New("")
	for i := range filenames {
		r, err := t.cache.Get(filenames[i])
		if err != nil {
			return nil, err
		}
		tmpl, err = tmpl.AddParseTree(filenames[i], r.(*htmlTemplate.Template).Tree)
		if err != nil {
			return nil, err
		}
	}
	return tmpl, nil
}

// ParseFiles cached wrapper function over text/template ParseFiles
func (t templateT) ParseFiles(filenames ...string) (*textTemplate.Template, error) {
	tmpl := textTemplate.New("")
	for i := range filenames {
		r, err := t.cache.Get(filenames[i])
		if err != nil {
			return nil, err
		}
		tmpl, err = tmpl.AddParseTree(filenames[i], r.(*textTemplate.Template).Tree)
		if err != nil {
			return nil, err
		}
	}
	return tmpl, nil
}
