package goBasic

import (
	"fmt"
	"unsafe"
)

func SampleForStringByteRune() {
	// byte占1个字节，与uint8没有本质区别。同样很明显中文字符超过256，是无法用byte表示的
	// 8进制写法： var c byte = '\101' 其中\是8进制固定前缀
	// 16进制写法： var c byte = '\x41' 其中\x是16进制固定前缀
	// %c格式化输出字符而不是像%d般输出数字
	var a byte = 65
	var b uint8 = 66
	var c byte = 'C'
	var d uint8 = 'D'
	fmt.Printf("a: %c, b: %c, c: %c, d: %c\n", a, b, c, d)

	// rune类型占4个字节，与uint32没有本质区别。它表示一个unicode字符（unicode是一个可以表示世界范围内的绝大部分字符的编码规范）
	// rune与byte都只能用单引号，不能用双引号，他们之间本质区别就是rune范围大,byte范围小。
	// byte常用于单个ASCII字符，rune可表示单个ASCII字符或单个中文字符
	var r1 rune = 'A'
	var r2 rune = '中' // 不可以用双引号，如"中"。双引号是给string用的
	fmt.Printf("sizeof r1, r2: %d, %d\n", unsafe.Sizeof(r1), unsafe.Sizeof(r2))
	fmt.Printf("r1: %c, r2: %c\n", r1, r2)

	// string类型表示字符串，多个byte或rune类型可组成数组，就组成了字符串。不过，注意字符串的底层代表形式是如下的：
	// type stringStruct struct {
	//	 str unsafe.Pointer // 这个是数组的指针
	// 	 len int
	// }
	// 从上可以看出，sizeof(string变量)肯定为16，注意sizeof大小与字符串长度本身的区别
	var str1 string = "hello" // len(str1) = 5，而unsafe.Sizeof(str1) = 16，详看上面string的定义
	var str2 [5]byte = [5]byte{'h', 'e', 'l', 'l', 'o'} // len(str2) = 5,unsafe.Sizeof(str2) = 5
	var str3 [5]rune = [5]rune{104, 101, 108, 108, 111} // len(str3) = 5,unsafe.Sizeof(str3) = 20
	fmt.Printf("str1: %s, str2: %s, str3: %q\n", str1, str2, str3)
	fmt.Printf("len str1: %d, str2: %d, str3: %d\n", len(str1), len(str2), len(str3))
	fmt.Printf("sizeof str1: %d, str2: %d, str3: %d\n", unsafe.Sizeof(str1), unsafe.Sizeof(str2), unsafe.Sizeof(str3))

	// Go的string用的是utf-8编码的，中文字符或标点是占3个字节，英文字符或标点占1个字节
	var str4 string = "hello，中国" // len: 6 + 3 * 3 = 14
	fmt.Printf("len str4: %d\n", len(str4))
}
