package mypackage

import (
	"io"
	"io/fs"
)

/*
write the minimal amount of code for the test to run and check the failing test output
add the new field to our Post type so that the test will run
*/
type Post struct {
	Title string
}

// Write the minimal amount of code for the test to run and check the failing test output
/*
func NewPostsFromFS(fileSystem fstest.MapFS) []Post {
	return nil
}
*/

// Write enough code to make it pass
//https://deniseyu.github.io/leveling-up-tdd/ (we could slime this to make it pass)
/*
Sliming is useful for giving a “skeleton” to your object. Designing an interface and
executing logic are two concerns, and sliming tests strategically lets you focus on one
at a time.
*/
/*
Refactor1: We can't use our new package outside of this context, because it is coupled to a concrete implementation
(fstest.MapFS).but it doesn't have to be. Change the argument to our NewPostsFromFS function
to accept the interface from the standard library.
*/
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	// fs.ReadDir reads a directory inside a given fs.FS return []DirEntry
	/*
	   Already our idealised view of the world has been foiled because errors can happen,
	   but remember now our focus is making the test pass, not changing design, so we'll ignore
	   the error for now.

	   The rest of the code is straightforward: iterate over the entries, create a Post for each
	   one and, return the slice.
	*/
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := loosenSeparatedGetPost(fileSystem, f)
		if err != nil {
			return nil, err // todo: needs clarification, should we totally fail if one file fails? or just ignore?
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// This code should be separated from opening file code and from parsing file contents code
// to make the code simpler to understand and work with
func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}

	// for now, we do not need any sophisticated parsing,
	// just cutting out the `Title:` text by slicing the string.
	post := Post{Title: string(postData)[7:]}
	return post, nil
}

/*
Think about coupling and cohesion. In this case you should ask youreslef:
Does newPost have to be coupled to an fs.File? Do we use all the methods and data from this
type? what do we really need?

In our case we only use it as an argument to io.ReadAll which needs an io.Reader. so we should
loosen the coupling in our function and ask for an io.Reader.

So we are going to loosen the function newPost by using the io.Reader type
*/
func separatedGetPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

func loosenSeparatedGetPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return loosenNewPost(postFile)
}
