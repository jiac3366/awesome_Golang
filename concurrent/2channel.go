package concurrent

import "fmt"

//func main() {
//
//	ch := make(chan int, 5)
//	ch <- 1
//
//	go func() {
//		for i := 0; i < 5; i++ {
//			v, _ := <-ch
//			fmt.Println(v)
//			//if ok {
//			//	fmt.Printf("v=%d\n", v)
//			//} else {
//			//	fmt.Printf("read out")
//			//}
//		}
//	}()
//	// 主 goroutine 关闭了 channel 之后，子goroutine 的 for range 才会结束
//	time.Sleep(1 * time.Second)
//	close(ch)
//	time.Sleep(3 * time.Second)
//
//}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := make([]int, 6)
	for i := 0; i < 6; i++ {
		s = append(s, i)
	}
	c := make(chan int)
	go func() {
		sum(s[:len(s)/2], c)
	}()

	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y)

}
