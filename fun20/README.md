# Fun20: Revisiting arrays and slices with generics

## Reduce pattern

Reduce captures the essence of the pattern, it's a funciton that takes a collection, an accumulating function,
an initial value, and returns a single value. There's no messy distractions around concrete types.

```go
func Reduce[A any](collection []A, f func(A, A) A, initialValue A) A {
    var result = initialValue
    for _, x := range collection {
        result = f(result, x)
    }
    return result
}
```
