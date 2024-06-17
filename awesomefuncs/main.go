package main

type FilterFunc func(int) bool
type MapperFunc func(int) int

func IsSquare(x int) bool {
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
	//TODO
	return false
}

func Abs(num int) int {
	//TODO
	return 0
}

func Cube(num int) int {
	//TODO
	return 0
}

func Filter(input []int, f FilterFunc) []int {
	//TODO
	return nil
}

func Map(input []int, m MapperFunc) []int {
	//TODO
	return nil
}
