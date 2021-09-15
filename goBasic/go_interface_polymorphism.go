package goBasic

import (
	"fmt"
)

// 接口：定义一个对象的行为。它指定了对象应该做什么，至于如何实现这个行为（实现细节），则由对象本身去确定
// Go语言中，接口就是方法签名的集合。当一个类型定义了某个接口中的所有方法，我们称它实现了该接口（这符合OOP说法）
// 接口指定了一个类型应该具有的方法，并由该类型确定如何实现这些方法。Go中的接口是实现多态的唯一好途径

// 定义一个接口Phone，并声明了一个Call方法。如果有一个结构体或类型实现了call方法，那么我们就称它实现了Phone接口。
// 其意思就是如果有一台机器，可以打电话给别人，那么我们就能把它叫做电话
// 它的变量可以等于任何所有实现了Call方法的结构，从而实现了多态
type Phone interface {
	Call() // Call方法没参数，没返回值
}

// 定义一个结构体，叫Nokia。同时实现call接口，所以它也是一部电话
type Nokia struct {
	name string
}
// 接口的实现是隐式的，不像Java中要用implements显式说明。其遵循鸭子类型定义，如果一个类型实现了在接口中定义的签名方法，则称该类型实现该接口
// 鸭子类型：只要你长得像鸭子，叫起来也像鸭子，那我认为你就是一个鸭子
// 类型要实现的方法名与接口中的方法签名必须一致，比如这里的Call()。而且，该类型必须实现接口声明的所有方法，否则会出错。
func (p Nokia) Call() {
	fmt.Printf("我是%s，是一部电话\n", p.name)
}

// 类型开关的例子
func explain(i interface{}) {
	switch i.(type) { // 类型开关，type是固定的关键字，返回空接口参数的动态类型
	case string:
		fmt.Printf("接口的动态类型为%T\n",i)
	case int:
		fmt.Printf("接口的动态类型为%T\n",i)
	default:
		fmt.Printf("不是很清楚这个是什么类型\n")
	}
}

// 一个商品的接口，包含了计算价格和订单信息方法。如果有一个类型，它提供了这两个方法，那么我们就可以称之为“商品”
type Good interface {
	settleAccount() int // 计算总金额
	orderInfo() // 订单信息
}
// 定义两个结构体，鞋和赠品，并且都实现了Good接口的settleAccount和orderInfo方法，那么它们都属于“商品”
type Shoe struct {
	name string
	quantity int
	price int
}
func (s *Shoe) settleAccount() int {
	return s.quantity * s.price
}
func (s *Shoe) orderInfo()  {
	fmt.Printf("您要购买的鞋子：%s，数量：%d双，单价：%d，总金额：%d\n", s.name, s.quantity, s.price, s.settleAccount())
}

type FreeGift struct {
	name string
	quantity int
	price int
}
func (f *FreeGift) settleAccount() int {
	return 0 // 既然是赠送品，那么应该是免费的
}
func (f *FreeGift) orderInfo()  {
	fmt.Printf("您得到的礼品：%s，数量：%d双，单价：%d，总金额：%d\n", f.name, f.quantity, f.price, f.settleAccount())
}

func caculateAllPrice(goods []Good) int {
	allPrice := 0
	for _, g := range goods {
		g.orderInfo()
		allPrice += g.settleAccount()
	}
	return allPrice
}

// 多接口，指的是一个类型可实现了多个不同的接口。用现实生活来描述就是，一个人可以掌握多种技能
type Shape interface {
	Area() int
}

type Object interface {
	Volume() int
}

type Cube struct { // 正方体既有面积也有体积，因此完全可既实现Shape接口也可实现Object接口
	side int
}
func (c *Cube) Area() int {
	return 6 * c.side * c.side
}
func (c *Cube) Volume() int {
	return c.side * c.side * c.side
}

// 接口嵌套，这里仅仅介绍。将Shape接口和Object接口形成一个新的接口Material
type Material interface {
	Shape // 注意，这是接口，可不是方法
	Object
}

func SampleForInterface() {
	var p Phone // 在这里，p的值是nil。因为它是一个静态接口，还不知道是谁会实现它。
	p = Nokia{"N93"} // 由于Nokia实现了Call方法，p当然可以接收Nokia类型。也就是p的类型已经动态变成了Nokia。一旦某个接口变量得到赋值，它就有了动态类型、动态值
	fmt.Printf("type of p is %T, value %v\n", p, p) // type of p is goBasic.Nokia, value {N93}
	p.Call() // 这里调用的固然是Nokia.Call方法，因为p的类型是Nokia。输出：我是N93，是一部电话

	// 空接口
	// 自创建空接口
	var eI interface{} // 什么方法都不需要实现就是空接口
	eI = 13 // 既然是空接口，相当于任何类型都实现了该接口，那么可以使用任意类型赋值给它
	fmt.Printf("type of p is %T\n", eI) // type of p is int
	// 使用空接口作为函数参数，同样也可以传参任意类型。
	explain("Hello world") // explain函数中包含了接口的类型开关使用例子
	explain(35)
	explain(byte(10))

	// 商品例子
	shoe := Shoe{
		name:     "鸿星尔克",
		quantity: 2,
		price:    300,
	}
	book := FreeGift{
		name:     "一万个为什么",
		quantity: 1,
		price:    300,
	}
	goods := []Good{&shoe, &book} // 既然实现了接口，那么自然也就可以把实例类型赋值给接口变量
	allPrice := caculateAllPrice(goods) // 输出各物品的信息，并计算总金额
	fmt.Printf("此订单的总金额：%d\n", allPrice)

	// 多接口，即一个类型实现了多个接口。如果将这些接口合成一个新的接口那么又可以称为接口嵌套
	var s Shape
	s = &Cube{3} // 由于Cube结构体的方法接收者都是指针，这里接口变量要存的必须是指针，否则会出错
	fmt.Printf("面积：%d\n", s.Area())
	o, ok := s.(Object) // 使用断言获取Object，这里的o其实是var o Object。如果Cube没有实现Object的Volume方法，那么这里ok会返回false
	if ok {
		fmt.Printf("体积：%d\n", o.Volume())
	}
}