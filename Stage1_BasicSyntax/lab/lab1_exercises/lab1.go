package main

import "fmt"

func main() {
	task1()
}

//task1:个人名片生成器
func task1() {
	var name string = "Juvin"
	var age int = 20
	is_active := true
	var salary float64 = 20000.00

	// "\t"制表符打印
	fmt.Printf(" Name\t| Age\t| IsActive\t| Salary\n")
	infoStr := fmt.Sprintf("%q\t| %d\t| %t    \t| %.2f", name, age, is_active, salary)
	fmt.Println(infoStr)

	fmt.Print("\n------------------------------------------\n\n")

	// 固定格数
	fmt.Printf("%-8s | %-4s | %-10s | %-10s\n", " Name", "Age", "IsActive", "Salary")
	// 内容：用 %-Ns 左对齐，控制每列宽度
	infoStr = fmt.Sprintf("%-8q | %-4d | %-10t | %-10.2f", name, age, is_active, salary)
	fmt.Println(infoStr)
}

//task2:超市收银系统
func task2() {

}

//task3:
