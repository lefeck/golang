# Golang 面向对象（封装、继承、多态）


Go语言并不像Java那样有类的概念，以及extends这样的关键字，但是可以用其特有的数据结构来实现类似面向对象的特性。主要有结构体实现封装，组合实现继承，接口实现多态。
封装可以隐藏类的实现细节并使代码具备模块化，继承可以扩展已存在的模块，多态的作用除了复用性外，还可以解决模块之间高耦合的问题。

## 一、结构体实现封装

在Go语言中，我们可以对结构体的字段进行封装，并通过结构体中的方法来操作内部的字段。如果结构体中字段名的首字母是小写字母，那么这样的字段是私有的，相当于private字段。外部包裹能直接访问，如果是在名的首字母是大写字母，那么这样的字段对外暴露的，相当于public字段。能够起的方法也是一样的，如果方法名首字母是大写字母，那么这样的方法对外暴露的。
下面案例体现了封装特性：

```go
package main

type Person struct {
	name string
	age int
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) SetAge(age int) {
	p.age = age
}

func (p *Person) GetName() string{
	return p.name
}

func (p *Person) GetAge() int {
	return p.age
}

```

## 二、组合实现继承

面向对象的继承特性实际上就是一种组件复用的机制。在Java中可以先定一个父类，然后通过继承特性来实现子类继承父类的功能。但Go语言中没有继承的关键字extends，而是采用组合的方式来实现继承的效果，组合和继承是有区别的，组合可以理解为has - a 关系，继承可以理解为is - a的关系。但他们都实现了代码复用的目的。组合相对于继承的优点有：

* 可以利用面向接口编程原则的一系列优点，封装性行耦合性低
* 相对于继承的编译期确定实现，组合的运行态指定实现更加灵活
* 组合是非侵入式的，继承是侵入式的

父类Person类：

```go
type Person struct {
	name string
	age int
}

func (p *Person) Eat() {
	fmt.println("Person Eat")
}

func (p *Person) Walk() {
	fmt.Println("Person Walk")
}

```

子类Student继承Person类：

```go
type Student struct {
	Person //组合Person，注意首字母大写，否则无法继承属性和方法
	school string
}

func (s *Student) study() {
	fmt.Println(s.name, "study") //调用了name，这里的name就是继承自person结构体的
}

//重写方法，会覆盖Person中的walk方法
func (s *Student) Walk() {
	fmt.Println(s.name, "walk")
}
```



## 三、接口实现多态

### 接口定义与实现

接口是一种对约定标准进行定义的，如果一个结构体嵌入了接口类型，那么任何其他类型实现了该接口都可以与之进行组合调用。接口实现最明显的优点就是实现了类和接口的分离，在切换实现类的时候不用更换接口功能。
在Go语言中定义接口的语法如下：

```go
type 接口名 interface {
	方法
}
```

在Go语言中对接口的实现只需要某个类型T实现了接口中的方法，就相当于实现了该接口。定义接口并实现的示例如下：

```go
import "fmt"

//定义一个考试接口
type exam interface {
	exam()
}

type Student struct {
	Person //组合Person，注意首字母大写，否则无法继承属性和方法
	school string
}

//实现了exam接口
func (s *Student) exam() {
	fmt.Println(s.name, "exam")
}
```

### 接口实现多态

在面向对象的语言中，接口的多种不同实现方式即为多态。多态的示例如下：

```go
package main

import "fmt"

//Person接口
type Person interface {
	ToSchool()
}

//学生类
type Student struct {
	work string
}

//学生类实现Person接口
func (this *Student) ToSchool() {
	fmt.Println("Student ", this.work)
}

//老师类
type Teacher struct {
	work string
}

//老师类实现Person接口
func (this *Teacher) ToSchool() {
	fmt.Println("Teacher ", this.work)
}

//工厂模式函数，根据传入工作不同动态返回不同类型
func Factory(work string) Person {
	switch work {
	case "study":
		return &Student{work: "study"}
	case "teach":
		return &Teacher{work: "teach"}
	default:
		panic("no such profession")
	}
}

func main() {
	person := Factory("study")
	person.ToSchool() //Student  study

	person = Factory("teach")
	person.ToSchool() //Teacher  teach

}
```



