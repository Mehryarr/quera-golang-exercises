package main

import (
	"strconv"
)

type FilterFunc func(int) bool
type MapperFunc func(int) int

func IsSquare(x int) bool {
	if x == 0 {
		return true
	}
	for i := 1; ; i = i + 2 {
		x -= i
		if x == 0 {
			return true
		} else if x <= 0 {
			return false
		}
	}
}

func IsPalindrome(x int) bool {

	y := strconv.Itoa(x)

	if x < 0 {
		y = y[1:]
	}

	for i := 0; i < len(y)/2; i++ {
		if y[i] != y[len(y)-i-1] {
			return false
		}
	}
	return true
}

func Abs(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}

func Cube(num int) int {
	return num * num * num
}

func Filter(input []int, f FilterFunc) []int {
	res := make([]int, 0, len(input))
	for _, v := range input {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

func Map(input []int, m MapperFunc) []int {
	var res []int
	for _, v := range input {
		res = append(res, m(v))
	}
	return res
}
