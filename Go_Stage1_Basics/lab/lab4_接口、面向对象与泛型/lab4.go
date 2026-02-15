package main

import "fmt"

// Go 是编译型语言，代码按从上到下顺序解析
// 提前声明 taskX
func task1()
func task2()
func task3()
func task4()
func task5()
func task6()

/*
1.Go 仅对函数支持 “先声明签名、后定义实现”，目的是解决函数循环调用问题；
2.结构体、泛型类型、方法不支持这种拆分，必须在使用前写完整定义（字段 / 实现）；
*/

type TaskFunc func()

func main() {
	// Go语言当中函数是一等公民，可以赋值给变量，可以作为参数传递，可以作为返回值

	// 创建任务列表
	tasks := []TaskFunc{task1, task2, task3, task4, task5, task6}
	for _, task := range tasks {
		task()
		fmt.Println()
	}
}
