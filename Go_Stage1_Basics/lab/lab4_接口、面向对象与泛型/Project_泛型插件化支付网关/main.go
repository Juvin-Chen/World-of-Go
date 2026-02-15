package main

import (
	"fmt"
)

// 泛型插件化支付网关,实现一个模拟企业级的支付系统。

// 订单流水号
const recordID = 1001

func main() {
	fmt.Println("泛型插件化支付网关小程序欢迎您！")
	var sign int
	fmt.Println("登录请按1，退出请按0")
	sign = fmt.Scanner(&sign)
	switch sign {
	case 1:
		var user1 User
		user1.Login()
	case 0:
		f
		return
	}
}
