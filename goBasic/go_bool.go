package goBasic

import "fmt"

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func int2bool(i int) bool {
	return i != 0
}

func SampleForBool() {
	// 关于布尔值，无非就两个值：true和false。不过，在不同语言中，这两个值可能不同
	// 在python中，真值用True表示，与1相等；假值用False表示，与0相等
	// 但在Go中，真值用true表示，不但不与1相等，且更加严格，不同类型无法进行比较；假值用false表示，同样不与0相等
	var isMale bool = true
	var isPolice bool = false
	fmt.Printf("type: %T, value: %t\n", isMale, isMale) // type: bool, value: true

	// Go确实不如python那样灵活，bool与int不能直接转换，如果要转换，需要自行实现函数
	// python使用not对逻辑值取反，而go和C语言一样，取反用!
	fmt.Printf("%t\n", !isMale)

	// 多个判断条件，python用and和or，而go和C语言一样，使用&&（且）和||（或）。
	// 并且短路行为（左边表达式已经可以确认整个表达式的值，那么右边将不会再被求值）
	if isMale && isPolice {
		fmt.Printf("He is a policeman!\n")
	}
}
