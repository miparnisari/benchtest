package fib

import (
	"math/rand"
	"sync"
	"time"
)

var _m sync.Map

// Fib returns nth Fibonnaci sequence number.
func Fib(n uint) uint {
	if n <= 1 {
		return 1
	}

	n2, ok := _m.Load(n - 2)
	if !ok {
		n2 = Fib(n - 2)
		_m.Store(n-2, n2)
	}
	n1, ok := _m.Load(n - 1)
	if !ok {
		n1 = Fib(n - 1)
		_m.Store(n-1, n1)
	}

	dur := 500*time.Microsecond + time.Duration(rand.Intn(500))*time.Microsecond
	time.Sleep(dur)

	return n2.(uint) + n1.(uint)
}
