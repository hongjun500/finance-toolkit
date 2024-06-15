// @author hongjun500
// @date 2024-06-15 21:50
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: account.go

package account

// Account 账户中心
type Account struct {
	ID                int    `json:"id"`                 // 账户 ID
	EncryptedPassword string `json:"encrypted_password"` // 加密密码
}
