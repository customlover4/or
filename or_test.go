package or

import (
	"testing"
	"time"
)

var sig = func(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func TestOr(t *testing.T) {
	start := time.Now()
	<-Or(sig(time.Second*5), sig(time.Second*10))
	sc := time.Since(start).Seconds()
	if !(sc > 4.9 && sc < 5.1) {
		t.Fail()
	}
}

func TestOrOneChan(t *testing.T) {
	start := time.Now()
	<-Or(sig(time.Second*1))
	sc := time.Since(start).Seconds()
	if !(sc > 0.9 && sc < 1.1) {
		t.Fail()
	}
}

func TestOrZeroChan(t *testing.T) {
	start := time.Now()
	<-Or()
	sc := time.Since(start).Seconds()
	if !(sc > 0.0 && sc < 0.1) {
		t.Fail()
	}
}