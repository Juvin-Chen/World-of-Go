package payment

import "fmt"

// 定义一个接口，规定所有支付插件必须具备的能力
type PaymentStrategy interface {
	Pay(amount float64) (string, error)
}

// 基础支付结构体
type BasePayment struct {
	Paytype   string
	PaymentID string  // 支付AppID号，如："WeChat1001 / AilPay1001"
	Balance   float64 // 账户余额
}

// 所有的支付方式都需要记录日志
func (b *BasePayment) Log(msg string) {
	fmt.Printf("[%s%s] 支付日志：%s\n", b.Paytype, b.PaymentID, msg)
}

// 具体支付插件1 Ailpay
type AilPay struct {
	BasePayment
}

func (a *AilPay) Pay(amount float64) (string, error) {
	a.Log(fmt.Sprintf("调用支付宝支付插件，发起支付，金额：%.2f元", amount))
	if a.Balance < amount {
		return "", fmt.Errorf("支付宝余额不足：当前%.2f元，需支付%.2f元", a.Balance, amount)
	}
	// 扣减余额
	a.Balance -= amount
	a.Log(fmt.Sprintf("支付成功，剩余余额：%.2f元", a.Balance))
	return "Alipay Success", nil
}

// 具体支付插件2 WeChat
type WeChat struct {
	BasePayment
}

func (w *WeChat) Pay(amount float64) (string, error) {
	w.Log(fmt.Sprintf("发起支付，金额：%.2f元", amount))
	if w.Balance < amount {
		return "", fmt.Errorf("微信余额不足：当前%.2f元，需支付%.2f元", w.Balance, amount)
	}
	w.Balance -= amount
	w.Log(fmt.Sprintf("支付成功，剩余余额：%.2f元", w.Balance))
	return "WeChat Success", nil
}
