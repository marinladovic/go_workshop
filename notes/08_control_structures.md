# Control Structures

## if - else

It can have multiple conditions (There is no ternary operator in Go)

```go
if user != nil {
} else {
}

if message := "hello"; user != nil {
} else {
}
```

## switch (reloaded!)

This is a simple switch operation. No `break` needed; you can `fallthrough` to the next case though

```go
switch day {
    case "Monday":
        fmt.Println("It's Monday!")
    case "Saturday":
        fallthrough
    case "Sunday":
        fmt.Println("It's Weekend!")
    default:
        fmt.Println("It's another working day")
}
```

It can replace large `if`s:

```go
switch {
    case user == nil:

    case user.active == false:

    case user.deleted == true:

    default:
}
```

## for

A multipurpose loop control structure

```go
// Classic for
for i := 0; i < len(collection); i++ {}

// For range, similar to "for in" in JS
for index := range collection {}

// For range, similar to "foreach"
for key, value := range map {}
```

We can emulate a `while` just by using a boolean expression

```go
endOfGame := false
for endOfGame {
    // process Game loop
}

for count < 10 {
    count += 1
}

// infinite loop
for {}
```

There is no `while` or `do - while`

- No parenthesis are needed for bool conditions or values
- Only one type of equality operator `==`
- Other operators `!=`, `<`, `>`, `<=`, `>=`, `!`