package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

}

/*------------------------------------------------------------------*/
/*
demo1 方法 （methods）
- 方法只是一个带 接收者 (Receiver) 的函数。
- 需要重点理解 值接收者 和 指针接收者 的区别。
*/

type Vertex struct { // vertex：点，顶点，端点
	X, Y float64
}

// 1. 值接收者 (Value Receiver)
// 语法：func (接收者变量 接收者类型) 方法名(参数列表) 返回值列表
// 特点：调用时会复制一份 Vertex，方法内修改不会影响原对象
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 2. 指针接收者 (Pointer Receiver) -- 推荐常用
// 特点：直接操作内存地址，方法内修改会影响原对象，且避免大对象复制，效率更高
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func demo1() {
	v := Vertex{3, 4}

	// 调用值接收者方法
	fmt.Println("Before Scale:", v.Abs())

	// 调用指针接收者方法
	// Go 语法糖：虽然 v 是值，但 Go 会自动解释为 (&v).Scale(10)
	v.Scale(10)

	fmt.Println("After Scale:", v.Abs()) // 50 (原对象被修改了)
}

/*------------------------------------------------------------------*/
/*
demo2 接口
Go 的接口是隐式实现的。
只要一个类型实现了接口要求的所有方法，它就自动实现了该接口，不需要 implements 关键字
*/

// 定义一个接口：只要能“叫(Speak)”，就是一种 Animal
type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

// Dog 实现了 Speak 方法 -> Dog 自动实现了 Animal 接口
func (d Dog) Speak() string {
	return "Woof"
}

type Cat struct {
	Name string
}

// Cat 也实现了 Speak 方法
func (c Cat) Speak() string {
	return "Meow!"
}

// 空接口 interface{}：没有任何方法要求，所以任何类型都实现了空接口
// 类似于 Java 的 Object 或 C++ 的 void*
/*
它是一个特殊的接口，没有要求实现任何方法；
Go 里的所有类型（int、string、数组、结构体、指针等）都「自动实现」了这个空接口；
因此，任何类型的变量都可以赋值给 interface{} 类型的参数 / 变量。
你可以把 interface{} 理解成一个「万能容器」—— 不管你往里面装什么类型的东西，它都能装下，还能记住这个东西原本的类型和值。
*/
func printAny(v interface{}) {
	fmt.Println("Type %T, Value:%v\n", v, v)
}

func demo2() {
	// 多态的使用
	var animals []Animal
	animals = append(animals, Dog{"Buddy"})
	animals = append(animals, Cat{"Kitty"})

	for _, a := range animals {
		fmt.Println(a.Speak())
	}

	// 空接口的使用
	printAny(100)
	printAny("hello")
	printAny(1.09)
	printAny(Dog{"旺旺"})
}

/*------------------------------------------------------------------*/
/*
demo3 Go语言类型转换
类型转换用于将一种数据类型的变量转换为另外一种类型的变量。
Go 语言类型转换基本格式如下：
type_name(expression)

将整型转换为浮点型：
var a int = 10
var b float64 = float64(a)

字符串类型转换
将一个字符串转换成另一个类型，可以使用以下语法：
var str string = "10"
var num int
num, _ = strconv.Atoi(str)
注意，strconv.Atoi 函数返回两个值，第一个是转换后的整型值，第二个是可能发生的错误，我们可以使用空白标识符 _ 来忽略这个错误
*/

func demo3() {
	// 1.字符串转换为整数
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换错误：", err)
	} else {
		fmt.Println("字符串 '%s' 转换为整数为：%d\n", str, num)
	}

	// 2.整数转换为字符串
	num2 := 123
	str2 := strconv.Itoa(num2)
	fmt.Println("整数 %d  转换为字符串为：'%s'\n", num2, str2)

	// 3.字符串转换为浮点数
	str3 := "3.14"
	num3, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转为浮点型为：%f\n", str3, num3)
	}

	// 4.浮点数转换为字符串
	num4 := 3.14
	str4 := strconv.FormatFloat(num4, 'f', 2, 64) //
	fmt.Printf("浮点数 %f 转为字符串为：'%s'\n", num4, str4)
}

/*
demo4 接口类型转换
接口类型转换有两种情况：类型断言和类型转换。

1.类型断言
类型断言用于将接口类型转换为指定类型，其语法为：
value.(type)  或者  value.(T)
- 其中 value 是接口类型的变量，type 或 T 是要转换成的类型。
- 如果类型断言成功，它将返回转换后的值和一个布尔值，表示转换是否成功。


2.类型转换
类型转换用于将一个接口类型的值转换为另一个接口类型，其语法为：
T(value)
T 是目标接口类型，value 是要转换的值。
在类型转换中，我们必须保证要转换的值和目标接口类型之间是兼容的，否则编译器会报错。
*/

type Writer interface {
	Write([]byte) (int, error)
}

// 实现 Writer 接口的结构体 StringWriter
type StringWriter struct {
	str string
}

// 实现 Write 方法
func (sw *StringWriter) Write(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}

func demo4() {
	var i interface{} = "hello world"
	str, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}
	/*
		以上实例中，我们定义了一个接口类型变量 i，并将它赋值为字符串 "Hello, World"。然后，我们使用类型断言将 i 转换为字符串类型，并将转换后的值赋值给变量 str。
		最后，我们使用 ok 变量检查类型转换是否成功，如果成功，我们打印转换后的字符串；否则，我们打印转换失败的消息。
	*/

	// 创建一个 StringWriter 实例并赋值给 Writer 接口变量
	var w Writer = &StringWriter{}

	// 将 Writer 接口类型转换为 StringWriter 类型
	sw := w.(*StringWriter)

	// 修改 StringWriter 的字段
	sw.str = "Hello, World"

	// 打印 StringWriter 的字段值
	fmt.Println(sw.str)
}
