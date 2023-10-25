package concurrent

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mu1, mu2 sync.RWMutex // 如果改成Mutex, 且都调用 Lock，会发生死锁

	wg.Add(2)
	go func() {
		defer wg.Done()
		mu1.RLock()
		defer mu1.RUnlock()
		time.Sleep(100 * time.Millisecond)

		mu2.RLock()
		defer mu2.RUnlock()

	}()

	go func() {
		defer wg.Done()
		mu2.RLock()
		defer mu2.RUnlock()
		time.Sleep(100 * time.Millisecond)

		mu1.RLock()
		defer mu1.RUnlock()
	}()

	wg.Wait()
}
