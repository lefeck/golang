# golang 基础练习



1. **整形转字符串类型**

```go
 // 输入一个整数，返回一个字符串
func StringtoInt(n int) string {
	a := strconv.Itoa(n)
	fmt.Printf( "type is :%T, value %v\n",a, a)
	return a
}
```

formatInt 函数将输入的 int64 类型整形转换为指定进制的字符串。格式如下所示：

```go
FormatInt(i int64, base int)
// 以下函数将i转为二进制后返回它的字符串类型
FormatInt(i int64, 2)
```

**Int64 转换成10进制**

```go
// 输入int64类型的1234，设置为10进制
func FormattoString(n int) string  {
	return  strconv.FormatInt(int64(n),10)
}
```

2. **整形转数组**

```go
/*
Go 语言中 strconv 包的 itoa 函数输入一个 int 类型，返回转换后的字符串。而 strings 包下的 Split 函数输入一个字符串和分隔符，返回分割后的切片
 */
func Int(n int) []int  {
	//使用strconv.Itoa来转换给定的数字为字符串
	s := strconv.Itoa(n)
	//使用strings.Split()来分割转换后的字符串，并返回一个字符串切片。
	d := make([]int,len(s))
	for i, l :=range strings.Split(s,"") {
		d[i],_ =strconv.Atoi(l)
	}
	return d
}
```

3. **连接集合中的所有元素**

```go
/*
遍历整形切片的每个元素，然后把每个整形转为字符后添加分隔符后加入到字符串中.
 */
func JoinStr(params ...interface{}) string  {
	//使用relect.ValueOf()来获取整个slice，还有字符串。 第2个参数: 是要传入的string类型的分隔符到字符串中
	arr, sp := reflect.ValueOf(params[0]),reflect.ValueOf(params[1]).String()
	fmt.Println(arr, sp)
	// 用来创建一个合适长度的字符串切片。
	ars :=make([]string,arr.Len())
	//使用 for 循环遍历每个元素，fmt.Sprintf()用来整形转字符类型。
	for i :=0; i<arr.Len();i++ {
		ars[i] = fmt.Sprintf("%v",arr.Index(i))
	}
	//strings.Join()的功能是使用提供的分隔符来组合字符串
	return strings.Join(ars,sp)
}
```

4. **两个集合的所有组合**

```go
/*
为了获得两个集合的所有集合，我们可以先创建一个二维切片列表，然后 for 循环输出并且保存到二维切片列表中。
 */

func ConSlice(params ...interface{}) [][]interface{}  {
	//使用reflect.ValueOf()来获取切片或者数组。
	a,b := reflect.ValueOf(params[0]),reflect.ValueOf(params[1])
	l := a.Len() * b.Len()
	//初始化二维切片长度
	r := make([][]interface{},l)
	for i:=0; i<l; i++ {
		//把a，b 取到的值作为slice，放到新的切片当中。
		r[i] = []interface{}{
			//获取a中对应的value
			a.Index(i % a.Len()).Interface(),
			//获取b中对应的value
			b.Index(i / a.Len() % b.Len()).Interface(),
		}
	}
	return r
}
```

5. **删除切片重复元素**

```go
/*
删除重复切片的一个方法是创建一个字典来存储字符是否已经出现过。
 */

func DedupeInts(arr []int) []int {
	m, uniq := make(map[int]bool), make([]int,0)
	//for 循环遍历传入切片的每一个元素，并进行判断是否在 map 中重复，如果不重复则加入到 map 中。
	for _,v :=range arr {
		if value, ok := m[v]; !ok {
			m[v]=true
			uniq =append(uniq,v)
			fmt.Println(value)
		}
	}
	return uniq
}
```

6. **返回带有索引的切片**

