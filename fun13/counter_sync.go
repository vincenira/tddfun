package fun13

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

/*
Sometimes people forget that embedding types means the methods of that type become
part of the public interface; and you often will not want that. Remember that we should
be very careful with our public APIs, the moment we make something public is the moment
other code can couple themselves to it. We always want to avoid unnecessary coupling.

Exposing Lock and Unlock is at best confusing but at worst potentially very harmful to your software
if callers of your type start calling these methods.
sync/v2/sync_test.go:16: call of assertCounter copies lock value: v1.Counter contains sync.Mutex
sync/v2/sync_test.go:39: assertCounter passes lock by value: v1.Counter contains sync.Mutex
*/
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
