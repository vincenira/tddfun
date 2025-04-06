package fun8

import (
	"fmt"
	"io"
	"net/http"
)

func Musalimie(writer io.Writer, jina string) {
	fmt.Fprintf(writer, "Jambo, %s", jina)
}

// Where else we can send information to print to a different media

// Handler handles the response to the request to Musalimie
func MusalimieHandler(w http.ResponseWriter, r http.Request) {
	Musalimie(w, "dunia")
}

/*
func main(){
	 Musalimie(os.Stdout, "Mufasa")
   log.Fatal(http.ListenAndServe(":5001", http.HnadlerFunc(MusalimieHandler)))
}
*/
