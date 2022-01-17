# gcf

gcf (Go Colletion Framework) is a library that provides various collection operations using Generics.  
By operating on the collection using a common interface, you can easily composite the operations.

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
r := itb.ToSlice(itb)
```

This example is meant to show how to use it briefly.  
Replacing inline processing with gcf will significantly reduce the performance, so it is not recommended to rewrite processing that can be easily implemented and managed inline using gcf.

## Environment

- Go 1.18 RC (or Beta)

Since gcf uses Generics, this version is currently in Beta status, but you need version 1.18 to use it.  
We also have a container usage environment for vscode that you can use if you do not want to install Go 1.18 in your local environment.  
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

## Performance

The performance of gcf has the following characteristics.

- It takes a processing time proportional to the number of elements in the collection and the number of processes to be combined.
- Overwhelmingly slower than in-line processing (about 70 times)
- About 4 times slower than function call without allocation
- Overwhelmingly faster than channel processing (about 60 times)

Due to the characteristics of the library, it is processed repeatedly, so it is not recommended to use it for processing that requires severe processing speed.

Please refer to the [Benchmark README](bench/README.md) for details.

## Function

### Implemented

The following functions are implemented.  
See function comments for feature details.  
There are some implementations for which comments have not been described, but we will add them in the future.

- `FromSlice`
  - Also implemented as an immutable version, `FromSliceImmutable`.
- `Filter`
- `Map`
- `Concat`
- `Repeat`
- `RepeatIterable`

### To Be

- `Range`
  - Specify from, to, step and return the values in order
  - Only numeric type will be provided
- `Reverse`
  - Returns in reverse order
- `Distinct`
  - Returns unique elements
- `FlatMap`
  - Maps and returns flatten iterator
- `Sort`
  - Sort collections
- `OrderBy`
  - Sort collections by specified criteria
- `Take`
  - Returns only the specified number from the beginning
- `Last`
  - Returns only the specified number from the end
- `Skip`
  - Excluding the number specified from the beginning
- `SkipLast`
  - Excluding the specified number from the end
- channel function
  - Create Iterable from channel
  - Get the result of Iterable on channel
