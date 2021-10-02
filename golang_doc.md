# 变量和常量

变量和常量是编程中必不可少的部分，也是很好理解的一部分。

### 标识符与关键字

#### 标识符

在编程语言中标识符就是程序员定义的具有特殊意义的词，比如变量名、常量名、函数名等等。 Go语言中标识符由字母数字和_(下划线）组成，并且只能以字母和_开头。 举几个例子：abc, _, _123, a123。

#### 关键字

关键字是指编程语言中预先定义好的具有特殊含义的标识符。 关键字和保留字都不建议用作变量名。

Go语言中有25个关键字：

    break        default      func         interface    select
    case         defer        go           map          struct
    chan         else         goto         package      switch
    const        fallthrough  if           range        type
    continue     for          import       return       var

此外，Go语言中还有37个保留字。

    Constants:    true  false  iota  nil
    
        Types:    int  int8  int16  int32  int64  
                  uint  uint8  uint16  uint32  uint64  uintptr
                  float32  float64  complex128  complex64
                  bool  byte  rune  string  error
    
    Functions:   make  len  cap  new  append  copy  close  delete
                 complex  real  imag
                 panic  recover

### 变量

#### 变量的来历

程序运行过程中的数据都是保存在内存中，我们想要在代码中操作某个数据时就需要去内存上找到这个变量，但是如果我们直接在代码中通过内存地址去操作变量的话，代码的可读性会非常差而且还容易出错，所以我们就利用变量将这个数据的内存地址保存起来，以后直接通过这个变量就能找到内存上对应的数据了。

#### 变量类型

变量（Variable）的功能是存储数据。不同的变量保存的数据类型可能会不一样。经过半个多世纪的发展，编程语言已经基本形成了一套固定的类型，常见变量的数据类型有：整型、浮点型、布尔型等。

Go语言中的每一个变量都有自己的类型，并且变量必须经过声明才能开始使用。

#### 变量声明

Go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。 并且Go语言的变量声明后必须使用。

#### 标准声明

Go语言的变量声明格式为：

```go
var 变量名 变量类型
```

变量声明以关键字`var`开头，变量类型放在变量的后面，行尾无需分号。 举个例子：

```go
var name string
var age int
var isOk bool
```

#### 批量声明

每声明一个变量就需要写`var`关键字会比较繁琐，go语言中还支持批量变量声明：

```go
var (
    a string
    b int
    c bool
    d float32
)
```

#### 变量的初始化

Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。每个变量会被初始化成其类型的默认值，例如： 整型和浮点型变量的默认值为`0`。 字符串变量的默认值为`空字符串`。 布尔型变量默认为`false`。 切片、函数、指针变量的默认为`nil`。

当然我们也可在声明变量的时候为其指定初始值。变量初始化的标准格式如下：

```go
var 变量名 类型 = 表达式
```

举个例子：

```go
var name string = "Q1mi"
var age int = 18
```

或者一次初始化多个变量

```go
var name, age = "Q1mi", 20
```

##### 类型推导

有时候我们会将变量的类型省略，这个时候编译器会根据等号右边的值来推导变量的类型完成初始化。

```go
var name = "Q1mi"
var age = 18
```

##### 短变量声明

在函数内部，可以使用更简略的 `:=` 方式声明并初始化变量。

```go
package main

import (
	"fmt"
)
// 全局变量m
var m = 100

func main() {
	n := 10
	m := 200 // 此处声明局部变量m
	fmt.Println(m, n)
}
```

**总结:** 声明变量的方式：

```
1.必须使用显示初始化；
2.不提供数据类型时，编译器会自动推导；
3.只能在函数内部使用简短模式；
```



##### 匿名变量

在使用多重赋值时，如果想要忽略某个值，可以使用`匿名变量（anonymous variable）`。 匿名变量用一个下划线`_`表示，例如：

```go
func foo() (int, string) {
	return 10, "Q1mi"
}
func main() {
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
```

匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。 (在`Lua`等编程语言里，匿名变量也被叫做哑元变量。)

注意事项：

1. 函数外的每个语句都必须以关键字开始（var、const、func等）
2. `:=`不能使用在函数外。
3. `_`多用于占位，表示忽略值。

### 常量

相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。 常量的声明和变量声明非常类似，只是把`var`换成了`const`，常量在定义的时候必须赋值。

```go
const pi = 3.1415
const e = 2.7182
```

声明了`pi`和`e`这两个常量之后，在整个程序运行期间它们的值都不能再发生变化了。

多个常量也可以一起声明：

```go
const (
    pi = 3.1415
    e = 2.7182
)
```

const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。 例如：

```go
const (
    n1 = 100
    n2
    n3
)
```

上面示例中，常量`n1`、`n2`、`n3`的值都是100。

#### iota

`iota`是go语言的常量计数器，只能在常量的表达式中使用。

`iota`在const关键字出现时将被重置为0。const中每新增一行常量声明将使`iota`计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。

举个例子：

```go
const (
		n1 = iota //0
		n2        //1
		n3        //2
		n4        //3
	)
```

#### 几个常见的`iota`示例:

使用`_`跳过某些值

```go
const (
		n1 = iota //0
		n2        //1
		_
		n4        //3
	)
```

`iota`声明中间插队

```go
const (
		n1 = iota //0
		n2 = 100  //100
		n3 = iota //2
		n4        //3
	)
	const n5 = iota //0
```

定义数量级 （这里的`<<`表示左移操作，`1<<10`表示将1的二进制表示向左移10位，也就是由`1`变成了`10000000000`，也就是十进制的1024。同理`2<<2`表示将2的二进制表示向左移2位，也就是由`10`变成了`1000`，也就是十进制的8。）

```go
const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
```

多个`iota`定义在一行

```go
const (
		a, b = iota + 1, iota + 2 //1,2
		c, d                      //2,3
		e, f                      //3,4
	)
```



# 数据类型

Go语言中有丰富的数据类型，除了基本的整型、浮点型、布尔型、字符串外，还有数组、切片、结构体、函数、map、通道（channel）等。Go 语言的基本类型和其他语言大同小异。

## 基本数据类型
### 整型
整型分为以下两个大类： 按长度分为：int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64

其中，uint8就是我们熟知的byte型，int16对应C语言中的short型，int64对应C语言中的long型。

类型 | 描述
---|---
uint8    |  无符号 8位整型 (0 到 255)
uint16   |  无符号 16位整型 (0 到 65535) 
uint32   |  无符号 32位整型 (0 到 4294967295) 
uint64   |  无符号 64位整型 (0 到 18446744073709551615) 
int8	 |  有符号 8位整型 (-128 到 127)  
int16    |  有符号 16位整型 (-32768 到 32767) 
int32    |  有符号 32位整型 (-2147483648 到 2147483647)   
int64	 |  有符号 64位整型 (-9223372036854775808 到 9223372036854775807)


### 特殊整型
类型 | 描述
---|---
uint     |  32位操作系统上就是uint32，64位操作系统上就是uint64
int      |  32位操作系统上就是int32，64位操作系统上就是int64 
uintptr  |  无符号整型，用于存放一个指针
注意：在使用int和uint类型时，不能假定它是32位或64位的整型，而是考虑int和uint可能在不同平台上的差异。

注意事项: 获取对象的长度的内建len()函数返回的长度可以根据不同平台的字节长度进行变化。实际使用中,切片或map的元素数量等都可以用int来表示。在涉及到二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用int和 uint。

### 数字字面量语法（Number literals syntax）
Go1.13版本之后引入了数字字面量语法，这样便于开发者以二进制、八进制或十六进制浮点数的格式定义数字，例如：

v := 0b00101101， 代表二进制的 101101，相当于十进制的 45。 v := 0o377，代表八进制的 377，相当于十进制的 255。 v := 0x1p-2，代表十六进制的 1 除以 2²，也就是 0.25。 而且还允许我们用 _ 来分隔数字，比如说：

v := 123_456 等于 123456。

我们可以借助fmt函数来将一个整数以不同进制形式展示。
```
package main
 
import "fmt"
 
func main(){
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a)  // 10
	fmt.Printf("%b \n", a)  // 1010  占位符%b表示二进制
 
	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b)  // 77
 
	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c)  // ff
	fmt.Printf("%X \n", c)  // FF
}
```
### 浮点型
Go语言支持两种浮点型数：float32和float64。这两种浮点型数据格式遵循IEEE 754标准： float32 的浮点数的最大范围约为 3.4e38，可以使用常量定义：math.MaxFloat32。 float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64。

打印浮点数时，可以使用fmt包配合动词%f，代码如下：
```
package main
import (
        "fmt"
        "math"
)
func main() {
        fmt.Printf("%f\n", math.Pi)
        fmt.Printf("%.2f\n", math.Pi)
}
```
### 复数

complex64和complex128
```
var c1 complex64
c1 = 1 + 2i
var c2 complex128
c2 = 2 + 3i
fmt.Println(c1)
fmt.Println(c2)
```
复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。

### 布尔值
Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。

注意：

    1.布尔类型变量的默认值为false。
    2.Go 语言中不允许将整型强制转换为布尔型.
    3.布尔型无法参与数值运算，也无法与其他类型进行转换。

### 字符串
Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。 Go 语言里的字符串的内部实现使用UTF-8编码。 字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符，例如：
```
s1 := "hello"
s2 := "你好"
```
### 字符串转义符
Go 语言的字符串常见转义符包含回车、换行、单双引号、制表符等，如下表所示。

转义符	| 	含义
	 ---|---
\r		| 	回车符（返回行首）
\n		| 	换行符（直接跳到下一行的同列位置）
\t		| 	制表符
\'		| 	单引号
\"		| 	双引号
\\		| 	反斜杠
举个例子，我们要打印一个Windows平台下的一个文件路径：
```
package main
import (
    "fmt"
)
func main() {
    fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")
}
```
### 多行字符串
Go语言中要定义一个多行字符串时，就必须使用反引号字符：
```
s1 := `第一行
第二行
第三行
`
fmt.Println(s1)
```
反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。

### 字符串的常用操作
方法	| 介绍
---|---
len(str)	          						|  求长度
+或fmt.Sprintf	      						|  拼接字符串
strings.Split	      						|  分割
strings.contains	  						|  判断是否包含
strings.HasPrefix,strings.HasSuffix			|  前缀/后缀判断
strings.Index(),strings.LastIndex()			|  子串出现的位置
strings.Join(a[]string, sep string)	  	    |  Join 操作
### byte和rune类型
组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来，如：
```
var a := '中'
var b := 'x'
```
Go 语言的字符有以下两种：

uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
rune类型，代表一个 UTF-8字符。
当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32。

Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾。
```
// 遍历字符串
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}
输出：

104(h) 101(e) 108(l) 108(l) 111(o) 230(æ) 178(²) 153() 230(æ) 178(²) 179(³) 
104(h) 101(e) 108(l) 108(l) 111(o) 27801(沙) 27827(河) 
```
因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串，否则就会出现上面输出中第一行的结果。

字符串底层是一个byte数组，所以可以和[]byte类型相互转换。字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。 rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。

### 修改字符串
要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
```
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}
```
### 类型转换
Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。

强制类型转换的基本语法如下：
> 
T(表达式)
其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等.

比如计算直角三角形的斜边长时使用math包的Sqrt()函数，该函数接收的是float64类型的参数，而变量a和b都是int类型的，这个时候就需要将a和b强制类型转换为float64类型。
```
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
```
### 练习题

编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。
编写代码统计出字符串"hello沙河小王子"中汉字的数量。



# 流程控制

流程控制是每种编程语言控制逻辑走向和执行次序的重要部分，流程控制可以说是一门语言的“经脉”。

Go语言中最常用的流程控制有if和for，而switch和goto主要是为了简化代码、降低重复代码而生的结构，属于扩展类的流程控制。

## if else(分支结构)

### if条件判断基本写法

Go语言中if条件判断的格式如下：

```
if 表达式1 {
    分支1
} else if 表达式2 {
    分支2
} else{
    分支3
}
```

当表达式1的结果为true时，执行分支1，否则判断表达式2，如果满足则执行分支2，都不满足时，则执行分支3。 if判断中的else if和else都是可选的，可以根据实际需要进行选择。

Go语言规定与if匹配的左括号{必须与if和表达式放在同一行，{放在其他位置会触发编译错误。 同理，与else匹配的{也必须与else写在同一行，else也必须与上一个if或else if右边的大括号在同一行。

举个例子：

```
func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
```

### if条件判断特殊写法

if条件判断还有一种特殊的写法，可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断，举个例子：

```
func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
```

思考题： 上下两种写法的区别在哪里？

## for(循环结构)

Go 语言中的所有循环类型均可以使用for关键字来完成。

for循环的基本格式如下：

```
for 初始语句;条件表达式;结束语句{
    循环体语句
}
```

条件表达式返回true时循环体不停地进行循环，直到条件表达式返回false时自动退出循环。

```
func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
```

for循环的初始语句可以被忽略，但是初始语句后的分号必须要写，例如：

```
func forDemo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}
```

for循环的初始语句和结束语句都可以省略，例如：

```
func forDemo3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
```

这种写法类似于其他编程语言中的while，在while后添加一个条件表达式，满足条件表达式时持续循环，否则结束循环。

## 无限循环

```
for {
    循环体语句
}
```

for循环可以通过break、goto、return、panic语句强制退出循环。

## for range(键值循环)

Go语言中可以使用for range遍历数组、切片、字符串、map 及通道（channel）。 通过for range遍历的返回值有以下规律：

1. 数组、切片、字符串返回索引和值。
2. map返回键和值。
3. 通道（channel）只返回通道内的值。

## switch case

使用switch语句可方便地对大量的值进行条件判断。

```
func switchDemo1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}
```

Go语言规定每个switch只能有一个default分支。

一个分支可以有多个值，多个case值中间使用英文逗号分隔。

```
func testSwitch3() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}
```

分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。例如：

```
func switchDemo4() {
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}
```

fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。

```
func switchDemo5() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
```

输出：

```
a
b
```



## goto(跳转到指定标签)

goto语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。Go语言中使用goto语句能简化一些代码的实现过程。 例如双层嵌套的for循环要退出时：

```
func gotoDemo1() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag {
			break
		}
	}
}
```

使用goto语句能简化代码：

```
func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}
```



## break(跳出循环)

break语句可以结束for、switch和select的代码块。

break语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的for、switch和 select的代码块上。 举个例子：

```
func breakDemo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}
```



## continue(继续下次循环)

continue语句可以结束当前循环，开始下一次的循环迭代过程，仅限在for循环内使用。

在 continue语句后添加标签时，表示开始标签对应的循环。例如：

```
func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}
```

### 练习题

编写代码打印9*9乘法表。

```
func main() {
	//顺序打印99乘法表	
for i := 1; i < 10; i++ {
	for j := 1; j <= i; j++ {
		fmt.Printf("%d * %d = %d\t",  j, i, i*j)
	}
	fmt.Println()
}
	//倒序打印99乘法表
for i := 9; i >= 1; i-- {
	for j := 1; j <= i; j++ {
		fmt.Printf("%d * %d = %d\t",  j,  i,  i*j)
	}
	fmt.Println()
}
}
```

### 斐波那契数列的几种实现方法

斐波纳契数列，又称黄金分割数列，指的是这样一个数列：1、1、2、3、5、8、13、……在数学上，斐波纳契数列以如下被以递归的方法定义：F0=0，F1=1，Fn=F(n-1)+F(n-2)（n>=2，n∈N*）

在学习go语言基础的过程中，学习了斐波那契数列的几种实现方法，总的可以分为递归和非递归实现。本文按照用到的方法侧重点不同，现细分如下：

```
递归实现
递归实现改进
数组递归实现
闭包实现
```

#### 递归实现

```
package main

import "fmt"

const LIM = 40

func main() {
    //result := 0
    //var array []int
    var array [LIM]int  //声明一个数组array并初始化
    for i := 0; i < LIM; i++ {
        array[i] = fibonacci(i)
        //result = fibonacci(i)
        //fmt.Println(result)
        //array = append(array, result)
        //fmt.Printf("fibonacci(%d) is: %d\n", i, result)
    }
    fmt.Println(array)
}

func fibonacci(n int) (res int) {
    if n <= 1 {
        res = 1
    } else {
        res = fibonacci(n-1) + fibonacci(n-2)
    }
    return
}
```

### 递归实现改进

```
package main

import "fmt"

const LIM = 40

var fibs [LIM]uint64

func main() {
    //var result uint64 = 0
    var array [LIM]uint64
    for i := 0; i < LIM; i++ {
        array[i] = fibonacci(i)
        //result = fibonacci(i)
        //array = append(array, result)
        //fmt.Printf("fibonacci(%d) is: %d\n", i, result)
    }
    fmt.Println(array)
}

func fibonacci(n int) (res uint64) {
    // memoization: check if fibonacci(n) is already known in array:
    if fibs[n] != 0 {
        res = fibs[n]
        return
    }
    if n <= 1 {
        res = 1
    } else {
        res = fibonacci(n-1) + fibonacci(n-2)
    }
    fibs[n] = res
    return
}
```

### 数组递归实现

```
package main

import "fmt"

const LIM = 40

func main() {
    fmt.Println(fibarray(LIM))
}

func fibarray(term int) []int {
    farr := make([]int, term)
    farr[0], farr[1] = 1, 1

    for i:= 2; i < term; i++ {
        farr[i] = farr[i-1] + farr[i-2]
    }
    return farr
}
```

### 闭包实现

```
package main

import "fmt"

const LIM = 40

func main() {
    f := fibonacci() //返回一个闭包函数
    var array [LIM]int
    for i := 0; i < LIM; i++ {
        array[i] = f()
    }
    fmt.Println(array)
}

func fibonacci() func() int {
    back1, back2 := 0, 1
    return func() int {
        // 重新赋值
        back1, back2 = back2, (back1 + back2)
        return back1
    }
}
```

### 小结

总体来说，用time包可以测试几种实现方法所用时间，对比得出递归实现会比闭包实现耗时更长，效率远远低于非递归实现。



# 数组

本文主要介绍Go语言中数组（array）及它的基本使用。

## Array(数组)

数组是同一种数据类型元素的集合。 在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。 基本语法：

> // 定义一个长度为3元素类型为int的数组a
> var a [3]int

### 数组定义：

var 数组变量名 [元素数量]T
比如：var a [5]int， 数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，长度不能变。 [5]int和[10]int是不同的类型。

```
var a [3]int
var b [4]int
a = b //不可以这样做，因为此时a和b是不同的类型
```

数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1，访问越界（下标在合法范围之外），则触发访问越界，会panic。

### 数组的初始化

数组的初始化也有很多方式。

#### 方法一

初始化数组时可以使用初始化列表来设置数组元素的值。

```
func main() {
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]
}
```

#### 方法二

按照上面的方法每次都要确保提供的初始值和数组长度一致，一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度，例如：

```
func main() {
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)                          //[0 0 0]
	fmt.Println(numArray)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	fmt.Println(cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string
}
```

#### 方法三

我们还可以使用指定索引值的方式来初始化数组，例如:

```
func main() {
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int
}
```

### 数组的遍历

遍历数组a有以下两种方法：

```
func main() {
	var a = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法2：for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}
```

### 多维数组

Go语言是支持多维数组的，我们这里以二维数组为例（数组中又嵌套数组）。

#### 二维数组的定义

```
func main() {
	a := [3][2]string{   //有3个元素，每一个元素都是长度为2的string数组
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a) //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆
}
```

#### 二维数组的遍历

```
func main() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}
```

输出：

```
北京	上海	
广州	深圳	
成都	重庆
```

注意： 多维数组只有第一层可以使用...来让编译器推导数组长度。例如：

```
//支持的写法
a := [...][2]string{
	{"北京", "上海"},
	{"广州", "深圳"},
	{"成都", "重庆"},
}

//不支持多维数组的内层使用...
b := [3][...]string{
	{"北京", "上海"},
	{"广州", "深圳"},
	{"成都", "重庆"},
}
```

#### 数组是值类型

数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。

```
func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}
func main() {
	a := [3]int{10, 20, 30}
	modifyArray(a) //在modify中修改的是a的副本x
	fmt.Println(a) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
}
```

注意：

数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的。
[n]*T表示指针数组，*[n]T表示数组指针 。

### 练习题

1. 求数组[1, 3, 5, 7, 8]所有元素的和
2. 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。



# 切片


因为数组的长度是固定的并且数组长度属于类型的一部分，所以数组有很多的局限性。 例如：

```
func arraySum(x [3]int) int{
    sum := 0
    for _, v := range x{
        sum = sum + v
    }
    return sum
}
```

这个求和函数只能接受[3]int类型，其他的都不支持。 再比如，

> a := [3]int{1, 2, 3}

数组a中已经有三个元素了，我们不能再继续往数组a中添加新元素了。

### 切片

切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。

切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。

### 切片的定义

声明切片类型的基本语法如下：

> var name []T

其中，

* name: 表示变量名
* T: 表示切片中的元素类型

举个例子：

```
func main() {
	// 声明切片类型
	var a []string              //声明一个字符串切片，没有分配内存，==nil
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[]
	fmt.Println(c)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false
	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
}
```



### **`nil`切片**

```
var s []int
fmt.Println(s == nil)   // 输出 true
fmt.Println(len(s),cap(s))   // 输出：0 0
```

上面这段代码声明了一个`nil`切片`s`，其实，切片的零值就是`nil`。为什么？通过[上一节](https://mp.weixin.qq.com/s?__biz=MzUxNjk4MDY1NA==&mid=2247483856&idx=1&sn=18bf041f7c548cd55b7d29cdfb1a332d&chksm=f99e6b11cee9e207766dc3b1a8e85c498c9da0f59b2d011041690ccf7cd6c0ee9fe7309248ea&token=556176193&lang=zh_CN&scene=21#wechat_redirect)我们知道，因为切片就是一个数组的引用。切片的类型在初始化时已经确认，就是`[]Type`，上面的代码就声明了`[]int`类型的`nil`切片`s`。`nil`切片的指向底层数组的指针为`nil`。

### **空切片**

如何声明空切片？有两种方式：

```
// 1、使用 make 创建空的整型切片
s := make([]int, 0)

// 2、使用切片字面量创建空的整型切片
s := []int{}
fmt.Println(s)   // 输出：[]
fmt.Println(len(s),cap(s))   // 输出：0 0
```

通过上面代码可以得出，与`nil`切片一样，空切片的长度和容量也都是 0，说明切片底层的数组大小为 0，是一个空数组（没有分配任何的存储空间）。

> 不管是使用 `nil` 切片还是空切片，对其调用内置函数`append`、`len`和`cap`的效果都是一样的。

### 切片的长度和容量

切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。

### 基于数组定义切片

由于切片的底层就是一个数组，所以我们可以基于数组定义切片。

```
func main() {
	// 基于数组定义切片
	a := [5]int{55, 56, 57, 58, 59}
	/* 打印切片从索引1(包含) 到索引4(不包含)*/
	b := a[1:4]                     //基于数组a创建切片，包括元素a[1],a[2],a[3]
	fmt.Println(b)                  //[56 57 58]
	fmt.Printf("type of b:%T\n", b) //type of b:[]int
}
```

还支持如下方式：

```
c := a[1:] //[56 57 58 59] /* 打印切片从索引1(包含)到最后一个索引结束*/
d := a[:4] //[55 56 57 58] /*索引是从0(包含)开始到索引4(不包含)*/
e := a[:]  //[55 56 57 58 59] /*所有索引的结果*/
```

### 切片再切片

除了基于数组得到切片，我们还可以通过切片来得到切片。

```
func main() {
	//切片再切片
	a := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	fmt.Printf("a:%v type:%T len:%d  cap:%d\n", a, a, len(a), cap(a))
	/* 打印切片从索引1(包含) 到索引3(不包含)*/
	b := a[1:3]
	fmt.Printf("b:%v type:%T len:%d  cap:%d\n", b, b, len(b), cap(b))
	c := b[1:5]
	fmt.Printf("c:%v type:%T len:%d  cap:%d\n", c, c, len(c), cap(c))
}
```

输出：

```
a:[北京 上海 广州 深圳 成都 重庆] type:[6]string len:6  cap:6
b:[上海 广州] type:[]string len:2  cap:5
c:[广州 深圳 成都 重庆] type:[]string len:4  cap:4
```

注意： 对切片进行再切片时，索引不能超过原数组的长度，否则会出现索引越界的错误。



### 使用make()函数构造切片

我们上面都是基于数组来创建的切片，如果需要动态的创建一个切片，我们就需要使用内置的make()函数，格式如下：

> make([]T, size, cap)

其中：

​	T: 切片的元素类型
​	size: 切片中元素的数量
​	cap: 切片的容量
举个例子：

```
func main() {
	a := make([]int, 2, 10)
	fmt.Println(a)      //[0 0]
	fmt.Println(len(a)) //2
	fmt.Println(cap(a)) //10
}
```

上面代码中a的内部存储空间已经分配了10个，但实际上只用了2个。 容量并不会影响当前元素的个数，所以len(a)返回2，cap(a)则返回该切片的容量。

切片的本质
切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。

举个例子，现在有一个数组a := [8]int{0, 1, 2, 3, 4, 5, 6, 7}，切片s1 := a[:5]，相应示意图如下。
**![img](https://www.liwenzhou.com/images/Go/slice/slice_01.png)**

slice_01切片s2 := a[3:6]，相应示意图如下：slice_02
**![img](https://www.liwenzhou.com/images/Go/slice/slice_02.png)**

### 判断切片是否为空

​		要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。

切片不能直接比较
		切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。 切片唯一合法的比较操作是和nil比较。 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。但是我们不能说一个长度和容量都是0的切片一定是nil，例如下面的示例：

```
var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
```

所以要判断一个切片是否是空的，要是用len(s) == 0来判断，不应该使用s == nil来判断。

### 切片的赋值拷贝

下面的代码中演示了拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容，这点需要特别注意。

```
func main() {
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	s2[0] = 100
	fmt.Println(s1) //[100 0 0]
	fmt.Println(s2) //[100 0 0]
}
```

### 切片遍历

切片的遍历方式和数组是一致的，支持索引遍历和for range遍历。

```
func main() {
	s := []int{1, 3, 5}

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	for index, value := range s {
		fmt.Println(index, value)
	}
}
```

### append()方法为切片添加元素

Go语言的内建函数append`()`可以为切片动态添加元素。 append() 的第二个参数不能直接使用 slice，需使用 … 操作符，将一个切片追加到另一个切片上：append(s1,s2…)。或者直接跟上元素，形如：append(s1,1,2,3)。。

```
func main(){
	var s []int
	s = append(s, 1)        // [1]
	s = append(s, 2, 3, 4)  // [1 2 3 4]
	s2 := []int{5, 6, 7}  
	s = append(s, s2...)    // [1 2 3 4 5 6 7]
}
```

注意：通过var声明的零值切片可以在append()函数直接使用，无需初始化。

```
var s []int
s = append(s, 1, 2, 3)
```

没有必要像下面的代码一样初始化一个切片再传入append()函数使用，

```
s := []int{}  // 没有必要初始化
s = append(s, 1, 2, 3)

var s = make([]int)  // 没有必要初始化
s = append(s, 1, 2, 3)
```

每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。“扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值。

举个例子：

```go
func main() {
	//append()添加元素和切片扩容
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
}
```

输出：

```
[0]  len:1  cap:1  ptr:0xc0000a8000
[0 1]  len:2  cap:2  ptr:0xc0000a8040
[0 1 2]  len:3  cap:4  ptr:0xc0000b2020
[0 1 2 3]  len:4  cap:4  ptr:0xc0000b2020
[0 1 2 3 4]  len:5  cap:8  ptr:0xc0000b6000
[0 1 2 3 4 5]  len:6  cap:8  ptr:0xc0000b6000
[0 1 2 3 4 5 6]  len:7  cap:8  ptr:0xc0000b6000
[0 1 2 3 4 5 6 7]  len:8  cap:8  ptr:0xc0000b6000
[0 1 2 3 4 5 6 7 8]  len:9  cap:16  ptr:0xc0000b8000
[0 1 2 3 4 5 6 7 8 9]  len:10  cap:16  ptr:0xc0000b8000
```

从上面的结果可以看出：
append()函数将元素追加到切片的最后并返回该切片。
切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍。
append()函数还支持一次性追加多个元素。 例如：

```
var citySlice []string
// 追加一个元素
citySlice = append(citySlice, "北京")
// 追加多个元素
citySlice = append(citySlice, "上海", "广州", "深圳")
// 追加切片
a := []string{"成都", "重庆"}
citySlice = append(citySlice, a...)
fmt.Println(citySlice) //[北京 上海 广州 深圳 成都 重庆]
```

下面的代码是有问题的，请说明原因。

```
func main()  {
  var p []interface{}

  b := []string{"hello","world"}

  o := append(p,b...)
  fmt.Println(o)
}
```
**原因:** 
​ string或int类型的切片,不能append到interface类型的切片, 只允许interface类型的切片,append到interface类型的切片。即 同类型的切片才能append到同类型的切片中。



### 切片的扩容策略

可以通过查看$GOROOT/src/runtime/slice.go源码，其中扩容相关代码如下：

```go
newcap := old.cap
doublecap := newcap + newcap
if cap > doublecap {
	newcap = cap
} else {
	if old.len < 1024 {
		newcap = doublecap
	} else {
		// Check 0 < newcap to detect overflow
		// and prevent an infinite loop.
		for 0 < newcap && newcap < cap {
			newcap += newcap / 4
		}
		// Set newcap to the requested cap when
		// the newcap calculation overflowed.
		if newcap <= 0 {
			newcap = cap
		}
	}
}
```

从上面的代码可以看出以下内容：

首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如int和string类型的处理方式就不一样。



### 使用copy()函数复制切片

首先我们来看一个问题：

```
func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(b) //[1 2 3 4 5]
	b[0] = 1000
	fmt.Println(a) //[1000 2 3 4 5]
	fmt.Println(b) //[1000 2 3 4 5]
}
```

**结论:**

​		**由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化。**



Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中，copy()函数的使用格式如下：

> copy(destSlice, srcSlice []T)

其中：

​	srcSlice: 数据源切片
​	destSlice: 目标切片
举个例子：

```
func main() {
	// copy()复制切片
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 5)
	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1 2 3 4 5]
	c[0] = 1000
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1000 2 3 4 5]
}
```


### 从切片中删除元素

Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素。 代码如下：

```go
func main() {
	// 从切片中删除元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a) //[30 31 33 34 35 36 37]
}
```

**总结：**

​		**要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)**

### 练习题

1.请写出下面代码的输出结果。

```
func main() {
	var a = make([]string, 5, 10)
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}
```

2.请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序（附加题，自行查资料解答）。



# map

Go语言中提供的映射关系容器为map，其内部使用散列表（hash）实现。

### map

> map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。

### map定义

Go语言中 map的定义语法如下：

> map[KeyType]ValueType

其中，

- KeyType:表示键的类型。
- ValueType:表示键对应的值的类型。


map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：

> make(map[KeyType]ValueType, [cap])

其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。

### map基本使用

map中的数据都是成对出现的，map的基本使用示例代码如下：

```
func main() {
    var scoreMap map[string]int
    fmt.Println(scoreMap == nil)  //true 还没初始化
	scoreMap = make(map[string]int, 8) //make初始化才能访问
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Println(scoreMap["小红"]) //如果key值不存在返回的是value对应类型的零值
	fmt.Printf("type of a:%T\n", scoreMap)
		fmt.Println(scoreMap["tom"]) //如果不存在key值，则返回对应类型的0值
}
```

输出：

```
true
map[小明:100 张三:90]
100
0
type of a:map[string]int
0
```

map也支持在声明的时候填充元素，例如：

```
func main() {
	userInfo := map[string]string{
		"username": "沙河小王子",
		"password": "123456",
	}
	fmt.Println(userInfo) //
}
```

### 判断某个键是否存在

Go语言中有个判断map中键是否存在的特殊写法，如果map函数
返回的是bool类型，我们通常用ok接受它。格式如下:

> value, ok := map[key]

举个例子：

```
func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}
```

### map的遍历

Go语言中使用for range遍历map。

```
func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}
```

只遍历key，写法：

```
func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	for k := range scoreMap {
		fmt.Println(k)
	}
}
```

注意： 遍历map时的元素顺序与添加键值对的顺序无关。

### 使用delete()函数删除键值对

使用delete()内建函数从map中删除一组键值对，delete()函数的格式如下：

> delete(map, key)

其中，

- map:表示要删除键值对的map
- key:表示要删除的键值对的键
  示例代码如下：

```
func main(){
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	delete(scoreMap, "小红") //删除的key不存在什么都不干
	delete(scoreMap, "小明")//将小明:100从map中删除
	for k,v := range scoreMap{
		fmt.Println(k, v)
	}
}
```

### 按照指定顺序遍历map

```
func main() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
```

### 元素为map类型的切片

下面的代码演示了切片中的元素为map类型时的操作：

```
func main() {
	//元素类型为map的切片
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}
```

### 值为切片类型的map

下面的代码演示了map中值为切片类型的操作：

```go
func main() {
  //构造值为slice类型的map
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
  fmt.Println(sliceMap)
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
```

### 练习题

写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
观察下面代码，写出最终的打印结果。

```go
func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
```





# 函数

函数是组织好的、可重复使用的、用于执行指定任务的代码块。本文介绍了Go语言中函数的相关内容。

## 函数

Go语言中支持函数、匿名函数和闭包，并且函数在Go语言中属于“一等公民”。

### 函数定义

Go语言中定义函数使用func关键字，具体格式如下：

> func 函数名(参数)(返回值){
> 函数体
> }

其中：

- 函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名（包的概念详见后文）。
- 参数：参数由参数变量和参数变量的类型组成，多个参数之间使用,分隔。
- 返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用()包裹，并用,分隔。
- 函数体：实现指定功能的代码块。

我们先来定义一个求两个数之和的函数：

```
func intSum(x int, y int) int {
	return x + y
}
```

函数的参数和返回值都是可选的，例如我们可以实现一个既没有参数也没有返回值的函数：

```
func sayHello() {
	fmt.Println("Hello tom")
}
```

### 函数的调用

定义了函数之后，我们可以通过函数名()的方式调用函数。 例如我们调用上面定义的两个函数，代码如下：

```
func main() {
	sayHello()
	ret := intSum(10, 20)
	fmt.Println(ret)
}
```

注意，调用有返回值的函数时，可以不接收其返回值。

### 参数

#### 类型简写

函数的参数中如果相邻变量的类型相同，则可以省略类型，例如：

```
func intSum(x, y int) int {
	return x + y
}
```

上面的代码中，intSum函数有两个参数，这两个参数的类型均为int，因此可以省略x的类型，因为y后面有类型说明，x参数也是该类型。

#### 可变参数

可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识。

注意：可变参数通常要作为函数的最后一个参数。

举个例子：

```
func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
```

调用上面的函数：

```
ret1 := intSum2()
ret2 := intSum2(10)
ret3 := intSum2(10, 20)
ret4 := intSum2(10, 20, 30)
fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 60
```

固定参数搭配可变参数使用时，可变参数要放在固定参数的后面，示例代码如下：

```
func intSum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}
```

调用上述函数：

```
ret5 := intSum3(100)
ret6 := intSum3(100, 10)
ret7 := intSum3(100, 10, 20)
ret8 := intSum3(100, 10, 20, 30)
fmt.Println(ret5, ret6, ret7, ret8) //100 110 130 160
```

本质上，函数的可变参数是通过切片来实现的。

### 返回值

Go语言中通过return关键字向外输出返回值。

#### 多返回值

Go语言中函数支持多返回值，函数如果有多个返回值时必须用()将所有返回值包裹起来。

举个例子：

```
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}
```

#### 命名返回值

函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回。

例如：

```
//sum 相当于函数内的局部变量，被初始化为零
func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return  //return sum 的简写模式
 //sum := x + y //如果是 sum := x + y，则相当于新声明一个 sum 变量命名返回变量sum覆盖
 //return sum //最后需要豆式地调用 return sum
}
```

#### 返回值补充

当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。

```
func someFunc(x string) []int {
	if x == "" {
		return nil // 没必要返回[]int{}
	}
	...
}
```



# 函数进阶

## 变量作用域

### 全局变量

全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。 在函数中可以访问到全局变量。

```
package main

import "fmt"

//定义全局变量num
var num int64 = 10

func testGlobalVar() {
	fmt.Printf("num=%d\n", num) //函数中可以访问全局变量num
}
func main() {
	testGlobalVar() //num=10
}
```

### 局部变量

局部变量又分为两种： 函数内定义的变量无法在该函数外使用，例如下面的示例代码main函数中无法使用testLocalVar函数中定义的变量x：

```
func testLocalVar() {
	//定义一个函数局部变量x,仅在该函数内生效
	var x int64 = 100
	fmt.Printf("x=%d\n", x)
}

func main() {
	testLocalVar()
	fmt.Println(x) // 此时无法使用变量x
}
```

函数查找变量的顺序：先在函数内部查找，找不到就去全局变量中查找。

```
package main

import "fmt"

//定义全局变量num
var num int64 = 10

func testNum() {
	num := 100
	fmt.Printf("num=%d\n", num) // 函数中优先使用局部变量
}
func main() {
	testNum() // num=100
}
```

接下来我们来看一下语句块定义的变量，通常我们会在if条件判断、for循环、switch语句上使用这种定义变量的方式。

```
func testLocalVar2(x, y int) {
	fmt.Println(x, y) //函数的参数也是只在本函数中生效
	if x > 0 {
		z := 100 //变量z只在if语句块生效
		fmt.Println(z)
	}
	//fmt.Println(z)//此处无法使用变量z
}
```

还有我们之前讲过的for循环语句中定义的变量，也是只在for语句块中生效：

```
func testLocalVar3() {
	for i := 0; i < 10; i++ {
		fmt.Println(i) //变量i只在当前for语句块中生效
	}
	//fmt.Println(i) //此处无法使用变量i
}
```

### 函数类型与变量

#### 定义函数类型

我们可以使用type关键字来定义一个函数类型，具体格式如下：

> type calculation func(int, int) int

上面语句定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。

简单来说，凡是满足这个条件的函数都是calculation类型的函数，例如下面的add和sub是calculation类型。

```
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
```

add和sub都能赋值给calculation类型的变量。

```
var c calculation
c = add
```

### 函数类型变量

我们可以声明函数类型的变量并且为该变量赋值：

```
func main() {
	var c calculation               // 声明一个calculation类型的变量c
	c = add                         // 把add赋值给c
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	fmt.Println(c(1, 2))            // 像调用add一样调用c

	f := add                        // 将函数add赋值给变量f1
	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	fmt.Println(f(10, 20))          // 像调用add一样调用f
}
```

### 高阶函数

高阶函数分为函数作为参数和函数作为返回值两部分。

#### 函数作为参数

函数可以作为参数：

```go
func add(x, y int) int {
	return x + y
}
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func main() {
	ret2 := calc(10, 20, add)
	fmt.Println(ret2) //30
}
```

#### 函数作为返回值

函数也可以作为返回值：

```go
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}
```

### 匿名函数和闭包

#### 匿名函数

函数当然还可以作为返回值，函数内部没办法声明带名字的函数，只能定义匿名函数。匿名函数就是没有函数名的函数，匿名函数的定义格式如下：

> func(参数)(返回值){
> 函数体
> }

匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数:

```go
func main() {
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
```

匿名函数多用于实现回调函数和闭包。

### 闭包

闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境。 首先我们来看一个例子：

```go
/*
go语言支持函数式编程:
	支持讲一个函数作为另一个函数的参数
	也支持将一个函数作为另一个函数的返回值
	
	闭包:
	  一个外层函数中,有内层函数,该内层函数中,会操作外层函数的局部变量(外层函数中的个参数,或者外岑函数中直接定义的变量),并且该外层函数的返回值就是这个内层函数.这个内层函数和外层函数的局部变量,统称为闭包结构.
		局部变量的生命周期会发生变化,正常的局部变量随着函数调用而创建,随着函数的结束而销毁
		但是闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁,因为内层函数还要继续调用
*/
//func(int) int是外层函数的返回值数据类型等于内层函数,
func adder() func(int) int {// 外层函数
	//定义一个局部变量
	x := 1
	//定义一个匿名函数,给变量自增并返回
	fun := func(y int) int {//内层函数
		x += y
		return x
	}
	return fun
	//第二种方式: f = fun
	//return func(y int) int {
	//	x += y
	//	return x
	//}
}

func main() {
	f := adder()  //f = fun
	fmt.Printf("%T\n", f) //func(int) int
	fmt.Println(f)        //0x109ee60
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60

	f1 := adder()
	fmt.Println(f1(40)) //40
	fmt.Println(f1(50)) //90
}
```

闭包实例说明：

变量f调用了adder这样一个函数，指向的是fun内层函数,它的值是内层函数的函数体地址,所以f()传参数即可调用.如下：

```go
	//f = fun
	fun := func(y int) int {
		x += y
		return x
	}
	return fun
```

这个函数f会引用外部作用域中的x变量。

##### 总结：

变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。 在f的生命周期内，变量x也一直有效。 

闭包进阶示例1：

```go
//函数作为返回值
func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	var f = adder2(10)
	fmt.Println(f(10)) //20
	fmt.Println(f(20)) //40
	fmt.Println(f(30)) //70

	f1 := adder2(20)
	fmt.Println(f1(40)) //60
	fmt.Println(f1(50)) //110
}
```

闭包进阶示例2：

```go
//函数作为返回值
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
 	fmt.Println(jpgFunc("test1.jpg")) //原样输出test1.jpg
	fmt.Println(txtFunc("test")) //test.txt
}
```

返回的匿名函数和 makeSuffix (suffix string) 的 suffix 变量组合成一个闭包,因为 返回的函数引用 到 suffix 这个变量

闭包进阶示例3：

```go
// 函数名calc 参数base int 返回值func(int) int, func(int) int 是2个匿名返回值，接受一个int类型的参数和返回值
//base是add和sub共同作用域中的变量
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}
```

上述实例分析：

首先变量f1是一个函数，它接受一个int类型的参数和返回值的函数，如下：

```
func(i int) int {
	base += i
	return base
}
```

首先变量f2也是一个函数，它也接受一个int类型的参数和返回值的函数，如下：

```
func(i int) int {
	 base -= i
   	return base
}
```

这2个函数会引用外部作用域中的base变量。

闭包进阶示例4

```go
//如何将name为string类型的参数传到fb中，参数类型不匹配，可以把fa包装成一个
func fa(name string) {
	fmt.Println("hello", name)
}
//带参数的函数
func fb(f func()) {
	f()
}
//有参数又有返回值的函数
func fc(f func(string), name string) func() {
	return func(){
		f(name)
	}
}

func main() {
	fd :=fc(fa, "tom")
	fb(fd)     //hello tom
	fmt.Println(fd)
}
```

实例分析：

变量fd是一个函数，它接受变量也是一个函数，且返回值参数也是一个函数。

```go
func(){
		f(name)
	}
}
```

这个函数会引用外部作用域中的name变量

闭包其实并不复杂，只要牢记闭包=函数+引用环境。

### defer语句

Go语言中的defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。

举个例子：

```go
func main() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}
```

输出结果：

```
start
end
3
2
1
```

由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。

### defer执行时机

在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。具体如下图所示：defer执行时机

### defer经典案例

阅读下面的代码，写出最后的打印结果。

```go
func f1() int {
	x := 5
	defer func() {
		x++     //修改的是x不是返回值
	}()
	return x    //1.返回值赋值 2.defer 3.执行RET指令
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5    //返回值=x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ 	//修改的是x
	}()
	return x 	// 返回值 = y = x 2. defer修改的是x 3.真正的返回
}
func f4() (x int) {
	defer func(x int) {
		x++		//改变的是函数中x的副本
		return x
	}(x)
	return 5	//返回值=x=5
}
func f5() (x int) {
	defer func(x int) int {
		x++		//改变的是函数中x的副本
	}(x)
	return 5	//返回值=x=5
}

//传一个x的指针到匿名函数中
func f6() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)  //内存地址
	return 5	//返回值=x=5 2. defer x=6 3 RET指令
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
	fmt.Println(f6())
}
```

defer面试题

```
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
```

问，上面代码的输出结果是？（提示：defer注册要延迟执行的函数时该函数所有的参数都需要确定其值）

### 内置函数介绍

| 内置函数       | 介绍                                                         |
| -------------- | ------------------------------------------------------------ |
| close          | 主要用来关闭channel                                          |
| len            | 用来求长度，比如string、array、slice、map、channel           |
| new            | 用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针 |
| make           | 用来分配内存，主要用来分配引用类型，比如chan、map、slice     |
| append         | 用来追加元素到数组、slice中                                  |
| panic和recover | 用来做错误处理                                               |

#### panic/recover

Go语言中目前（Go1.12）是没有异常机制，但是使用panic/recover模式来处理错误。 panic可以在任何地方引发，但recover只有在defer调用的函数中有效。 首先来看一个例子：

```go
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	funcA()
	funcB()
	funcC()
}
```

输出：

```go
func A
panic: panic in B

goroutine 1 [running]:
main.funcB(...)
        .../code/func/main.go:12
main.main()
        .../code/func/main.go:20 +0x98
```

程序运行期间funcB中引发了panic导致程序崩溃，异常退出了。这个时候我们就可以通过recover将程序恢复回来，继续往后执行。

```
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	funcA()
	funcB()
	funcC()
}
```

注意：
		recover()必须搭配defer使用。
		defer一定要在可能引发panic的语句之前定义。

### 练习题

你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数

```go
 var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {
//拿到每个人的名字
for _, name := range users {
	//拿到一个人名根据分金币规则去分金币
	for _, c := range name {
		switch c {
		case 'e', 'E':
			//满足调剂分得1枚金币
			distribution[name]++
			coins--
		case 'i', 'I':
			//满足调剂分得2枚金币
			distribution[name] +=2
			coins -=2
		case 'o', 'O':
			//满足调剂分得3枚金币
			distribution[name] +=3
			coins -= 3
		case 'u', 'U':
			//满足调剂分得4枚金币
			distribution[name] +=4
			coins -=4
		}
	}
}
left = coins
return
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	//打印每个人分得的金币数目
	for k, v := range distribution {
		fmt.Printf("%s:%d\n", k, v)
	}
}
```



# 指针

任何程序数据载入内存后，在内存都有他们的地址，这就是指针。而为了保存一个数据在内存中的地址，我们就需要指针变量。

比如，“永远不要高估自己”这句话是我的座右铭，我想把它写入程序中，程序一启动这句话是要加载到内存（假设内存地址0x123456），我在程序中把这段话赋值给变量`A`，把内存地址赋值给变量`B`。这时候变量`B`就是一个指针变量。通过变量`A`和变量`B`都能找到我的座右铭。

Go语言中的指针不能进行偏移和运算，因此Go语言中的指针操作非常简单，我们只需要记住两个符号：`&`（取地址）和`*`（根据地址取值）。

## 指针地址和指针类型

每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go语言中使用`&`字符放在变量前面对变量进行“取地址”操作。 Go语言中的值类型（int、float、bool、string、array、struct）都有对应的指针类型，如：`*int`、`*int64`、`*string`等。

取变量指针的语法如下：

```go
ptr := &v    // v的类型为T
```

其中：

- v:代表被取地址的变量，类型为`T`
- ptr:用于接收地址的变量，ptr的类型就为`*T`，称做T的指针类型。*代表指针。

举个例子：

```go
func main() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b)                    // 0xc00000e018
}
```

我们来看一下`b := &a`的图示：![取变量地址图示](https://www.liwenzhou.com/images/Go/pointer/ptr.png)

## 指针取值

在对普通变量使用&操作符取地址后会获得这个变量的指针，然后可以对指针使用*操作，也就是指针取值，代码如下。

```go
func main() {
	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}
```

输出如下：

```go
type of b:*int
type of c:int
value of c:10
```

**总结：** 取地址操作符`&`和取值操作符`*`是一对互补操作符，`&`取出地址，`*`根据地址取出地址指向的值。

变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：

- 指针就是地址(&a的内存地址0xc00001a078)，地址就是内存单元的编号。
- 对变量进行取地址（&）操作，可以获得这个变量的指针变量。b = &a
- 指针变量就是存放内存地址的变量。
- 对指针变量b进行取值（*）操作，可以获得指针变量指向的原变量的值。 *b =10

**指针传值示例：**

```go
func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func main() {
	a := 10
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a) // 100
}
```

## new和make

我们先来看一个例子：

```go
func main() {
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["沙河娜扎"] = 100
	fmt.Println(b)
}
```

执行上面的代码会引发panic，为什么呢？ 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。要分配内存，就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存。

### new

new是一个内置的函数，它的函数签名如下：

```go
func new(Type) *Type
```

其中，

- Type表示类型，new函数只接受一个参数，这个参数是一个类型
- *Type表示类型指针，new函数返回一个指向该类型内存地址的指针。

new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。举个例子：

```go
func main() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false
}	
```

本节开始的示例代码中`var a *int`只是声明了一个指针变量a但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值。应该按照如下方式使用内置的new函数对a进行初始化之后就可以正常对其赋值了：

```go
func main() {
	var a *int
	a = new(int)
	*a = 10
	fmt.Println(*a)
}
```

### make

make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。make函数的函数签名如下：

```go
func make(t Type, size ...IntegerType) Type
```

make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作。这个我们在上一章中都有说明，关于channel我们会在后续的章节详细说明。

本节开始的示例中`var b map[string]int`只是声明变量b是一个map类型的变量，需要像下面的示例代码一样使用make函数进行初始化操作之后，才能对其进行键值对赋值：

```go
func main() {
	var b map[string]int
	b = make(map[string]int, 10)
	b["沙河娜扎"] = 100
	fmt.Println(b)
}
```

### new与make的区别

1. 二者都是用来做内存分配的。
2. make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
3. 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。



nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量。



# 结构体

Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。

## 类型别名和自定义类型

### 自定义类型

在Go语言中有一些基本的数据类型，如string、整型、浮点型、布尔等数据类型， Go语言中可以使用type关键字来定义自定义类型。

自定义类型是定义了一个全新的类型。我们可以基于内置的基本类型定义，也可以通过struct定义。例如：

//将MyInt定义为int类型

> type MyInt int

通过type关键字的定义，MyInt就是一种新的类型，它具有int的特性。

### 类型别名

类型别名是Go1.9版本添加的新功能。

类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。就像一个孩子小时候有小名、乳名，上学后用学名，英语老师又会给他起英文名，但这些名字都指的是他本人。

> type TypeAlias = Type

我们之前见过的rune和byte就是类型别名，他们的定义如下：

```
type byte = uint8
type rune = int32
```

### 类型定义和类型别名的区别

类型别名与类型定义表面上看只有一个等号的差异，我们通过下面的这段代码来理解它们之间的区别。

```go
//类型定义
type NewInt int

//类型别名
type MyInt = int

func main() {
	var a NewInt
	var b MyInt
	
	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int
}
```

结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型。

## 结构体

Go语言中的基础数据类型可以表示一些事物的基本属性，但是当我们想表达一个事物的全部或部分属性时，这时候再用单一的基本数据类型明显就无法满足需求了，Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型叫结构体，英文名称struct。 也就是我们可以通过struct来定义自己的类型了。

Go语言中通过struct来实现面向对象。

### 结构体的定义

使用type和struct关键字来定义结构体，具体代码格式如下：

```
type 类型名 struct {
    字段名 字段类型
    字段名 字段类型
    …
}
```

其中：

- 类型名：标识自定义结构体的名称，在同一个包内不能重复。
- 字段名：表示结构体字段名。结构体中的字段名必须唯一。
- 字段类型：表示结构体字段的具体类型。
  举个例子，我们定义一个Person（人）结构体，代码如下：

```
type person struct {
	name string
	city string
	age  int8
}
```

同样类型的字段也可以写在一行，

```
type person1 struct {
	name, city string
	age        int8
}
```

这样我们就拥有了一个person的自定义类型，它有name、city、age三个字段，分别表示姓名、城市和年龄。这样我们使用这个person结构体就能够很方便的在程序中表示和存储人信息了。

语言内置的基础数据类型是用来描述一个值的，而结构体是用来描述一组值的。比如一个人有名字、年龄和居住城市等，本质上是一种聚合型的数据类型

## 结构体实例化

只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。

结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。

> var 结构体实例 结构体类型

### 基本实例化

举个例子：

```
type person struct {
	name string
	city string
	age  int8
	hobboy []string
}

func main() {
    //声明一个person类型的变量p
	var p1 person
	//通过字段赋值 
	p1.name = "娜扎"
	p1.city = "北京"
	p1.age = 18
	p.hobboy = []string{"basketball", "football", "reader"}
	//访问变量p的字段
	fmt.Printf("%T\n", p1)       //main.person
	fmt.Println(p1.name)         //娜扎
	fmt.Printf("p1=%v\n", p1)  //p1={娜扎 北京 18 [basketball football reader]}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"娜扎", city:"北京", age:18，hobboy:[]string{"basketball", "football", "reader"}}
}
```

结构体指针访问结构体字段仍然使用“.”点操作符,例如p1.name和p1.age等。

### 匿名结构体

在定义一些临时数据结构等场景下还可以使用匿名结构体。

```go
package main
     
import (
    "fmt"
)
     
func main() {
    var user struct{
   	 	Name string
    	Age int
    }
    user.Name = "小王子"
    user.Age = 18
    fmt.Printf("%#v\n", user)
}
```

### 创建指针类型结构体

我们还可以通过使用new关键字对结构体进行实例化，得到的是结构体的地址。 格式如下：

```go
var p2 = new(person)
fmt.Printf("%T\n", p2)     //*main.person
fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}
```

从打印的结果中我们可以看出p2是一个结构体指针。

需要注意的是在Go语言中支持对结构体指针直接使用.来访问结构体的成员。

```go
var p2 = new(person)
p2.name = "小王子"
p2.age = 28
p2.city = "上海"
fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"小王子", city:"上海", age:28}
```

### 取结构体的地址实例化

使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。

```go
p3 := &person{}
fmt.Printf("%T\n", p3)     //*main.person
fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
p3.name = "七米"
p3.age = 30
p3.city = "成都"
fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"七米", city:"成都", age:30}
```

p3.name = "七米"其实在底层是(*p3).name = "七米"，这是Go语言帮我们实现的语法糖。

### 结构体初始化

没有初始化的结构体，其成员变量都是对应其类型的零值。

```go
type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p4 person
	fmt.Printf("p4=%#v\n", p4) //p4=main.person{name:"", city:"", age:0}
}
```

### 使用键值对初始化

使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。

```go
p5 := person{
	name: "小王子",
	city: "北京",
	age:  18,
}
fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"小王子", city:"北京", age:18}
```

也可以对结构体指针进行键值对初始化，例如：

```go
p6 := &person{
	name: "小王子",
	city: "北京",
	age:  18,
}
fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"小王子", city:"北京", age:18}
```

当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。

```
p7 := &person{
	city: "北京",
}
fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}
```

### 使用值的列表初始化

初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值：

```
p8 := &person{
	"沙河娜扎",
	"北京",
	28,
}
fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"沙河娜扎", city:"北京", age:28}
```

使用这种格式初始化时，需要注意：

1. 必须初始化结构体的所有字段。
2. 初始值的填充顺序必须与字段在结构体中的声明顺序一致。
3. 该方式不能和键值初始化方式混用。

### 结构体内存布局

结构体占用一块连续的内存。

```go
//同类型变量的结构体内存地址是连续的
type x struct {
	a int8 //8bit => 1byte
	b int8
	c string
	d int8
}
func main() {
	m := x {
		a: int8(10),
		b: int8(20),
		c: "嘿嘿",
		d: int8(30),
	}
	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))
	fmt.Printf("%p\n", &(m.d))
}
```

输出：

```go
0xc00000c060
0xc00000c061
0xc00000c068
0xc00000c078
```

【进阶知识点】关于Go语言中的内存对齐推荐阅读:在 Go 中恰到好处的内存对齐

### 空结构体

空结构体是不占用空间的。

```go
var v struct{}
fmt.Println(unsafe.Sizeof(v))  // 0
```

### 面试题

请问下面代码的执行结果是什么？

```go
type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
```

### 构造函数

Go语言的结构体没有构造函数，我们可以自己实现。 例如，下方的代码就实现了一个person的构造函数。 因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。

```
//当结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
type person struct {
	name string
	city string
	age int
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
```

调用构造函数

```
p9 := newPerson("张三", "沙河", 90)
fmt.Printf("%#v\n", p9) //&main.person{name:"张三", city:"沙河", age:90}
```

方法和接收者
Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self。

## 方法

一般的，我们把普通定义的函数叫做函数，定义给结构体的函数叫做方法。

### 方法的定义格式如下：

```
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
```

其中，

- 接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，而不是self、this之类的命名。例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
- 接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
- 方法名、参数列表、返回参数：具体格式与函数定义相同。
  举个例子：

```go
//Person 结构体
type Person struct {
	name string
	age  int8
}

//NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream Person做梦的方法
// 接受者表示的是调用该方法好的具体类型变量，多用类型字母大小写表示
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func main() {
	p1 := NewPerson("小王子", 25)
	p1.Dream()
}
```

方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。

### 指针类型的接收者

指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。这种方式就十分接近于其他语言中面向对象中的this或者self。 例如我们为Person添加一个SetAge方法，来修改实例变量的年龄。

```
// SetAge 设置p的年龄
// 使用指针接收者，传内存地址进去
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}
```

调用该方法：

```
func main() {
	p1 := NewPerson("小王子", 25)
	fmt.Println(p1.age) // 25
	p1.SetAge(30)
	fmt.Println(p1.age) // 30
}
```

### 值类型的接收者

当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。

```
// SetAge2 设置p的年龄
// 使用值接收者，传拷贝进去
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

func main() {
	p1 := NewPerson("小王子", 25)
	p1.Dream()
	fmt.Println(p1.age) // 25
	p1.SetAge2(30) // (*p1).SetAge2(30)
	fmt.Println(p1.age) // 25
}
```

#### 什么时候应该使用指针类型接收者

需要修改接收者中的值
接收者是拷贝代价比较大的大对象
保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

### 任意类型添加方法

在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。 举个例子，我们基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。

```
//MyInt 将int定义为自定义MyInt类型
type MyInt int

//SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}
func main() {
	var m1 MyInt
	m1.SayHello() //Hello, 我是一个int。
	m1 = 100
	fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt
}
```

注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。

### 结构体的匿名字段

结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。

```
//Person 结构体Person类型
type Person struct {
	string
	int
}

func main() {
	p1 := Person{
		"小王子",
		18,
	}
	fmt.Printf("%#v\n", p1)        //main.Person{string:"北京", int:18}
	fmt.Println(p1.string, p1.int) //北京 18
}
```

匿名字段默认采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。

### 嵌套结构体

一个结构体中可以嵌套包含另一个结构体或结构体指针。

```
//Address 地址结构体
type Address struct {
	Province string
	City     string
}

//User 用户结构体
type User struct {
	Name    string
	Gender  string
	Addr    Address
}

func main() {
	user1 := User{
		Name:   "小王子",
		Gender: "男",
		Addr: Address{
			Province: "山东",
			City:     "威海",
		},
	}
	fmt.Printf("user1=%#v\n", user1)//user1=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}
	fmt.Println(user1.Name, user1.Addr.City)  //小王子 威海
}
```

### 匿名嵌套结构体

先在自己结构体找这个字段，找不到就去匿名嵌套结构体中查找该字段

```
//Address 地址结构体
type Address struct {
	Province string
	City     string
}

//User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address //匿名结构体
}

func main() {
	var user2 User
	user2.Name = "小王子"
	user2.Gender = "男"
	user2.Address.Province = "山东"    //通过匿名结构体.字段名访问
	user2.City = "威海"                //直接访问匿名结构体的字段名
	fmt.Printf("user2=%#v\n", user2) //user2=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}
	fmt.Println(user2.Name, user2.City)  //小王子 威海
}
```

当访问结构体成员时会先在结构体中查找该字段，找不到再去匿名结构体中查找。

### 嵌套结构体的字段名冲突

嵌套结构体内部可能存在相同的字段名。这个时候为了避免歧义需要指定具体的内嵌结构体的字段。

```
//Address 地址结构体
type Address struct {
	Province   string
	City       string
	CreateTime string
}

//Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

//User 用户结构体
type User struct {
	Name   string
	Gender string
	Address
	Email
}

func main() {
	var user3 User
	user3.Name = "沙河娜扎"
	user3.Gender = "男"
	// user3.CreateTime = "2019" //ambiguous selector user3.CreateTime
	user3.Address.CreateTime = "2000" //指定Address结构体中的CreateTime
	user3.Email.CreateTime = "2000"   //指定Email结构体中的CreateTime
    fmt.Printf("user3.Address.CreateTime")  //2000
    fmt.Printf("user3.Email.CreateTime")    //2000
}
```

### 结构体的“继承”

Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。

```
//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}
```

### 结构体字段的可见性

结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

### 结构体与JSON序列化

JSON(JavaScript Object Notation) 是一种轻量级的数据交换格式。易于人阅读和编写。同时也易于机器解析和生成。JSON键值对是用来保存JS对象的一种方式，键/值对组合中的键名写在前面并用双引号""包裹，使用冒号:分隔，然后紧接着值；多个键值之间使用英文,分隔。

```go
//序列化：把go语言中的结构体变量--> json格式的字符串
//反序列化：json格式的字符串 --> go语言中能够识别的结构体变量
//Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string
	Students []*Student
}

func main() {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1) //传指针是为了能在json.unmarshal函数内部修改c1的值
}
```

### 结构体标签（Tag）

Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。 Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：

> `key1:"value1" key2:"value2"`

结构体标签由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。同一个结构体字段可以设置多个键值对标签（Tag），不同的键值对标签之间使用空格分隔。

> 注意事项： 为结构体编写Tag时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。例如不要在key和value之间添加空格。

例如我们为Student结构体的每个字段定义json序列化时使用的Tag：

```go
//Student 学生
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func main() {
	s1 := Student{
		ID:     1,
		Gender: "男",
		name:   "沙河娜扎",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"男"}
}
```

### 结构体和方法补充知识点

因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意。我们来看下面的例子：

```go
type Person struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person) SetDreams(dreams []string) {
	p.dreams = dreams
}

func main() {
	p1 := Person{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams)  // ?
}
```

正确的做法是在方法中使用传入的slice的拷贝进行结构体赋值。

```go
func (p *Person) SetDreams(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}
```

同样的问题也存在于返回值slice和map的情况，在实际编码过程中一定要注意这个问题。

### 练习题

-  使用“面向对象”的思维方式编写一个学生信息管理系统。
   1. 学生有id、姓名、年龄、分数等信息
   2. 程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能

主配置文件名 main.go

```go
package main

import (
	"fmt"
	"os"

)
//学生管理系统
var smr studentMgr  //声明一个全局的变量学生管理对象：smr

//有一个物件
// 1.它保存了一些数据 --- > 结构体的字段 比如：人的属性
// 2. 有三个功能  --- > 结构体的方法，比如；人能干啥
//菜单函数
func showMenu () {
	fmt.Println("----------welcom sms!-------------")
	fmt.Println(`
	1. 查看所有学生
	2. 添加学生
	3. 修改学生
	4. 删除学生
	5. 退出
	`)
}

func main() {
	//结构体初始化
	smr = studentMgr{ //修改全局的那个变量
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		//等待用户输入选项
		fmt.Print("输入序号：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你输入的是:", choice)
		//方法的实现
		switch choice {
		case 1:
			smr.showStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.editStudent()
		case 4:
			smr.deleteStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("go")
		}
	}	
}

```

文件名 student_mgr.go

```go
package main

import (
	"fmt"
	//"studentMgr"
)
	//学生管理系统
type student struct {
	id int64
	name string
}
//造一个学生的管理者
type studentMgr struct {
	allStudent map[int64]student
}
//查看学生
func (s studentMgr) showStudent() {
	//从s.allstudent这个map中把所有的学生逐个拿出来
	for _, stu := range s.allStudent {
		fmt.Printf("id：%d  name:%s\n", stu.id, stu.name)
	}
}
//增加学生
func (s studentMgr) addStudent() {
	//根据用户输入的内容创建一个新的学生
	//
	var (
		stuID int64
		stuName string
	)
	//获取用户输入
	fmt.Print("input id:")
	fmt.Scanln(&id)
	fmt.Print("please input name:")
	fmt.Scanln(&name)
	newStu := student {
		id: stuID,
		name: stuName,
	}
	//把新的学生放到s.allstudent这个map中
	s.allStudent[newStu.id] = newStu
	fmt.Println("add is ok")
}
//修改学生
func (s studentMgr) editStudent() {
	//获取用户输入的学号
	var (
		stuID int64
	)
	//获取用户输入
	fmt.Print("input id:")
	fmt.Scanln(&id)
	//展示该学号对应的学生信息，如果没有提示无此人
	stuObj, ok := s.allstudent[stuID] //用变量stuObj接受值,ok=ture表示有这个key，ok=false表示没有这key，如果没有可以返回的是value对应类型的0值，string对应类型的0值为空，空串
	if !ok {
		fmt.Println("The user does not exist")
		return
	}
	fmt.Printf("you need to modify student information blew follow: id:%d name:%s\n",stuObj.id, stuObj.name)
	//请输入修改后的学生名
	fmt.Print("please input student name")
	var newName string
	fmt.Scanln($newName)
	//更新学生的名
	s.allstudent[stuID] = stuObj // 更新map中的学生

//删除学生
func (s studentMgr) deleteStudent() {
	//输入要删除的用户
	var (
		stuID int64
	)
	fmt.Print("please input your delete id:")
	fmt.Scanln(&stuID)
	//去map中查看有没有这个学生，如果没有打印，无此人信息
	_, ok := s.allstudent[stuID]
	if !ok {
		fmt.Println("please input student name")
		return
	}
	//如果有就删除，如何从map中删除键值对
	delete(s.allstudent,stuID)
	fmt.Println("delete is ok")
}
```



# 接口

接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。



### 接口类型

在Go语言中接口（interface）是一种类型，一种抽象的类型。

interface是一组method的集合，是duck-type programming的一种体现。接口做的事情就像是定义一个协议（规则），只要一台机器有洗衣服和甩干的功能，我就称它为洗衣机。不关心属性（数据），只关心行为（方法）。

接口类型具体描述了一系列方法的集合，一个实现了这些方法的具体类型是这个接口类型的实例。

io.Writer类型是用得最广泛的接口之一，因为它提供了所有类型的写入bytes的抽象，包括文件类 型，内存缓冲区，网络链接，HTTP客户端，压缩工具，哈希等等。io包中定义了很多其它有用的接口 类型。Reader可以代表任意可以读取bytes的类型，Closer可以是任意可以关闭的值，例如一个文 件或是网络链接。

```go
package io
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Closer interface {
    Close() error
}
```

我们可以用这种方式以一个简写命名一个接口，而不用声明它所有 的方法。这种方式称为接口内嵌。

```go
type ReadWriter interface {
    Reader
		Writer
}
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```

我们还可以像下面这样，不使用接口内嵌来声明 io.ReadWriter接口。

```go
type ReadWriter interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
}
```

或者甚至使用一种混合的风格

```go
type ReadWriter interface {
		Read(p []byte) (n int, err error)
		Writer
}
```

上面3种定义方式都是一样的效果。方法顺序的变化也没有影响，唯一重要的就是这个集合里面的方 法。



为了保护你的Go语言职业生涯，请牢记接口（interface）是一种类型。

### 为什么要使用接口

```
type Cat struct{}

func (c Cat) Say() string { 
		return "喵喵喵" 
}

type Dog struct{}

func (d Dog) Say() string { 
		return "汪汪汪"
}

func main() {
	c := Cat{}
	fmt.Println("猫:", c.Say())
	d := Dog{}
	fmt.Println("狗:", d.Say())
}
```

上面的代码中定义了猫和狗，然后它们都会叫，你会发现main函数中明显有重复的代码，如果我们后续再加上猪、青蛙等动物的话，我们的代码还会一直重复下去。那我们能不能把它们当成“能叫的动物”来处理呢？

像类似的例子在我们编程过程中会经常遇到：

比如一个网上商城可能使用支付宝、微信、银联等方式去在线支付，我们能不能把它们当成“支付方式”来处理呢？

比如三角形，四边形，圆形都能计算周长和面积，我们能不能把它们当成“图形”来处理呢？

比如销售、行政、程序员都能计算月薪，我们能不能把他们当成“员工”来处理呢？

Go语言中为了解决类似上面的问题，就设计了接口这个概念。接口区别于我们之前所有的具体类型，接口是一种抽象的类型。当你看到一个接口类型的值时，你不知道它是什么，唯一知道的是通过它的方法能做什么。

### 接口的定义

Go语言提倡面向接口编程。

每个接口由数个方法组成，接口的定义格式如下：

```
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
```

其中：

接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。
举个例子：

```
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

当你看到这个接口类型的值时，你不知道它是什么，唯一知道的就是可以通过它的Write方法来做一些事情。

### 实现接口的条件

一个对象`(类型)`只要全部实现了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个需要实现的方法列表。

我们来定义一个Sayer接口：

```
// Sayer 接口
type Sayer interface {
	say()
}
```

定义dog和cat两个结构体：

```
type dog struct {}

type cat struct {}
```

因为Sayer接口里只有一个say方法，所以我们只需要给dog和cat 分别实现say方法就可以实现Sayer接口了。

```
// dog实现了Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}

// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}
```

接口的实现就是这么简单，只要实现了接口中的所有方法，就实现了这个接口。

### 接口类型变量

那实现了接口有什么用呢？

接口类型变量能够存储所有实现了该接口的实例。 例如上面的示例中，Sayer类型的变量能够存储dog和cat类型的变量。

```go
func main() {
	var x Sayer // 声明一个Sayer类型的变量x
	a := cat{}  // 实例化一个cat
	b := dog{}  // 实例化一个dog
	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵
	x = b       // 可以把dog实例直接赋值给x
	x.say()     // 汪汪汪
}
```

### 值接收者和指针接收者实现接口的区别

使用值接收者实现接口和使用指针接收者实现接口有什么区别呢？接下来我们通过一个例子看一下其中的区别。

我们有一个Mover接口和一个dog结构体。

```
type Mover interface {
	move()
}

type dog struct {}
```

### 值接收者实现接口

```
func (d dog) move() {
	fmt.Println("狗会动")
}
```

此时实现接口的是dog类型：

```go
func main() {
	var x Mover          //定义一个接口类型的变量x
	var wangcai = dog{} // 旺财是dog类型
	x = wangcai         // x可以接收dog类型
	var fugui = &dog{}  // 富贵是*dog类型
	x = fugui           // x可以接收*dog类型
	x.move()
}
```

从上面的代码中我们可以发现，使用值接收者实现接口之后，不管是dog结构体还是结构体指针\*dog类型的变量都可以赋值给该接口变量。因为Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui。

### 指针接收者实现接口

同样的代码我们再来测试一下使用指针接收者有什么区别：

```go
func (d *dog) move() {
	fmt.Println("狗会动")
}
func main() {
	var x Mover
	var wangcai = dog{} // 旺财是dog类型
	x = wangcai         // x不可以接收dog类型
	var fugui = &dog{}  // 富贵是*dog类型
	x = fugui           // x可以接收*dog类型
}
```

此时实现Mover接口的是\*dog类型，所以不能给x传入dog类型的wangcai，此时x只能存储*dog类型的值。



### 适配器模式

适配器：就是一个类的实例通过适配器功能可以实现一个它没有的功能

```go
package main

import(
	"fmt"
)
//B的充电头
type ChargeHeadB struct{
	name string
	size int
}

type TypeB interface{
	ChargeB()
}

//充电头b
func (this *ChargeHeadB)ChargeB(){
	fmt.Println("Head B ready to charge...")
}

//首先，需要实现想转换成的类型接口，也就是客户所期望看到的接口，C的接口适配器
type ChargeIfC struct{
	t TypeB //适配B
}
//适配器B转换成C
//func NewHead()TypeB{
//	return &ChargeHeadB{}
//}
////TypeB是接口
//func AdaptBtoC(b TypeB)(t TypeB){
//	return &ChargeIfC{
//		t : b,
//	}
//}
func NewHead(t TypeB) (*ChargeIfC){
	return &ChargeIfC{ t}
}

func (this *ChargeIfC)ChargeB() {
	this.ChargeC()
}
//B转换成C的适配器
func (this *ChargeIfC)ChargeC()  {
	fmt.Println("TypeC is charging...")
}

//func (this *ChargeIfC)ChargeB(){ //C的适配器调整到适合C的去充电
//	this.ChargeC()
//}
//
//func (this *ChargeIfC)ChargeC(){
//	fmt.Println("TypeC is charging...")
//}


func main()  {
	chargeB :=&ChargeHeadB{name: "apple", size: 16}
	chargeB.ChargeB()
	b := NewHead(chargeB)///创建新的充电头B
	b.ChargeC()
	//b.ChargeB()
	//c := AdaptBtoC(b)
	//c.ChargeB()
}
```



## 面试题

注意：这是一道你需要回答“能”或者“不能”的题！

首先请观察下面的这段代码，然后请回答这段代码能不能通过编译？

```
type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "hello" {
		talk = "hello world"
	} else {
		talk = "世界,你好"
	}
	return
}

func main() {
	var peo People = Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
```

## 类型与接口的关系

### 一个类型实现多个接口

一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。 例如，狗可以叫，也可以动。我们就分别定义Sayer接口和Mover接口，如下： Mover接口。

```go
// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}
dog既可以实现Sayer接口，也可以实现Mover接口。

type dog struct {
	name string
}

// 实现Sayer接口
func (d dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

// 实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会动\n", d.name)
}

func main() {
	var x Sayer
	var y Mover

	var a = dog{name: "旺财"}
	x = a
	y = a
	x.say()
	y.move()
}
```

### 多个类型实现同一接口

Go语言中不同的类型还可以实现同一接口 首先我们定义一个Mover接口，它要求必须由一个move方法。

```
// Mover 接口
type Mover interface {
	move()
}
```

例如狗可以动，汽车也可以动，可以使用如下代码实现这个关系：

```
type dog struct {
	name string
}

type car struct {
	brand string
}

// dog类型实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会跑\n", d.name)
}

// car类型实现Mover接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}
```

这个时候我们在代码中就可以把狗和汽车当成一个会动的物体来处理了，不再需要关注它们具体是什么，只需要调用它们的move方法就可以了。

```
func main() {
	var x Mover
	var a = dog{name: "旺财"}
	var b = car{brand: "保时捷"}
	x = a
	x.move()
	x = b
	x.move()
}
```

上面的代码执行结果如下：

```
旺财会跑
保时捷速度70迈
```

并且一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。

```go
// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}
```

### 接口嵌套

接口与接口间可以通过嵌套创造出新的接口。

```go
// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}
```

嵌套得到的接口的使用与普通接口一样，这里我们让cat实现animal接口：

```go
type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (c cat) move() {
	fmt.Println("猫会动")
}

func main() {
	var x animal
	x = cat{name: "花花"}
	x.move()
	x.say()
}
```

### 空接口

### 空接口的定义

空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。

> interface{} //空接口，没必要起名字

空接口类型的变量可以存储任意类型的变量。

```go
func main() {
	// 定义一个空接口x
	var x interface{}
	s := "Hello 沙河"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
}
```

### 空接口的应用

### 空接口作为函数的参数

使用空接口实现可以接收任意类型的函数参数。

```go
// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
```

### 空接口作为map的值

使用空接口实现可以保存任意值的字典。

```go
// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	studentInfo["hobby"] = [...]string{"sing", "dance", "rap"}
	fmt.Println(studentInfo)
```

### 类型断言

空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢？

### 接口值

一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值。

我们来看一个具体的例子：

```go
var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)
w = nil
```

请看下图分解：接口值图解

想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：

> x.(T)

其中：

​		x：表示类型为interface{}的变量
​		T：表示断言x可能是的类型。
该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。

举个例子：

```
func main() {
	var x interface{}
	x = "Hello 沙河"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
```

上面的示例中如果要断言多次就需要写多个if判断，这个时候我们可以使用switch语句来实现：

```
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
```

因为空接口可以存储任意类型值的特点，所以空接口在Go语言中的使用十分广泛。

关于接口需要注意的是，只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗。



# 反射

本文介绍了Go语言反射的意义和基本使用。

## 变量的内在机制

Go语言中的变量是分为两部分的:

```
类型信息：预先定义好的元信息。
值信息：程序运行过程中可动态变化的。
```

## 反射介绍

反射是指在程序运行期对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。

支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。

Go程序在运行期使用reflect包访问程序的反射信息。

在上一篇博客中我们介绍了空接口。 空接口可以存储任意类型的变量，那我们如何知道这个空接口保存的数据是什么呢？ 反射就是在运行时动态的获取一个变量的类型信息和值信息。

## reflect包

在Go语言的反射机制中，任何接口值都由是一个具体类型和具体类型的值两部分组成的(我们在上一篇接口的博客中有介绍相关概念)。 在Go语言中反射的相关功能由内置的reflect包提供，任意接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成，并且reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type。

### TypeOf

在Go语言中，使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息。

```
package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}
func main() {
	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64
}
```

### type name和type kind

在反射中关于类型还划分为两种：类型（Type）和种类（Kind）。因为在Go语言中我们可以使用type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型，但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）。 举个例子，我们定义了两个指针类型和两个结构体类型，通过反射查看它们的类型和种类。

```
package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func main() {
	var a *float32 // 指针
	var b myInt    // 自定义类型
	var c rune     // 类型别名
	reflectType(a) // type: kind:ptr
	reflectType(b) // type:myInt kind:int64
	reflectType(c) // type:int32 kind:int32

	type person struct {
		name string
		age  int
	}
	type book struct{ title string }
	var d = person{
		name: "沙河小王子",
		age:  18,
	}
	var e = book{title: "《跟小王子学Go语言》"}
	reflectType(d) // type:person kind:struct
	reflectType(e) // type:book kind:struct
}
```

Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。

在reflect包中定义的Kind类型如下：

```
type Kind uint
const (
    Invalid Kind = iota  // 非法类型
    Bool                 // 布尔型
    Int                  // 有符号整型
    Int8                 // 有符号8位整型
    Int16                // 有符号16位整型
    Int32                // 有符号32位整型
    Int64                // 有符号64位整型
    Uint                 // 无符号整型
    Uint8                // 无符号8位整型
    Uint16               // 无符号16位整型
    Uint32               // 无符号32位整型
    Uint64               // 无符号64位整型
    Uintptr              // 指针
    Float32              // 单精度浮点数
    Float64              // 双精度浮点数
    Complex64            // 64位复数类型
    Complex128           // 128位复数类型
    Array                // 数组
    Chan                 // 通道
    Func                 // 函数
    Interface            // 接口
    Map                  // 映射
    Ptr                  // 指针
    Slice                // 切片
    String               // 字符串
    Struct               // 结构体
    UnsafePointer        // 底层指针
)
```

### ValueOf

reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息。reflect.Value与原始值之间可以互相转换。

reflect.Value类型提供的获取原始值的方法如下：

| 方法                     | 说明                                                         |
| ------------------------ | ------------------------------------------------------------ |
| Interface() interface {} | 将值以 interface{} 类型返回，可以通过类型断言转换为指定类型  |
| Int() int64              | 将值以 int 类型返回，所有有符号整型均可以此方式返回          |
| Uint() uint64            | 将值以 uint 类型返回，所有无符号整型均可以此方式返回         |
| Float() float64          | 将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回 |
| Bool() bool              | 将值以 bool 类型返回                                         |
| Bytes() []bytes          | 将值以字节数组 []bytes 类型返回                              |
| String() string          | 将值以字符串类型返回                                         |

通过反射获取值

```
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}
func main() {
	var a float32 = 3.14
	var b int64 = 100
	reflectValue(a) // type is float32, value is 3.140000
	reflectValue(b) // type is int64, value is 100
	// 将int类型的原始值转换为reflect.Value类型
	c := reflect.ValueOf(10)
	fmt.Printf("type c :%T\n", c) // type c :reflect.Value
}
```

### 通过反射设置变量的值

想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。而反射中使用专有的Elem()方法来获取指针对应的值。

```
package main

import (
	"fmt"
	"reflect"
)

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
func main() {
	var a int64 = 100
	// reflectSetValue1(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&a)
	fmt.Println(a)
}
```

### isNil()和isValid()

#### isNil()

> func (v Value) IsNil() bool

IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。

#### isValid()

> func (v Value) IsValid() bool

IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。

#### 举个例子

IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。

```
func main() {
	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
}
```

## 结构体反射

### 与结构体相关的方法

任意值通过reflect.TypeOf()获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象（reflect.Type）的NumField()和Field()方法获得结构体成员的详细信息。

reflect.Type中与获取结构体成员相关的的方法如下表所示。

| 方法                                                        | 说明                                                         |
| ----------------------------------------------------------- | ------------------------------------------------------------ |
| Field(i int) StructField                                    | 根据索引，返回索引对应的结构体字段的信息。                   |
| NumField() int                                              | 返回结构体成员字段数量。                                     |
| FieldByName(name string) (StructField, bool)                | 根据给定字符串返回字符串对应的结构体字段的信息。             |
| FieldByIndex(index []int) StructField                       | 多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息。 |
| FieldByNameFunc(match func(string) bool) (StructField,bool) | 根据传入的匹配函数匹配需要的字段。                           |
| NumMethod() int                                             | 返回该类型的方法集中方法的数目                               |
| Method(int) Method                                          | 返回该类型方法集中的第i个方法                                |
| MethodByName(string)(Method, bool)                          | 根据方法名返回该类型方法集中的方法                           |

### StructField类型

StructField类型用来描述结构体中的一个字段的信息。

StructField的定义如下：

```
type StructField struct {
    // Name是字段的名字。PkgPath是非导出字段的包路径，对导出字段该字段为""。
    // 参见http://golang.org/ref/spec#Uniqueness_of_identifiers
    Name    string
    PkgPath string
    Type      Type      // 字段的类型
    Tag       StructTag // 字段的标签
    Offset    uintptr   // 字段在结构体中的字节偏移量
    Index     []int     // 用于Type.FieldByIndex时的索引切片
    Anonymous bool      // 是否匿名字段
}
```



### 结构体反射示例

当我们使用反射得到一个结构体数据之后可以通过索引依次获取其字段信息，也可以通过字段名去获取指定的字段信息。

```go
type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}
```

接下来编写一个函数printMethod(s interface{})来遍历打印s包含的方法。

// 给student添加两个方法 Study和Sleep(注意首字母大写)

```go
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}
```

### 反射是把双刃剑

反射是一个强大并富有表现力的工具，能让我们写出更灵活的代码。但是反射不应该被滥用，原因有以下三个。
    1. 基于反射的代码是极其脆弱的，反射中的类型错误会在真正运行的时候才会引发panic，那很可能是在代码写完的很长时间之后。
    2. 大量使用反射的代码通常难以理解。
    3. 反射的性能低下，基于反射实现的代码通常比正常代码运行速度慢一到两个数量级。

### 练习题

编写代码利用反射实现一个ini文件的解析器程序。



# 并发性

## 1.1 多任务

​		怎么来理解多任务呢？其实就是指我们的操作系统可以同时执行多个任务。举个例子，你一边听音乐，一边刷微博，一边聊QQ，一边用Markdown写作业，这就是多任务，至少同时有4个任务正在运行。还有很多任务悄悄地在后台同时运行着，只是界面上没有显示而已。



## 1.2 什么是并发

​		Go是并发语言，而不是并行语言。在讨论如何在Go中进行并发处理之前，我们首先必须了解什么是并发，以及它与并行性有什么不同。(Go is a concurrent language and not a parallel one. )

​		**并发性Concurrency是同时处理许多事情的能力。**

​		举个例子，假设一个人在晨跑。在晨跑时，他的鞋带松了。现在这个人停止跑步，系鞋带，然后又开始跑步。这是一个典型的并发性示例。这个人能够同时处理跑步和系鞋带，这是一个人能够同时处理很多事情。

​		什么是并行性parallelism，它与并发concurrency有什么不同?
并行就是同时做很多事情。这听起来可能与并发类似，但实际上是不同的。

​		让我们用同样的慢跑例子更好地理解它。在这种情况下，我们假设这个人正在慢跑，并且使用它的手机听音乐。在这种情况下，一个人一边慢跑一边听音乐，那就是他同时在做很多事情。这就是所谓的并行性(parallelism)。



并发性和并行性

​		假设我们正在编写一个web浏览器。web浏览器有各种组件。其中两个是web页面呈现区域和下载文件从internet下载的下载器。**假设我们以这样的方式构建了浏览器的代码，这样每个组件都可以独立地执行。当这个浏览器运行在单个核处理器中时，处理器将在浏览器的两个组件之间进行上下文切换。它可能会下载一个文件一段时间，然后它可能会切换到呈现用户请求的网页的html。这就是所谓的并发性。**并发进程从不同的时间点开始，它们的执行周期重叠。在这种情况下，下载和呈现从不同的时间点开始，它们的执行重叠。



**假设同一浏览器运行在多核处理器上。在这种情况下，文件下载组件和HTML呈现组件可能同时在不同的内核中运行。这就是所谓的并行性。**

![WX20190730-100944](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190730-100944.png)



​		并行性Parallelism不会总是导致更快的执行时间。这是因为并行运行的组件可能需要相互通信。例如，在我们的浏览器中，当文件下载完成时，应该将其传递给用户，比如使用弹出窗口。这种通信发生在负责下载的组件和负责呈现用户界面的组件之间。这种通信开销在并发concurrent 系统中很低。当组件在多个内核中并行concurrent 运行时，这种通信开销很高。因此，并行程序并不总是导致更快的执行时间!



![t](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/t.png)



也就是说从某个时间段来讲,并行是在执行多个程序,并发是在执行一个程序



## 1.3 进程、线程、协程

**进程(Process)，线程(Thread)，协程(Coroutine，也叫轻量级线程)**

**进程**

​		进程是一个程序在一个数据集中的一次动态执行过程，可以理解为“正在执行的程序”，它是CPU资源分配和调度的独立单位。 
​		进程一般由程序、数据集、进程控制块三部分组成。我们编写的程序用来描述进程要完成哪些功能以及如何完成；数据集则是程序在执行过程中所需要使用的资源；进程控制块用来记录进程的外部特征，描述进程的执行变化过程，系统可以利用它来控制和管理进程，它是系统感知进程存在的唯一标志。 **进程的局限是创建、撤销和切换的开销比较大。**



**线程**

​		线程是在进程之后发展出来的概念。 线程也叫轻量级进程，它是一个基本的CPU执行单元，也是程序执行过程中的最小单元，由线程ID、程序计数器、寄存器集合和堆栈共同组成。一个进程可以包含多个线程。 
线程的优点是减小了程序并发执行时的开销，提高了操作系统的并发性能，缺点是线程没有自己的系统资源，只拥有在运行时必不可少的资源，但同一进程的各线程可以共享进程所拥有的系统资源，如果把进程比作一个车间，那么线程就好比是车间里面的工人。不过对于某些独占性资源存在锁机制，处理不当可能会产生“死锁”。



**协程**

​		协程是一种用户态的轻量级线程，又称微线程，英文名Coroutine，协程的调度完全由用户控制。人们通常将协程和子程序（函数）比较着理解。子程序调用总是一个入口，一次返回，一旦退出即完成了子程序的执行。 

​		**与传统的系统级线程和进程相比，协程的最大优势在于其"轻量级"，可以轻松创建上百万个而不会导致系统资源衰竭，而线程和进程通常最多也不能超过1万的。这也是协程也叫轻量级线程的原因。**

> 协程与多线程相比，其优势体现在：协程的执行效率极高。因为子程序切换不是线程切换，而是由程序自身控制，因此，没有线程切换的开销，和多线程比，线程数量越多，协程的性能优势就越明显。

**Go语言对于并发的实现是靠协程，Goroutine**



## Go语言的并发模型

Go 语言相比Java等一个很大的优势就是可以方便地编写并发程序。Go 语言内置了 goroutine 机制，使用goroutine可以快速地开发并发程序， 更好的利用多核处理器资源。接下来我们来了解一下Go语言的并发原理。

### 2.1 线程模型

​		在现代操作系统中，线程是处理器调度和分配的基本单位，进程则作为资源拥有的基本单位。每个进程是由私有的虚拟地址空间、代码、数据和其它各种系统资源组成。线程是进程内部的一个执行单元。 每一个进程至少有一个主执行线程，它无需由用户去主动创建，是由系统自动创建的。 用户根据需要在应用程序中创建其它线程，多个线程并发地运行于同一个进程中。

​		我们先从线程讲起，无论语言层面何种并发模型，到了操作系统层面，一定是以线程的形态存在的。而操作系统根据资源访问权限的不同，体系架构可分为用户空间和内核空间；内核空间主要操作访问CPU资源、I/O资源、内存资源等硬件资源，为上层应用程序提供最基本的基础资源，用户空间呢就是上层应用程序的固定活动空间，用户空间不可以直接访问资源，必须通过“系统调用”、“库函数”或“Shell脚本”来调用内核空间提供的资源。

​		我们现在的计算机语言，可以狭义的认为是一种“软件”，它们中所谓的“线程”，往往是用户态的线程，和操作系统本身内核态的线程（简称KSE），还是有区别的。

Go并发编程模型在底层是由操作系统所提供的线程库支撑的，因此还是得从线程实现模型说起。

​		线程可以视为进程中的控制流。一个进程至少会包含一个线程，因为其中至少会有一个控制流持续运行。因而，一个进程的第一个线程会随着这个进程的启动而创建，这个线程称为该进程的主线程。当然，一个进程也可以包含多个线程。这些线程都是由当前进程中已存在的线程创建出来的，创建的方法就是调用系统调用，更确切地说是调用pthread create函数。拥有多个线程的进程可以并发执行多个任务，并且即使某个或某些任务被阻塞，也不会影响其他任务正常执行，这可以大大改善程序的响应时间和吞吐量。另一方面，线程不可能独立于进程存在。它的生命周期不可能逾越其所属进程的生命周期。

​		**线程的实现模型主要有3个，分别是:用户级线程模型、内核级线程模型和两级线程模型。**它们之间最大的差异就在于线程与内核调度实体( Kernel Scheduling Entity,简称KSE)之间的对应关系上。顾名思义，内核调度实体就是可以被内核的调度器调度的对象。在很多文献和书中，它也称为内核级线程，是操作系统内核的最小调度单元。



#### 2.1.1 内核级线程模型

​		用户线程与KSE是1对1关系(1:1)。大部分编程语言的线程库(如linux的pthread，Java的java.lang.Thread，C++11的std::thread等等)都是对操作系统的线程（内核级线程）的一层封装，创建出来的每个线程与一个不同的KSE静态关联，因此其调度完全由OS调度器来做。这种方式实现简单，直接借助OS提供的线程能力，并且不同用户线程之间一般也不会相互影响。但其创建，销毁以及多个线程之间的上下文切换等操作都是直接由OS层面亲自来做，在需要使用大量线程的场景下对OS的性能影响会很大。

![moxing2](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/moxing2.jpg)

每个线程由内核调度器独立的调度，所以如果一个线程阻塞则不影响其他的线程。

**优点：**

​		在多核处理器的硬件的支持下，内核空间线程模型支持了真正的并行，当一个线程被阻塞后，允许另一个线程继续执行，所以并发能力较强。

**缺点:** 

​		每创建一个用户级线程都需要创建一个内核级线程与其对应，这样创建线程的开销比较大，会影响到应用程序的性能。



#### 2.1.2 用户级线程模型

​		用户线程与KSE是多对1关系(M:1)，这种线程的创建，销毁以及多个线程之间的协调等操作都是由用户自己实现的线程库来负责，对OS内核透明，一个进程中所有创建的线程都与同一个KSE在运行时动态关联。现在有许多语言实现的 **协程** 基本上都属于这种方式。这种实现方式相比内核级线程可以做的很轻量级，对系统资源的消耗会小很多，因此可以创建的数量与上下文切换所花费的代价也会小得多。但该模型有个致命的缺点，如果我们在某个用户线程上调用阻塞式系统调用(如用阻塞方式read网络IO)，那么一旦KSE因阻塞被内核调度出CPU的话，剩下的所有对应的用户线程全都会变为阻塞状态（整个进程挂起）。 
所以这些语言的**协程库**会把自己一些阻塞的操作重新封装为完全的非阻塞形式，然后在以前要阻塞的点上，主动让出自己，并通过某种方式通知或唤醒其他待执行的用户线程在该KSE上运行，从而避免了内核调度器由于KSE阻塞而做上下文切换，这样整个进程也不会被阻塞。

![moxing1](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/moxing1.jpg)

**优点：**

​		这种模型的好处是线程上下文切换都发生在用户空间，避免的模态切换（mode switch），从而对于性能有积极的影响。

**缺点：**

​		所有的线程基于一个内核调度实体即内核线程，这意味着只有一个处理器可以被利用，在多处理器环境下这是不能够被接受的，本质上，用户线程只解决了并发问题，但是没有解决并行问题。如果线程因为 I/O 操作陷入了内核态，内核态线程阻塞等待 I/O 数据，则所有的线程都将会被阻塞，用户空间也可以使用非阻塞而 I/O，但是不能避免性能及复杂度问题。



#### 2.1.3 两级线程模型

​		用户线程与KSE是多对多关系(M:N)，这种实现综合了前两种模型的优点，为一个进程中创建多个KSE，并且线程可以与不同的KSE在运行时进行动态关联，当某个KSE由于其上工作的线程的阻塞操作被内核调度出CPU时，当前与其关联的其余用户线程可以重新与其他KSE建立关联关系。当然这种动态关联机制的实现很复杂，也需要用户自己去实现，这算是它的一个缺点吧。Go语言中的并发就是使用的这种实现方式，Go为了实现该模型自己实现了一个运行时调度器来负责Go中的"线程"与KSE的动态关联。此模型有时也被称为 **混合型线程模型**，**即用户调度器实现用户线程到KSE的“调度”，内核调度器实现KSE到CPU上的调度**。

![moxing3](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/moxing3.jpg)



### 2.2 Go并发调度: G-P-M模型

​		在操作系统提供的内核线程之上，Go搭建了一个特有的两级线程模型。goroutine机制实现了M : N的线程模型，goroutine机制是协程（coroutine）的一种实现，golang内置的调度器，可以让多核CPU中每个CPU执行一个协程。

#### 2.2.1 调度器是如何工作的

有了上面的认识，我们可以开始真正的介绍Go的并发机制了，先用一段代码展示一下在Go语言中新建一个“线程”(Go语言中称为Goroutine)的样子：

```go
// 用go关键字加上一个函数（这里用了匿名函数）
// 调用就做到了在一个新的“线程”并发执行任务
go func() { 
    // do something in one new goroutine
}()
```

功能上等价于Java8的代码:

```java
new java.lang.Thread(() -> { 
    // do something in one new thread
}).start();
```

理解goroutine机制的原理，关键是理解Go语言scheduler的实现。

Go语言中支撑整个scheduler实现的主要有4个重要结构，分别是M、G、P、Sched， 前三个定义在runtime.h中，Sched定义在proc.c中。

- Sched结构就是调度器，它维护有存储M和G的队列以及调度器的一些状态信息等。
- M结构是Machine，系统线程，它由操作系统管理的，goroutine就是跑在M之上的；M是一个很大的结构，里面维护小对象内存cache（mcache）、当前执行的goroutine、随机数发生器等等非常多的信息。
- P结构是Processor，处理器，它的主要用途就是用来执行goroutine的，它维护了一个goroutine队列，即runqueue。Processor是从N:1调度到M:N调度的重要部分。
- G是goroutine实现的核心结构，它包含了栈，指令指针，以及其他对调度goroutine很重要的信息，例如其阻塞的channel。

> Processor的数量是在启动时被设置为环境变量GOMAXPROCS的值，或者通过运行时调用函数GOMAXPROCS()进行设置。Processor数量固定意味着任意时刻只有GOMAXPROCS个线程在运行go代码。

分别用三角形，矩形和圆形表示Machine Processor和Goroutine。

![moxing4](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/moxing4.jpg)

​		在单核处理器的场景下，所有goroutine运行在同一个M系统线程中，每一个M系统线程维护一个Processor，任何时刻，一个Processor中只有一个goroutine，其他goroutine在runqueue中等待。一个goroutine运行完自己的时间片后，让出上下文，回到runqueue中。 多核处理器的场景下，为了运行goroutines，每个M系统线程会持有一个Processor。

![moxing5](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/moxing5.jpg)





  	 在正常情况下，scheduler会按照上面的流程进行调度，但是线程会发生阻塞等情况，看一下goroutine对线程阻塞等的处理。



#### 2.2.2 系统调用

​		P的个数默认等于CPU核数，每个M必须持有一个P才可以执行G，一般情况下M的个数会略大于P的个数，这多出来的M将会在G产生系统调用时发挥作用。类似线程池，Go也提供一个M的池子，需要时从池子中获取，用完放回池子，不够用时就再创建一个。

当M运行的某个G产生系统调用时，如下图所示：

![moxing6](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/moxing6.jpg)

如图所示，当G0即将进入系统调用时，M0将释放P，进而某个空闲的M1获取P，继续执行P队列中剩下的G。而M0由于陷入系统调用而进被阻塞，M1接替M0的工作，只要P不空闲，就可以保证充分利用CPU。

M1的来源有可能是M的缓存池，也可能是新建的。当G0系统调用结束后，根据M0是否能获取到P，将会将G0做不同的处理：

1. 如果有空闲的P，则获取一个P，继续执行G0。
2. 如果没有空闲的P，则将G0放入全局队列，等待被其他的P调度。然后M0将进入缓存池睡眠。

#### 2.2.3 工作量窃取

当其中一个Processor的runqueue为空，没有goroutine可以调度。它会从另外一个上下文偷取一半的goroutine。

![moxing7](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/moxing7.jpg)

> 其图中的G，P和M都是Go语言运行时系统（其中包括内存分配器，并发调度器，垃圾收集器等组件，可以想象为Java中的JVM）抽象出来概念和数据结构对象：
> G：Goroutine的简称，上面用go关键字加函数调用的代码就是创建了一个G对象，是对一个要并发执行的任务的封装，也可以称作用户态线程。属于用户级资源，对OS透明，具备轻量级，可以大量创建，上下文切换成本低等特点。
> M：Machine的简称，在linux平台上是用clone系统调用创建的，其与用linux pthread库创建出来的线程本质上是一样的，都是利用系统调用创建出来的OS线程实体。M的作用就是执行G中包装的并发任务。**Go运行时系统中的调度器的主要职责就是将G公平合理的安排到多个M上去执行**。其属于OS资源，可创建的数量上也受限了OS，通常情况下G的数量都多于活跃的M的。
> P：Processor的简称，逻辑处理器，主要作用是管理G对象（每个P都有一个G队列），并为G在M上的运行提供本地化资源。

​		从两级线程模型来看，似乎并不需要P的参与，有G和M就可以了，那为什么要加入P这个东东呢？
​		其实Go语言运行时系统早期(Go1.0)的实现中并没有P的概念，Go中的调度器直接将G分配到合适的M上运行。但这样带来了很多问题，例如，不同的G在不同的M上并发运行时可能都需向系统申请资源（如堆内存），由于资源是全局的，将会由于资源竞争造成很多系统性能损耗，为了解决类似的问题，后面的Go（Go1.1）运行时系统加入了P，让P去管理G对象，M要想运行G必须先与一个P绑定，然后才能运行该P管理的G。这样带来的好处是，我们可以在P对象中预先申请一些系统资源（本地资源），G需要的时候先向自己的本地P申请（无需锁保护），如果不够用或没有再向全局申请，而且从全局拿的时候会多拿一部分，以供后面高效的使用。就像现在我们去政府办事情一样，先去本地政府看能搞定不，如果搞不定再去中央，从而提供办事效率。
​		而且由于P解耦了G和M对象，这样即使M由于被其上正在运行的G阻塞住，其余与该M关联的G也可以随着P一起迁移到别的活跃的M上继续运行，从而让G总能及时找到M并运行自己，从而提高系统的并发能力。
​		Go运行时系统通过构造G-P-M对象模型实现了一套用户态的并发调度系统，可以自己管理和调度自己的并发任务，所以可以说Go语言**原生支持并发**。**自己实现的调度器负责将并发任务分配到不同的内核线程上运行，然后内核调度器接管内核线程在CPU上的执行与调度。**

可以看到Go的并发用起来非常简单，用了一个语法糖将内部复杂的实现结结实实的包装了起来。其内部可以用下面这张图来概述：

![goroutine2](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/goroutine2.png)

写在最后，Go运行时完整的调度系统是很复杂，很难用一篇文章描述的清楚，这里只能从宏观上介绍一下，让大家有个整体的认识。

```go
// Goroutine1
func task1() {
    go task2()
    go task3()
}
```

假如我们有一个G(Goroutine1)已经通过P被安排到了一个M上正在执行，在Goroutine1执行的过程中我们又创建两个G，这两个G会被马上放入与Goroutine1相同的P的本地G任务队列中，排队等待与该P绑定的M的执行，这是最基本的结构，很好理解。 关键问题是:

**a.如何在一个多核心系统上尽量合理分配G到多个M上运行，充分利用多核，提高并发能力呢？**

如果我们在一个Goroutine中通过**go**关键字创建了大量G，这些G虽然暂时会被放在同一个队列, 但如果这时还有空闲P（系统内P的数量默认等于系统cpu核心数），Go运行时系统始终能保证至少有一个（通常也只有一个）活跃的M与空闲P绑定去各种G队列去寻找可运行的G任务，该种M称为**自旋的M**。一般寻找顺序为：自己绑定的P的队列，全局队列，然后其他P队列。如果自己P队列找到就拿出来开始运行，否则去全局队列看看，由于全局队列需要锁保护，如果里面有很多任务，会转移一批到本地P队列中，避免每次都去竞争锁。如果全局队列还是没有，就要开始玩狠的了，直接从其他P队列偷任务了（偷一半任务回来）。这样就保证了在还有可运行的G任务的情况下，总有与CPU核心数相等的M+P组合 在执行G任务或在执行G的路上(寻找G任务)。

**b. 如果某个M在执行G的过程中被G中的系统调用阻塞了，怎么办？**

在这种情况下，这个M将会被内核调度器调度出CPU并处于阻塞状态，与该M关联的其他G就没有办法继续执行了，但Go运行时系统的一个监控线程(sysmon线程)能探测到这样的M，并把与该M绑定的P剥离，寻找其他空闲或新建M接管该P，然后继续运行其中的G，大致过程如下图所示。然后等到该M从阻塞状态恢复，需要重新找一个空闲P来继续执行原来的G，如果这时系统正好没有空闲的P，就把原来的G放到全局队列当中，等待其他M+P组合发掘并执行。



**c. 如果某一个G在M运行时间过长，有没有办法做抢占式调度，让该M上的其他G获得一定的运行时间，以保证调度系统的公平性?**

我们知道linux的内核调度器主要是基于时间片和优先级做调度的。对于相同优先级的线程，内核调度器会尽量保证每个线程都能获得一定的执行时间。为了防止有些线程"饿死"的情况，内核调度器会发起抢占式调度将长期运行的线程中断并让出CPU资源，让其他线程获得执行机会。当然在Go的运行时调度器中也有类似的抢占机制，但并不能保证抢占能成功，因为Go运行时系统并没有内核调度器的中断能力，它只能通过向运行时间过长的G中设置抢占flag的方法温柔的让运行的G自己主动让出M的执行权。 
说到这里就不得不提一下Goroutine在运行过程中可以动态扩展自己线程栈的能力，可以从初始的2KB大小扩展到最大1G（64bit系统上），因此在每次调用函数之前需要先计算该函数调用需要的栈空间大小，然后按需扩展（超过最大值将导致运行时异常）。Go抢占式调度的机制就是利用在判断要不要扩栈的时候顺便查看以下自己的抢占flag，决定是否继续执行，还是让出自己。
运行时系统的监控线程会计时并设置抢占flag到运行时间过长的G，然后G在有函数调用的时候会检查该抢占flag，如果已设置就将自己放入全局队列，这样该M上关联的其他G就有机会执行了。但如果正在执行的G是个很耗时的操作且没有任何函数调用(如只是for循环中的计算操作)，即使抢占flag已经被设置，该G还是将一直霸占着当前M直到执行完自己的任务。



## runtime包

官网文档对runtime包的介绍：

```
Package runtime contains operations that interact with Go's runtime system, such as functions to control goroutines. It also includes the low-level type information used by the reflect package; see reflect's documentation for the programmable interface to the run-time type system.
```

尽管 Go 编译器产生的是本地可执行代码，这些代码仍旧运行在 Go 的 runtime（这部分的代码可以在 runtime 包中找到）当中。这个 runtime 类似 Java 和 .NET 语言所用到的虚拟机，它负责管理包括内存分配、垃圾回收（第 10.8 节）、栈处理、goroutine、channel、切片（slice）、map 和反射（reflection）等等。

### 3.1 常用函数

**关于 `runtime` 包几个方法:**

- **NumCPU**：返回当前系统的 `CPU` 核数量

- **GOMAXPROCS**：设置最大的可同时使用的 `CPU` 核数

  通过runtime.GOMAXPROCS函数，应用程序何以在运行期间设置运行时系统中得P最大数量。但这会引起“Stop the World”。所以，应在应用程序最早的调用。并且最好是在运行Go程序之前设置好操作程序的环境变量GOMAXPROCS，而不是在程序中调用runtime.GOMAXPROCS函数。

  无论我们传递给函数的整数值是什么值，运行时系统的P最大值总会在1~256之间。

> go1.8后，默认让程序运行在多个核上,可以不用设置了
> go1.8前，还是要设置一下，可以更高效的利益cpu

- **Gosched**：让当前线程让出 `cpu` 以让其它线程运行,它不会挂起当前线程，因此当前线程未来会继续执行

  这个函数的作用是让当前 `goroutine` 让出 `CPU`，当一个 `goroutine` 发生阻塞，`Go` 会自动地把与该 `goroutine` 处于同一系统线程的其他 `goroutine` 转移到另一个系统线程上去，以使这些 `goroutine` 不阻塞。

- **Goexit**：退出当前 `goroutine`(但是`defer`语句会照常执行)

- **NumGoroutine**：返回正在执行和排队的任务总数

  runtime.NumGoroutine函数在被调用后，会返回系统中的处于特定状态的Goroutine的数量。这里的特指是指Grunnable\Gruning\Gsyscall\Gwaition。处于这些状态的Groutine即被看做是活跃的或者说正在被调度。

  注意：垃圾回收所在Groutine的状态也处于这个范围内的话，也会被纳入该计数器。

- **GOOS**：目标操作系统

- **runtime.GC**:会让运行时系统进行一次强制性的垃圾收集

  1. 强制的垃圾回收：不管怎样，都要进行的垃圾回收。
  2. 非强制的垃圾回收：只会在一定条件下进行的垃圾回收（即运行时，系统自上次垃圾回收之后新申请的堆内存的单元（也成为单元增量）达到指定的数值）。

- **GOROOT** :获取goroot目录

- **GOOS** : 查看目标操作系统 
  很多时候，我们会根据平台的不同实现不同的操作，就而已用GOOS了：




### 3.2 示例代码：

1. 获取goroot和os：

 ```go
//获取goroot目录：
fmt.Println("GOROOT-->",runtime.GOROOT())
   
//获取操作系统
fmt.Println("os/platform-->",runtime.GOOS) // GOOS--> darwin，mac系统
   
 ```

2. 获取CPU数量，和设置CPU数量：

```go
func init(){
	//1.获取逻辑cpu的数量
	fmt.Println("逻辑CPU的核数：",runtime.NumCPU())
	//2.设置go程序执行的最大的：[1,256]
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n)
}
```

![WX20190806-103956](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190806-103956.png)

3. Gosched()：

```go
func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine。。。")
		}

	}()

	for i := 0; i < 4; i++ {
		//让出时间片，先让别的协议执行，它执行完，再回来执行此协程
		runtime.Gosched()
		fmt.Println("main。。")
	}
}

