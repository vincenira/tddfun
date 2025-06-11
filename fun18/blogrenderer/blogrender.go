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
type PostViewModel struct {
	Title, SanitisedTitle, Description, Body string
	Tags                                     []string
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
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
	indexTemplate := `<ol>{{range .}}<li><a href="/post/{{sanitizeTitle .Title}}">{{.Title}}</a></li>{{end}}</ol>`
	templ, err := template.New("index").Parse(indexTemplate)
	if err != nil {
		return err
	}

	if err := templ.ExecuteTemplate(w, "index.gohtml", posts); err != nil {
		return err
	}
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
