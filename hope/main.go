package main

import "fmt"

func printhope() {
	fmt.Printf("Hope ")
	return
}

func main() {
	var num1, num2 int
	j := 1
	fmt.Scanf("%d %d", &num1, &num2)

	for i := 1; i <= num2; i++ {
		if i%num1 == 0 {
			for x := 0; x < j; x++ {
				printhope()
			}
			fmt.Println()
			j++
		} else {
			fmt.Println(i)
		}
	}
}