```

![WX20190806-104235](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190806-104235.png)



4. Goexit的使用（终止协程）

```go
func main() {
	//创建新建的协程
	go func() {
		fmt.Println("goroutine开始。。。")

		//调用函数
		fun()

		fmt.Println("goroutine结束。。")
	}() //立即执行函数

	//sleep 3s
	time.Sleep(3*time.Second)
}

func fun() {
	defer fmt.Println("defer。。。")

	//return           //终止此函数
	runtime.Goexit() //终止所在的协程

	fmt.Println("fun函数。。。")
}
```

执行结果:

![WX20190806-105752](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190806-105752.png)



## 临界资源安全问题

### 4.1 临界资源

**临界资源:** 指并发环境中多个进程/线程/协程共享的资源。

但是在并发编程中对临界资源的处理不当， 往往会导致数据不一致的问题。

示例代码：

```go
package main

import (
	"fmt"
	"time"
)

func main()  {
	a := 1
	go func() {
		a = 2
		fmt.Println("子goroutine。。",a)
	}()
	a = 3
	time.Sleep(1)
	fmt.Println("main goroutine。。",a)
}
```

我们通过终端命令来执行：

![WX20190806-155844](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190806-155844.png)

能够发现一处被多个goroutine共享的数据。



### 4.2 临界资源安全问题


并发本身并不复杂，但是因为有了资源竞争的问题，就使得开发出好的并发程序变得复杂起来，因为会引起很多莫名其妙的问题。

如果多个goroutine在访问同一个数据资源的时候，其中一个线程修改了数据，那么这个数值就被修改了，对于其他的goroutine来讲，这个数值可能是不对的。

举个例子，我们通过并发来实现火车站售票这个程序。一共有100张票，4个售票口同时出售。

我们先来看一下示例代码：

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//全局变量
var ticket = 10 // 10张票

func main() {
	/*
	4个goroutine，模拟4个售票口，4个子程序操作同一个共享数据。
	*/
	go saleTickets("售票口1") // g1,10
	go saleTickets("售票口2") // g2,10
	go saleTickets("售票口3") //g3,10
	go saleTickets("售票口4") //g4,10

	time.Sleep(5*time.Second)
}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	//for i:=1;i<=100;i++{
	//	fmt.Println(name,"售出：",i)
	//}
	for { //ticket=1
		if ticket > 0 { //g1,g3,g2,g4
			//睡眠
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			// g1 ,g3, g2,g4
			fmt.Println(name, "售出：", ticket)  // 1 , 0, -1 , -2
			ticket--   //0 , -1 ,-2 , -3
		} else {
			fmt.Println(name,"售罄，没有票了。。")
			break
		}
	}
}

```

