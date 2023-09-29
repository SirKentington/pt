package pt

import (
	"testing"
)

func Test_dimensionalOverlap(t *testing.T) {
	type args struct {
		d11 float64
		d12 float64
		d21 float64
		d22 float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{d11: 10, d12: 10, d21: 10, d22: 10}, want: true},
		{args: args{d11: 0, d12: 1, d21: 0, d22: 1}, want: true},
		{args: args{d11: 0, d12: 1, d21: 2, d22: 3}, want: false},
		{args: args{d11: 2, d12: 3, d21: 0, d22: 1}, want: false},
		{args: args{d11: -2, d12: -2, d21: -2, d22: -2}, want: true},
		{args: args{d11: -10, d12: 10, d21: -2, d22: 2}, want: true},
		{args: args{d11: -1, d12: 1, d21: -2, d22: 2}, want: true},
		{args: args{d11: 0, d12: 1, d21: 1, d22: 2}, want: true},
		{args: args{d11: 1, d12: 0, d21: 2, d22: 1}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dimensionalOverlap(tt.args.d11, tt.args.d12, tt.args.d21, tt.args.d22); got != tt.want {
				t.Errorf("dimensionalOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAreColliding(t *testing.T) {

	type args struct {
		h1 Hitbox
		h2 Hitbox
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{
			h1: Hitbox{BV1: Vec{}, BV2: Vec{}},
			h2: Hitbox{BV1: Vec{}, BV2: Vec{}}},
			want: true},
		{args: args{
			h1: Hitbox{BV1: Vec{1, 1, 1}, BV2: Vec{2, 2, 2}},
			h2: Hitbox{BV1: Vec{3, 3, 3}, BV2: Vec{4, 4, 4}}},
			want: false},
		{args: args{
			h1: Hitbox{BV1: Vec{}, BV2: Vec{1, 1, 1}},
			h2: Hitbox{BV1: Vec{}, BV2: Vec{-1, -1, -1}}},
			want: true},
		{args: args{
			h1: Hitbox{BV1: Vec{}, BV2: Vec{1, 1, 1}},
			h2: Hitbox{BV1: Vec{-1, -1, -1}, BV2: Vec{}}},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AreColliding(tt.args.h1, tt.args.h2); got != tt.want {
				t.Errorf("AreColliding() = %v, want %v", got, tt.want)
			}
		})
	}
}
