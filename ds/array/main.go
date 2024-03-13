package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	array := []int{0, 1, 2, 3, 4}
	_ = array[4]
	PrintSlice(&array)

	s1 := append(array[:1], array[2:]...)

	PrintSlice(&s1)
	PrintSlice(&array)

}

func PrintSlice(s *[]int) {
	ss := (*reflect.SliceHeader)(unsafe.Pointer(s))
	fmt.Printf("slice struct %+v, slice is %v\n", ss, s)
}
