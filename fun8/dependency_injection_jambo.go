package fun8

import (
	"fmt"
	"io"
)

func Musalimie(writer io.Writer, jina string) {
	fmt.Fprintf(writer, "Jambo, %s", jina)
}