为了更好的观察临界资源问题，每个goroutine先睡眠一个随机数，然后再售票，运行结果，还可以卖出编号为负数的票。

**执行结果:**

```txt
$ go run saleticket.go
售票口2 售出: 10
售票口1 售出: 9
售票口1 售出: 8
售票口4 售出: 7
售票口4 售出: 6
售票口2 售出: 5
售票口3 售出: 4
售票口4 售出: 3
售票口2 售出: 2
售票口4 售出: 1
售罄,没有票了.
售票口1 售出: 0
售罄,没有票了.
售票口3 售出: -1
售罄,没有票了.
售票口2 售出: -2
售罄,没有票了.
```

**分析：**

​		卖票逻辑是先判断票数是否为负数，如果大于0，就进行卖票，前提是卖票前先睡眠，然后再卖，假如说此时已经卖票到只剩最后1张了，某一个goroutine持有了CPU的时间片，那么它再该时间片段是否有票，条件是成立的，所以它可以卖票编号为1的最后一张票。但是因为它在卖之前，先睡眠了，那么其他的goroutine就会持有CPU的时间片，而此时这张票还没有被卖出，那么第二个goroutine再判断是否有票的时候，条件也是成立的，那么它可以卖出这张票，然而它也进入了睡眠。。其他的第三个第四个goroutine都是这样的逻辑，当某个goroutine醒来的时候，不会再判断是否有票，而是直接售出，这样就卖出最后一张票了，然而其他的goroutine醒来的时候，就会依次卖出第0张，-1张，-2张。

