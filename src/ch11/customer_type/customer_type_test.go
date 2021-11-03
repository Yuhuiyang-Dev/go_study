package customer_type_test

import (
	"fmt"
	"testing"
	"time"
)

type intConv func(op int) int

func timeSpent(inner intConv) intConv {
	return func(op int) int {
		start := time.Now()
		i := inner(op)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return i
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	t.Log(timeSpent(slowFun)(2))
}
