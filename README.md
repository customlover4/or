# Example

Or function return signal to channel when, one of all channels closing

```go
var sig = func(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

<-Or(sig(time.Second*5), sig(time.Second*10))
```