```go
/*
我们可以使用 for 循环来遍历输入的参数，然后把索引和相关的值保存到 map 中。
 */
func WithIndex(params ...interface{}) map[int]interface{}  {
	//使用reflect.ValueOf()来获取切片或者数组。
	arr := reflect.ValueOf(params[0])
	m := make(map[int]interface{})
	//使用 for 循环来遍历数组，通过索引获取到每个元素添加到字典
	for i:=0;i<arr.Len();i++ {
		m[i] = arr.Index(i).Interface()
	}
	return m
}

func main()  {
	r := WithIndex([]int{4,5,7,8})
	fmt.Println(r)

	d := WithIndex([]string{"hello","merry","creck"})
	fmt.Println(d)
}
```

7. **复制指定条件的切片**

```go
/*
Go 中可以把函数作为参数传入，所有我们可以对传入的参数进行遍历，每个元素通过传入的函数来进行判断.
 */

//创建一个条件函数，来作为参数传入主功能函数中
//int
func FilterInt(arr []int, f func(int)bool) []int {
   it:=make([]int,0)
   //For循环遍历每个元素，使用传入的函数进行判断
   for _,v:=range arr {
      if f(v) {
         it = append(it,v)
      }
   }
   return it
}

//float64
func FilterFloat64(arr []float64, f func(float64)bool) []float64  {
   ft := make([]float64,0)
   for _, v :=range arr {
      ft = append(ft,v)
   }
   return ft
}
//string
func FilterString(arr []string, f func(string)bool) []string  {
   fs := make([]string,0)
   for _, v :=range arr {
      fs = append(fs,v)
   }
   return fs
}
func main()  {
	intCheck := func(x int) bool {return x > 1 }
	p := FilterInt([]int{1,2,3,4,5,6},intCheck)
	fmt.Println(p)
}

```

8. **返回集合中符合条件的第一个元素**

```go
/*
我们可以使用 for 循环来遍历输入的集合，然后通过传入的判断函数来进行判断，
 */

//使用 range 来正向遍历切片，用 f 函数来判断是否符合条件。
func FilterIndex(arr []int, f func(x int)bool) int  {
    //var v int = 0
	for _,v :=range arr {
		if f(v) {
		    return v
			//return i
		}
	}
	return 1
}

func main()  {
	It := func(x int) bool {return x%2==0}
	f := FilterIndex([]int{3,5,6,2},It)
	fmt.Println(f)
}
```

9. **返回集合中满足指定条件的最后一个元素**

```go
/*
我们可以使用 for 循环倒序遍历来进行查看，因为 Go 中参数可以为函数，所有我们使用判断函数来作为判断的参数传入主函数中

使用 for 循环逆序遍历给定的集合
使用函数 f 来判断是否是符合条件的元素
如果符合条件，则输出元素和 nil，否则输出 0 和错误信息
使用fmt.Error()来生成错误

*/
func FilterLastValue(arr []int,f func(x int)bool) (int,error)  {
	for i:= len(arr)-1; i>=0; i-- {
		if f(arr[i]) {
			return arr[i],nil
		}
	}
	return 0,fmt.Errorf("No matchs found")
}

func main()  {
	Id := func(x int) bool {return x%2==0}
	d,err := FilterLastValue([]int{3,5,10,19},Id)
	fmt.Println(d,err)
}
```







20. 是否包含空格

```go
/*
对于判断给定的字符串是否包含空格，我们可以使用 regexp 包下的函数来实现，下面是一个例子。
 */
func StringEmpty(str string) bool  {
	//使用 regexp.MustCompile 来安全初始化正则表达式。
	re := regexp.MustCompile(`\s`)
	//使用 Regexp.MatchString()函数来判断给定的字符串是否包含空格。
	return re.MatchString(str)
}
```

21. 替换指定范围的字符

```go
/*

将使用 Repeat 函数来用指定的掩码字符替换最后 n 个字符以外的所有字符。
*/

//cc 输入的字符串, 被替换的n个字符  m替换后的字符
func Mask(cc string, n int, m rune)string  {
	//使用 strings.Repeat()来对函数进行替换。
	return strings.Repeat(string(m),len(cc)-n)+ cc[len(cc)-n:]
}
```

22. 反转字符串

