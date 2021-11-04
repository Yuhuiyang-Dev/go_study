package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			//fmt.Println("Producing", i)
			ch <- i
			//fmt.Println("Produced", i)
		}
		close(ch)
		//ch <- 2 //引发 panic
		wg.Done()
	}()

}

func dataReceiver(ch chan int, wg *sync.WaitGroup, name string) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Printf("Receiver-%s %d\n", name, data)
			} else {
				break
			}
		}
		wg.Done()
	}()

}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg, "1")
	//wg.Add(1)
	//dataReceiver(ch, &wg, "2")
	wg.Wait()

}
