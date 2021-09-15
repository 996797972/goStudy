package goBasic

import "fmt"

func SampleForVariable() {
	// 常用的三种方式
	var v1 string = "GO for it"
	var v2 = "GO for it" // 由编译器自动推导出类型
	v3 := 100 // 变量声明加初始化，注意其作用域只限定在函数体或结构语句中（如for）。作用域外该变量无法识别
	fmt.Printf("v1: %s\n", v1)
	fmt.Printf("v2: %s\n", v2)
	fmt.Printf("v3: %d\n", v3)

	// 多个变量同时声明，默认都是0值
	var (
		name string
		age int
		gender string
	)
	name = "Mike"
	age = 23
	gender = "Male"
	fmt.Printf("%s %d %s\n", name, age, gender)

	// 使用new内建函数，返回的是Type类型的指针，值为0
	ptr := new(int)
	*ptr = 100
	fmt.Printf("ptr, type: %T, address: %p, value: %d\n", ptr, ptr, *ptr)
}