```go
/*
反转字符串我们可以使用 for 循环加切片来实现，下面是一个例子。
*/

func ReverString(s string) string {
	//使用 make 来生成一个 rune 类型的切片。
	o := make([]rune,len(s))
	//使用 range 和 len 函数来倒序遍历字符串然后将值添加到结果中。
	//使用下标方式索引字符串s，得到的是s[i]的字节值。可以先把字符串转为rune数组,再通过索引s来读取，s[0]。
	for i,v :=range s {
		fmt.Println(i)
		o[len(s)-i-1]= v
	}
	//使用 string()来将 rune 切片转为字符串。
	return  string(o)
}
```

23. 使用 map 来创建一个 set 类型

```go
/*
使用 map 来创建一个 set 类型
 */

func Set( set map[string]struct{}) map[string]struct{} {
	//s := make(map[string]struct{})
	set["a"] = struct{}{}
	set["b"] = struct{}{}
	set["c"] = struct{}{}
	fmt.Println(set)

	if _, ok := set["a"]; ok {
		fmt.Println("exist in set")
	}
	delete(set,"a")
	fmt.Println(set)
	return set
}

func main()  {
	s := make(map[string]struct{})
	Set(s)
}
```



24. 度转换为弧度

```go
/*
使用 math.Pi 常量和转换公式来将度转换为弧度
πx/180=x°
 */

func Rads(d float64)float64  {
	return d * math.Pi/180
}
```

25. 限制数

```go
/*
25.限制数指的是如果输入的数在给定的范围内，则返回该数，否则返回离该数范围最近的范围数。
*/

//在函数Clamp(n, a, b float64)中如果 n 在 a-b 的范围内，则返回 n。否则返回 ab 中最接近 n 的值。
func Clamp(n,a,b float64)float64  {
	return math.Max(math.Min(n,math.Max(a,b)),math.Min(a,b))
}

func main()  {
	c := Clamp(4.0,3.0,5.0)
	c1 := Clamp(1.0,3.0,5.0)
	c2 := Clamp(7.0,3.0,5.0)
	fmt.Println(c,c1,c2)
}
```

26. 数字是否为二的幂

```go
//使用按位二进制 AND 运算符(&)来确定 n 是否为 2 的幂。此外检查 n 是否为零。
/*
0011  0100
0010  0011
0010  0000
 */
func IsPowerOf2(n int)bool  {
	return n > 0 && n&(n-1)==0
}
```

27. 数是否为奇数

```go
//当给定数字与 2 的模为 1 时，就说该数字为奇数
func IsOdd(n int)bool  {
	//找使用模运算符%来判断给定的数字能否得到整除
	if n < 0 {
		//转换成正数
		n *=-1
	}
	return n%2==1
}
```

28. 华氏度之间相互摄氏度

```go
/*
华氏转摄氏度
*/
func FahrenheitToCelsius(n float64) float64  {
	return (n - 32) / 1.8
}
/*
摄氏转华氏度
*/
func CelsiusToFahrenheit(n float64) float64  {
	return  32+ n * 1.8
}
```

30. 截断字符串

```go
/*

判断字符串长度是否大于给定的范围，然后判断范围是否大于 3，然后进行裁剪字符串并返回。
 */

func TruncateString(s string,l int) string  {
	r := s
	if len(s) > l {
		if l > 3 {
			l -=3
		}
		r = s[0:l]+ "..."
	}
	return r
}
```

31. 字符串是否为大写字符串

```go
/*
判断字符串是否为大写字符串
 */
func IsUpper(s string)bool  {
	//if strings.ToUpper(s)==s {
	//	return true
	//}
	//return false
	//使用strings.ToUpper()将字符串转换为大写字符串并且与原先的字符串进行比较
	return strings.ToUpper(s)==s
}
```

32. 函数来将给定字符串的第一个字母转换为小写

```GO
//函数来将给定字符串的第一个字母转换为小写
func IsLower(s string)string  {
	return strings.ToLower(s[0:1])+s[1:]
}

func test(s string)string  {
	return s[0:] //hello world
	//return s[0:1]+s[2:3] //hl
	//return s[1:] //ello world

}
```

