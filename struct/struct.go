package _struct

import (
	"fmt"
	"unsafe"
)

// Myint 自定义类型
type Myint int

// NewInt 类型别名
type NewInt = int

// 结构体定义
// 类型名：标识自定义结构体的名称，在同一个包内不能重复。
// 字段名：表示结构体字段名。结构体中的字段名必须唯一。
// 字段类型：表示结构体字段的具体类型。
type Person struct {
	username, city string
	age            int
}

// 空结构体
type Empty struct {
}

type Student struct {
	name string
	age  int
}

// Go语言的结构体没有构造函数，我们可以自己实现。
// 例如，下方的代码就实现了一个person的构造函数。
// 因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型
func NewStudent(name string, age int) *Student {
	return &Student{name: name, age: age}
}

func (s Student) GetName() string {
	return s.name
}

func (s Student) GetAge() int {
	return s.age
}

func (s *Student) SetName(name string) {
	s.name = name
}

func (s Student) SetAge(age int) {
	s.age = age
}

// Address 地址结构体
type Address struct {
	Province string
	City     string
}

// User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}

func Demo() {
	// 类型定义和类型别名的区别
	var a Myint
	var b NewInt
	fmt.Printf("type of a:%T\n", a) // a的类型是main.NewInt，表示main包下定义的NewInt类型
	fmt.Printf("type of b:%T\n", b) // MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型

	// 结构体初始化
	var person1 Person // 没有初始化的结构体，其成员变量都是对应其类型的零值
	fmt.Println(person1)

	// 使用键值对初始化时键对应结构体的字段，值对应该字段的初始值
	person2 := Person{
		username: "eric",
		city:     "chengdu",
		age:      18,
	}
	fmt.Printf("%v", person2)

	// 也可以对结构体指针进行键值对初始化。
	// 当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值
	person3 := &Person{
		username: "eric",
		city:     "chengdu",
	}
	fmt.Printf("%v", person3)

	// 使用值的列表初始化
	// 初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值
	//
	// 使用这种格式初始化时，需要注意：
	//
	// 1.必须初始化结构体的所有字段。
	// 2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	// 3.该方式不能和键值初始化方式混用。
	person4 := Person{"bob", "USA", 20}
	fmt.Printf("%v \n", person4)

	/*========== 结构体内存布局 ===========*/

	// 结构体占用一块连续的内存。
	fmt.Printf("%p \n", &person4.username) // 0xc000098240 在不同的系统中该值不一样
	fmt.Printf("%p \n", &person4.city)     // 0xc000098250
	fmt.Printf("%p \n", &person4.age)      // 0xc000098240

	// 空结构体 空结构体是不占用空间的
	var e Empty
	fmt.Println(unsafe.Sizeof(e)) // 0

	// 面试题
	m := make(map[string]*Student)
	stus := []Student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	// 根本原因在于for-range会使用同一块内存去接收循环中的值
	for _, stu := range stus {
		m[stu.name] = &stu
		fmt.Println(m[stu.name])
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	/*========== 构造函数 ===========*/
	zhangsan := NewStudent("张三", 18)
	fmt.Println(zhangsan)

	/*========== 方法和接收者 ===========*/
	lisi := NewStudent("李四", 18)
	fmt.Printf("lisi name: %s, age: %d \n", lisi.GetName(), lisi.GetAge())

	// 指针类型的接收者
	// 指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。
	wangwu := NewStudent("王武", 18)
	wangwu.SetName("王武1")
	fmt.Println(wangwu.name) // 王武1

	// 值类型的接收者
	// 当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。
	// 在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。
	// 什么时候应该使用指针类型接收者?
	// 1.需要修改接收者中的值
	// 2.接收者是拷贝代价比较大的大对象
	// 3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
	lisan := NewStudent("李三", 18)
	lisan.SetAge(20)
	fmt.Println(lisan.age) // 18

	/*========== 嵌套结构体 ===========*/
	user1 := User{
		Name:   "user1",
		Gender: "man",
		Address: Address{
			Province: "sc",
			City:     "cd",
		},
	}
	fmt.Printf("%v", user1)
}
