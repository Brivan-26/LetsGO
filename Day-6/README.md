# Day 6 Notes: Dependency Injection

<hr />

## Syntax

- both `os.Stdout` and `bytes.Buffer` implement `io.Writer`
  

## Best practices

- Dependency Injection allows to write general purposes functions.
- If a function can't be tested easily, it's usually because of dependencies *hard-wired* into a function, so we use DI to solve that.
