package main

import "fmt"

func main() {
	var x, r float32
	fmt.Scanln(&x)
	switch {
	case x <= 100:
		r = 0.05 * x
	case x > 100 && x <= 500:
		r = 5 + (x-100)*0.1
	case x > 500 && x <= 1000:
		r = 45 + (x-500)*0.15
	case x > 1000:
		r = 120 + (x-1000)*0.2
	}

	fmt.Println(int(r))

}
