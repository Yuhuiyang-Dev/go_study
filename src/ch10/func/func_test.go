package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
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
	a, _ := returnMultiValues()
	t.Log(a)
	t.Log(timeSpent(slowFun)(2))
}

func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret = ret + op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3, 4))
	t.Log(sum(1, 2, 3, 4, 5))
}
func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	//panic("err")
}
