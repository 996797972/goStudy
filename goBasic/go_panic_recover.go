package goBasic

import "fmt"

// 一般情况下，一些常规性的错误如语法错误等等，直接在编译的时候就可以发现。但是有一些是运行态才会发现，如内存错误访问、数组越界等，此时通常程序会crash
// 但是除了上面系统的检测机制，有时候我们需要自动触发或者捕捉一些异常。此时就会涉及到异常抛出和异常捕捉
// 异常抛出：是为了自己主动触发程序crash，进而抛出异常。相关函数panic
// 异常捕捉：是为了程序出现异常时，我们能处理一些异常或透过异常信息来调试开发中的程序，继而让程序继续运行上去。相关函数recover，同时必须结合defer使用
func directPanic(msg string)  {
	panic(msg)
}

func set_data(x int) {
	defer func() {
		if err := recover(); err != nil { // recover一定要结合defer来使用
			fmt.Println(err)
		}
	}()
	fmt.Printf("panic前的会执行\n")
	var arr [10]int
	arr[x] = 121 // 故意造成越界在这里会panic，然后直接执行到defer
	fmt.Printf("panic后的不会执行\n")
}

func SampleForPanicRecover() {
	// 直接触发panic，抛出异常
	// directPanic("Dead here")
	set_data(20)
	fmt.Printf("准备退出\n")
}
