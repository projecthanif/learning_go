## Variables in Go

In this lesson, I learned how to declare variables in Go using the `var` keyword.

Example from `main.go`:

```go
var username string = "projecthanif"
var isLoggedin bool = false
var count int = 0
var score float32 = 255.9173738393828292
var score64 float64 = 255.9173738393828292
var anotherVariable int
var website = "https://projecthanif.1hanif.click"
fullName := "Ibrahim Mustapha"
```

This shows that a variable declaration in Go can include:

- the variable name
- the data type
- the value assigned to it

General pattern:

```go
var variableName dataType = value
```

Go can also infer the type when using `var`:

```go
var website = "https://projecthanif.1hanif.click"
```

In this case, Go automatically understands that `website` is a `string`.

In this file, I used these data types:

- `string` for text values
- `bool` for true or false values
- `int` for whole numbers
- `float32` for decimal numbers
- `float64` for decimal numbers with more precision

I also used:

```go
fmt.Printf("Variable is of type: %T\n", username)
```

The `%T` format verb prints the type of a variable. This is useful when learning because it helps confirm what type Go is using.

## Short Variable Declaration

Go also supports short variable declaration using `:=`.

Example:

```go
name := "projecthanif"
age := 25
```

This is a shorter way to declare and initialize variables, and Go will infer the type automatically.

Important:

- `:=` can only be used inside functions
- it cannot be used for package-level variables

That is why this works inside `main()`:

```go
fullName := "Ibrahim Mustapha"
```

But this does not work at the top level of the file:

```go
// whoami := "nice"
```

## Zero Values

If a variable is declared without assigning a value, Go gives it a default zero value.

Examples:

```go
var name string
var loggedIn bool
var count int
var price float64
```

Their zero values are:

- `string` -> `""`
- `bool` -> `false`
- `int` -> `0`
- `float64` -> `0`

This matches your code:

```go
var anotherVariable int
```

Since no value was assigned, Go gives it the zero value for `int`, which is `0`.

## Constants

Go also supports constants using the `const` keyword.

Example from your code:

```go
const JwtToken string = "projecthanif"
```

Constants are used for values that should not change while the program is running.

General pattern:

```go
const constantName dataType = value
```

You usually use constants for fixed values like:

- tokens
- configuration values
- fixed messages
- mathematical values

## Float Precision Note

`float32` and `float64` both store decimal numbers, but `float64` keeps more precision.

Example:

```go
var score float32 = 255.9173738393828292
var score64 float64 = 255.9173738393828292
```

`float32` may lose some precision, while `float64` preserves more of the original value.

What I learned here:

- Go is a statically typed language
- each variable has a specific type
- I can print both the value and the type of a variable
- Go can infer types with `var` when the value is provided
- Go can infer types with `:=`
- `:=` only works inside functions
- constants are declared with `const`
- different number types can store values with different precision

One example:

```go
var username string = "projecthanif"
fmt.Println(username)
fmt.Printf("Variable is of type: %T\n", username)
```

This prints the value of `username` and then shows that its type is `string`.
