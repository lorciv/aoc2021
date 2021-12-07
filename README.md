# Advent of Code 2021 🎄

Here are my solutions to the [Advent of Code 2021](https://adventofcode.com/2021). I am writing my solutions in Go. I try to keep them as simple as possible.

Below are some comments.

## Day 7

I wrote the `fuel` function as recursive, however I suspect that an iterative version may be more efficient. Also, with the given input, the function can call itself more than 1000 times, which is not so far from the limit of around 8000 that is set by default on my computer. An iterative version would avoid the risk of stack overflow.