​		**这就是临界资源的不安全问题。某一个goroutine在访问某个数据资源的时候，按照数值，已经判断好了条件，然后又被其他的goroutine抢占了资源，并修改了数值，等这个goroutine再继续访问这个数据的时候，数值已经改变了。**



### 4.3 临界资源安全问题的解决

**要想解决临界资源安全的问题，很多编程语言的解决方案都是同步。通过上锁的方式，某一时间段，只能允许一个goroutine来访问这个共享数据，当前goroutine访问完毕，解锁后，其他的goroutine才能来访问。**

可以借助于sync包下的锁机制。

示例代码：

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

//全局变量
var ticket = 10 // 100张票

var wg sync.WaitGroup //创建同步等待组对象
var matex sync.Mutex // 创建锁对象

func main() {
	/*
	4个goroutine，模拟4个售票口，4个子程序操作同一个共享数据。
	 */
	wg.Add(4)
	go saleTickets("售票口1") // g1,100
	go saleTickets("售票口2") // g2,100
	go saleTickets("售票口3") //g3,100
	go saleTickets("售票口4") //g4,100
	wg.Wait()              // main要等待。。。

	//time.Sleep(5*time.Second)
}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
  
	for { //ticket=1
		matex.Lock()
		if ticket > 0 { //g1,g3,g2,g4
			//睡眠
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			// g1 ,g3, g2,g4
			fmt.Println(name, "售出：", ticket) // 1 , 0, -1 , -2
			ticket--                         //0 , -1 ,-2 , -3
		} else {
			matex.Unlock() //解锁
			fmt.Println(name, "售罄，没有票了。。")
			break
		}
		matex.Unlock() //解锁
	}
}
```

运行结果：

```
$ go run saleticket.go
售票口1 售出： 10
售票口1 售出： 9
售票口4 售出： 8
售票口2 售出： 7
售票口3 售出： 6
售票口1 售出： 5
售票口4 售出： 4
售票口2 售出： 3
售票口3 售出： 2
售票口1 售出： 1
售票口1 售罄，没有票了。。
售票口3 售罄，没有票了。。
售票口4 售罄，没有票了。。
售票口2 售罄，没有票了。。
```

在Go语言中并不鼓励用锁保护共享状态的方式在不同的Goroutine中分享信息(以共享内存的方式去通信)。而是鼓励通过**channel**将共享状态或共享状态的变化在各个Goroutine之间传递（以通信的方式去共享内存），这样同样能像用锁一样保证在同一的时间只有一个Goroutine访问共享状态。

当然，在主流的编程语言中为了保证多线程之间共享数据安全性和一致性，都会提供一套基本的同步工具集，如锁，条件变量，原子操作等等。Go语言标准库也毫不意外的提供了这些同步机制，使用方式也和其他语言也差不多。

## sync包

官网文档对sync包的介绍：

```
Package sync provides basic synchronization primitives such as mutual exclusion locks. Other than the Once and WaitGroup types, most are intended for use by low-level library routines. Higher-level synchronization is better done via channels and communication.
```

![WX20190807-101109](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190807-101109.png)

sync是synchronization同步这个词的缩写，所以也会叫做同步包。这里提供了基本同步的操作，比如互斥锁等等。这里除了Once和WaitGroup类型之外，大多数类型都是供低级库例程使用的。更高级别的同步最好通过channel通道和communication通信来完成

### 5.1 WaitGroup

WaitGroup，同步等待组。

在类型上，它是一个结构体。一个WaitGroup的用途是等待一个goroutine的集合执行完成。主goroutine调用了Add()方法来设置要等待的goroutine的数量。然后，每个goroutine都会执行并且执行完成后调用Done()这个方法。与此同时，可以使用Wait()方法来阻塞，直到所有的goroutine都执行完成。

![WX20190807-101436](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190807-101436.png)

### 5.1.1 Add()方法：

Add这个方法，用来设置到WaitGroup的计数器的值。可以理解为每个waitgroup中都有一个计数器
用来表示这个同步等待组中要执行的goroutine的数量。

如果计数器的数值变为0，那么就表示等待时被阻塞的goroutine都被释放，如果计数器的数值为负数，那么就会引发panic，程序就报错了。

![WX20190807-102137](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190807-102137.png)



### 5.1.2 Done()方法

Done()方法，就是当WaitGroup同步等待组中的某个goroutine执行完毕后，设置这个WaitGroup的counter数值减1。

![WX20190807-102843](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190807-102843.png)

其实Done()的底层代码就是调用了Add()方法：

```go
// Done decrements the WaitGroup counter by one.
func (wg *WaitGroup) Done() {
	wg.Add(-1)
}
```



### 5.1.3 Wait()方法

Wait()方法，表示让当前的goroutine等待，进入阻塞状态。一直到WaitGroup的计数器为零。才能解除阻塞，
这个goroutine才能继续执行。

![WX20190807-103015](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190807-103015.png)



### 5.1.4 示例代码：

创建并启动两个goroutine，来打印数字和字母，并在main goroutine中，将这两个子goroutine加入到一个WaitGroup中，同时让main goroutine进入Wait()，让两个子goroutine先执行。当每个子goroutine执行完毕后，调用Done()方法，设置WaitGroup的counter减1。当两条子goroutine都执行完毕后，WaitGroup中的counter的数值为零，解除main goroutine的阻塞。

示例代码：

```go
package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup // 创建同步等待组对象
func main()  {
	/*
	WaitGroup：同步等待组
		可以使用Add(),设置等待组中要执行的子goroutine的数量，
		
		在main 函数中，使用wait(),让主程序处于等待状态。直到等待组中子程序执行完毕。解除阻塞

		子gorotuine对应的函数中。wg.Done()，用于让等待组中的子程序的数量减1
	 */
	//启动一个goroutine登记,设置等待组中，要执行的goroutine的数量
	wg.Add(2)
	go fun1()
	go fun2()
	fmt.Println("main进入阻塞状态。。。等待wg中的子goroutine结束。。")
	wg.Wait() //表示main goroutine进入等待，意味着阻塞
	fmt.Println("main，解除阻塞。。")

}
func fun1()  {
	for i:=1;i<=10;i++{
		fmt.Println("fun1.。。i:",i)
	}
	wg.Done() //给wg等待中的执行的goroutine数量减1.同Add(-1)
}
func fun2()  {
	defer wg.Done()
	for j:=1;j<=10;j++{
		fmt.Println("\tfun2..j,",j)
	}
}
```

运行结果：

```
$ go run saleticket.go
fun1.。。i: 1
fun1.。。i: 2
fun1.。。i: 3
fun1.。。i: 4
fun1.。。i: 5
fun1.。。i: 6
fun1.。。i: 7
fun1.。。i: 8
fun1.。。i: 9
fun1.。。i: 10
main进入阻塞状态。。。等待wg中的子goroutine结束。。
	fun2..j, 1
	fun2..j, 2
	fun2..j, 3
	fun2..j, 4
	fun2..j, 5
	fun2..j, 6
	fun2..j, 7
	fun2..j, 8
	fun2..j, 9
	fun2..j, 10
