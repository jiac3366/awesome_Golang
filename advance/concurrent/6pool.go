package concurrent

import (
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() any {
			return &Student{
				Name: "chase",
				Age:  11,
			}
		},
	}
	st := pool.Get().(*Student)
	st2 := pool.Get().(*Student)
	fmt.Println(st)
	fmt.Printf("point of st: %p\n", st)

	fmt.Println(st2)
	fmt.Printf("point of st2: %p\n", st2)

	st.Name = "chase2222"
	pool.Put(st)
	st = pool.Get().(*Student)
	fmt.Println(st)
	fmt.Printf("point of st updated: %p\n", st)

	st = pool.Get().(*Student)
	fmt.Println(st)
	fmt.Printf("point of st: %p\n", st)

}
