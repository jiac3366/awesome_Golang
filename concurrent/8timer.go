package concurrent

import (
	"fmt"
	"time"
)

func main1() {
	ch := make(chan bool, 1)
	select {
	case <-ch:
		fmt.Println("ch")
	case <-time.After(time.Second * 2):
		fmt.Println("timeout")

	}

	fmt.Println("jjj")

}

// ====================

func TickerWatcher() chan struct{} {
	ch := make(chan struct{})
	fmt.Printf("%p\n", ch)

	ticker := time.NewTicker(2 * time.Second)

	go func(ticker *time.Ticker) {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Println("watch!!")
			case <-ch:
				fmt.Println("end")
				return
			}
		}

	}(ticker)

	return ch
}

func main() {

	ch := TickerWatcher()
	fmt.Printf("%p\n", ch)

	time.Sleep(5 * time.Second)

	ch <- struct{}{}

	time.Sleep(5 * time.Second)
	close(ch)

}