main，解除阻塞。。

Process finished with exit code 0

```



### 5.2 Mutex(互斥锁)

我们知道在并发程序中，会存在临界资源问题。就是当多个协程来访问共享的数据资源，那么这个共享资源是不安全的。为了解决协程同步的问题我们使用channel，但是Go语言也提供了传统的同步工具。

**什么是锁呢？就是某个协程（线程）在访问某个资源时先锁住，防止其它协程的访问，等访问完毕解锁后其他协程再来加锁进行访问。一般用于处理并发中的临界资源问题。**

**Go语言包中的 sync 包提供了两种锁类型：sync.Mutex 和 sync.RWMutex。**

Mutex 是最简单的一种锁类型，互斥锁，同时也比较暴力，当一个 goroutine 获得了 Mutex 后，其他 goroutine 就只能等到这个 goroutine 释放该 Mutex。

每个资源都对应于一个可称为 “互斥锁” 的标记，这个标记用来保证在任意时刻，只能有一个协程（线程）访问该资源。其它的协程只能等待。

互斥锁是传统并发编程对共享资源进行访问控制的主要手段，它由标准库sync中的Mutex结构体类型表示。sync.Mutex类型只有两个公开的指针方法，Lock和Unlock。Lock锁定当前的共享资源，Unlock进行解锁。

在使用互斥锁时，一定要注意：对资源操作完成后，一定要解锁，否则会出现流程执行异常，死锁等问题。通常借助defer。锁定后，立即使用defer语句保证互斥锁及时解锁。


![WX20190807-101436](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190808-092409.png)

部分源码：

```go
/ A Mutex is a mutual exclusion lock.
// The zero value for a Mutex is an unlocked mutex.
//
// A Mutex must not be copied after first use.
type Mutex struct {
	state int32 //互斥锁上锁状态枚举值如下所示
	sema  uint32 //信号量，向处于Gwaitting的G发送信号
}

