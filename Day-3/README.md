# Day 3 Notes: Structs, Methods, Interfaces

<hr />

## Syntax

- in `%.2f`: `f` stands for `float64` and the `.2` means print 2 decimal places.
- `%g` prints more precise decimal number
- Methods are very similar to functions but *they are called by invoking them on an instance of a particular type*
- The syntax of declaring methods: `func (receiverName ReceiverType) MethodName(args)`
- In Go **interface resolution is implicit**. If the type passed in matches what the interface is asking for, it will compile.

## Testing

- `Table driven tests` are useful when you want to build a list of test cases that can be tested in the same manner.
  
## Best practices

- It is a convention in GO to have the received variable in methods the first letter of the type. E.g: `c Circle`
