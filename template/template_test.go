package template

import (
	"bytes"
	htmlTemplate "html/template"
	"testing"
	textTemplate "text/template"
	"time"
)

func TestTemplateH_ParseFiles(t *testing.T) {
	resultMust := "Hello, Golang!"
	tmpl := NewCacheHTMLTemplate("test_data", 10, time.Minute)
	tp, err := tmpl.ParseFiles("1.tmpl", "2.tmpl")
	if err != nil {
		t.Fatal(err)
	}
	//for _, ti := range tp.Templates() {
	//	t.Logf("%+v", ti.Name())
	//}
	buf := bytes.NewBuffer([]byte{})
	err = tp.Lookup("1.tmpl").Execute(buf, map[string]string{"Name": "Golang"})
	if err != nil {
		t.Error(err)
	}
	if buf.String() != resultMust {
		t.Errorf("result: '%s' must be: '%s'", buf.String(), resultMust)
	}
}

func TestTemplateT_ParseFiles(t *testing.T) {
	resultMust := "Hello, Golang!"
	tmpl := NewCacheTextTemplate("test_data", 10, time.Minute)
	tp, err := tmpl.ParseFiles("1.tmpl", "2.tmpl")
	if err != nil {
		t.Error(err)
	}
	buf := bytes.NewBuffer([]byte{})
	err = tp.Lookup("1.tmpl").Execute(buf, map[string]string{"Name": "Golang"})
	if err != nil {
		t.Error(err)
	}
	if buf.String() != resultMust {
		t.Errorf("result: '%s' must be: '%s'", buf.String(), resultMust)
	}
}

func BenchmarkTemplateHTML(b *testing.B) {
	w := bytes.NewBuffer([]byte{})
	for i := 0; i < b.N; i++ {
		t, err := htmlTemplate.ParseFiles("test_data/1.tmpl", "test_data/2.tmpl")
		if err != nil {
			b.Error(err)
		}
		err = t.Lookup("1.tmpl").Execute(w, map[string]string{"Name": "Golang"})
		if err != nil {
			b.Error(err)
		}
		w.Reset()
	}
}

func BenchmarkTemplateH(b *testing.B) {
	tmpl := NewCacheHTMLTemplate("test_data", 10, time.Minute)
	w := bytes.NewBuffer([]byte{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t, err := tmpl.ParseFiles("1.tmpl", "2.tmpl")
		if err != nil {
			b.Error(err)
		}
		err = t.Lookup("1.tmpl").Execute(w, map[string]string{"Name": "Golang"})
		if err != nil {
			b.Error(err)
		}
		w.Reset()
	}
}

func BenchmarkTemplateText(b *testing.B) {
	w := bytes.NewBuffer([]byte{})
	for i := 0; i < b.N; i++ {
		t, err := textTemplate.ParseFiles("test_data/1.tmpl", "test_data/2.tmpl")
		if err != nil {
			b.Error(err)
		}
		err = t.Lookup("1.tmpl").Execute(w, map[string]string{"Name": "Golang"})
		if err != nil {
			b.Error(err)
		}
		w.Reset()
	}
}

func BenchmarkTemplateT(b *testing.B) {
	tmpl := NewCacheTextTemplate("test_data", 10, time.Minute)
	w := bytes.NewBuffer([]byte{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t, err := tmpl.ParseFiles("1.tmpl", "2.tmpl")
		if err != nil {
			b.Error(err)
		}
		err = t.Lookup("1.tmpl").Execute(w, map[string]string{"Name": "Golang"})
		if err != nil {
			b.Error(err)
		}
		w.Reset()
	}
}

func BenchmarkTemplateHTMLOnlyParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := htmlTemplate.ParseFiles("test_data/1.tmpl", "test_data/2.tmpl")
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkTemplateHOnlyParse(b *testing.B) {
	tmpl := NewCacheHTMLTemplate("test_data", 10, time.Minute)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tmpl.ParseFiles("1.tmpl", "2.tmpl")
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkTemplateTextOnlyParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := textTemplate.ParseFiles("test_data/1.tmpl", "test_data/2.tmpl")
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkTemplateTOnlyParse(b *testing.B) {
	tmpl := NewCacheTextTemplate("test_data", 10, time.Minute)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tmpl.ParseFiles("1.tmpl", "2.tmpl")
		if err != nil {
			b.Error(err)
		}
	}
}
