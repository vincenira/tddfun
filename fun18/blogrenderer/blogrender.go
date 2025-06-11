package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown/parser"
	"github.com/hashicorp/hcl/json/parser"
)

// if you're continuing from the read files chapte, you shouldn't redefine this

type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
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

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	return &PostRenderer{templ: templ, mdParser: parser}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", newPostVM(p, r)); err != nil {
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

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(p Post, r *PostRenderer) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}
