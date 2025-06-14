# Fun with Generics

## Is a generic function with T any the same as interface{}?

consider two functions

`func GenericFoo[T any](x, y T)`
`func InterfaceyFoo(x, y interface{})`

what's the point of generics here? Doesn't `any` describe... anything?

In terms of constraints, `any` does mean "anything" and so does `interface{}`. In fact, any was
added in 1.18 an is just an alias for `interface{}`.

The difference with the generic version is you're still describing a specific type and what tha means
is we've still constrained this function to only work with one type.

what this means is you can call `InterfaceyFoo` with any combination of types (e.g`InterfaceyFoo(apple, orange)`).
However `GenericFoo` still offers some constraints because we've said that it only works with one type, `T`.

valid:

- GenericFoo(apple1, appl2)
- GenericFoo(orange1, orange2)
- GenericFoo(1, 2)
- GenericFoo("one", "two")

Not valid(fails compilation):

- GenericFoo(apple1, orange1)
- GenericFoo("1", 1)

If you function returns the generic type, the caller can also use the type as it was, rather than having
to make a type assertion because when a function returns `interface{}` the compiler cannot make
any guarantees about the type.