// A Locker represents an object that can be locked and unlocked.
type Locker interface {
	Lock()
	Unlock()
}

const (
	mutexLocked = 1 << iota // mutex is locked  ，1 互斥锁是锁定的
	mutexWoken // 2 唤醒锁
	mutexStarving
	mutexWaiterShift = iota // 统计阻塞在这个互斥锁上的goroutine数目需要移位的数值
	starvationThresholdNs = 1e6
)

```



### 5.2.1 Lock()方法：

Lock()这个方法，锁定m。如果该锁已在使用中，则调用goroutine将阻塞，直到互斥锁可用。

![WX20190807-102137](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190808-104517.png)



### 5.2.2 Unlock()方法

Unlock()方法，解锁m。如果m未在要解锁的条目上锁定，则为运行时错误。

锁定的互斥体不与特定的goroutine关联。允许一个goroutine锁定互斥体，然后安排另一个goroutine解锁互斥体。

![WX20190807-102843](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190808-104744.png)



### 5.2.3 示例代码：

使用goroutine，模拟4个售票口出售火车票的案例。4个售票口同时卖票，会发生临界资源数据安全问题。使用互斥锁解决一下。(Go语言推崇的是使用Channel来实现数据共享，但是也还是提供了传统的同步处理方式)

示例代码：

```go
package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

//全局变量，表示票
var ticket = 10 //10张票


var mutex sync.Mutex //创建锁头

var wg sync.WaitGroup //同步等待组对象
func main() {
	/*
	4个goroutine，模拟4个售票口，

	在使用互斥锁的时候，对资源操作完，一定要解锁。否则会出现程序异常，死锁等问题。
	defer语句
	 */

	wg.Add(4)
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")

	wg.Wait() //main要等待
	fmt.Println("程序结束了。。。")

	//time.Sleep(5*time.Second)
}

func saleTickets(name string){
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for{
		//上锁
		mutex.Lock() //g2
		if ticket > 0{ //ticket 1 g1
			time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
			fmt.Println(name,"售出：",ticket) // 1
			ticket-- // 0
		}else{
			mutex.Unlock() //条件不满足，也要解锁
			fmt.Println(name,"售罄，没有票了。。")
			break
		}
		mutex.Unlock() //解锁
	}
}

```

运行结果：

```
售票口4 售出： 10
售票口4 售出： 9
售票口2 售出： 8
售票口1 售出： 7
售票口3 售出： 6
售票口4 售出： 5
售票口2 售出： 4
售票口1 售出： 3
售票口3 售出： 2
售票口4 售出： 1
售票口2 售罄，没有票了。。
售票口1 售罄，没有票了。。
售票口3 售罄，没有票了。。
售票口4 售罄，没有票了。。
程序结束了。。。

Process finished with exit code 0

```



### 5.3 RWMutex(读写锁)

通过对互斥锁的学习，我们已经知道了锁的概念以及用途。主要是用于处理并发中的临界资源问题。

Go语言包中的 sync 包提供了两种锁类型：sync.Mutex 和 sync.RWMutex。其中RWMutex是基于Mutex实现的，只读锁的实现使用类似引用计数器的功能。

RWMutex是读/写互斥锁。锁可以由任意数量的读取器或单个编写器持有。RWMutex的零值是未锁定的mutex。

如果一个goroutine持有一个rRWMutex进行读取，而另一个goroutine可能调用lock，那么在释放初始读取锁之前，任何goroutine都不应该期望能够获取读取锁。特别是，这禁止递归读取锁定。这是为了确保锁最终可用；被阻止的锁调用会将新的读卡器排除在获取锁之外。


![WX20190807-101436](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190808-160432.png)



我们怎么理解读写锁呢？当有一个 goroutine 获得写锁定，其它无论是读锁定还是写锁定都将阻塞直到写解锁；当有一个 goroutine 获得读锁定，其它读锁定仍然可以继续；当有一个或任意多个读锁定，写锁定将等待所有读锁定解锁之后才能够进行写锁定。所以说这里的读锁定（RLock）目的其实是告诉写锁定。我们可以将其总结为如下三条：

1. 同时只能有一个 goroutine 能够获得写锁定。
2. 同时可以有任意多个 gorouinte 获得读锁定。
3. 同时只能存在写锁定或读锁定（读和写互斥）。

所以，RWMutex这个读写锁，该锁可以加多个读锁或者一个写锁，其经常用于读次数远远多于写次数的场景。

读写锁的写锁只能锁定一次，解锁前不能多次锁定，读锁可以多次，但读解锁次数最多只能比读锁次数多一次，一般情况下我们不建议读解锁次数多余读锁次数。

基本遵循两大原则：

	1、可以随便读，多个goroutine同时读。
	
	2、写的时候，啥也不能干。不能读也不能写。

读写锁即是针对于读写操作的互斥锁。它与普通的互斥锁最大的不同就是，它可以分别针对读操作和写操作进行锁定和解锁操作。读写锁遵循的访问控制规则与互斥锁有所不同。在读写锁管辖的范围内，它允许任意个读操作的同时进行。但是在同一时刻，它只允许有一个写操作在进行。

并且在某一个写操作被进行的过程中，读操作的进行也是不被允许的。也就是说读写锁控制下的多个写操作之间都是互斥的，并且写操作与读操作之间也都是互斥的。但是，多个读操作之间却不存在互斥关系。

### 5.3.1 RLock()方法

```go
func (rw *RWMutex) RLock()
```

读锁，当有写锁时，无法加载读锁，当只有读锁或者没有锁时，可以加载读锁，读锁可以加载多个，所以适用于“读多写少”的场景。

![WX20190807-102843](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190809-101020.png)





### 5.3.2 RUnlock()方法

```go
func (rw *RWMutex) RUnlock()
```

读锁解锁，RUnlock 撤销单次RLock调用，它对于其它同时存在的读取器则没有效果。若rw并没有为读取而锁定，调用RUnlock就会引发一个运行时错误。



![WX20190807-102843](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190809-101051.png)



### 5.3.3 Lock()方法：

```go
func (rw *RWMutex) Lock()
```

写锁，如果在添加写锁之前已经有其他的读锁和写锁，则Lock就会阻塞直到该锁可用，为确保该锁最终可用，已阻塞的Lock调用会从获得的锁中排除新的读取锁，即写锁权限高于读锁，有写锁时优先进行写锁定。

![WX20190807-102137](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190809-100627.png)



### 5.3.4 Unlock()方法

```go
func (rw *RWMutex) Unlock()
```

写锁解锁，如果没有进行写锁定，则就会引起一个运行时错误。

![WX20190807-102843](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190809-100753.png)



### 5.3.5 示例代码：

示例代码：

```go
package main

