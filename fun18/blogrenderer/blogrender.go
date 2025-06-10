package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

// if you're continuing from the read files chapte, you shouldn't redefine this

type PostRenderer struct {
	templ *template.Template
}
type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}
	return nil
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return nil
}

//go:embed "templates/*"
var postTemplates embed.FS

func Render(w io.Writer, p Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}
	return nil
}
