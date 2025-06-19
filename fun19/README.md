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

## The problem with throwing out type safety

The first Problem is the same as we saw with our `AssertEqual` we've lost type safety. I can now `Push`
apples onto a stack of oranges.

Enven if we have the discipline not to do this, the code is still unpleasant to work with because when
methods _return interface{} they are horrible to work with._
Add the following test,

```go
t.Run("interface stack DX is horrid", func(t *testing.T) {
	myStackOfInts := new(StackOfInts)

	myStackOfInts.Push(1)
	myStackOfInts.Push(2)
	firstNum, _ := myStackOfInts.Pop()
	secondNum, _ := myStackOfInts.Pop()
	AssertEqual(t, firstNum+secondNum, 3)
})
```

You get a compiler error, showing the weakness of losing type-safety:
`invalid operation: operator + not defined on firstNum(variable of type interface{})`

When `Pop` returns `interface{}` it means the compiler has no information about what the data is
and therefore severely limits what we can do. It can't know that it should be an integer, so it does not
let us use the `+` operator.

To get around this, the caller has to do a [type assertion](https://golang.org/ref/spec#Type_assertions) for value

```go
t.Run("interface stack dx is horrid", func(t *testing.T) {
	myStackOfInts := new(StackOfInts)

	myStackOfInts.Push(1)
	myStackOfInts.Push(2)
	firstNum, _ := myStackOfInts.Pop()
	secondNum, _ := myStackOfInts.Pop()

	// get our ints from out interface{}
	reallyFirstNum, ok := firstNum.(int)
	AssertTrue(t, ok) // need to check we definitely got an int out of the interface{}

	reallySecondNum, ok := secondNum.(int)
	AssertTrue(t, ok) // and again!

	AssertEqual(t, reallyFirstNum+reallySecondNum, 3)
})
```

The Unpleasantness radiating from this test would be repeated for every potential user of our `Stack`
implementation, Yuck.

Using a generic data type we have:

- Reduced duplication of important logic.
- Made `Pop` return `T` so that if we create a `Stack[int]` we in practice get back `int` from `Pop`;
  we can now use `+` without the need for type assertion gymnastics.
- Prevented misuse at compile time. You cannot `Push` oranges to an apple stack
