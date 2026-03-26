# Go Variables

Variables are containers for storing data values.

## Types of Variables in Go

In Go, different types of variables store different kinds of data:

* `int` stores integers, such as `123` or `-123`
* `float32` stores decimal numbers, such as `19.99` or `-19.99`
* `string` stores text, such as `"Hello World"`
* `bool` stores values with two states, `true` or `false`

## Declaring Variables

Declaration means giving a variable a name and a type.

### Using the `var` keyword

Use the `var` keyword, followed by the variable name, type, and value:

```go id="8z2lq1"
var variablename type = value
```

You can declare a single variable or multiple variables.

## Go Variable Naming Rules

* A variable name must start with a letter or an underscore `_`
* A variable name cannot start with a digit
* A variable name can only contain letters, numbers, and underscores `(a-z, A-Z, 0-9, _)`
* Variable names are case-sensitive (`age`, `Age`, and `AGE` are different)
* There is no limit to the length of a variable name
* A variable name cannot contain spaces
* A variable name cannot be a Go keyword
