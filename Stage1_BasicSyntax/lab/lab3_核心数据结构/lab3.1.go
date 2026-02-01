package main

func main() {

}

func demo1() {
	arr := [5]int{10, 20, 30, 40, 50}
	doubleArrayByVal(arr)
}

// 错误示范：值传递
func doubleArrayByVal(arr [5]int) {
	for _, v := range arr {
		v *= 2
	}
}

// 正确示范：指针传递
func task2() {

}
