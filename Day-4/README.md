# Day 4 Notes: Pointers, Errors

<hr />

## Syntax

- In Go if a symbol (variables, types, functions et al) starts with a lowercase symbol then it is private outside the package it's defined in
- In Go, **when calling a function or a method, the arguments are *copied***
- To get the memory address of something, we add `&` at the beginning of the symbol
- In Go, we can create new types from existing ones using the following syntax: `type newName originalName`
- The `var` keyword allows to define values globally on the package

## Testing

- `t.Fatal` will stop the test if it is called
  
## Best practices

- When a function returns a pointer to something, we need to make sure we check if it's nil or we might raise a runtime exception - **the compiler won't help here**.
- **Do not just check for errors, handle them perfectly**
