package main

import (
	"fmt"
	"pay-project/internal/user"
)

// 泛型插件化支付网关,实现一个模拟企业级的支付系统。

// 订单流水号
const recordID = 1001

func main() {
	// 1. 用户登录
	var u user.User
	fmt.Println("泛型插件化支付网关欢迎您！")

	var sign int
	// 循环输入：直到输入1/0为止
	for {
		fmt.Println("\n登录请按1，退出请按0")
		_, err := fmt.Scan(&sign)
		if err != nil {
			fmt.Println("输入错误：请输入数字（1/0），", err)
			// 清空输入缓冲区，避免死循环
			fmt.Scanln()
			continue
		}

		switch sign {
		case 1:
			// 登录逻辑：登录失败则继续循环，成功则退出循环
			if u.Login() {
				fmt.Println("=== 登录成功，进入支付环节 ===")
				goto PayStep // 跳转到支付选择环节
			} else {
				fmt.Println("登录失败，请重新尝试！")
				continue
			}
		case 0:
			fmt.Println("退出支付系统")
			return // 直接退出程序
		default:
			fmt.Println("无效输入，请输入1（登录）或0（退出）！")
		}
	}

	// 支付方式选择环节（通过goto跳转）
PayStep:
	// 2. 选择支付方式
	var payChoice int
	// 同样循环输入：直到输入1/2为止
	for {
		fmt.Println("\n请选择支付方式：")
		fmt.Println("1. 支付宝")
		fmt.Println("2. 微信支付")
		fmt.Print("输入序号：")
		_, err := fmt.Scanln(&payChoice)
		if err != nil {
			fmt.Println("输入错误：请输入数字（1/2），", err)
			fmt.Scanln() // 清空缓冲区
			continue
		}

		switch payChoice {
		case 1:
			fmt.Println("你选择了【支付宝】支付，后续将调用支付宝支付插件")
			// 这里后续可添加调用支付宝构造函数的逻辑
			return
		case 2:
			fmt.Println("你选择了【微信支付】，后续将调用微信支付插件")
			// 这里后续可添加调用微信构造函数的逻辑
			return
		default:
			fmt.Println("无效输入，请输入1（支付宝）或2（微信支付）！")
		}
	}
}
