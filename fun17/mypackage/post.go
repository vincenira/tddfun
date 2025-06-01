package mypackage

import (
	"io"
	"io/fs"
)

type Post struct {
	Title       string
	Description string
}

/*
From now on, most of our efforts can be neatly contained within newPost. The concerns of opening and iterating over files are done, and now we can focus on extracting the data for our Post type. Whilst not technically necessary, files are a nice way
to logically group related things together, so I moved the Post type and newPost into a new post.go file.
*/

// Tightly coupled to the fs.File type(is it necessary?)
func newPost(postFile fs.File) (Post, error) {
	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}
	post := Post{Title: string(postData)[7:]}
	return post, nil
}

// loosen version of newPost function from fs.File typ argument to io.Reader
func loosenNewPost(postFile io.Reader) (Post, error) {
	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}
	post := Post{Title: string(postData)[7:]}
	return post, nil
}
