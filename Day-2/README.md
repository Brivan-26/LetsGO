# Day 2 Notes: Arrays, Slices

<hr />

## Syntax

- Arrays have a fixed capacity which must be defined on variable declaration
- `%v` to print the *default* format (works well with arrays)
- `range` is a way to iterate over an array. On each iterations, it returns *two values*: the **index** and the **value**
- `_` is used to ignore a return value.
- Slices can be sliced via the syntax: `slice[low:high]`. ***Low* inclusive, and *high* exclusive**.
- A function can be assigned to a variable. *The order of declaration matters*
-**modifying the elements (not the slice itself) of a re-slice modifies the elements of the original slice**


## Testing

- We can not compare directly two slices, we can use `reflect.DeepEqual` to compare *any two variables*


## Best practices

- It should not be a goal to have as many tests as possible, but rather to have as much confidence as possible
- Assigning a function to a variable is a technique that can be used to hide the function to be used from outside functions
- `reflect.DeepEqual` reduces the *type-safety*