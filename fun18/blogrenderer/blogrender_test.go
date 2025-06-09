package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	"fun18/blogrenderer"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	aPost := blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}
	t.Run("It converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := blogrenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

/*
To see the impact of not doing this parsing over and over, we can use the benchmarking tool
to sse how fact our function is.
*/
func BenchmarkRender(b *testing.B) {
	aPost := blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blogrenderer.Render(io.Discard, aPost)
	}
}
