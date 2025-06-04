package blogrenderer

import "io"

// if you're continuing from the read files chapte, you shouldn't redefine this

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func Render(w io.Writer, p Post) error {
	return nil
}
