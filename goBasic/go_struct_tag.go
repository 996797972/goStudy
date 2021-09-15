package goBasic

import (
	"fmt"
	"reflect"
)

// 在结构体成员额外再加一个属性，用反引号包含的字符串，这种就是TAG（标签）。其语法就是`key01:"value01" key02:"value02" key03:"value03"`
// 标签常用于将结构体的对象转化为json字符串
// 如何从结构体中取出TAG呢？答案就是通过反射机制。
// TAG作为结构体成员的一部分，当然也可以透过reflect的Type或者Value对象里的方法获取到，与通过Field()或者FieldByName方法获取成员对应的Field一样
// 大致分为以下三步：
// 1.获取结构体成员的Field：field := reflect.TypeOf(obj).Field(i)/FieldByName(name)，除了TypeOf当然也可以透过ValueOf去获取
// 2.通过Field，获得Field中的TAG：tag := fileld.TAG
// 3.通过TAG，获取对应key的值：tag.Get(key)/LookUp(key)
// 另外，空TAG和不设置TAG是一样的

// 以下Person结构体中，每个成员都有一个TAG（非空的）。TAG中有两个key：一个是label，一个是default。这两个key的值都可以用Tag.Get()或Tag.LookUp()取得
type Person struct {
	Name	string	`label:"His/Her name is" default:"Jone"`
	Age		int		`label:"His/Her age is"`
	Addr	string	`label:"His/Her addr is" default:"China"`
}

func SampleForStructTAG() {
	p := Person{
		Age: 14,
	}
	v := reflect.ValueOf(&p).Elem() // 获取Value对象，因为这里传进的是结构体变量的指针，因此还要加上Elem()去取出结构体对象

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag // 根据结构体变量的成员的索引来获取Tag。这里先透过v.Type()得到结构体变量的Type对象，然后Field(i)获取StructField，再获取Tag
		label := tag.Get("label") // 根据Tag的键来获取值，结构体的成员的每个Tag中都有两个key，一个label，一个default
		defaultValue := tag.Get("default")

		valString := fmt.Sprintf("%v", v.Field(i)) // 获取结构体变量p当前索引的成员的值。Sprintf会按格式去返回string类型
		if valString == "" { // 如果结构体变量p的此索引的成员的值没有初始化或者为空，那么将Tag中default的值作为成员值
			valString = defaultValue
		}
		fmt.Printf("%s %s\n", label, valString)
	}
}