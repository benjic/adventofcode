package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_digits(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		args    args
		wantSum int
	}{
		// Official examples
		{args{[]int{1, 1, 2, 2}}, 3},
		{args{[]int{1, 1, 1, 1}}, 4},
		{args{[]int{1, 2, 3, 4}}, 0},
		{args{[]int{9, 1, 2, 1, 2, 1, 2, 9}}, 9},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.args.input), func(t *testing.T) {
			if gotSum := sumOfRepeatedDigits(tt.args.input); gotSum != tt.wantSum {
				t.Errorf("digits() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func Test_opposite(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		args    args
		wantSum int
	}{
		// Official examples
		{args{[]int{1, 2, 1, 2}}, 6},
		{args{[]int{1, 2, 2, 1}}, 0},
		{args{[]int{1, 2, 3, 4, 2, 5}}, 4},
		{args{[]int{1, 2, 3, 1, 2, 3}}, 12},
		{args{[]int{1, 2, 1, 3, 1, 4, 1, 5}}, 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.args.input), func(t *testing.T) {
			if gotSum := sumOfOpposingRepeatedDigits(tt.args.input); gotSum != tt.wantSum {
				t.Errorf("opposite() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func Test_solve(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    captcha
		wantErr bool
	}{
		{
			"puzzle solution",
			args{puzzleInput},
			captcha{1393, 1292},
			false,
		},
		{
			"bad input",
			args{"abcd"},
			captcha{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
