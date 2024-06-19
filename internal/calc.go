// @author hongjun500
// @date 2024-06-16 21:57
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: calc.go

package internal

import "math/big"

// Precision 保留两位小数
const Precision = 2

// Add 加法
func Add(a, b big.Float) *big.Float {
	sum := new(big.Float).Add(&a, &b)
	return Round(sum, Precision)
}

// Sub 减法
func Sub(a, b big.Float) *big.Float {
	diff := new(big.Float).Sub(&a, &b)
	return Round(diff, Precision)
}

// Mul 乘法
func Mul(a, b big.Float) *big.Float {
	prod := new(big.Float).Mul(&a, &b)
	return Round(prod, Precision)
}

// Div 除法
func Div(a, b big.Float) *big.Float {
	quotient := new(big.Float).Quo(&a, &b)
	return Round(quotient, Precision)
}

// Round rounds a big.Float to the specified number of decimal places.
func Round(x *big.Float, places int) *big.Float {
	ten := big.NewFloat(10)
	multiplier := new(big.Float).SetPrec(uint(places * 4)).Copy(ten)
	for i := 1; i < places; i++ {
		multiplier.Mul(multiplier, ten)
	}

	// Multiply x by the rounding multiplier
	tmp := new(big.Float).Mul(x, multiplier)

	// Add 0.5 for rounding
	half := big.NewFloat(0.5)
	tmp.Add(tmp, half)

	// Round down to the nearest integer
	rounded, _ := tmp.Int(nil)

	// Convert rounded result back to a float and divide by the rounding multiplier
	result := new(big.Float).SetInt(rounded)
	result.Quo(result, multiplier)

	return result
}
