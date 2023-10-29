# Goroutines ans Channels

- A goroutine is the the Go way of using threads
- We open a goroutine just by invoking any function with a go prefix
- go `functionCall()`
- Goroutines can communicate through `channels` (a special type of variable)
- A `channel` contains a value of any kind
- a routine can define a value for a channel and other routine can wait for that value
- Channels can be buffered or not

## Using Channels

```go
// define
var m1 chan string

// construct
m2 := make(chan string)

// write in
m2 <- "hello"

// saving
message := <- m2
```

### Buffer Channel

```go
logs := make(chan string, 2) // sending 2 messages at the same time
logs <- "hello"
logs <- "world"
fmt.Println(<-messages)
fmt.Println(<-messages)
```

### WARNING

To avoid deadlocks you have to close the channels before ending the program with `close(chan)`
