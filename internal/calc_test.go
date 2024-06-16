// @author hongjun500
// @date 2024-06-16 22:07
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: calc_test.go

package internal

import (
	"math/big"
	"testing"
)

func TestDiv(t *testing.T) {

	div := Div(*big.NewFloat(10), *big.NewFloat(3))
	t.Logf("div = %v", div.String())

}
