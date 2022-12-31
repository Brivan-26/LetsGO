# Day 5 Notes: Maps

<hr />

## Syntax

- The syntax `map[key_type]value_type` is used to declare a map.
  - The `key_type` **MUST** be comparable type. 
  - The `value_type` can be any type, even an other map.
- The `map[key]` returns two values: `value, found` where `found` is a boolean value that indicates whether the key exists or not
- If a *key-value paires* are attempted to be added on a map, and the **key already exists**, **The old value will be overrided by the new one**
- `delete` is a built-in function that used to delete a key from a map, it accepts two arguments: the **map**, and the **key** to be deleted
  
## Testing

- **Errors** can be converted to strings via the syntax `.Error()`

## Best practices

- It is better to set the error messages as **constants** so the tests won't fail after changing them.
- **Never initialize an empty map variable: `var m map[key_type]value_type`**, instead, one of the syntaxes must be used (both initialize an **empty** hash maps):
  - `var m = map[key_type]value_type{}` 
  - `var m = make(map[key_type]value_type)`
