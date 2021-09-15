package goBasic

import "fmt"

func SampleForMap() {
	// MAP的各种声明和初始化
	// 1.声明但没初始化
	var sMap1 map[string]int // 注意此时sMap1是空值不是零值，如果直接赋值会panic，必须先make分配空间再赋值
	if sMap1 == nil {
		sMap1 = make(map[string]int) // 先make
	}
	sMap1["English"] = 85 // 再赋值
	sMap1["Chinese"] = 90
	// 上面的方法太麻烦了，直接声明并初始化一步到位简单化，如：var sMap = map[string]int{"English": 85, "Chinese": 90}
	// 2. 直接初始化
	sMap2 := map[string]int{"English": 80, "Chinese": 82}
	// 3. make分配内存后，再初始化
	sMap3 := make(map[string]int) // make直接默认零值初始化。此时sMap3是零值，不是空值，可直接赋值
	sMap3["English"] = 83
	sMap3["Math"] = 92
	fmt.Println(sMap1, sMap2, sMap3) // map[Chinese:90 English:85] map[Chinese:82 English:80] map[English:83 Math:92]

	// 添加元素直接赋值即可。另，如果某个key的值已经存在，则直接覆盖
	sMap3["English"] = 94
	fmt.Println(sMap3, len(sMap3)) // map[English:94 Math:92], 2
	// 如果某个key不存在，那么会返回零值，不会出错
	fmt.Printf("%d\n", sMap3["Biology"]) // 0

	// 删除元素，使用内建delete函数，该函数只能用于map。如果要删除的key本身不存在，delete则什么都不会做也不会报错
	delete(sMap2, "English")
	fmt.Println(sMap2) // map[Chinese:82]

	// 判断对应的key-value是否存在。因为不存在也返回0，而值本身也有可能为0，因此不能直接以返回值为0判断某key是否存在。而应该如下
	if biology, ok := sMap2["Biology"]; ok { // 直接访问会返回两者值value和布尔值。这里ok也可以另起一行写
		fmt.Printf("Biology Score: %d\n", biology)
	} else {
		fmt.Printf("Biology is not exist!\n")
	}

	// Go没有提供类似python的keys()、values()函数，需要自行循环遍历，可使用for-range遍历
	// range会默认按顺序返回两个值，先是key然后是value
	for key, value := range sMap1 {
		fmt.Printf("Score Map1 %s: %d\n", key, value)
	}
	// 如果使用一个变量接收那么只返回key
	for key := range sMap2 {
		fmt.Printf("Score Map2 %s: %d\n", key, sMap2[key])
	}
	// 如果不关心key而只想输出value，那么必须使用_将排在前面的key忽略
	for _, value := range sMap3 {
		fmt.Printf("Score Map3: %d\n", value)
	}
}
