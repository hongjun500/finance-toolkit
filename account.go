// @author hongjun500
// @date 2024-06-15 21:50
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: 账户中心

package main

import "math/big"

// Account 账户中心
type Account struct {
	ID                int      `json:"id"`                 // 账户 ID
	EncryptedPassword string   `json:"encrypted_password"` // 加密密码
	Stocks            []*Stock `json:"stocks"`             // 股票
}

// NewAccount 创建账户
func NewAccount(id int, encryptedPassword string) *Account {
	return &Account{
		ID:                id,
		EncryptedPassword: encryptedPassword,
	}
}

// Stock 股票
type Stock struct {
	Code         string    // 股票代码
	Name         string    // 股票名称
	CostPrice    big.Float // 成本价
	CurrentPrice big.Float // 当前价
	Quantity     int       // 数量
	*TradeFee              // 交易费用
}

// TradeFee 交易费用
type TradeFee struct {
	Commission big.Float // 佣金
	StampDuty  big.Float // 印花税
	Transfer   big.Float // 过户费
}
