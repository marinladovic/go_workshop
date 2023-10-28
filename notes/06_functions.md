# Functions

## Introduction

- Similar to other languages
- It can receive arguments
- Arguments can have default values
- The last argument can be of variable length 
- The tricky part:
    - Functions can return more than one value (at once)
    - Functions can return labeled variables
    - Functions receive arguments always by value

## We need to define types for arguments and return variables

```go
func save() {}
func save(text string) {}

func add(a int, b int) int {
    return a + b
}

func addAndSubtract(a int, b int) (int, int) {
    return a+b, a-b
}
```

## Functions receiving references

We need to receive pointers instead of the value

```go
func increment(x *int) {
    *x++
}

func main() {
    i := 1
    increment(&i)
}
```