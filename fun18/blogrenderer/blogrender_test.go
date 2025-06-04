package blogrenderer_test

import (
	"bytes"
	"testing"

	"fun17/blogrenderer"
)

func TestRender(t *testing.T) {
	aPost := blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is description",
		Tags:        []string{"go", "tdd"},
	}
	t.Run("It converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world</h1>`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
