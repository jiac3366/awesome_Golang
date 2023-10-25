package concurrent

import (
	"fmt"
	"sync/atomic"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	st1 := Student{
		Name: "zhangsan",
		Age:  19,
	}

	st2 := Student{
		Name: "lisi",
		Age:  20,
	}

	st3 := Student{
		Name: "wangwu",
		Age:  20,
	}

	var v atomic.Value
	v.Store(st1)
	fmt.Println(v.Load().(Student))

	old := v.Swap(st2)
	fmt.Println(v.Load().(Student))
	fmt.Println(old)

	swapped := v.CompareAndSwap(st1, st3)
	fmt.Println("st3 could update depends on st1 == v, st1 == v？", swapped)
	fmt.Println(swapped)

	swapped = v.CompareAndSwap(st2, st3)
	fmt.Println("st3 could update depends on st2 == v, st2 == v？", swapped)
	fmt.Println(swapped)
}
