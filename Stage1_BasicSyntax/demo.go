package main

import "fmt"

/*
全局变量是允许声明但不使用的。 同一类型的多个变量可以声明在同一行，如：
var a, b, c int
多变量可以在同一行进行赋值，如：
var a, b int
var c string
a, b, c = 5, 7, "abc"
*/

//这种因式分解关键字的写法一般用于声明全局变量
var (
	global_val int
)

/*----------------------------------------------------------------------------------------*/

/*
demo1 Go语言数据类型

Go 语言的基础组成有以下几个部分：
1 包声明
2 引入包
3 函数
4 变量
5 语句 & 表达式
6 注释
*/
func demo1() {
	/*this is my first program*/
	fmt.Println("hello go")
	fmt.Print("hello go\n")
	//Printf 是可以控制格式化的输入输出的，若没有传入占位符，它就会跟Print用法变得一样
	fmt.Printf("hello %s", "go") // 输出：hello go
	fmt.Printf("hello go")
}

/*
关于demo1的说明
1 必须在源文件的非注释的第一行指明这个文件属于哪个包，如:package main表示一个可独立执行的程序，每一个go应用程序都应该包含一个名为main的package
2 import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
3 func main() 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
4 fmt.Println(...) 可以将字符串输出到控制台，并在最后自动增加换行字符 \n,使用 fmt.Print("hello, world\n") 可以得到相同的结果。Print 和 Println 这两个函数也支持使用变量，如：fmt.Println(arr)。如果没有特别指定，它们会以默认的打印格式将变量 arr 输出到控制台。
5 Go 语言通过标识符（变量、函数、结构体等）的大小写直接控制可见性：大写开头的标识符（如 fmt.Println 中的 Println）属于对外导出（类似 public），其他包导入后即可调用，就像公司的对外公开产品；小写开头的标识符仅在当前包内可见（类似 private），外部包无法访问，就像公司的内部机密工具。
*/

/*----------------------------------------------------------------------------------------*/

/*
demo2 Go语言基础语法

编译器是按照标记拆分的
Go 程序可以由多个标记组成，可以是关键字，标识符，常量，字符串，符号。如以下 GO 语句由 6 个标记组成：
fmt.Println("Hello, World!")

6 个标记是(每行一个)：
1. fmt
2. .
3. Println
4. (
5. "Hello, World!"
6. )
*/

/*
行分隔符
`在 Go 程序中，一行代表一个语句结束。每个语句不需要像 C 家族中的其它语言一样以分号 ; 结尾，因为这些工作都将由 Go 编译器自动完成。
`如果你打算将多个语句写在同一行，它们则必须使用 ; 人为区分，但在实际开发中我们并不鼓励这种做法

标识符的第一个字符只能是字母/下划线，还有也不能是go语言里面的关键词
*/
func demo2() {
	fmt.Println("sentence1")
	fmt.Println("sentence2")

	//Go里面变量的几种定义方式
	//1. var 关键字（通用型，包级 / 函数内都能用）
	var age int //0
	var name string = "Go"
	var score = 95.5          //不写数据类型（可自动推导）
	var a, b, c int = 1, 2, 3 //同类型变量可批量定义
	// 定义不同类型变量也正确：自动推导不同类型
	//var a, b, c = 10, 10.8, "r"

	//2. := 短声明（仅函数内可用，快捷型）
	msg := "Hello"
	/*
		这里需要打注释才能不报错是因为上面定义过abc
		a,b := 10,"test"  //可同时声明不同类型变量
		a,c := 20,true //左边至少有 1 个新变量（如 c 是新的），否则报错
	*/
	fmt.Println(name, age, score, a, b, c, msg)

	str1, str2 := "Google", "Gemini"
	fmt.Println(str1 + str2) //str连接可用+实现

	//常量的定义
	const PI float64 = 3.1415926
	const MAX_NUM = 100 //（识别为 int）
	/*
			批量定义方式：
			const (
		    MAX_NUM = 100
		    PI = 3.14159
			)
	*/
	if a > 0 {
		fmt.Println("a>0")
	} else {
		fmt.Println(PI)
	}
	a, b = 5, 5 // 这里不是声明而是需要修改变量的值，所以写的是 = 而不是 :=
	//函数调用
	fmt.Print("a add b = result = ")
	result := add(a, b)
	fmt.Println(result)

	/*
		格式化字符串
		Go 语言中使用 fmt.Sprintf 或 fmt.Printf 格式化字符串并赋值给新串：
		`Sprintf 根据格式化参数生成格式化的字符串并返回该字符串。
		`Printf 根据格式化参数生成格式化的字符串并写入标准输出。
	*/
	//Sprintf
	var stockcode, enddate, url = 123, "2025-1-18", "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
	//Printf
	stockcode = 123
	enddate = "2025-1-18"
	url = "Code=%d&endDate=%s"
	fmt.Printf(url, stockcode, enddate)
}

func add(x, y int) int {
	fmt.Print("x add y = ")
	return x + y
}

