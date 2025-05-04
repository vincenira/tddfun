# property based tests

## what is a Roman Numeral Kata

Kata s a small, focused exercise designed to help developers practice specific skills
and improve their coding abilities through repetition and refinement.

The main idea for software developers

> As this book stresses, a key skill for software developers is to try and identify "thin vertical slices"
> of useful functionality and then iterating. The TDD workflow helps facilitate iterative development

Therefore, we should start by solving 1 in 1984

## Exercises

Can you think of a way of making it so it's impossible for someone to call our code with a number
greater than 3999?

- you could return an error
- Or create a new type that cannot represent > 3999
  - what do you think is best?

## Property based tests

- Built into the standard library
- If you can think of ways to describe your domain rules in code, they are an excellent tool for giving
  you more confidence
- Force you to think about your domain deeply
- Potentially a nice complement to your test suite
