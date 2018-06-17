[![Build Status](https://travis-ci.org/direvus/sudoku.svg?branch=master)](https://travis-ci.org/direvus/sudoku)
[![codecov](https://codecov.io/gh/direvus/sudoku/branch/master/graph/badge.svg)](https://codecov.io/gh/direvus/sudoku)


# Sudoku - すどく

https://github.com/direvus/sudoku

A small sudoku library written in Go.

I wrote this library in order to learn the Go programming language.  It is not
intended to be useful, or expected to be novel.

The original inspiration for the library and the input format comes from
Andrew Gerrand's November 2015 Go Challenge
[http://golang-challenge.org/go-challenge8/], although this library is not
written as a submission for that challenge.


## Build instructions

- Install go v1.10.x
- Checkout source
- Execute `make build`
- Copy the resulting binaries wherever you like
- Use them!


## Usage instructions

### sudoku-solve

The `sudoku-solve` executable takes a sudoku puzzle on stdin, attempts to solve it,
and produces the result on stdout.

The input format is one line per row, each line terminated by a single newline
character (0x0a).  Cells within the row are delimited by a single space
character (0x20).  Unknown cells are indicated by the underscore character
(0x5f).

So, for example:

	_ _ 8 _ _ 6 2 5 _
	_ _ _ _ 7 _ _ 3 _
	_ _ _ _ 1 2 9 8 _
	_ _ 5 _ _ 3 _ _ _
	_ 2 _ 7 _ 1 _ 6 _
	_ _ _ 8 _ _ 1 _ _
	_ 3 6 2 8 _ _ _ _
	_ 7 _ _ 9 _ _ _ _
	_ 8 2 1 _ _ 4 _ _


## License

This library is released under the terms of the BSD 2-clause license, a copy of
which can be found in the file 'LICENSE' at the root directory of this
repository.
