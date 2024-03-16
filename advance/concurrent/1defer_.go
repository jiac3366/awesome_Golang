package main

import (
	"errors"
	"fmt"
	"sync"
)

func withGoroutine(handlers ...func() error) (err error) {
	var wg sync.WaitGroup
	for _, _handler := range handlers {
		wg.Add(1)
		go func(handler func() error) {
			defer func() {
				if e := recover(); e != nil {
					fmt.Printf("recover:%v\n", e)
				}
				wg.Done()
			}()
			e := handler()
			if err == nil && e != nil {
				err = e
			}

		}(_handler)
	}
	wg.Wait()

	return
}

func main() {

	handler1 := func() error {
		//panic("handler1 fail ")
		return errors.New("hei!")
	}
	handler2 := func() error {
		panic("handler2 fail ")
		return nil
	}
	err := withGoroutine(handler1, handler2)
	if err != nil {
		fmt.Println("！！！err is:%v", err)
	}

}
