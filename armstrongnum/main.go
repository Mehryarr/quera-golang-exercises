package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func checkarmstrong(num int) bool {
	numstr := strconv.Itoa(num)
	originalnum := num
	sum := 0
	for i := 0; i < len(numstr); i++ {
		sum += int(math.Pow(float64(num%10), float64(len(numstr))))
		num = num / 10
	}
	if originalnum == sum {
		return true
	} else {
		return false
	}
}

func main() {
	//our string
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	// using regular exp to find all sequences of digits
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(input.Text(), -1)

	sum := 0
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			fmt.Println(err)
		}
		sum += num
	}

	//check
	if checkarmstrong(sum) == true {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
