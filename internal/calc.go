// @author hongjun500
// @date 2024-06-16 21:57
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: calc.go
// 用途：计算

package internal

import "math/big"

// Add 加法
func Add(a, b big.Float) big.Float {
	sum := a.Add(&a, &b)
	return *sum
}

// Sub 减法
func Sub(a, b big.Float) big.Float {
	diff := a.Sub(&a, &b)
	return *diff
}

// Mul 乘法
func Mul(a, b big.Float) big.Float {
	prod := a.Mul(&a, &b)
	return *prod
}

// Div 除法
func Div(a, b big.Float) big.Float {
	quotient := a.Quo(&a, &b).SetPrec(0).SetMode(big.ToNearestEven)
	return *quotient
}

// Round 四舍五入
func Round(a big.Float) big.Float {
	return *a.SetMode(big.ToNearestEven).SetPrec(0)
}
