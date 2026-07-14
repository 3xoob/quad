# quad

A small Go package that draws quadrilateral (rectangle) shapes to standard output, using different border and fill characters for each variant.

## Overview

The repository is a Go module named `piscine` that implements five functions — `QuadA` through `QuadE` — each of which prints a rectangle of a given width and height to stdout, character by character, via the `github.com/01-edu/z01` package. A `test/main.go` file is included as a manual entry point that imports the `piscine` package and calls one of the functions directly.

## Features

- `QuadA(x, y int)` — draws a rectangle using `o` corners, `-` for the top/bottom edges, and `|` for the sides, with spaces filling the interior.
- `QuadB(x, y int)` — draws a rectangle using `/` and `\` corners, `*` for edges, and spaces for the interior.
- `QuadC(x, y int)` — draws a rectangle using `A`/`B` on top, `C`/`B` on bottom, and `B` for the sides (symmetric corners).
- `QuadD(x, y int)` — draws a rectangle using `A`/`C` corners on both the top and bottom rows, with `B` for edges/sides.
- `QuadE(x, y int)` — draws a rectangle using `A`/`C` on top and `C`/`A` on bottom (diagonal corner pattern), with `B` for edges/sides.
- Each function is a no-op unless both `x > 0` and `y > 0`.
- Each function handles the `x == 1` case by omitting the closing corner/edge character on each row.

## Technologies

- Go (module declares `go 1.21.4` in `go.mod`)
- [`github.com/01-edu/z01`](https://github.com/01-edu/z01) — external dependency used for `z01.PrintRune`, the character-printing primitive all five functions rely on.

## Project Structure

- `quadA.go`, `quadB.go`, `quadC.go`, `quadD.go`, `quadE.go` — the five function implementations, all in `package piscine`.
- `test/main.go` — a standalone `package main` program that imports `piscine` and calls `QuadE(1, 5)`.
- `go.mod`, `go.sum` — Go module definition and dependency checksums.
- `LICENSE`, `COPYRIGHT.md` — copyright notice.

## Requirements

- Go 1.21.4 or later.

## Installation

```bash
git clone https://github.com/3xoob/quad.git
cd quad
go mod download
```

## Usage

The root of the module is a library (`package piscine`) with no `main` package, so it is meant to be imported and its functions called directly. The included `test/main.go` demonstrates this:

```bash
go run ./test
```

This runs `test/main.go`, which calls `piscine.QuadE(1, 5)`.

To use the functions from another Go program within this module, import the package by its module path (`piscine`) and call any of `QuadA` through `QuadE` with a width (`x`) and height (`y`):

```go
package main

import "piscine"

func main() {
	piscine.QuadA(4, 3)
}
```

## License

This project includes a `LICENSE` file and a `COPYRIGHT.md` file. The source code is publicly available for portfolio and viewing purposes only; see those files for the full copyright terms.
