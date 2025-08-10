# Acceptance test

## When you have grace

what we want to do is listen for SIGTERM, and rather than instantly killing the server, we want to:

- Stop listening to any more requests
- Allow any in-flight requests to finish
- Then terminate the process

## How to have grace

[net/http/Server.Shutdown](https://pkg.go.dev/net/http#Server.Shutdown).
Shutdown gracefully shuts down the server without interrupting any active connections.
Shutdown works by first closing all open listeners, then closing all idle connections, and then
waiting indefinitely for connections to return to idle and then shut down. if the provided context
expires before the shutdown is complete, shutdown returns the context's error, otherwise it returns
any error returned from closing the Server's underlying Listener(s).

To handle `SIGTERM` we can use [os/signal.Notify](https://pkg.go.dev/os/signal#Notify), which will send any incoming signals
to a channel we provide.
With those two features from the standard library, you can listen for `SIGTERM` and shutdown gracefully

```go
func main() {
	var (
		ctx        = context.Background()
		httpServer = &http.Server{Addr: ":8080", Handler: http.HandlerFunc(acceptancetests.SlowHandler)}
		server     = gracefulshutdown.NewServer(httpServer)
	)

	if err := server.ListenAndServe(ctx); err != nil {
		// this will typically happen if our responses aren't written before the ctx deadline, not much can be done
		log.Fatalf("uh oh, didn't shutdown gracefully, some responses may have been lost %v", err)
	}

	// hopefully, you'll always see this instead
	log.Println("shutdown gracefully! all responses were sent")
}
```

## what are they?

Acceptance tests sometimes also referred to as "functional tests" are a kind of "black-box test". They should
exercise the system as a user of the system would.
The term "black-box" refers to the idea that the test code has no access to the internals of the system, it
can only use its public interface and make assertions on the behaviours it observes. This means they can only
test the system as a whole.
This is an advantageous trait because it means the tests exercise the system the same as a user would,
it can't use any special workarounds that could make a test pass, but not actually prove what you
need to prove. This is similar to the principle of preferring your unit test files to live inside
a separate test package, for example, package `mypkg_test` rather than package `mypkg`.

## Benefits of acceptance tests:

- When they pass, you know your entire system behaves how you want it to.
- They are more accurate, quicker, and require less effort than manual testing.
- When written well, they act as accurate, verified documentation of your system. It doesn't fall into the trap of documentation that diverges from the real behaviour of the system.
- No mocking! It's all real.

## Potential drawbacks vs unit tests

- They are expensive to write
- they take longer to run.
- they are dependent on the design of the system.
- When they fail, they typically don't give you a root cause, and can be difficult to debug.
- They don't give you feedback on the internal quality of your system. You could write total garbage and still
  make an acceptance test pass
- Not all scenarios are practical to exercise due to the black-box nature.

For this reason, it is foolish to only rely on acceptance tests. They do not have many of the
qualities unit tests have, and a system with a large number of acceptance tests will tend
to suffer in terms of maintenance costs and poor [lead time](## Lead time).

## Lead Time?

Lead time refers to how long it takes from a commit being merged into your main branch to it being
deployed in production. This number can vary from weeks and even months for some teams to a matter
of minutes.
A balanced testing approach is required for a reliable system with excellent lead time, and this is
usually describe in terms of the [Test Pyramid](https://martinfowler.com/articles/practical-test-pyramid.html)
