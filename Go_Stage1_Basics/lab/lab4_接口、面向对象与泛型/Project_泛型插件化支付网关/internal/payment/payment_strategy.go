package payment

import "fmt"

// 定义一个接口，规定所有支付插件必须具备的能力
type PaymentStrategy interface {
	Pay(amount float64) string
}

// 基础组件，基类
type BasePayment struct {
	AppID string // 每种支付方式App的ID号
}

// 所有的支付方式都需要记录日志
func (b *BasePayment) Log(msg string) {
	fmt.Printf("支付日志：%s\n", msg) // 有实际意义
}

// 具体支付插件1 Ailpay
type AilPay struct {
	BasePayment
}

func (a *AilPay) Pay(amount float64) string {
	fmt.Println("调用支付宝支付API...")

	return "AliPay Success"
}

// 具体支付插件2 WeChat
type WeChat struct {
	BasePayment
}

func (w *WeChat) Pay(amount float64) string {
	fmt.Println("调用微信支付API...")
	return "WeChatId Success"
}
