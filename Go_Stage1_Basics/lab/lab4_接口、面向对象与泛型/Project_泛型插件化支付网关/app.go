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
	sign, err := fmt.Scan(&sign)
	if err != nil {
		fmt.Println("输入错误：", err)
		return
	}
	switch sign {
	case 1:
		var user1 User
		user1.Login()
	case 0:
		fmt.Println("退出支付系统")
		return
	default: // 逻辑补充：处理用户输入非1/0的情况
		fmt.Println("无效输入，请输入1或0")

	}
}
