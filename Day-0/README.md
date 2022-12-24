# Day 0 Notes: First GO program, and first test
<hr />

## Syntax

- A Short hand of declaring variables: **[name] := [value]**. The type is captured during initialization
- If statements are much similar to other programming languages
- Switch syntax is similar to other programming languages:
  - `switch(param){
     case "value1":
        ...
     default:
        ...
    }`

## Testing

Writing test is just like writing a function with few rules
- Test file must be started with `xxx_test.go`
- The test function must be started with `Test`, e.g: *TestHello*
- The test function takes one param only `t *testing.T` (importing "testing")
- The **f** on `t.Errorf` stands for *format*, `%q` allows us to put strings inside double quotes.
- `t.Helper()` tells the test suite that the invoking function is a helper, so it will print the original code error line.


## Best practices
- It is good to seperate the *domain* code from the *side-effects*. E.g: The `fmt.Prinln` is a side effect, and the `string we send` is our domain
- Commiting code can be done when we have working software backed by a test
- Pushing to master MUST be after refactoring
- TDD Life Cycle:
  - Write a test
  - Make the compiler pass
  - Run the test, if it fails check the error message is meaningful
  - Write enough code to make the test pass
  - Refactor

**Wow, a lot to learn from `Hello, world`**