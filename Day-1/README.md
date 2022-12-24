# Day 1 Notes: Integers, Iterations
<hr />

## Syntax

- If a function accepts many arguments with the same type, we can shorten the type. E.g: `(x, y int)` instead of `(x int, y int)`
- **%d** to format integers
- In GO, there are no `while`, `do`, `until`. To do things repeatedly, there's only `for`
- There are no parentheses surrounding the three components of the for statement.

## Testing

- Writing **Benchmarks** in Go is another first-class feature of the language, and it is very similar to writing tests, it gives details on how a function would take time to be executed over N iterations.
- To run a benchmark test: `go test -bench=.`


## Best practices

- *Named return values* should generally be used when the meaning of the result is not clear from the context. But for add example, **Add** clearly means that we're adding two integers, so no need for named return variables.
- Comments will be appear on Go DOC, it's better to include them
- Go examples are executed just like tests, they are compiled (and optionally executed) as part of a package's test suite. They also appear in the documentation inside godoc.
- The example function will not be executed if we remove the comment **// Output: ..**
