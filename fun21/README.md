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
