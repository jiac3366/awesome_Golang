package main

import "fmt"

// 函数接受多个参数，这些参数可以是任意类型，并打印它们的值和类型
func printValues(values ...interface{}) {
	for _, value := range values {
		// 使用类型断言确定参数的实际类型
		switch v := value.(type) {
		case int:
			fmt.Printf("Integer: %d\n", v)
		case float64:
			fmt.Printf("Float: %f\n", v)
		case string:
			fmt.Printf("String: %s\n", v)
		default:
			fmt.Println("Unknown type")
		}
	}
}

func main() {
	// 调用printValues函数并传递多种类型的参数
	printValues(10, 3.14, "Hello", true)
}
