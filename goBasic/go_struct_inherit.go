package goBasic

import "fmt"

type Company struct {
	companyName string
	companyAddr string
}

type Profile struct {
	name string
	age int
	gender string
	mother *Profile
	father *Profile
}
// 结构体的方法访问权限与普通函数或变量保持一致，大写时这个方法对于所有包是Public的，随意调用；而小写时只能本包内可见其他包无法调用，相当于Private
// 这里有个小细节，如果接收器是值方式而不是指针方式，那么涉及修改结构体成员的值操作都不会起效。因此接收器必须用指针，而调用的实例可以是指针也可以是值
func (p *Profile) IncreaseAge() {
	p.age += 1
}
// 为了统一，定义结构体的方法时最好接收器全部使用指针方式，事实上系统的代码也是这么写的。至于调用时的实例可以是指针也可以是值
func (p *Profile) PersonInfo() {
	fmt.Printf("姓名：%s，年龄：%d，性别：%s\n", p.name, p.age, p.gender)
}

// Staff由Profile和Company组合而成，继承了Profile
type Staff struct {
	p *Profile // 是有名成员，那么对于Staff实例，访问成员或调用方法需要指明。如s.p.PersonInfo()
	*Company // 作为匿名成员，那么对于Staff的实例，是直接继承Company成员以及方法。比如s.Company.companyName和s.companyName都可以访问，是一样的
	Control func() // 可以用函数变量作为结构体成员变量，用法和C语言函数指针类似
}

func test() {
	fmt.Printf("test \n")
}

func SampleForStructInherit()  {
	p1 := &Profile{ // 其实这里如果是实例的值，同样也可以调用方法（虽然定义方法时接收器为指针）。不过，为了语义正确，这里还是最好用指针
		name:   "jone",
		gender: "male",
		age:    32,
	}
	p1.PersonInfo()
	p1.IncreaseAge() // 调用的实例可以是指针也可以是值，由于方法定义的接收器是指针，所以修改结构体内部成员的值是可行和有效的
	p1.PersonInfo()

	// 使用继承，对继承的结构体进行赋值
	s := Staff{
		p:       p1,
		Company: &Company{
			companyName: "Google",
			companyAddr: "American",
		},
		Control: test, // 对结构体的函数变量赋值
	}
	s.p.IncreaseAge()
	s.p.PersonInfo()
	s.Control() // 调用结构体中的函数变量成员
	fmt.Printf("Company: %s\n", s.Company.companyAddr) // 由于Company是Staff的匿名成员，所以这里简写成s.companyAddr，两者没有任何区别
	fmt.Printf("Company: %s\n", s.companyAddr) // 对于结构体作匿名成员，外部实例直接访问内部成员或方法即可
}
