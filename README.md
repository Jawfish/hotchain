hotchain is a Go package that offers functional utilities for slices, allowing for a more declarative approach to data manipulation.

## Features

- `All(fn)`: Checks if all elements in the slice satisfy `fn`.
- `Any(fn)`: Checks if any element in the slice satisfies `fn`.
- `Each(fn)`: Applies `fn` to each element in the slice purely for its side effects (does not produce a new slice).
- `Filter(fn)`: Filters the slice, keeping only those elements that satisfy `fn`.
- `Pfilter(fn)`: Like `Filter`, but operations are parallelized using goroutines.
- `Find(fn)`: Finds the first element in the slice that satisfies `fn`.
- `Map(fn)`: Applies `fn` to each element in the slice, producing a new slice with the results.
- `Pmap(fn)`: Like `Map`, but operations are parallelized using goroutines.
- `Reduce(fn, initial)`: Performs a reduction on the slice elements using `fn`, starting from an initial value.

`Pmap` and `Pfilter` are guaranteed to preserve the order of the elements in the slice. Use these judiciously, as they will often not be faster than their synchronous counterparts. Use them if `fn` is computationally expensive or otherwise takes a long time to execute (e.g. network requests, database queries, etc.)



## Usage

```go
import "hotchain"

func main() {
    chain := hotchain.Chain[int]{1, 2, 3, 4, 5}
    result := chain.
        Map(func(x int) int { return x * x }). // Squares each number
        Filter(func(x int) bool { return x > 10 }). // Keeps only numbers greater than 10
        Reduce(func(x, y int) int { return x + y }, 0) // Sums up the remaining numbers
    fmt.Println(result) // Outputs: 41

    allEven := chain.All(func(x int) bool { return x%2 == 0 }) // Checks if all numbers are even
    fmt.Println(allEven) // Outputs: false

    hasOdd := chain.Any(func(x int) bool { return x%2 != 0 }) // Checks if there's any odd number
    fmt.Println(hasOdd) // Outputs: true

    chain.Each(func(x int) { fmt.Println(x) }) // Prints each number to stdout

    concurrentResult := chain.
        Pmap(func(x int) int { return x * x }). // Squares each number in parallel
        Pfilter(func(x int) bool { return x > 10 }) // Keeps only numbers greater than 10 in parallel
    fmt.Println(concurrentResult) // Outputs: [16 25]
}
```
