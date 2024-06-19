// @author hongjun500
// @date 2024/6/19 16:52
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: calc_test.go

package internal

import (
	"math/big"
	"reflect"
	"testing"
)

type Args struct {
	a big.Float
	b big.Float
}

func TestAdd(t *testing.T) {

	tests := []struct {
		name string
		args Args
		want *big.Float
	}{
		{
			name: "Test Add 1 + 2",
			args: Args{
				a: *big.NewFloat(1),
				b: *big.NewFloat(2),
			},
			want: big.NewFloat(3),
		},
		{
			name: "Test Add 1.2 + 2.2",
			args: Args{
				a: *big.NewFloat(1.2),
				b: *big.NewFloat(2.2),
			},
			want: big.NewFloat(3.40),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.args.a, tt.args.b)
			gotFloat64, _ := got.Float64()
			wantFloat64, _ := tt.want.Float64()

			// if gotFloat64 != wantFloat64 {
			if !reflect.DeepEqual(gotFloat64, wantFloat64) {
				t.Errorf("Add() = %v, want %v", got.String(), tt.want.String())
			}
		})
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		name string
		args Args
		want *big.Float
	}{
		{
			name: "Test Div 10 / 3",
			args: Args{
				a: *big.NewFloat(10),
				b: *big.NewFloat(3),
			},
			want: big.NewFloat(3.33),
		},
		{
			name: "Test Div 2 / 3",
			args: Args{
				a: *big.NewFloat(2),
				b: *big.NewFloat(3),
			},
			want: big.NewFloat(0.67),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := Div(tt.args.a, tt.args.b); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("tt.name = %v, Div() = %v, want %v", tt.name, got.String(), tt.want.String())
			}
		})
	}
}

func TestMul(t *testing.T) {
	type args struct {
		a big.Float
		b big.Float
	}
	tests := []struct {
		name string
		args args
		want *big.Float
	}{
		{
			name: "Test Mul 1 * 2",
			args: args{
				a: *big.NewFloat(1),
				b: *big.NewFloat(2),
			},
			want: big.NewFloat(2.00),
		},
		{
			name: "Test Mul 1.2 * 2.2",
			args: args{
				a: *big.NewFloat(1.2),
				b: *big.NewFloat(2.2),
			},
			want: big.NewFloat(2.64),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mul(tt.args.a, tt.args.b); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("tt.name = %v, Mul() = %v, want %v", tt.name, got.String(), tt.want.String())
			}
		})
	}
}

func TestRound(t *testing.T) {
	type roundArgs struct {
		x      *big.Float
		places int
	}
	tests := []struct {
		name string
		args roundArgs
		want *big.Float
	}{
		{
			name: "Test Round 1.2345 to 2 places",
			args: roundArgs{
				x:      big.NewFloat(1.2345),
				places: 2,
			},
			want: big.NewFloat(1.23),
		},
		{
			name: "Test Round 1.2356 to 2 places",
			args: roundArgs{
				x:      big.NewFloat(1.2356),
				places: 2,
			},
			want: big.NewFloat(1.24),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.x, tt.args.places); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("tt.name = %v, Round() = %v, want %v", tt.name, got.String(), tt.want.String())
			}
		})
	}
}

func TestSub(t *testing.T) {

	tests := []struct {
		name string
		args Args
		want *big.Float
	}{
		{
			name: "Test Sub 2 - 1",
			args: Args{
				a: *big.NewFloat(2),
				b: *big.NewFloat(1),
			},
			want: big.NewFloat(1),
		},
		{
			name: "Test Sub 2.2 - 1.2",
			args: Args{
				a: *big.NewFloat(2.2),
				b: *big.NewFloat(1.2),
			},
			want: big.NewFloat(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sub(tt.args.a, tt.args.b); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("tt.name = %v, Sub() = %v, want %v", tt.name, got.String(), tt.want.String())
			}
		})
	}
}
