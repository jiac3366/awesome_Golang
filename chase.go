package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := `"你好"，世界！` // 包含6个字符，但占用18个字节（每个汉字占3个字节）

	fmt.Printf("字符串: %s\n", str)
	fmt.Printf("字符数: %d\n", len(str))                           // 输出字节数
	fmt.Printf("Unicode字符数: %d\n", utf8.RuneCountInString(str)) // 输出Unicode字符数
}
