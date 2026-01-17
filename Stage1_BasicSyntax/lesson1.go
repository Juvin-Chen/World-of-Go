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
	/*this is my first program*/
	fmt.Println("hello go")
	fmt.Print("hello go\n")
	//Printf 是可以控制格式化的输入输出的，若没有传入占位符，它就会跟Print用法变得一样
	fmt.Printf("hello %s", "go") // 输出：hello go
	fmt.Printf("hello go")
}

/*
关于这个demo的说明
1 必须在源文件的非注释的第一行指明这个文件属于哪个包，如:package main表示一个可独立执行的程序，每一个go应用程序都应该包含一个名为main的package

2 import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。

3 func main() 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。

4 fmt.Println(...) 可以将字符串输出到控制台，并在最后自动增加换行字符 \n,使用 fmt.Print("hello, world\n") 可以得到相同的结果。Print 和 Println 这两个函数也支持使用变量，如：fmt.Println(arr)。如果没有特别指定，它们会以默认的打印格式将变量 arr 输出到控制台。

5 Go 语言通过标识符（变量、函数、结构体等）的大小写直接控制可见性：大写开头的标识符（如 fmt.Println 中的 Println）属于对外导出（类似 public），其他包导入后即可调用，就像公司的对外公开产品；小写开头的标识符仅在当前包内可见（类似 private），外部包无法访问，就像公司的内部机密工具。
*/
