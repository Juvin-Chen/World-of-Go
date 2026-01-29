package main

import (
	"fmt"
)

func main() {
	demo1()
}

/*--------------------------------------------------------------------------------------------------*/
/*
demo1 go语言数组
Go 语言提供了数组类型的数据结构。
数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整型、字符串或者自定义类型。

1.数组的声明: var arrayName [size]datatype
var balance [10]float32

2.在声明时，数组中的每个元素都会根据其数据类型进行默认初始化，对于整数类型，初始值为 0。
var numbers [5]int

3.还可以使用初始化列表来初始化数组的元素
var numbers_2 = [5]int{1, 2, 3, 4, 5}

4.:=的方式
numbers_3 := [5]int{1, 2, 3, 4, 5}

5.在 Go 语言中，数组的大小是类型的一部分，因此不同大小的数组是不兼容的，也就是说 [5]int 和 [10]int 是不同的类型。

- 以下定义了数组 balance 长度为 5 类型为 float32，并初始化数组的元素：
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

- 我们也可以通过字面量在声明数组的同时快速初始化数组：
balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

- 如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度：
var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
或
balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

- 如果设置了数组的长度，我们还可以通过指定下标来初始化元素：
//  将索引为 1 和 3 的元素初始化
balance := [5]float32{1:2.0,3:7.0}
var balance = [5]float63{1:9.08, 3:5.0}

- 初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小

- balance[4] = 50.0
以上实例读取了第五个元素。数组元素可以通过索引（位置）来读取（或者修改），索引从 0 开始，第一个元素索引为 0，第二个索引为 1，以此类推。

6.访问数组元素
数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值。例如：
var salary float32 = balance[9]
*/

func demo1() {
	// Element[%d]=%d 以此格式打印一个数组
	var n [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d]=%d\n", j, n[j])
	}

	// 几个打印示例
	var i_, j_, k int
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for i_ = 0; i_ < 5; i_++ {
		fmt.Printf("balance[%d] = %f\n", i_, balance[i_])
	}

	balance2 := [...]float32{100.0, 2.0, 3.4, 7.0}
	for j_ = 0; j_ < 4; j_++ {
		fmt.Printf("balance2[%d] = %f\n", j_, balance2[j_])
	}

	balance3 := [5]float32{1: 2.0, 3: 4.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}

	/* 补充部分：二维数组 & 向函数传递数组 */
	matrix := [2][3]int{
		{1, 2, 3}, // 第一行
		{4, 5, 6}, // 第二行
	}
	fmt.Println("开始遍历二维数组：")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
		}
	}

	// 可以让编译器自动推导行数，但不能推导列数
	animals := [...][2]string{
		{"Cat", "Dog"},
	}
	fmt.Println("\n推导后的数组animals:", animals)

	/*
		C++ 的习惯：在 C++ 中，当你把数组传给函数时，实际上传递的是首地址指针。所以在函数里修改数组，外面的数组也会变。
		Go 的特性：在 Go 中，数组是值类型 (Value Type)。
		当你把数组传给函数时，Go 会完整地复制一份数组给函数（Deep Copy）。
		在函数里修改数组，只会改那个副本，原本的数组不会变。
		硬伤：Go 语言中，[5]int 和 [10]int 是完全不同的两种数据类型。如果你的函数定义为 func demo(arr [5]int)，那你绝对不能传一个长度为 10 的数组进去，编译器会报错。
	*/

	// 关于向函数传递数组的测试
	myarr := [3]int{111, 222, 333}
	modifyArr(myarr)
	fmt.Println("Not Pointer:", myarr)
	modifyArrByPointer(&myarr)
	fmt.Println("With Pointer:", myarr)
	/*
		Go 的数组其实挺笨重的：1.传参会发生大量内存拷贝（性能低）。2.长度写死在类型里，函数很难通用（一个函数不能同时处理长度为 5 和 10 的数组）。
		实际代码的使用上我们一般用Slice切片，后面会学，暂时只做一个简单的了解
	*/
}

// 这种方式是值传递，跟C++不一样
func modifyArr(arr [3]int) {
	arr[0] = 999
}

// 如果你想在函数里真的修改外面的数组，必须传指针（和C++类似）
func modifyArrByPointer(arr *[3]int) {
	arr[0] = 888
}

