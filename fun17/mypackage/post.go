package mypackage

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

/*
From now on, most of our efforts can be neatly contained within newPost. The concerns of opening and iterating over files are done, and now we can focus on extracting the data for our Post type. Whilst not technically necessary, files are a nice way
to logically group related things together, so I moved the Post type and newPost into a new post.go file.
*/

// Tightly coupled to the fs.File type(is it necessary?)
func oldnewPost(postFile fs.File) (Post, error) {
	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}
	post := Post{Title: string(postData)[7:]}
	return post, nil
}

// loosen version of newPost function from fs.File typ argument to io.Reader
func loosenNewPost(postFile io.Reader) (Post, error) {
	/*
		Let's use the standard library bufio.Scanner to scan line by line through data
	*/
	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}
	post := Post{Title: string(postData)[7:]}
	return post, nil
}

// using bufio.scanner to read line by line
const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
)

func bufioScannerNewPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)
	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	title := readLine()[len(titleSeparator):]
	description := readLine()[len(descriptionSeparator):]
	return Post{Title: title, Description: description}, nil
}

func newPost(postBody io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postBody)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tags := strings.Split(readMetaLine(tagSeparator), ", ")

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprint(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
