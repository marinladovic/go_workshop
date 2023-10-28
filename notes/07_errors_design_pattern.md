# Errors Design Pattern

We don't have exceptions in Go, this is the typical design pattern when an action may trigger an error.

```go
func readUser(id int) (user, err) {
    // ... we proceed with the reading and see a bool ok value
    if ok {
        return user, nil
    } else {
        return nil, errorDetails
    }
}

func main {
    user, err := readUser(2)
}
```