/*--------------------------------------------------------------------------------------------------*/
/*
demo2 go语言指针
我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。
Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。

1.什么是指针
一个指针变量指向了一个值的内存地址。
类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：
var var_name *var-type
var ip *int         指向整型
var fp *float32     指向浮点型

2. 如何使用指针
指针使用流程：
-定义指针变量。
-为指针变量赋值。
-访问指针变量中指向地址的值。
--在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。

Go 空指针
当一个指针被定义后没有分配到任何变量时，它的值为 nil。
nil 指针也称为空指针。
nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
空指针判断：
if(ptr != nil)     ptr 不是空指针
if(ptr == nil)     ptr 是空指针
*/

func demo2() {
	var a int = 42
	var ip *int
	ip = &a
	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	// 空指针测试
	var ptr_ *int
	fmt.Println("ptr的值为%x", ptr_)

	/* 补充部分：指针数组 & 二级指针** & 向函数传递指针参数（同上） */
	x, y, z := 7, 8, 9
	var ptrArr [3]*int
	ptrArr[0] = &x
	ptrArr[1] = &y
	ptrArr[2] = &z
	for i := 0; i < 3; i++ {
		fmt.Printf("ptrArr[%d] 指向的值是：%d\n", i, *ptrArr[i])
	}
	*ptrArr[0] = 999
	fmt.Printf("修改后变量 x 的值: %d\n", x) // 输出 999

	// 二级指针
	var ptr *int
	var pptr **int
	ptr = &x    // ptr 存 a 的地址
	pptr = &ptr // pptr 存 ptr 的地址
	fmt.Printf("变量 x = %d\n", a)
	fmt.Printf("指针 ptr 指向的值 (*ptr) = %d\n", *ptr)
	fmt.Printf("二级指针 pptr 最终指向的值 (**pptr) = %d\n", **pptr)
	/* C++ 差异提示： 虽然 Go 支持二级指针，但在 Go 的实际业务开发（尤其是 Web 后端）中，极少会像 C++ 那样用到三级、四级指针。通常一级指针处理结构体传递就足够了。 */

	// 向函数传递指针参数
	f, g := 100, 200
	fmt.Println("Before Swap(f, g):(%d, %d)", f, g)
	swap(&f, &g)
	fmt.Println("After Swap(f, g):(%d, %d)", f, g)
	/*
		the same as
		a, b = b, a
	*/
}

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

/*--------------------------------------------------------------------------------------------------*/
/*
demo3 go语言结构体
Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
结构体表示一项记录，比如保存图书馆的书籍记录，每本书有以下属性：
- Title ：标题
- Author ： 作者
- Subject：学科
- ID：书籍ID

1.定义结构体
结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}

一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：
variable_name := structure_variable_type {value1, value2...valuen}
或
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}

2.访问结构体成员
如果要访问结构体成员，需要使用点号 . 操作符，格式为：
结构体.成员名"

3.结构体作为函数参数
你可以像其他数据类型一样将结构体类型作为参数传递给函数。

4.结构体指针
你可以定义指向结构体的指针类似于其他指针变量，格式如下：
var struct_pointer *Books

以上定义的指针变量可以存储结构体变量的地址。查看结构体变量地址，可以将 & 符号放置于结构体变量前：
struct_pointer = &Book1

使用结构体指针访问结构体成员，使用 "." 操作符：
struct_pointer.title
*/
type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

func demo3() {
	// 创建一个新的结构体
	fmt.Println(Book{"Go_Study", "www.book.com", "Go_language", 1001})
	// 也可以使用 key => value 格式
	fmt.Println(Book{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})
	// 忽略的字段为 0 或 空
	fmt.Println(Book{title: "Go 语言", author: "www.runoob.com"})

	var Book1, Book2 Book
	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	printBook_(&Book1)

	/* 打印 Book2 信息 */
	printBook(Book2)
}

// 结构体作为函数参数,值传递
func printBook(book Book) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

// 结构体指针作为函数参数
func printBook_(book *Book) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

/*--------------------------------------------------------------------------------------------------*/
/*
demo4 go语言切片
Go 语言切片是对数组的抽象。
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。


*/

/*--------------------------------------------------------------------------------------------------*/
/*
demo5 go语言范围
*/

/*--------------------------------------------------------------------------------------------------*/
/*
demo6 go语言Map(集合)
*/