33. 求slice的和

```go
/*
求slice的和
 */
func Sum(nums []float64) float64  {
	sum := float64(0)
	for _,v :=range nums {
		sum +=v
	}
	return sum
}
```

34. 判断一个数是否在给定的范围

```go
/*
判断一个数是否在给定的范围
 */

func IsRange(n,a,b float64)bool  {
	//n 大于最小的，小于最大的，即是在给定去范围之内
	if n < math.Max(a,b) && n>=  math.Min(a,b) {
		return true
	}
	return false
}
```

35. 给定的字符串是否为小写字符串

```go
/*
给定的字符串是否为小写字符串
 */

func IsLowers(s string)bool  {
	return strings.ToLower(s) ==s
}
```

36.输出今天星期几，并且判断是否今天是星期三。

```go
/*
输出今天星期几，并且判断是否今天是星期三
 */

func Today()  {
	now := time.Now()
	fmt.Println("today is:", now.Weekday())
	if now.Weekday()== time.Tuesday {
		fmt.Println("today is Wednesday")
	}
}
```



37. 输出更好效果json

```go
/*
37.输出 json
*/

type Student struct {
	Name    string
	Age     int
	Lessons []string
}
func main() {
    s := Student{
        Name: "John",
        Age:  "17",
        Lessons: []string{
            "Mathematics",
            "Computer science",
            "Philosophy",
        },
    }

    jsonBytes, err := json.Marshal(s)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("\nUgly print:\n%s\n", jsonBytes)

    jsonBytes, err = json.MarshalIndent(s, "", "\t")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("\nPretty print:\n%s\n", jsonBytes)
}
```

39. 字符的 sha256 校验

```go
/*
39.字符的 sha256 校验
 */
func Sha256CheckSum(s string) string  {
	//sha256.Sum256()来获得字符的 sha256 校验。
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}


```

40. 字符的 md5 校验

```go
/*
 同上字符的 sha1 校验
 */

func md5Checksum(s string)string  {
	return fmt.Sprintf("%x",sha1.Sum([]byte(s)))
	//return fmt.Sprintf("%x",md5.Sum([]byte(s)))
}
```

41.翻转slice

```go
/*
Golang中数组反转的实现方式
 */
//方法一：
func ReverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i,j:=0,size-1; i<j; i,j = i+1,j-1 {
		swap(i,j)
	}
	fmt.Println(s)
	return
}
//方式二：
func ReverseSlice1(p []int)  {
	for i,j :=0, len(p)-1; i<j; i,j = i+1,j-1 {
		p[i],p[j] = p[j],p[i]
	}
}
//方式三：
func ReverseSlice2(s []interface{})  {
	sort.SliceStable(s, func(i, j int) bool {
		return true
	})
}
```

42. 删除切片指定元素

```go
/*
删除切片指定元素
 */
func DeleteSlice(s []int ,i int)([]int, error)  {

	// Determine whether the length of the slice is zero
	if len(s)== 0 {
		return nil,errors.New("Cannot delete an element from a nil or empty slice")
	}
	// Determine whether the given value is less than zero,  more than the length of the slice, return error
	if i < 0 || i > len(s) -1 {
		return nil, errors.New("Index out of bounds")
	}
	//
	fmt.Println(s[:i], s[i+1:])

	return append(s[:i],s[i+1:]...),nil
}
```

43. 切片分割成块

```go
//切片分割成块
func SplitSliceInChunks(a []int, chuckSize int,) ([][]int,error)  {
	// determine where the chucksize size less than 1, return error
	if chuckSize<1 {
		return nil, errors.New("chuckSize must be greater that zero")
	}
	chunks := make([][]int, 0, (len(a)+chuckSize-1)/chuckSize)
	//遍历给定的切片，把指定长度的切片添加到新建的切片中。
	for chuckSize < len(a) {
		a, chunks = a[chuckSize:],append(chunks,a[0:chuckSize])
	}
	chunks = append(chunks,a)
	return chunks,nil
}

```

