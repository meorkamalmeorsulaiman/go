# Primitive Types and Declaration

## Literals

- Integer
- Floating point
- Rune - char that surronded by single quotes

## Boolean
- True or False

## Numeric Types
- Interger

### Integer Operators
- `+, -, *, /` or `%`
- Combining operator `+=, -+, *=, /=` or `%=`

Example:

```go
var x int = 10
x *= 2
```
`x` will be 20, I guess

### Floating 
- float32 or float64

### Strings nad Runes

- String compare with `==` or `!=`
- Ordering with `>, >=, <` or `<=`
- Concate with `+`


### var vs :=

- The verbose god level `var x int = 10`
- Cheap `var x = 10` which you can leave the type, and default or integer literal is `int`
- Zero vaule declaration `var x int`
- Declare multiple variable at once `var x, y int = 30, 10` or `var x, y int` or `var x, y = 10, "hello"`
- Declaring multiple variable at once:
```go
var (
    x   int
    y   = 20
    z   int - 30
    d, e    = 40, "hello"
    f, g    string
)
```

- Go short declaration format `var x = 10` is equal to `x := 10`
- Declaring multi-var `x, y := 10, "hi"`
- Use `:=` withing function, use `var` when declaring zero values