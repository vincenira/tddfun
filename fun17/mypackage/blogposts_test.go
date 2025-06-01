package mypackage_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"fun17/mypackage"
)

/*
Add new test where we inject a failing fs.FS test-double to make fs.ReadDir return an error
*/
type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

// later
//_, err := mypackage.NewPostsFromFS(StubFailingFS{})

/*
Notice that the package of our test is blogposts_test. Remember, when TDD is practiced well
we take a consumer-driven approach: we don't want to test internal details
because consumers don't care about them. By appending _test to our intended package
name, we only access exported members from our package - just like a real user of
our package.
*/
func TestNewPosts(t *testing.T) {
	// We've import testing/fstest which gives us access  to the fstest.MapFS type.
	// Our fake file system will pass fstest.MapFS to our package
	/*
	   A MapFS is a simple in-memory file system for use in tests, represented as a map from path
	   names (arguments to Open) to information about he files or directories they represent
	*/
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	/*
	   To acknowledge that errors can happen when working with files. let's handle error
	*/

	posts, err := mypackage.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}

// Write the test first
func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
		Description: Description 1`
		secondBody = `Title: Post 2
		Description: Description 2`
	)
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := mypackage.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, posts[0], mypackage.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body:        `HelloWorld`,
	})
}

func assertPost(t *testing.T, got mypackage.Post, want mypackage.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
