# Maths

## State the problem.

Make an SVG of a clock. Not a digital clock. Just function that takes a Time from the time package
and spits out an SVG of a clock with all the hands - hour, minute and second pointing in
the right direction

SVGs are fantastic image format to manipulate programmatically because they're written as a series of
shapes, described in XML.

## An Acceptance Test

The idea is that you write a really high level test to describe what you're trying to achieve - a user clicks a button
on a website, and they see a complete list of the Pokémon they've caught, for instance.
When we've written that test, we can then write more tests - unit tests - that build towards
a working system that will pass the acceptance test.
So for our example these tests might be about rendering a webpage with a button, testing route
handlers on a web server, performing database look ups, etc.
All of these things will be TDD'd, and all of them will go towards making the original acceptance test pass.
