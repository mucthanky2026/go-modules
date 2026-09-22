# Go Sorting Algorithms

Small, dependency-free implementations of common sorting algorithms in Go.

## Requirements

- Go 1.23.1 or newer

## Installation

Import the package from the module path:

```bash
go get github.com/mucthanky2026/go-modules
```

## Usage

```go
package main

import (
	"fmt"

	sortalg "github.com/mucthanky2026/go-modules/sort"
)

func main() {
	values := []int{5, 2, 8, 1, 3}
	sorted := sortalg.MergeSort(values)

	fmt.Println(sorted) // [1 2 3 5 8]
}
```

The package name is `sort`, so the import is aliased above to avoid confusion
with Go's standard library `sort` package.

## Algorithms

| Function | Average time | Worst-case time | Extra space | Notes |
| --- | ---: | ---: | ---: | --- |
| `BubbleSort` | $O(n^2)$ | $O(n^2)$ | $O(1)$ | Sorts the input slice in place |
| `InsertionSort` | $O(n^2)$ | $O(n^2)$ | $O(1)$ | Sorts the input slice in place |
| `SelectionSort` | $O(n^2)$ | $O(n^2)$ | $O(1)$ | Sorts the input slice in place |
| `MergeSort` | $O(n \log n)$ | $O(n \log n)$ | $O(n)$ | Returns a sorted result using additional storage |
| `QuickSort` | $O(n \log n)$ | $O(n^2)$ | $O(n)$ | Uses the middle element as the pivot |

All functions accept and return `[]int` values in ascending order.

## Development

Run the module's tests with:

```bash
go test ./...
```

Format the source code with:

```bash
gofmt -w sort/*.go
```