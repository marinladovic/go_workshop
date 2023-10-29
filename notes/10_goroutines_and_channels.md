# Goroutines ans Channels

- A goroutine is the the Go way of using threads
- We open a goroutine just by invoking any function with a go prefix
- go `functionCall()`
- Goroutines can communicate through `channels` (a special type of variable)
- A `channel` contains a value of any kind
- a routine can define a value for a channel and other routine can wait for that value
- Channels can be buffered or not
