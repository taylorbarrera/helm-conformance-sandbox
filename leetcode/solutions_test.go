package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{name: "pair", nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{name: "duplicate values", nums: []int{3, 3}, target: 6, want: []int{0, 1}},
		{name: "no pair", nums: []int{1, 2}, target: 4},
		{name: "empty input", target: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{input: "()[]{}", want: true},
		{input: "([{}])", want: true},
		{input: "(]", want: false},
		{input: "(((", want: false},
		{input: ")", want: false},
		{input: "", want: true},
		{input: "a", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := IsValidParentheses(tt.input); got != tt.want {
				t.Fatalf("IsValidParentheses(%q) = %t, want %t", tt.input, got, tt.want)
			}
		})
	}
}
