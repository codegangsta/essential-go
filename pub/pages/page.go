package pages

import (
	"io"
	"io/ioutil"
	"text/template"

	"github.com/theplant/blackfriday"
)

type Page struct {
	Title  string
	Author string
	HTML   string
	File   string
}

func NewPage(filename string) (*Page, error) {
	p := &Page{File: filename}
	return p, p.load()
}

func (p *Page) Render(layout string, w io.Writer) error {
	t, err := template.ParseFiles(layout)
	if err != nil {
		return err
	}

	return t.Execute(w, p)
}

func (p *Page) load() error {
	c, err := LoadConfig("config.json")
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(p.File)
	if err != nil {
		return nil
	}

	p.HTML = string(blackfriday.MarkdownCommon(data))
	p.Title = c.Name()
	p.Author = c.Author()
	return nil
}
