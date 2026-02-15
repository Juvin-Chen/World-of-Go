package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetPaymentBalance(paymentID string) (float64, error) {
	filePath := "configs/payment_balance.json"

	// 1. 打开JSON文件
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("打开余额文件失败：%w", err) // %w包装原始错误，便于排查
	}
	defer file.Close()

	// 2. 定义与JSON结构匹配的结构体
	type PaymentBalance struct {
		PaymentID string  `json:"PaymentID"`
		Balance   float64 `json:"Balance"`
	}

	// 3. 解析JSON文件到切片
	var balanceList []PaymentBalance
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&balanceList); err != nil {
		return 0, fmt.Errorf("解析余额文件失败：%w", err)
	}

	// 4. 遍历切片，查找对应PaymentID的余额
	for _, item := range balanceList {
		if item.PaymentID == paymentID {
			return item.Balance, nil // 找到，返回余额
		}
	}

	// 5. 未找到对应ID的情况
	return 0, fmt.Errorf("未找到支付ID【%s】的余额记录", paymentID)
}
