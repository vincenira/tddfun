package fun17_test

import (
	"testing"
	"testing/fstest"
)

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
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts := fun17.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