import (
	"fmt"
	"sync"
	"time"
)


var rwMutex *sync.RWMutex
var wg *sync.WaitGroup
func main() {
	rwMutex = new(sync.RWMutex)
	wg = new (sync.WaitGroup)

	//wg.Add(2)
	//
	////多个同时读取
	//go readData(1)
	//go readData(2)

	wg.Add(3)
	go writeData(1)
	go readData(2)
	go writeData(3)

	wg.Wait()
	fmt.Println("main..over...")
}


func writeData(i int){
	defer wg.Done()
	fmt.Println(i,"开始写：write start。。")
	rwMutex.Lock()//写操作上锁
	fmt.Println(i,"正在写：writing。。。。")
	time.Sleep(3*time.Second)
	rwMutex.Unlock()
	fmt.Println(i,"写结束：write over。。")
}

func readData(i int) {
	defer wg.Done()

	fmt.Println(i, "开始读：read start。。")

	rwMutex.RLock() //读操作上锁
	fmt.Println(i,"正在读取数据：reading。。。")
	time.Sleep(3*time.Second)
	rwMutex.RUnlock() //读操作解锁
	fmt.Println(i,"读结束：read over。。。")
}
```

运行结果：

![WX20190807-103748](/Users/jinhuaiwang/Desktop/Golang/Day16-20(Go语言基础进阶)/img/WX20190809-112822.png)



```
GOROOT=/usr/local/go #gosetup
GOPATH=/Users/ruby/go #gosetup
/usr/local/go/bin/go build -i -o /private/var/folders/kt/nlhsnpgn6lgd_q16f8j83sbh0000gn/T/___go_build_demo07_rwmutex_go /Users/ruby/go/src/l_goroutine/demo07_rwmutex.go #gosetup
/private/var/folders/kt/nlhsnpgn6lgd_q16f8j83sbh0000gn/T/___go_build_demo07_rwmutex_go #gosetup
3 开始写：write start
3 正在写：writing
2 开始读：read start
1 开始写：write start
3 写结束：write over
2 正在读：reading
2 读结束：read over
1 正在写：writing
1 写结束：write over
main..over...

Process finished with exit code 0

```



最后概括：

1. 读锁不能阻塞读锁
2. 读锁需要阻塞写锁，直到所有读锁都释放
3. 写锁需要阻塞读锁，直到所有写锁都释放
4. 写锁需要阻塞写锁







# 模块管理



## 为什么需要依赖管理

最早的时候，Go所依赖的所有的第三方库都放在GOPATH这个目录下面。这就导致了同一个库只能保存一个版本的代码。如果不同的项目依赖同一个第三方的库的不同版本，应该怎么解决？

## godep

Go语言从v1.5开始开始引入`vendor`模式，如果项目目录下有vendor目录，那么go工具链会优先使用`vendor`内的包进行编译、测试等。

`godep`是一个通过vender模式实现的Go语言的第三方依赖管理工具，类似的还有由社区维护准官方包管理工具`dep`。

### 安装

执行以下命令安装`godep`工具。

```go
go get github.com/tools/godep
```

### 基本命令

安装好godep之后，在终端输入`godep`查看支持的所有命令。

```bash
godep save     将依赖项输出并复制到Godeps.json文件中
godep go       使用保存的依赖项运行go工具
godep get      下载并安装具有指定依赖项的包
godep path     打印依赖的GOPATH路径
godep restore  在GOPATH中拉取依赖的版本
godep update   更新选定的包或go版本
godep diff     显示当前和以前保存的依赖项集之间的差异
godep version  查看版本信息
```

使用`godep help [command]`可以看看具体命令的帮助信息。

### 使用godep

在项目目录下执行`godep save`命令，会在当前项目中创建`Godeps`和`vender`两个文件夹。

其中`Godeps`文件夹下有一个`Godeps.json`的文件，里面记录了项目所依赖的包信息。 `vender`文件夹下是项目依赖的包的源代码文件。

### vender机制

Go1.5版本之后开始支持，能够控制Go语言程序编译时依赖包搜索路径的优先级。

例如查找项目的某个依赖包，首先会在项目根目录下的`vender`文件夹中查找，如果没有找到就会去`$GOAPTH/src`目录下查找。

### godep开发流程

1. 保证程序能够正常编译
2. 执行`godep save`保存当前项目的所有第三方依赖的版本信息和代码
3. 提交Godeps目录和vender目录到代码库。
4. 如果要更新依赖的版本，可以直接修改`Godeps.json`文件中的对应项

## go module

`go module`是Go1.11版本之后官方推出的版本管理工具，并且从Go1.13版本开始，`go module`将是Go语言默认的依赖管理工具。

### GO Modules 相关属性

* 1个开关环境变量:GO111MODULE
* 5个辅助环境变量:GOPROXY,GONOPROXY,GOSUMDB,GONOSUMDB,GOPRIVATE
* 2个辅助概念: go module proxy和 go checksum database
* 2个主要文件: go.sum和go.mod

* 1个主要管理子命令:go.sum

* 内置在几乎所有其它子命令中:go get, go install, go list, go test, go run, go build...

* Global Caching:不同项目的相同模块版本只会在你的电脑上缓存一份

### GO111MODULE

要启用`go module`支持首先要设置环境变量`GO111MODULE`，通过它可以开启或关闭模块支持，它有三个可选值：`off`、`on`、`auto`，默认值是`auto`。

1. `GO111MODULE=off`禁用模块支持，编译时会从`GOPATH`和`vendor`文件夹中查找包。
2. `GO111MODULE=on`启用模块支持，编译时会忽略`GOPATH`和`vendor`文件夹，只根据 `go.mod`下载依赖。
3. `GO111MODULE=auto`，当项目在`$GOPATH/src`外且项目根目录有`go.mod`文件时，开启模块支持。

简单来说，设置`GO111MODULE=on`之后就可以使用`go module`了，以后就没有必要在GOPATH中创建项目了，并且还能够很好的管理项目依赖的第三方包信息。

使用 go module 管理依赖后会在项目根目录下生成两个文件`go.mod`和`go.sum`。

### GOPROXY

这个环境变量主要是用于设置 Go 模块代理，主要如下：

Go1.11之后设置GOPROXY命令为：

```bash
export GOPROXY=https://goproxy.cn
```

作用：用于使 Go 在后续拉取模块版本时能够脱离传统的 VCS 方式从镜像站点快速拉取。Go1.13之后它拥有一个默认：`https://proxy.golang.org,direct`，但很可惜 `proxy.golang.org` 在中国无法访问，故而建议使用 `goproxy.cn` 作为替代，可以执行语句：

```bash
go env -w GOPROXY=https://goproxy.cn,direct
```

**注意:**

​		值列表中的 “direct” 为特殊指示符，用于指示 Go 回源到模块版本的源地址去抓取(比如 GitHub 等)，当值列表中上一个 Go module proxy 返回 404 或 410 错误时，Go 自动尝试列表中的下一个，遇见 “direct” 时回源，遇见 EOF 时终止并抛出类似 “invalid version: unknown revision...” 的错误。



设置为 “off” ：禁止 Go 在后续操作中使用任 何 Go module proxy。



### GOSUMDB



它的值是一个 Go checksum database，用于使 Go 在拉取模块版本时(无论是从源站拉取还是通过 Go module proxy 拉取)保证拉取到的模块版本数据未经篡改，也可以是“off”即禁止 Go 在后续操作中校验模块版本

- 格式 1：`<SUMDB_NAME>+<PUBLIC_KEY>`。
- 格式 2：`<SUMDB_NAME>+<PUBLIC_KEY> <SUMDB_URL>`。
- 拥有默认值：`sum.golang.org` (之所以没有按照上面的格式是因为 Go 对默认值做了特殊处理)。
- 可被 Go module proxy 代理 (详见：Proxying a Checksum Database)。
- `sum.golang.org` 在中国无法访问，故而更加建议将 GOPROXY 设置为 `goproxy.cn`，因为 `goproxy.cn` 支持代理 `sum.golang.org`。



### Global Caching

这个主要是针对 Go modules 的全局缓存数据说明，如下：

- 同一个模块版本的数据只缓存一份，所有其他模块共享使用。
- 目前所有模块版本数据均缓存在 `$GOPATH/pkg/mod`和 `$GOPATH/pkg/sum` 下，未来或将移至 `$GOCACHE/mod `和`$GOCACHE/sum` 下( 可能会在当 `$GOPATH` 被淘汰后)。
- 可以使用 `go clean -modcache` 清理所有已缓存的模块版本数据。

另外在 Go1.11 之后 GOCACHE 已经不允许设置为 off 了，这也是为了模块数据缓存移动位置做准备，因此大家应该尽快做好适配。



### go mod命令

常用的`go mod`命令如下：

```
go mod download    下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
go mod edit        编辑go.mod文件
go mod graph       打印模块依赖图
go mod init        初始化当前文件夹, 创建go.mod文件
go mod tidy        增加缺少的module，删除无用的module
go mod vendor      将依赖复制到vendor下
go mod verify      校验依赖
go mod why         解释为什么需要依赖
```

### go.mod

go.mod文件记录了项目所有的依赖信息，其结构大致如下：

```go
module github.com/Q1mi/studygo/blogger

go 1.12

require (
	github.com/DeanThompson/ginpprof v0.0.0-20190408063150-3be636683586
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/satori/go.uuid v1.2.0
	google.golang.org/appengine v1.6.1 // indirect
)

exclude gopkg.in/ini.v1 v1.60.2
//在国内访问golang.org/x的各个包都需要翻墙，你可以在go.mod中使用replace替换成github上对应的库。
replace (
	golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
	golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)

```

go.mod 是启用了 Go moduels 的项目所必须的最重要的文件，它描述了当前项目（也就是当前模块）的元信息，每一行都以一个动词开头，目前有以下 5 个动词:

- module：用于定义当前项目的模块路径。
- go：用于设置预期的 Go 版本。
- require：用于设置一个特定的模块版本。
- exclude：用于从使用中排除一个特定的模块版本。
- replace：用于将一个模块版本替换为另外一个模块版本。

- `indirect`表示间接引用

### go.sum

go.sum 是类似于比如 dep 的 Gopkg.lock 的一类文件，它详细罗列了当前项目直接或间接依赖的所有模块版本，并写明了那些模块版本的 SHA-256 哈希值以备 Go 在今后的操作中保证项目所依赖的那些模块版本不会被篡改。

```
github.com/gin-gonic/gin v1.6.3 h1:ahKqKTFpO5KTPHxWZjEdPScmYaGtLo8Y4DMHoEsnp14=
github.com/gin-gonic/gin v1.6.3/go.mod h1:75u5sXoLsGZoRN5Sgbi1eraJ4GU3++wFwWzhwvtwp4M=
github.com/go-playground/assert/v2 v2.0.1/go.mod h1:VDjEfimB/XKnb+ZQfWdccd7VUvScMdVu0Titje2rxJ4=
```

我们可以看到一个模块路径可能有如下两种：

```
github.com/gin-gonic/gin v1.6.3 h1:ahKqKTFpO5KTPHxWZjEdPScmYaGtLo8Y4DMHoEsnp14=
github.com/gin-gonic/gin v1.6.3/go.mod h1:75u5sXoLsGZoRN5Sgbi1eraJ4GU3++wFwWzhwvtwp4M=
```

前者为 Go modules 打包整个模块包文件 zip 后再进行 hash 值，而后者为针对 go.mod 的 hash 值。他们两者，要不就是同时存在，要不就是只存在 go.mod hash。

那什么情况下会不存在 zip hash 呢，就是当 Go 认为肯定用不到某个模块版本的时候就会省略它的 zip hash，就会出现不存在 zip hash，只存在 go.mod hash 的情况。



### go get

在项目中执行`go get`命令可以下载依赖包，并且还可以指定下载的版本。

1. 运行`go get -u`将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
2. 运行`go get -u=patch`将会升级到最新的修订版本
3. 运行`go get package@version`将会升级到指定的版本号version

如果下载所有依赖可以使用`go mod download`命令。



### 整理依赖

我们在代码中删除依赖代码后，相关的依赖库并不会在`go.mod`文件中自动移除。这种情况下我们可以使用`go mod tidy`命令更新`go.mod`中的依赖关系。



### go mod edit

#### 格式化

因为我们可以手动修改go.mod文件，所以有些时候需要格式化该文件。Go提供了一下命令：

```bash
go mod edit -fmt
```

#### 添加依赖项

```bash
go mod edit -require=golang.org/x/text
```

#### 移除依赖项

如果只是想修改`go.mod`文件中的内容，那么可以运行`go mod edit -droprequire=package path`，比如要在`go.mod`中移除`golang.org/x/text`包，可以使用如下命令：

```bash
go mod edit -droprequire=golang.org/x/text
```

关于`go mod edit`的更多用法可以通过`go help mod edit`查看。



- 可以使用 `go clean -modcache` 清理所有已缓存的模块版本数据。

另外在 Go1.11 之后 GOCACHE 已经不允许设置为 off 了，我想着这也是为了模块数据缓存移动位置做准备，因此大家应该尽快做好适配。



## 快速迁移项目至 Go Modules

- 第一步: 升级到 Go 1.13。
- 第二步: 让 GOPATH 从你的脑海中完全消失，早一步踏入未来。
  - 修改 GOBIN 路径（可选）：`go env -w GOBIN=$HOME/bin`。
  - 打开 Go modules：`go env -w GO111MODULE=on`。
  - 设置 GOPROXY：`go env -w GOPROXY=https://goproxy.cn,direct` # 在中国是必须的，因为它的默认值被墙了。
- 第三步(可选): 按照你喜欢的目录结构重新组织你的所有项目。
- 第四步: 在你项目的根目录下执行 `go mod init <OPTIONAL_MODULE_PATH>` 以生成 go.mod 文件。

## 迁移后 go get 行为的改变

- 用 `go help module-get` 和 `go help gopath-get`分别去了解 Go modules 启用和未启用两种状态下的 go get 的行为
- 用 `go get` 拉取新的依赖
  - 拉取最新的版本(优先择取 tag)：`go get golang.org/x/text@latest`
  - 拉取 `master` 分支的最新 commit：`go get golang.org/x/text@master`
  - 拉取 tag 为 v0.3.2 的 commit：`go get golang.org/x/text@v0.3.2`
  - 拉取 hash 为 342b231 的 commit，最终会被转换为 v0.3.2：`go get golang.org/x/text@342b2e`
  - 用 `go get -u` 更新现有的依赖
  - 用 `go mod download` 下载 go.mod 文件中指明的所有依赖
  - 用 `go mod tidy` 整理现有的依赖
  - 用 `go mod graph` 查看现有的依赖结构
  - 用 `go mod init` 生成 go.mod 文件 (Go 1.13 中唯一一个可以生成 go.mod 文件的子命令)
- 用 `go mod edit` 编辑 go.mod 文件
- 用 `go mod vendor` 导出现有的所有依赖 (事实上 Go modules 正在淡化 Vendor 的概念)
- 用 `go mod verify` 校验一个模块是否被篡改过

这里我们注意到有两点比较特别，分别是：

- 第一点：为什么 “拉取 hash 为 342b231 的 commit，最终会被转换为 v0.3.2” 呢。这是因为虽然我们设置了拉取 @342b2e commit，但是因为 Go modules 会与 tag 进行对比，若发现对应的 commit 与 tag 有关联，则进行转换。
- 第二点：为什么不建议使用 `go mod vendor`，因为 Go modules 正在淡化 Vendor 的概念，很有可能 Go2 就去掉了。



## 在项目中使用go module

### 既有项目

如果需要对一个已经存在的项目启用`go module`，可以按照以下步骤操作：

1. 在项目目录下执行`go mod init`，生成一个`go.mod`文件。
2. 执行`go get`，查找并记录当前项目的依赖，同时生成一个`go.sum`记录每个依赖库的版本和哈希值。

### 新项目

对于一个新创建的项目，我们可以在项目文件夹下按照以下步骤操作：

1. 执行`go mod init 项目名`命令，在当前项目文件夹下创建一个`go.mod`文件。
2. 手动编辑`go.mod`中的require依赖项或执行`go get`自动发现、维护依赖。



# Docker部署Go Web应用



本文介绍了如何使用`Docker`以及`Docker Compose`部署我们的 Go Web 程序。

## 为什么需要Docker？

> 使用docker的主要目标是容器化。也就是为你的应用程序提供一致的环境，而不依赖于它运行的主机。

想象一下你是否也会遇到下面这个场景，你在本地开发了你的应用程序，它很可能有很多的依赖环境或包，甚至对依赖的具体版本都有严格的要求，当开发过程完成后，你希望将应用程序部署到web服务器。这个时候你必须确保所有依赖项都安装正确并且版本也完全相同，否则应用程序可能会崩溃并无法运行。如果你想在另一个web服务器上也部署该应用程序，那么你必须从头开始重复这个过程。这种场景就是Docker发挥作用的地方。

对于运行我们应用程序的主机，不管是笔记本电脑还是web服务器，我们唯一需要做的就是运行一个docker容器平台。从以后，你就不需要担心你使用的是MacOS，Ubuntu，Arch还是其他。你只需定义一次应用，即可随时随地运行。

## Docker部署示例

### 准备代码

这里我先用一段使用`net/http`库编写的简单代码为例讲解如何使用Docker进行部署，后面再讲解稍微复杂一点的项目部署案例。

```go
//demo.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	server := &http.Server{
		Addr: ":8888",
	}
  fmt.Println("server startup...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}

func hello(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("hello world!"))
}
```

上面的代码通过`8888`端口对外提供服务，返回一个字符串响应：`hello world!`。

### 创建Docker镜像

> 镜像（image）包含运行应用程序所需的所有东西——代码或二进制文件、运行时、依赖项以及所需的任何其他文件系统对象。

或者简单地说，镜像（image）是定义应用程序及其运行所需的一切。

### 编写Dockerfile

要创建Docker镜像（image）必须在配置文件中指定步骤。这个文件默认我们通常称之为`Dockerfile`。（虽然这个文件名可以随意命名它，但最好还是使用默认的`Dockerfile`。）

现在我们开始编写`Dockerfile`，具体内容如下：

**注意：某些步骤不是唯一的，可以根据自己的需要修改诸如文件路径、最终可执行文件的名称等**

```dockerfile
FROM golang:alpine

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 容器内的工作目录：/build, 如该目录不存在，WORKDIR会帮你建立目录；
WORKDIR /build

# 将本地代码到拷贝到容器中的工作目录中
COPY .  /build

# 将代码编译成二进制可执行文件app
RUN go build -o app ./demo.go

# 移动到用于存放生成的二进制文件的 /dist 目录
WORKDIR /dist

# 将二进制文件从 /build 目录复制到这里
RUN cp /build/app .

# 暴露容器内部端口
EXPOSE 8888

# 启动容器时运行的命令
CMD ["/dist/app"]
```



### 构建镜像

在项目目录下，执行下面的命令创建镜像，并指定镜像名称为`goweb_app`：

```bash
docker build . -t goweb_app
```

等待构建过程结束，输出如下提示：

```bash
...
Successfully built 90d9283286b7
Successfully tagged goweb_app:latest
```

现在我们已经准备好了镜像，但是目前它什么也没做。我们接下来要做的是运行我们的镜像，以便它能够处理我们的请求。运行中的镜像称为容器。

执行下面的命令来运行镜像：

```bash
docker run --name go-web -d  -p 8888:8888  goweb_app
```

标志位`-p`用来定义端口绑定。由于容器中的应用程序在端口8888上运行，我们将其绑定到主机端口也是8888。如果要绑定到另一个端口，则可以使用`-p $HOST_PORT:8888`。例如`-p 5000:8888`。

