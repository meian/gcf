# gcf


gcf (Go Colletion Framework) is a library that provides various collection operations using Generics.  
By operating on the collection using a common interface, you can easily composite the operations.

[![Go Reference](https://pkg.go.dev/badge/github.com/meian/gcf.svg)](https://pkg.go.dev/github.com/meian/gcf)
[![codecov](https://codecov.io/gh/meian/gcf/branch/main/graph/badge.svg?token=PDHAVSGE0E)](https://codecov.io/gh/meian/gcf)
[![Go Report Card](https://goreportcard.com/badge/github.com/meian/gcf)](https://goreportcard.com/report/github.com/meian/gcf)


## Motivation

I wanted a functions that allows Go to easily composite and use processes instead of using basic syntax such as for and if.

Until now, it was difficult to provide processing to multiple types with the same interface, but since the support for Generics in Go 1.18 made this implementation easier, gcf was actually built as a library.

## Example

Take as an example the process of `extracting only odd numbers from the elements of a slice and returning those numbers by 3 times`.

```golang
func Odd3(s []int) []int {
    var r []int
    for _, v := range s {
        if v%2 == 0 {
            continue
        }
        r := append(v*3)
    }
    return r
}
```

When using gcf, implement as follows.

```golang
// var s []int

itb := gcf.FromSlice(s)
itb = gcf.Filter(itb, func(v int) bool {
    return v%2 > 0
})
itb = gcf.Map(itb, func(v int) int {
    return v * 3
})

// Get the processing result as a slice
r := gcf.ToSlice(itb)
```

This example is meant to show how to use it briefly.  
Replacing inline processing with gcf will significantly reduce the performance, so it is not recommended to rewrite processing that can be easily implemented and managed inline using gcf.

## Environment

- Go 1.18 or 1.19

Since gcf uses Generics feature, version 1.18 or higher is required.  
We also have a container usage environment for vscode that you can use if you do not want to install new Go version in your local environment.  
(See below [.devcontainer](https://github.com/meian/gcf/tree/main/.devcontainer))

## Installation

Install by using `go get` on the directory under the control of the Go module.

```bash
go get -d github.com/meian/gcf
```

## Design

### Implements by Iterator

gcf is designed to composite processing with the `Iterator` pattern.  
Some processes may allocate memory internally, but most processes avoid unnecessary memory allocations in the middle of the process.

### Iterable + Iterator

Each function returns `Iterable[T]`, which only has the ability to generate `Iterator[T]` by `Iterator()`.  
`Iterator[T]` moves the element position next by `MoveNext()`, and gets current element by `Current()`.  
Functions is composited by `Iterable[T]`, and the state is keep only in the `Iterator[T]`, which makes it easy to reuse the generated composition.

### MoveNext + Current

In Iterator implementation, you may see set of functions that uses `HasNext()` to check for the next element and `Next()` to move to the next element and return the element.  
In gcf, we implemented that `MoveNext()` moves to the next element and returns move result, and `Current()` returns the current element.  
This is because we took advantage to get current value multiple times without changing, rather than providing next check without changing.

### Top-level functions

In libraries of collection operations in other languages, the processing is often defined by methods so that we can use method chain, but we implemented by top-level functions.  
This is because generics in Go cannot define type parameters at the method level, so some functions cannot be provided by methods, and if only some functions are provided by methods, the processing cannot be maintained consistency.  
If it is implemented to define type parameters at the method level as a result of future version upgrades of Go, we will consider providing method chain functions.  

### Thread Safety

The exported functions that generate Iterables are thread-safe and can be accessed from multiple go routines at the same time.  
The `Iterator()` in each Iterable is also thread-safe and can be called from multiple go routines at the same time.  
In `MoveNext()` and `Current()` in each Iterator, thread-safe is not guaranteed , so when sharing between go routines, separate the processing with mutex etc. as needs.  
The process of retrieving elements from Iterable, such as `ToSlice()`, is thread-safe.

It is not currently implemented, but it is undecided how the thread-safe will change when functions for channel is implemented.

## Performance

The performance of gcf has the following characteristics.

- It takes a processing time proportional to the number of elements in the collection and the number of processes to be combined.
- Overwhelmingly slower than in-line processing (about 70 times)
- About 4 times slower than function call without allocation
- Overwhelmingly faster than channel processing (about 60 times)

Due to the characteristics of the library, it is processed repeatedly, so it is not recommended to use it for processing that requires severe processing speed.

Please refer to the [Benchmark README](bench/README.md) for details.

## Under the consideration

- channel function implements
  - Create Iterable from channel
  - Get the result of Iterable on channel
  - It's suspended cause I don't understand good design.
- `Zip`
  - Combine multiple Iterable elements into one Iterable.
  - We plan to refer to other implementations for how to handle when the number of elements of each Iterable is different.

----


## Configure README Display Language

- [日本語](README.ja.md)