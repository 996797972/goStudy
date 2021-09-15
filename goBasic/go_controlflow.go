package goBasic

import "fmt"

// 多个defer语句遵循反序调用，类似栈一样，后进先出。并且传递的变量不会受到后续程序控制
// hello my world
func sampleDefer1()  {
	name := "world"
	defer fmt.Println(name)

	name = "my"
	defer fmt.Println(name)

	name = "hello"
	fmt.Println(name)
}

// 函数中使用defer，return先得到调用，才会调用该函数里的defer
var deferValue = "Shell" // 全局变量
func sampleDefer2() string {
	defer func() { // 匿名函数
		deferValue = "Go"
	}()
	fmt.Printf("sampleDefer1函数中的deferValue：%s\n", deferValue)
	return deferValue
}

func SampleForControlFlow() {
	// if语句
	year := 2000
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Printf("%d是闰年\n", year)
	}
	// switch-case语句
	// switch后的条件可以是变量、函数、表达式
	// case后面可以是常量或表达式。case后如涉及多个条件，则他们之间是或关系。另注意，case后面的常量不允许重复。
	// 优先从上至下按case匹配，只要匹配一个则执行对应代码块，然后退出switch-case
	// 某个case代码块可以接fallthrough（必须放在case代码块最后）。如果匹配此case，执行完此case语句块后，无条件也执行下一个case的代码块（不管下个case是否满足条件），如果下一个case也带fallthrough那么继续无条件执行下去
	month := 6
	switch month {
	case 3, 4, 5:
		fmt.Println("春天")
	case 6, 7, 8:
		fmt.Println("夏天")
		fallthrough // 如果此case命中，无条件执行下一个case，fallthrough关键字必须写在代码看最后
	case 9, 10, 11:
		fmt.Println("秋天")
	case 12, 1, 2:
		fmt.Println("冬天")
	default:
		fmt.Println("输入有误")
	}
	// for-range语句，这里不讲for循环只讲for-range。
	// for-range是用来迭代访问的，可以用在切片、数组、map、通道、字符串
	// 默认for-range返回两个值：第一个是索引，第二个是值。如果用一个变量接只返回索引
	countryMap := map[string]string{"Chinese": "Beijing", "Farance": "Paris", "American": "New-York", "Japan": "Tokyo"}
	for country, capital := range countryMap { // 这种情况无法用传统的for循环，只能使用for-range
		fmt.Printf("%s's capital: %s\n", country, capital)
	}

	// goto语句与C语言的goto一样
	fmt.Println("Ready to goto!")
	goto exit
	fmt.Println("This will not exec")
exit:
	fmt.Println("Goto here")

	// defer语句为延时语句，Go语言独有的关键字，能实现将defer后面的语句或函数调用延时到当前函数执行完后再执行
	sampleDefer1()
	// 涉及到函数return先得到调用，才会调用该函数里的defer
	newDeferValue := sampleDefer2()
	fmt.Printf("Main函数中的deferValue：%s\n", deferValue)
	fmt.Printf("Main函数中的newDeferValue:%s\n", newDeferValue)
}
