package concurrent

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	ch1 <- 45
	ch2 <- 46

	select {
	case v := <-ch1:
		fmt.Println("receive", v)
	case v := <-ch2:
		fmt.Println("receive", v)
	default:
		fmt.Println("default")
	}
}

// ==============================

var (
	ch chan string
	wg sync.WaitGroup
)

func Reader(index int, ch chan string) {
	defer wg.Done()
	for {
		time.Sleep(time.Second * 2)
		val := <-ch
		fmt.Printf("%d RRRRRRRRroutinue read, vaule is %s\n", index, val)
	}
}

func Write(index int, ch chan string) {
	defer wg.Done()
	for {
		time.Sleep(time.Second * 1)
		timer := time.NewTicker(time.Millisecond * 800) // 如果timeout 时间大于读写速度的差值，就不会 timeout; 在这里 timeout > 100ms 就不会报timeout
		select {
		case ch <- strconv.Itoa(index):
			fmt.Printf("%d WWWWWWWWWWroutinue write\n", index)
		case <-timer.C:
			log.Printf("send to handle chan timeout, req : %d", index)
			break //breack select 不是 for
		}
	}
}

func main() {

	ch = make(chan string, 3)
	wg.Add(6)
	for i := 0; i < 1; i++ {
		go Reader(i, ch)
	}

	for i := 0; i < 1; i++ {
		go Write(i, ch)
	}
	wg.Wait()

}
