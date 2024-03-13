package main

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	p1 := &Person{}
	var p2 *Person
	var p3 Person

	println(p1 == nil)
	println(p2 == nil)
	println(p3 == nil)
}

//p1 == nil：
//p1 是一个指向 Person 实例的指针，它不是 nil。
//因此，p1 == nil 的结果是 false。
//
//p2 == nil：
//p2 是一个空指针变量，没有指向任何有效的内存地址。但是p2 == nil可以执行原因是：nil 也用于表示指针类型的零值
//将一个指针类型的变量与 nil 进行比较时，Go 会隐式地将 nil 转换为相应指针类型的零值
//因此，p2 == nil 的结果是 true。

// p3==nil不能执行的原因是： Go 语言中，比较运算符 == 要求两边的操作数类型必须相同或者是可以相互转换的类型。
// 对于 p3 这个结构体变量和 nil 这个预定义的无类型常量，它们的类型是不同的，因此无法直接进行比较

// ps:
// zeroPerson := Person{}
// fmt.Println(p3 == zeroPerson) 输出--> true

// =====================================================
//%v：默认格式。用于打印值的默认格式。
//%+v：详细格式。用于打印值的详细格式，包括结构体字段的标签信息。
//%#v：Go 语法格式。用于打印值的 Go 语法表示。
//%T：类型。用于打印值的类型。
//%t：布尔。用于打印布尔值。
//%d：十进制整数。用于打印十进制整数。
//%b：二进制整数。用于打印二进制表示的整数。
//%o：八进制整数。用于打印八进制表示的整数。
//%x、%X：十六进制整数。用于打印十六进制表示的整数。
//%f、%F：浮点数。用于打印浮点数。
//%e、%E：科学计数法表示的浮点数。
//%s：字符串。用于打印字符串。
//%q：带引号的字符串。用于打印带双引号的字符串。
//%p：指针。用于打印指针的十六进制表示。
//%%：百分号。用于打印百分号（%）本身。
