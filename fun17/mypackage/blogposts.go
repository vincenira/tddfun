package mypackage

import (
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
	for range dir {
		posts = append(posts, Post{})
	}
	return posts, nil
}
