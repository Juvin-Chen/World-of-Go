package main

import (
	"fmt"
)

func main() {

}

/*------------------------------------------------------------------------*/

/*
demo1 Go语言条件语句
条件语句需要开发者通过指定一个或多个条件，并通过测试条件是否为 true 来决定是否执行指定语句，并在条件为 false 的情况在执行另外的语句。
Go 语言提供了以下几种条件判断语句：
` if 语句	if 语句 由一个布尔表达式后紧跟一个或多个语句组成。
` if...else 语句	if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。
` if 嵌套语句	你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。
` switch 语句	switch 语句用于基于不同条件执行不同动作。
` select 语句	select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
注意：Go 没有三目运算符，所以不支持 ?: 形式的条件判断。
*/

func demo1() {
	// switch使用示例，没有break，默认匹配到一个 case 执行完就自动结束
	// 如果想要像C++一样达到一个穿透的效果，必须显式加上 fallthrough 关键字。
	code := "A"
	switch code {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
		fallthrough
	case "C":
		fmt.Println("及格")
	default:
		fmt.Println("其他")
	}

	// Go支持任意类型 C++仅限整数/枚举 ，C++: switch(expression) 里的表达式必须是 Integral type (int, char, long, enum) 或者能隐式转为 int 的类
	color := "red"
	//Go: 支持几乎所有可比较类型。你可以直接 switch 字符串 (string)，甚至数组、结构体（只要它们是可比较的）。
	switch color {
	case "red":
		fmt.Println("Stop")
	case "green":
		fmt.Println("Go")
	}

	// Go可以是变量/逻辑 而C++只能是编译期常量
	score := 85
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}

	// 多个匹配值：Go 用逗号，C++ 堆叠 Case
	/*
			C++
			case 1:
			case 2:
			case 3:
				doSomething();
				break;
			Go
			case 1, 2, 3:
		    	doSomething()
	*/

	//select TODO: 需配合 Stage 5 Channel 章节实操]。
}

/*------------------------------------------------------------------------*/

/*
demo2 Go 语言循环语句
在不少实际问题中有许多具有规律性的重复操作，因此在程序中就需要重复执行某些语句。
循环控制语句
循环控制语句可以控制循环体内语句的执行过程。

GO 语言支持以下几种循环控制语句：
break 语句	经常用于中断当前 for 循环或跳出 switch 语句
continue 语句	跳过当前循环的剩余语句，然后继续进行下一轮循环。
goto 语句	将控制转移到被标记的语句。
*/

func demo2() {
	//无限循环
	for true {
		fmt.Printf("这是一个无线循环\n")
	}

	//主要是明确一下关于label这个标签的使用 通常和流程控制一起配合使用~

	// Go 的进阶用法：配合 Label 跳出多层循环
	// 给外层循环起个名字叫 Loop1
Loop1:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				// 直接跳出 Loop1，也就是连外层都结束了
				break Loop1
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
	// break Loop1 后，代码会直接运行到这里
	fmt.Println("Break Loop1 \n Done\n")

	/*
		Go 的进阶用法：配合 Label 控制外层循环
		同样，continue 也可以加标签。
		这通常用于：在内层循环里发现某种情况，想要直接触发外层循环的下一次迭代。
	*/
OuterLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j > i {
				// 这里的逻辑是：结束内层循环，直接让外层 i++，开始下一轮
				continue OuterLoop
			}
			fmt.Print(j)
		}
		fmt.Println()
	}

	/*
		Goto 语句 (跳转)
		用法（和 C++ 一样）： 无条件跳转到指定的标签行。
		Go: 虽然也不推荐滥用，但在 Go 的标准库源码中，goto 其实用得不少！
		常见场景：用于统一的错误处理或清理资源（虽然后面你会学到更优雅的 defer，但在某些高性能场景或复杂的状态机解析中，goto 依然有一席之地）。
	*/
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto ErrorHandler
		}
	}
ErrorHandler:
	fmt.Println("检测到错误，执行清理")
}

/*------------------------------------------------------------------------*/

/*
demo3 Go 语言标准库提供了多种可动用的内置的函数。
例如，len() 函数可以接受不同类型参数并返回该类型的长度。
如果我们传入的是字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含的元素个数。

func function_name( [parameter list] ) [return_types] {
   函数体
}

//函数参数
两种方式来传递参数：
值传递	值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
引用传递	引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。


函数用法
函数作为另外一个函数的实参	函数定义后可作为另外一个函数的实参数传入
闭包	闭包是匿名函数，可在动态编程中使用  == C++ lambda表达式
方法	方法就是一个包含了接受者的函数
*/

// 函数作为实参
func compute(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func add(x, y int) int { return x + y }

// 2 闭包函数
/*
`Go 概念：闭包是一个匿名函数，但它最神奇的地方在于它能“捕获”并“记住”其外部作用域的变量。
即使外部函数已经执行完了，闭包里引用的那个外部变量依然活着。
`C++ 类比：这完全对应 C++ 的 Lambda 表达式带捕获列表 [=] 或 [&]。
`底层差异：在 C++ 里，如果引用的局部变量销毁了，Lambda 再去访问就会野指针/崩溃。
但在 Go 里，编译器会进行逃逸分析 (Escape Analysis)，如果发现闭包用了局部变量，它会把这个变量自动搬到堆 (Heap) 上，保证它不会死掉。
*/

func demo3() {
	c := max(5, 6)
	fmt.Println(c)

}

// Go里面使用函数无需提前声明
func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

/*------------------------------------------------------------------------*/

/*

 */
