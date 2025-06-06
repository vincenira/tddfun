package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

// if you're continuing from the read files chapte, you shouldn't redefine this

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

//go:embed "templates/*"
var postTemplates embed.FS

func Render(w io.Writer, p Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.Execute(w, p); err != nil {
		return err
	}
	return nil
}