/*----------------------------------------------------------------------------------------*/
/*
demo3 Go语言数据类型

在 Go 编程语言中，数据类型用于声明函数和变量。
数据类型的出现是为了把数据分成所需内存大小不同的数据，编程的时候需要用大数据的时候才需要申请大内存，就可以充分利用内存。

Go 语言按类别有以下几种数据类型：
1 布尔型 var b bool = true / false

2 数字类型 Go语言里面像 int8、float64 这类带数字的类型，数字表示这个类型在内存中占用的二进制位数（1 byte = 8 bit）。
位数直接决定了两个关键特性：
`取值范围：位数越多，能存储的数值范围越大。
`内存占用：位数越多，占用的内存字节数越多（比如 int8 占 1 字节，int64 占 8 字节）。

为什么要设计「固定位数」的类型？
Go 语言追求性能可控和跨平台稳定性：
普通 int 在 32 位系统是 32 位，64 位系统是 64 位，跨平台时大小会变，可能导致二进制数据解析错误。
固定位数的类型（如 int32、uint8）在任何平台上大小都不变，因此在处理网络协议、二进制文件、硬件交互等需要精确内存布局的场景时，必须使用固定位数的类型。

`int8 int16 int32 int64 | uint8 uint16 uint32 uint64 (unsigned类型)
float32 float64 关键区别：float64 精度远高于 float32，且是 Go 的默认浮点类型，除非为了节省内存，否则优先用 float64
复数型（数学复数，平时很少用）
complex64	64	2 个 float32（实部 + 虚部）
complex128	128	2 个 float64（实部 + 虚部）

其他数字类型的补充：
1. byte 	=uint8的别名 专门用来处理「字节」（比如二进制数据、文件流、网络包）。
2. rune 	=int32的别名 专门用来处理「Unicode 字符」（比如中文、Emoji，因为 Unicode 码点是 32 位的）。
3. uint 	32位/64位
4. int  	与uint一样大小
```int 和 uint 是 Go 里的「懒人版整型」—— 它们的内存大小、取值范围不固定，会跟着你的操作系统 “自适应”；而 int32/uint8 这类带数字的是「精准版整型」，大小和范围永远固定。
类型	32 位系统下		64 位系统下		核心特点
int		32 位（4 字节）	64 位（8 字节）	有符号，大小随系统变
uint	32 位（4 字节）	64 位（8 字节）	无符号，大小随系统变
int32	32 位（4 字节）	32 位（4 字节）	有符号，大小永远固定
uint64	64 位（8 字节）	64 位（8 字节）	无符号，大小永远固定

5. uintptr 是一个无符号整数类型，用于存放一个指针，专门用来存储「指针的地址值」（也就是内存地址的数字形式，比如 0x123456）。


3 字符串类型
字符串就是一串固定长度的字符连接起来的字符序列。
Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。

4 派生类型:
(a) 指针类型（Pointer）
(b) 数组类型
(c) 结构化类型(struct)
(d) Channel 类型
(e) 函数类型
(f) 切片类型
(g) 接口类型（interface）
(h) Map 类型
*/

func demo3() {
	var char_demo byte = 'a'
	fmt.Println(char_demo)

}

/*----------------------------------------------------------------------------------------*/
/*
demo4 Go语言变量 | 此部分的部分内容可能与上面有重复
Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。
变量声明
第一种，指定变量类型，如果没有初始化，则变量默认为零值。
var v_name v_type
v_name = value
*/

func demo4() {
	var a = "oop"
	var b0 int
	var c bool
	fmt.Println(a, "\n", b0, "\n", c)
	/*
		1 数值类型（包括complex64/128）为 0
		2 布尔类型为 false
		3 字符串为 ""（空字符串）

		以下几种类型为 nil：
		nil 在 Go 中不是 “数字 0”，而是 “引用类型 / 接口类型的空值”，只能赋值给上述类型，不能赋值给 int、string、bool 等基础类型（比如var b int = nil会直接报错）；
		1 var a *int
		2 var a []int
		3 var a map[string] int
		4 var a chan int
		5 var a func(string) int
		6 var a error // error 是接口
	*/
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s) //主要目的是让空字符串输出时候能显式看到""
	/*
		不同占位符的含义
		%v	「通用占位符」：按变量的 “自然格式” 输出（Go 自动识别类型）	所有类型（int/float/bool/string 等）	int(0) → 输出0；bool(false) → 输出false；string("") → 输出空
		%q	「带引号的字符串占位符」：输出字符串时自动加双引号	仅 string 类型	空字符串"" → 输出""；"abc" → 输出"abc"
		%d	「十进制整数占位符」：仅输出整数	int/int32/uint 等整型	int(0) → 输出0；用在 bool/string 上会报错
		%s	「无引号字符串占位符」：输出字符串本身	仅 string 类型
	*/
	var intval int
	intval = 1 //== intval:=1
	fmt.Println(intval)
	/*
		//类型相同多个变量, 非全局变量
		var vname1, vname2, vname3 type
		vname1, vname2, vname3 = v1, v2, v3

		var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断

		vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误

		// 这种因式分解关键字的写法一般用于声明全局变量
		var (
			vname1 v_type1
			vname2 v_type2
		)
	*/

}

/*----------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------------------------*/
