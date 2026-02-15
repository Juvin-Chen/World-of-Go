package main

import (
	"fmt"
	"user"
)

// 泛型插件化支付网关,实现一个模拟企业级的支付系统。

// 订单流水号
const recordID = 1001

// 定义一个接口，规定所有支付插件必须具备的能力
type PaymentStrategy interface {
	Pay(ctx string, amount float64) string
}

// 基础组件，基类
type BasePayment struct {
}

// 所有的支付方式都需要记录日志
func (b *BasePayment) Log(msg string) {
	fmt.Println()
}

// 具体插件
type AilPay struct {
	BasePayment
}

func (a *AilPay) Pay(ctx string, amount float64) string {
	fmt.Println("调用支付宝支付API...")
	return "AliPay Success"
}

type WeChat struct {
	BasePayment
}

func (w *WeChat) Pay(ctx string, amount float64) string {
	fmt.Println("调用微信支付API...")
	return "WeChatId Success"
}

// 泛型结果容器,支付可能会返回不同的元数据（有时是字符串ID，有时是结构体）
type Result[T any] struct{ 
	Success bool 
	Data T 
	Timestamp int64
}

// 支付网关 (Gateway)
type Gateway struct{}
func (g *Gateway) ProcessPayment[T any](method PaymentStrategy, amount float64) Result[T]{
	method.Pay() 
	method.Log()
	return {true,"Alipat",20}
}

func main() {
	fmt.Println("泛型插件化支付网关小程序欢迎您！")
	var sign int 
	fmt.Println("登录请按1，退出请按0")
	sign=fmt.Scanner(&sign)
	switch sign {
	case 1:
		var user1 user 
		
	}
}