现在就可以测试下我们的web程序是否工作正常，打开浏览器输入`http://127.0.0.1:8888`就能看到我们事先定义的响应内容如下：

```bash
hello world!
```



## 分阶段构建示例

我们的Go程序编译之后会得到一个可执行的二进制文件，其实在最终的镜像中是不需要go编译器的，也就是说我们只需要一个运行最终二进制文件的容器即可。

Docker的最佳实践之一是通过仅保留二进制文件来减小镜像大小，为此，我们将使用一种称为多阶段构建的技术，这意味着我们将通过多个步骤构建镜像。

```dockerfile
FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 容器内创建的工作目录：/build
WORKDIR /build

# 将本地代码到拷贝到容器中的工作目录中
COPY .  /build

# 将代码编译成二进制可执行文件 app
RUN go build -o app ./demo.go

###################
# 接下来创建一个小镜像
###################
FROM scratch

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/app /

# 需要运行的命令
ENTRYPOINT ["./app"]
```

使用这种技术，我们剥离了使用`golang:alpine`作为编译镜像来编译得到二进制可执行文件的过程，并基于`scratch`生成一个简单的、非常小的新镜像。我们将二进制文件从命名为`builder`的第一个镜像中复制到新创建的`scratch`镜像中。有关scratch镜像的更多信息，请查看https://hub.docker.com/_/scratch



## 附带其他文件的部署示例

项目的Github仓库地址为：https://github.com/Q1mi/bubble。

如果项目中带有静态文件或配置文件，需要将其拷贝到最终的镜像文件中。

我们的bubble项目用到了静态文件和配置文件，具体目录结构如下：

```bash
bubble
├── README.md
├── bubble
├── conf
│   └── config.ini
├── controller
│   └── controller.go
├── dao
│   └── mysql.go
├── example.png
├── go.mod
├── go.sum
├── main.go
├── models
│   └── todo.go
├── routers
│   └── routers.go
├── setting
│   └── setting.go
├── static
│   ├── css
│   │   ├── app.8eeeaf31.css
│   │   └── chunk-vendors.57db8905.css
│   ├── fonts
│   │   ├── element-icons.535877f5.woff
│   │   └── element-icons.732389de.ttf
│   └── js
│       ├── app.007f9690.js
│       └── chunk-vendors.ddcb6f91.js
└── templates
    ├── favicon.ico
    └── index.html
```

我们需要将`templates`、`static`、`conf`三个文件夹中的内容拷贝到最终的镜像文件中。更新后的`Dockerfile`如下

```dockerfile
FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 bubble
RUN go build -o bubble ./main.go

###################
# 接下来创建一个小镜像
###################
FROM scratch

COPY ./templates /templates
COPY ./static /static
COPY ./conf /conf

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/bubble /

# 需要运行的命令
ENTRYPOINT ["/bubble", "conf/config.ini"]
```

**Tips：** 这里把COPY静态文件的步骤放在上层，把COPY二进制可执行文件放在下层，争取多使用缓存。



### 关联其他容器

又因为我们的项目中使用了MySQL，我们可以选择使用如下命令启动一个MySQL容器，它的别名为`mysql8019`；root用户的密码为`root1234`；挂载容器中的`/var/lib/mysql`到本地的`/Users/q1mi/docker/mysql`目录；内部服务端口为3306，映射到外部的13306端口。

```bash
docker run --name mysql8019 -p 13306:3306 -e MYSQL_ROOT_PASSWORD=root1234 -v /Users/q1mi/docker/mysql:/var/lib/mysql -d mysql:8.0.19
```

这里需要修改一下我们程序中配置的MySQL的host地址为容器别名，使它们在内部通过别名（此处为mysql8019）联通。

```ini
[mysql]
user = root
password = root1234
host = mysql8019
port = 3306
db = bubble
```

修改后记得重新构建`bubble_app`镜像：

```bash
docker build . -t bubble_app
```

这里运行`bubble_app`容器的时候需要使用`--link`的方式与上面的`mysql8019`容器关联起来，具体命令如下：

```bash
docker run  -d --name=bubble_app  --link=mysql8019 -p 9000:9000 bubble_app
```



## Docker Compose模式

除了像上面一样使用`--link`的方式来关联两个容器之外，我们还可以使用`Docker Compose`来定义和运行多个容器。

`Compose`是用于定义和运行多容器 Docker 应用程序的工具。通过 Compose，你可以使用 YML 文件来配置应用程序需要的所有服务。然后，使用一个命令，就可以从 YML 文件配置中创建并启动所有服务。

使用Compose基本上是一个三步过程：

1. 使用`Dockerfile`定义你的应用环境以便可以在任何地方复制。
2. 定义组成应用程序的服务，`docker-compose.yml` 以便它们可以在隔离的环境中一起运行。
3. 执行 `docker-compose up`命令来启动并运行整个应用程序。

我们的项目需要两个容器分别运行`mysql`和`bubble_app`，我们编写的`docker-compose.yml`文件内容如下：

```yaml
# yaml 配置
version: "3.7"
services:
  mysql8019:
    image: "mysql:8.0.19"
    ports:
      - "33061:3306"
    command: "--default-authentication-plugin=mysql_native_password --init-file /data/application/init.sql"
    environment:
      MYSQL_ROOT_PASSWORD: "root1234"
      MYSQL_DATABASE: "bubble"
      MYSQL_PASSWORD: "root1234"
    volumes:
      - ./init.sql:/data/application/init.sql
  bubble_app:
    build: .
    command: sh -c "./wait-for.sh mysql8019:3306 -- ./bubble ./conf/config.ini"
    depends_on:
      - mysql8019
    ports:
      - "8888:8888"
```

这个 Compose 文件定义了两个服务：`bubble_app` 和 `mysql8019`。其中：

##### bubble_app

使用当前目录下的`Dockerfile`文件构建镜像，并通过`depends_on`指定依赖`mysql8019`服务，声明服务端口8888并绑定对外8888端口。

##### mysql8019

mysql8019 服务使用 Docker Hub 的公共 mysql:8.0.19 镜像，内部端口3306，外部端口33061。

这里需要注意一个问题就是，我们的`bubble_app`容器需要等待`mysql8019`容器正常启动之后再尝试启动，因为我们的web程序在启动的时候会初始化MySQL连接。这里共有两个地方要更改，第一个就是我们`Dockerfile`中要把最后一句注释掉：

```dockerfile
# Dockerfile
...
# 需要运行的命令（注释掉这一句，因为需要等MySQL启动之后再启动我们的Web程序）
# ENTRYPOINT ["/bubble", "conf/config.ini"]
```

第二个地方是在`bubble_app`下面添加如下命令，使用提前编写的`wait-for.sh`脚本检测`mysql8019:3306`正常后再执行后续启动Web应用程序的命令：

```bash
command: sh -c "./wait-for.sh mysql8019:3306 -- ./bubble ./conf/config.ini"
```

当然，因为我们现在要在`bubble_app`镜像中执行sh命令，所以不能在使用`scratch`镜像构建了，这里改为使用`debian:stretch-slim`，同时还要安装`wait-for.sh`脚本用到的`netcat`，最后不要忘了把`wait-for.sh`脚本文件COPY到最终的镜像中，并赋予可执行权限哦。更新后的`Dockerfile`内容如下：

```dockerfile
FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 bubble
RUN go build -o bubble .

###################
# 接下来创建一个小镜像
###################
FROM debian:stretch-slim

COPY ./wait-for.sh /
COPY ./templates /templates
COPY ./static /static
COPY ./conf /conf


# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/bubble /

RUN set -eux; \
	apt-get update; \
	apt-get install -y \
		--no-install-recommends \
		netcat; \
        chmod 755 wait-for.sh

# 需要运行的命令
# ENTRYPOINT ["/bubble", "conf/config.ini"]
```

所有的条件都准备就绪后，就可以执行下面的命令跑起来了：

```bash
docker-compose up
```

完整版代码示例，请查看我的github仓库：https://github.com/Q1mi/deploy_bubble_using_docker。



## Go如何跨平台交叉编译

Golang 支持交叉编译，在一个平台上生成另一个平台的可执行程序，最近使用了一下，非常好用。

Mac 下编译 Linux 和 Windows 64位可执行程序

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```


Linux 下编译 Mac 和 Windows 64位可执行程序

```
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

GOOS：目标平台的操作系统（darwin、freebsd、linux、windows）
GOARCH：目标平台的体系架构（386、amd64、arm）
交叉编译不支持 CGO 所以要禁用它

上面的命令编译 64 位可执行程序，你当然应该也会使用 386 编译 32 位可执行程序。


## 总结

使用Docker容器能够极大简化我们在配置依赖环境方面的操作，但同时也对我们的技术储备提了更高的要求。对于Docker不管你是熟悉抑或是不熟悉，技术发展的车轮都滚滚向前。



# 部署Go语言项目的 N 种方法

本文以部署 Go Web 程序为例，介绍了在 CentOS7 服务器上部署 Go 语言程序的若干方法。

## 独立部署

Go 语言支持跨平台交叉编译，也就是说我们可以在 Windows 或 Mac 平台下编写代码，并且将代码编译成能够在 Linux amd64 服务器上运行的程序。

对于简单的项目，通常我们只需要将编译后的二进制文件拷贝到服务器上，然后设置为后台守护进程运行即可。

### 编译

编译可以通过以下命令或编写 makefile 来操作。

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/bluebell
```

下面假设我们将本地编译好的 bluebell 二进制文件、配置文件和静态文件等上传到服务器的`/data/app/bluebell`目录下。

补充一点，如果嫌弃编译后的二进制文件太大，可以在编译的时候加上`-ldflags "-s -w"`参数去掉符号表和调试信息，一般能减小20%的大小。

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/bluebell
```

如果还是嫌大的话可以继续使用 upx 工具对二进制可执行文件进行压缩。

我们编译好 bluebell 项目后，相关必要文件的目录结构如下：

```bash
├── bin
│   └── bluebell
├── conf
│   └── config.yaml
├── static
│   ├── css
│   │   └── app.0afe9dae.css
│   ├── favicon.ico
│   ├── img
│   │   ├── avatar.7b0a9835.png
│   │   ├── iconfont.cdbe38a0.svg
│   │   ├── logo.da56125f.png
│   │   └── search.8e85063d.png
│   └── js
│       ├── app.9f3efa6d.js
│       ├── app.9f3efa6d.js.map
│       ├── chunk-vendors.57f9e9d6.js
│       └── chunk-vendors.57f9e9d6.js.map
└── templates
    └── index.html
```

### nohup

nohup 用于在系统后台**不挂断**地运行命令，不挂断指的是退出执行命令的终端也不会影响程序的运行。

我们可以使用 nohup 命令来运行应用程序，使其作为后台守护进程运行。由于在主流的 Linux 发行版中都会默认安装 nohup 命令工具，我们可以直接输入以下命令来启动我们的项目：

```bash
sudo nohup ./bin/bluebell conf/config.yaml > nohup_bluebell.log 2>&1 &
```

其中：

- `./bluebell conf/config.yaml`是我们应用程序的启动命令
- `nohup ... &`表示在后台不挂断的执行上述应用程序的启动命令
- `> nohup_bluebell.log`表示将命令的标准输出重定向到 nohup_bluebell.log 文件
- `2>&1`表示将标准错误输出也重定向到标准输出中，结合上一条就是把执行命令的输出都定向到 nohup_bluebell.log 文件

上面的命令执行后会返回进程 id

```bash
[1] 6338
```

当然我们也可以通过以下命令查看 bluebell 相关活动进程：

```bash
ps -ef | grep bluebell
```

输出：

```bash
root      6338  4048  0 08:43 pts/0    00:00:00 ./bin/bluebell conf/config.yaml
root      6376  4048  0 08:43 pts/0    00:00:00 grep --color=auto bluebell
```

此时就可以打开浏览器输入`http://服务器ip:端口`查看应用程序的展示效果了。



### supervisor

Supervisor 是业界流行的一个通用的进程管理程序，它能将一个普通的命令行进程变为后台守护进程，并监控该进程的运行状态，当该进程异常退出时能将其自动重启。

首先使用 yum 来安装 supervisor：

如果你还没有安装过 EPEL，可以通过运行下面的命令来完成安装，如果已安装则跳过此步骤：

```bash
sudo yum install epel-release
```

安装 supervisor

```bash
sudo yum install supervisor
```

Supervisor 的配置文件为：`/etc/supervisord.conf` ，Supervisor 所管理的应用的配置文件放在 `/etc/supervisord.d/` 目录中，这个目录可以在 supervisord.conf 中的`include`配置。

```bash
[include]
files = /etc/supervisord.d/*.conf
```

启动supervisor服务：

```bash
sudo supervisord -c /etc/supervisord.conf
```

我们在`/etc/supervisord.d`目录下创建一个名为`bluebell.conf`的配置文件，具体内容如下。

```bash
[program:bluebell]  ;程序名称
user=root  ;执行程序的用户
command=/data/app/bluebell/bin/bluebell /data/app/bluebell/conf/config.yaml  ;执行的命令
directory=/data/app/bluebell/ ;命令执行的目录
stopsignal=TERM  ;重启时发送的信号
autostart=true  
autorestart=true  ;是否自动重启
stdout_logfile=/var/log/bluebell-stdout.log  ;标准输出日志位置
stderr_logfile=/var/log/bluebell-stderr.log  ;标准错误日志位置
```

创建好配置文件之后，重启supervisor服务

```bash
sudo supervisorctl update # 更新配置文件并重启相关的程序
```

查看bluebell的运行状态：

```bash
sudo supervisorctl status bluebell
```

输出：

```bash
bluebell                         RUNNING   pid 10918, uptime 0:05:46
```

最后补充一下常用的supervisr管理命令：

```bash
supervisorctl status       # 查看所有任务状态
supervisorctl shutdown     # 关闭所有任务
supervisorctl start 程序名  # 启动任务
supervisorctl stop 程序名   # 关闭任务
supervisorctl reload       # 重启supervisor
```

接下来就是打开浏览器查看网站是否正常了。



## 搭配nginx部署

在需要静态文件分离、需要配置多个域名及证书、需要自建负载均衡层等稍复杂的场景下，我们一般需要搭配第三方的web服务器（Nginx、Apache）来部署我们的程序。

### 正向代理与反向代理

正向代理可以简单理解为客户端的代理，你访问墙外的网站用的那个属于正向代理。

![正向代理](https://www.liwenzhou.com/images/Go/deploy_go_app/image-20200920002334065.png)

反向代理可以简单理解为服务器的代理，通常说的 Nginx 和 Apache 就属于反向代理。

![反向代理](https://www.liwenzhou.com/images/Go/deploy_go_app/image-20200920002443846.png)

Nginx 是一个免费的、开源的、高性能的 HTTP 和反向代理服务，主要负责负载一些访问量比较大的站点。Nginx 可以作为一个独立的 Web 服务，也可以用来给 Apache 或是其他的 Web 服务做反向代理。相比于 Apache，Nginx 可以处理更多的并发连接，而且每个连接的内存占用的非常小。

### 使用yum安装nginx

EPEL 仓库中有 Nginx 的安装包。如果你还没有安装过 EPEL，可以通过运行下面的命令来完成安装：

```bash
sudo yum install epel-release
```

安装nginx

```bash
sudo yum install nginx
```

安装完成后，执行下面的命令设置Nginx开机启动：

```bash
sudo systemctl enable nginx
```

启动Nginx

```bash
sudo systemctl start nginx
```

查看Nginx运行状态：

```bash
sudo systemctl status nginx
```

### Nginx配置文件

通过上面的方法安装的 nginx，所有相关的配置文件都在 `/etc/nginx/` 目录中。Nginx 的主配置文件是 `/etc/nginx/nginx.conf`。

默认还有一个`nginx.conf.default`的配置文件示例，可以作为参考。你可以为多个服务创建不同的配置文件（建议为每个服务（域名）创建一个单独的配置文件），每一个独立的 Nginx 服务配置文件都必须以 `.conf`结尾，并存储在 `/etc/nginx/conf.d` 目录中。

### Nginx常用命令

补充几个 Nginx 常用命令。

```bash
nginx -s stop    # 停止 Nginx 服务
nginx -s reload  # 重新加载配置文件
nginx -s quit    # 平滑停止 Nginx 服务
nginx -t         # 测试配置文件是否正确
```

### Nginx反向代理部署

我们推荐使用 nginx 作为反向代理来部署我们的程序，按下面的内容修改 nginx 的配置文件。

```bash
worker_processes  1;

events {
    worker_connections  1024;
}

http {
    include       mime.types;
    default_type  application/octet-stream;

    sendfile        on;
    keepalive_timeout  65;

    server {
        listen       80;
        server_name  localhost;

        access_log   /var/log/bluebell-access.log;
        error_log    /var/log/bluebell-error.log;

        location / {
            proxy_pass                 http://127.0.0.1:8084;
            proxy_redirect             off;
            proxy_set_header           Host             $host;
            proxy_set_header           X-Real-IP        $remote_addr;
            proxy_set_header           X-Forwarded-For  $proxy_add_x_forwarded_for;
        }
    }
}
```

执行下面的命令检查配置文件语法：

```bash
nginx -t
```

执行下面的命令重新加载配置文件：

```bash
nginx -s reload
```

接下来就是打开浏览器查看网站是否正常了。

当然我们还可以使用 nginx 的 upstream 配置来添加多个服务器地址实现负载均衡。

```bash
worker_processes  1;

events {
    worker_connections  1024;
}

http {
    include       mime.types;
    default_type  application/octet-stream;

    sendfile        on;
    keepalive_timeout  65;

    upstream backend {
      server 127.0.0.1:8084;
      # 这里需要填真实可用的地址，默认轮询
      #server backend1.example.com;
      #server backend2.example.com;
    }

    server {
        listen       80;
        server_name  localhost;

        access_log   /var/log/bluebell-access.log;
        error_log    /var/log/bluebell-error.log;

        location / {
            proxy_pass                 http://backend/;
            proxy_redirect             off;
            proxy_set_header           Host             $host;
            proxy_set_header           X-Real-IP        $remote_addr;
            proxy_set_header           X-Forwarded-For  $proxy_add_x_forwarded_for;
        }
    }
}
```

### Nginx分离静态文件请求

上面的配置是简单的使用 nginx 作为反向代理处理所有的请求并转发给我们的 Go 程序处理，其实我们还可以有选择的将静态文件部分的请求直接使用 nginx 处理，而将 API 接口类的动态处理请求转发给后端的 Go 程序来处理。

![分离静态文件请求图示](https://www.liwenzhou.com/images/Go/deploy_go_app/image-20200920002735894.png)

下面继续修改我们的 nginx 的配置文件来实现上述功能。

```bash
worker_processes  1;

events {
    worker_connections  1024;
}

http {
    include       mime.types;
    default_type  application/octet-stream;

    sendfile        on;
    keepalive_timeout  65;

    server {
        listen       80;
        server_name  bluebell;

        access_log   /var/log/bluebell-access.log;
        error_log    /var/log/bluebell-error.log;

		# 静态文件请求
        location ~ .*\.(gif|jpg|jpeg|png|js|css|eot|ttf|woff|svg|otf)$ {
            access_log off;
            expires    1d;
            root       /data/app/bluebell;
        }

        # index.html页面请求
        # 因为是单页面应用这里使用 try_files 处理一下，避免刷新页面时出现404的问题
        location / {
            root /data/app/bluebell/templates;
            index index.html;
            try_files $uri $uri/ /index.html;
        }

		# API请求
        location /api {
            proxy_pass                 http://127.0.0.1:8084;
            proxy_redirect             off;
            proxy_set_header           Host             $host;
            proxy_set_header           X-Real-IP        $remote_addr;
            proxy_set_header           X-Forwarded-For  $proxy_add_x_forwarded_for;
        }
    }
}
```

### 前后端分开部署

前后端的代码没必要都部署到相同的服务器上，也可以分开部署到不同的服务器上，下图是前端服务将 API 请求转发至后端服务的方案。

![前后端分开部署方案1](https://www.liwenzhou.com/images/Go/deploy_go_app/image-20200920003753373.png)

上面的部署方案中，所有浏览器的请求都是直接访问前端服务，而如果是浏览器直接访问后端API服务的部署模式下，如下图。

此时前端和后端通常不在同一个域下，我们还需要在后端代码中添加跨域支持。

![前后端分开部署方案2](https://www.liwenzhou.com/images/Go/deploy_go_app/image-20200920003335577.png)

这里使用[github.com/gin-contrib/cors](https://github.com/gin-contrib/cors)库来支持跨域请求。

最简单的允许跨域的配置是使用`cors.Default()`，它默认允许所有跨域请求。

```go
func main() {
	router := gin.Default()
	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	router.Use(cors.Default())
	router.Run()
}
```

此外，还可以使用`cors.Config`自定义具体的跨域请求相关配置项：

```go
package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.Run()
}
```

## 容器部署

容器部署方案可参照我之前的博客：[使用Docker和Docker Compose部署Go Web应用](https://www.liwenzhou.com/posts/Go/how_to_deploy_go_app_using_docker/)，这里就不再赘述了。
