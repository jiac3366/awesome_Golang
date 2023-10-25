package concurrent

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(4*time.Second))
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second) // 主线程等 6s,但这里 4s 就提前退出

	go Watch(ctx, "fun1")
	go Watch(ctx, "fun2")

	time.Sleep(6 * time.Second)
	fmt.Println("end watching")
	cancel()
	time.Sleep(1 * time.Second)

}

func Watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "exit")
			return
		default:
			fmt.Println(name, "default...")
			time.Sleep(time.Second)
		}

	}
}

//==================
func fun1(ctx context.Context) {
	fmt.Println("name is %s", ctx.Value("name").(string))
}

func main1() {
	ctx := context.WithValue(context.Background(), "name", "zhangsan")

	go fun1(ctx)
	time.Sleep(time.Second)

}
