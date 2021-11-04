package cancel

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int, ctx context.Context) {
			childTask(i, ctx, &wg)
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
			wg.Done()
		}(i, ctx)
	}
	cancel()
	wg.Wait()
	//time.Sleep(time.Second * 1)
}

func childTask(i int, ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func(i int, ctx context.Context) {
		for {
			if isCancelled(ctx) {
				break
			}
			time.Sleep(time.Millisecond * 5)
		}
		fmt.Println(i, "Cancelled-child")
		wg.Done()
	}(i, ctx)
}
