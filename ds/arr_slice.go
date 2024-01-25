package ds

import "fmt"

func modifySlice(s *[]int) {
	// 修改切片内部的底层数组
	(*s)[0] = 99

	// 重新分配切片，指向新的底层数组
	*s = append(*s, 4, 5, 6)
}

func main() {
	// 创建一个切片
	originalSlice := []int{1, 2, 3}

	// 调用修改切片的函数
	modifySlice(&originalSlice)

	// 打印原始切片
	fmt.Println(originalSlice) // 输出: [99 2 3]
}

// ============================================

import "fmt"

func modifySlice(s *[]int) {
	// 修改切片内部的底层数组
	(*s)[0] = 99

	// 通过指针修改原始切片，使其包含新的元素
	*s = append(*s, 4, 5, 6)
}

func main() {
	// 创建一个切片
	originalSlice := []int{1, 2, 3}

	// 调用修改切片的函数
	modifySlice(&originalSlice)

	// 打印修改后的原始切片
	fmt.Println(originalSlice) // 输出: [99 2 3 4 5 6]
}
