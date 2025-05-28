# Reading files

[source of the articles](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reading-files)

1. Example data

```md
Title: Hello, TDD world!

Description: First post on our wonderful blog
Tags: tdd, go

---

Hello world!

The body of posts starts after the `---`
```

2. Expected data

```go
type Post struct{
    Title, Description, Body string
    Tags                     []string
}
```

## Thinking about the kind of test we want to see

Let's remind ourselves of our mindset and goals when starting:

- _Write the test we want to see_. Think about how we'd like to use the code we're going to write
  from a consumer's point of view.
- Focus on what and why, but don't get distracted by how.

In this tutorial, we want that our package provides a function that can be pointed at function,
and return us some posts.

is it the what? To get some posts from a given folder.
why?: so that it can be posted to the webserver to be viewed

```go
var posts []blogposts.Post
posts = blogposts.NewPostsFromFS("some-folder")
```

To write test around this, we'd need some kind of test folder with some example posts in it.
There is nothing terribly wrong with this, but you are making some trade-offs:

- for each test you may need to create new files a particular behaviour
- some behaviour will be challenging to test, such as failing to load files
- the tests will run a little slower because they will need to access the file system

Note that we are also unnecessarily coupling ourselfves to a specific implementation of the filesystem

## File system abstractions introduced in Go 1.16

- [io/fs](https://golang.org/pkg/io/fs/) package gives us an abstraction for file system

note: Package fs defines basic interfaces to a file system. A file system can be provided by the host operating
system but also by other packages.

This lets us loosen our coupling to a specific file system, which will then let us inject different implementations
according to our needs
[On the producer side of the interface, the new embed.FS type implements fs.FS, as does zip.Reader. The new os.DirFS function provides an implementation of fs.FS backed by a tree of operating system files.](https://golang.org/doc/go1.16#fs)

If we use this interface, users of our package have a number of options baked-in to the standard library to use.
Learning to leverage interfaces defined in Go's standard library (e.g. io.fs, [io.Reader](https://golang.org/pkg/io/#Reader), [io.Writer](https://golang.org/pkg/io/#Writer)), is vital
to writing loosely coupled packages. These packages can then be re-used in contexts different to those you imagined,
with minimal fuss from your consumers.

In this case, maybe our consumer wants the posts to be embedded into Go binary rather than files in a "real" filesystem?
Either way, our code doesn't need to care.

Here, the package for test is [testing/fstest](https://golang.org/pkg/testing/fstest/) to implement [io/FS](https://golang.org/pkg/io/fs/#FS)
it is similar to the tools we're familiar in [net/http/httptest](https://golang.org/pkg/net/http/httptest/).

## Write the test first

Create a new project to work through this chapter.

- `mkdir blogposts`
- `cd blogposts`
- `go mod init github.com/{yourName}/blogposts`
- `touch blogposts_test.go`
