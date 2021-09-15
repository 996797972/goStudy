package goBasic

import (
	"fmt"
	"unsafe"
)

func SampleForArraySlice() {
	// Go数组的长度是固定的，所以比较少直接使用数组
	// 注意：[3]int和[4]int这是两种不同的类型
	var arr1 [2]int // 声明但不初始化，默认为0
	arr1[0] = 1
	arr1[1] = 2
	var arr2 [2]int = [2]int{} // 声明并初始化为0，可以不加type让编译器自动推导
	arr3 := [3]int{1, 2, 3} // 声明并初始化
	arr4 := [...]int{1, 2, 3, 4} // 用...代替长度，让编译器自动推导。如果不加元素则为0
	fmt.Println(arr1, arr2, arr3, arr4)
	fmt.Printf("sizeof, arr1: %d, arr2: %d, arr3: %d, arr4: %d\n", unsafe.Sizeof(arr1), unsafe.Sizeof(arr2), unsafe.Sizeof(arr3), unsafe.Sizeof(arr4))
	// 如果觉得麻烦可以使用，type关键字定义类型字面量，类似C语言中typedef
	type newArr [3]int // 定义一个容量为3的数组，别名为newArr
	arr5 := newArr{1, 100, 1000}
	fmt.Println(arr5)

	// Go切片，与数组一样，也是可以容纳若干类型相同元素的容器。
	// 切片的构造有三种：
	// 1）切片是对数组的一个连续片段的引用
	arr6 := [6]int{1, 2, 3, 4, 5, 6} // 这是一个数组
	fmt.Printf("type: %T, val: %d\n", arr6[2:4], arr6[2:4]) // type: []int, val: [3 4]，被截取下来会变长切片
	// 2）从头声明赋值
	strSlice := []string{"hello", "world"} // var strSlice []string = []string{"hello", "world"}
	var zeroSlice = []int{} // 零值切片，不是nil，只是元素值为0。zeroSlice == nil为false
	var emptySlice []int // 声明不赋值，则是空切片(nil)，emptySlice == nil为true，注意空切片和被赋0值的切片不是一回事
	fmt.Println(strSlice, zeroSlice == nil, emptySlice == nil)
	// 3）用make内建函数构造，make函数的格式make([]type, size, cap)
	arr7 := make([]int, 5) // 缺省cap参数则len和cap都为5
	fmt.Println(arr7)
	fmt.Printf("len: %d, cap: %d\n", len(arr7), cap(arr7))

	// 数组的容器大小固定，而切片本身是引用类型，可以追加元素
	// 注意：当切片的空余容量不足时，Go会直接申请分配更大的内存，然后将原内存的内容拷贝至新内存块，再在新内存进行追加。增大的空间一般为原容量*2。若余量充足直接添加
	arr8 := []int{1} // len: 1, cap: 1
	arr8 = append(arr8, 2) // 追加一个元素，返回值是新的切片值，并且必须要有变量接收否则报错。len: 2, cap: 2
	arr8 = append(arr8, 3, 4) // 追加多个元素。len: 4, cap: 4
	arr8 = append(arr8, []int{7, 8}...) // 追加一个长度为2的切片，不过要用...表示解包，不能省略。len: 6, cap: 8
	arr8 = append([]int{0}, arr8...) // 在第一个位置添加元素。len: 7, cap: 8。arr8 = [0, 1, 2, 3, 4, 7, 8]
	arr8 = append(arr8[:5], append([]int{5, 6}, arr8[5:]...)...) // 往中间添加两个元素，不要忘记...解包，否则出错。
	fmt.Println(arr8, len(arr8), cap(arr8)) // arr8 = [0 1 2 3 4 5 6 7 8]，len: 9, cap: 16

	// 切片间赋值，只是修改指针指向，并不会拷贝数据。
	// 切片截取时，有三个参数slice[low:high:max]，如arr9[3,6,7]。不过一般max都会缺省，如arr9[3:6]
	// low为截取的起始下标，high为截取的结束下标（不含high），max为新切片容量最大保留的最大下标（不含max）。注意，high和max不能超过原切片容量，否则出错
	// 新的切片，len = high - low，cap = max - low。如果缺省max，那么cap默认从low开始到原切片最后
	arr9 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	tmpSlice := arr9[3:6:8] // 从3开始，因此tmpSlice = [3 4 5]，len = 3, cap = 5，此时共享底层数组。如果修改，原数组内容会被修改
	fmt.Println(tmpSlice, len(tmpSlice), cap(tmpSlice))
	tmpSlice1 := arr9[3:6] // tmpSlice1 = [3 4 5]，len = 3，cap = 7。
	fmt.Println(tmpSlice1, len(tmpSlice1), cap(tmpSlice1))
	// 可以看到如果新切片在没有达到重新分配内存的情况下，是直接基于原切片修改的，其他由原切片截取的切片也会被修改。体现了切片底层数据结构的共享性
	tmpSlice1[2] = 100 // 这里不可以超出新切片的len，否则越界出错，如tmpSlice[3] = 100
	fmt.Println(tmpSlice1, len(tmpSlice1), cap(tmpSlice1)) // tmpSlice1[3 4 100], len: 3, cap: 7
	fmt.Println(arr9, len(arr9), cap(arr9)) // arr9 = [0 1 2 3 4 100 6 7 8 9], len: 10, cap: 10
	fmt.Println(tmpSlice, len(tmpSlice), cap(tmpSlice)) // tmpSlice[3 4 100]，len: 3, cap: 5
	tmpSlice = append(tmpSlice, 3, 4, 5) // 超出容量，重新申请内存，不影响原切片。这里如果只添加两个元素，不会触发重新申请内存，则原切片和其他切片内容会改变
	fmt.Println(tmpSlice, len(tmpSlice), cap(tmpSlice)) // tmpSlice[3 4 100 3 4 5], len: 6, cap: 10
	fmt.Println(tmpSlice1, len(tmpSlice1), cap(tmpSlice1)) // tmpSlice1[3 4 100], len: 3, cap: 7
	fmt.Println(arr9, len(arr9), cap(arr9)) // arr9 = [0 1 2 3 4 100 6 7 8 9], len: 10, cap: 10

	// copy函数
	arr10 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	tmpSlice2 := make([]int, len(arr10)-3) // make会重新申请内存，因此这两者并不共享底层信息
	copy(tmpSlice2, arr10[2:]) // tmpSlice = [2 3 4 5 6 7 8]， arr10 = [0 1 2 3 4 5 6 7 8 9]
	fmt.Println(tmpSlice2, arr10)
	tmpSlice2[3]=444 // tmpSlice = [2 3 4 444 6 7 8]， arr10 = [0 1 2 3 4 5 6 7 8 9]，并不会影响arr10
	fmt.Println(tmpSlice2, arr10)
}
