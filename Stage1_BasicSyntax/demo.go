/*
Go 语言的基础组成有以下几个部分：
1 包声明
2 引入包
3 函数
4 变量
5 语句 & 表达式
6 注释
*/

package main

import "fmt"

func main() {
	demo2()
}

/*--------------------------------------------------------*/

/*
demo1

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

/*--------------------------------------------------------*/

/*
demo2

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
