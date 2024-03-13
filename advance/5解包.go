package advance

import "fmt"

// 可变参数函数
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	// 定义一个切片
	numbers := []int{1, 2, 3, 4, 5}

	// 通过...解包切片传递给可变参数函数
	fmt.Println(sum(numbers...)) // 输出: 15
}
