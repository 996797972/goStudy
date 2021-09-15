package goBasic

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string
	age int
	gender string
}
// 这里方法的接收者都是指针，那么反射时的构造Type和Value对象时最好也以指针形式
func (p *MyStruct) SayHello() {
	fmt.Printf("Hello\n")
}
func (p *MyStruct) SayGoodBye() {
	fmt.Printf("GoodBye\n")
}
func (p *MyStruct) SellIntroduce(name string, age int) {
	fmt.Printf("My name is %s, age %d\n", name, age)
}

func SampleForReflect() {
	var i1 interface{} = 25
	fmt.Printf("%T %v\n", i1, i1) // int 25

	// 反射第一定律，接口变量转换成反射对象
	t1 := reflect.TypeOf(i1)      // 返回的是Type类型，Type是以接口形式存在的，这里t的值真实类型为*rtype类型
	v1 := reflect.ValueOf(i1)     // 返回的是Value类型，Value是结构体类型
	fmt.Printf("%T %T\n", t1, v1) // *reflect.rtype reflect.Value

	// 反射第二定律，反射对象转换成接口变量（这里说的反射对象只能是Value类型，而Type类型是无法转换成接口变量的），使用的方法是Value.Interface()方法
	rI1 := v1.Interface()   // 返回interface{}类型，该返回值是一个静态类型，所以不能直接赋值给原始类型的变量。比如这里var n int = rI就会出错。
	fmt.Printf("%T\n", rI1) // int
	var n int = rI1.(int)   // 这里指定了n是个int类型变量，那么需要将rI再进行类型断言取出原类型(这里是int)才可以输出
	fmt.Printf("n = %d\n", n) // n = 25

	// 反射第三定律，如果要修改“反射类型对象”，其类型必须是可写的（settable）。可写性遵循：
	// 1.不是接收变量的指针创建的反射对象，不具可写性
	// 2.是否具备可写性，可使用Value中的CanSet方法来获知，该方法返回bool值
	// 3.对不具备可写性的对象进行修改是没有意义且不合法的，会报错
	// 由上，要让反射对象具备可写性，需要注意：
	// 1. 创建反射对象时传入的必须是变量的指针
	// 2. 除了要传入指针外，还必须使用Elem()函数返回指针指向的数据（也是一个Value类型）
	i2 := 12
	v2 := reflect.ValueOf(i2) // 这里传进去的是值而不是指针，很显然，返回的v2不具备可写性。这里由于没传指针，如果v2调用Elem()会出错
	fmt.Printf("%t\n", v2.CanSet()) // false
	// 修改三步曲
	v2 = reflect.ValueOf(&i2) // 第一步，传进去的必须是变量的指针，否则调用Elem()也会报错
	fmt.Printf("%t %t\n", v2.CanSet(), v2.Elem().CanSet()) // false true
	v2Elem := v2.Elem() // 第二步，基于返回的Value再调用Elem()来获取指针指向的数据，注意Elem()规定调用其的实例(这里指v2)必须是是接口或指针创建的，否则调用报错
	v2Elem.SetInt(13) // 第三部，基于Elem返回的Value，才能调用SetXXX来修改，XXX对应原始类型，这里i2是Int类型
	fmt.Printf("i2的新值为：%d\n", i2) // i2的新值为：13

	// 搞定了三大定律，再看看反射的其他相关函数
	// Kind()函数，Type对象和Value对象都可以通过Kind()返回对应的变量的基础类型。关于基础类型，可关注/src/reflect/type.go，用const常量定义了多种基础类型
	// Kind()函数返回的是Kind类型，Kind类型的定义为：type Kind uint，这个返回值其实数值上等于const常量中定义的对应类型值，如ptr为22，int为2
	i3 := MyStruct{
		name:   "jone",
		age:    44,
		gender: "male",
	}
	t3 := reflect.TypeOf(&i3) // 也可以使用reflect.ValueOf，它们都可以调用Kind()方法。这里t3代表的是结构体的指针
	fmt.Println("Type: ", t3) // Type: *goBasic.MyStruct。从这里可以看出Type和Kind概念上还是有区别的
	fmt.Println("Kind: ", t3.Kind()) // Kind: ptr。如果传进来的是值（如reflect.Typeof(i3)），那么这里是struct。
	fmt.Printf("%d\n", t3.Kind()) // 22。即Ptr的常量值

	// 关于结构体的属性的操作
	// NumField()和Field()这两个方法无论是Type调用还是Value对象都有实现。不过注意，它们返回值不一定一样，比如t.Field返回的是StructField，v.Field返回的是Value
	// NumField()返回结构体的成员个数；Field()返回第i个属性的元素；FieldByName()根据结构体的成员名字返回对应元素
	v3 := reflect.ValueOf(&i3) // 这里v3代表的是结构体的指针
	fmt.Printf("结构体myStruct的字段个数：%d\n", v3.Elem().NumField()) // 结构体myStruct的字段数：3，也可以透过for循环遍历
	fmt.Printf("第1个字段名：%v，值：%v\n", t3.Elem().Field(0).Name, v3.Elem().Field(0)) // 第1个字段名：name，值：jone
	fmt.Printf("第2个字段名：%v，值：%v\n", t3.Elem().Field(1).Name, v3.Elem().Field(1)) // 第2个字段名：age，值：44
	fmt.Printf("第3个字段名：%v，值：%v\n", t3.Elem().Field(2).Name, v3.Elem().Field(2)) // 第3个字段名：gender，值：male
	// 关于结构体的方法的操作，最好使用指针级的对象不然有可能会panic，如这里的t3代表的是结构体指针
	// 与Field()类似，Type和Value各自的Field()返回值并不一样
	// NumMethod()返回结构体的方法个数（这里只包含大小开头的，可导出的。小写开头的，外面的包不可见因此不可导出）；Method()返回第i个方法的元素
	fmt.Printf("结构体myStruct的方法个数（只包含可导出的）：%d\n", t3.NumMethod()) // 结构体myStruct的方法个数（只包含可导出的）：3
	fmt.Printf("第1个方法名：%v\n", t3.Method(0).Name) // 第1个方法名：SayGoodBye（方法名Name是按照ASCII升序排序的）
	fmt.Printf("第2个方法名：%v\n", t3.Method(1).Name) // 第2个方法名：SayHello
	fmt.Printf("第3个方法名：%v\n", t3.Method(2).Name) // 第2个方法名：SellIntroduce
	// 结构体方法的调用
	// 调用无参数无返回值的函数，需要用Value对象的Call，Type对象没有该实现
	v3.Method(0).Call(nil) // GoodBye。调用的是第0个，也就是SayGoodBye方法。Call的参数为nil，代表没有参数要传给SayGoodBye
	// 使用另一种方式MethodByName，根据函数名而不是索引
	v3.MethodByName("SayHello").Call(nil) // hello。查找函数名SayHello并调用，无参数传入
	// 调用有参数的函数，需要先构造参数集切片，然后Call的时候将切片传进去（注意参数切片和方法参数的顺序要一致）
	name := reflect.ValueOf("Jone") // 每个参数是Value类型
	age := reflect.ValueOf(32)
	params := []reflect.Value{name, age} // 构造参数切片
	v3.MethodByName("SellIntroduce").Call(params)

	// Value中还有Int()、String()、Bool()、Float()、Slice()/Slice3()、Pointer()等转换函数可以将Value对象转换成对应的类型，并返回其值。
	// 注意调用前，一定要确保Value对象的类型是可行的，比如v.Int()那么v的原始类型必须是int类型而不能string、bool等，关于这些接口，这里不赘述

	// 关于切片的操作，除了上面的Value对象转切片类型的方法Slice()和Slice3()外，还有Append方法
	s1 := []int{1, 3, 4, 5, 7}
	v4 := reflect.ValueOf(&s1) // 因为涉及修改，一定要传入指针，这个应该毫无疑问
	v4Elem := v4.Elem() // v4是指针，因此要取出v4指向的元素（也就是s1）
	v4Elem.Set(reflect.Append(v4Elem, reflect.ValueOf(11))) // 这里分作两步看，先是Append得到一个更新后的切片的Value对象，然后Value.Set()方法更新对应的v4Elem。
	fmt.Printf("更新后的长度：%d\n", v4Elem.Len()) // 更新后的长度：6
	fmt.Println(v4Elem) // [1 3 4 5 7 11]
	fmt.Println(s1) // [1 3 4 5 7 11]。这里注意，上面一定要用Set方法不可直接v4Elem = reflect.Append(xxxxx)，不然s1不会被更新
}