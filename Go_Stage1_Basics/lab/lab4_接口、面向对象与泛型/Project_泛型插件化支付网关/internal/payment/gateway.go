package payment

import(
	"fmt" 
)

// 泛型结果容器,支付可能会返回不同的元数据（有时是字符串ID，有时是结构体）
type Result[T any] struct{ 
	Success bool 
	Data T 
	Timestamp int64
}

// 支付网关 (Gateway)
type Gateway struct{}
func (g *Gateway) ProcessPayment[T any](method PaymentStrategy, amount float64) Result[T]{
	method.Pay("ctx",) 
	method.Log()
	return {true,"Alipat",20}
}
