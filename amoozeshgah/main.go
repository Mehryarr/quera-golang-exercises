package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	teacherNames := []string{}
	scores := []float64{}

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		teacherNames = append(teacherNames, scanner.Text())
		scanner.Scan()
		scoreSlice := strings.Split(scanner.Text(), " ")

		sum := 0
		for _, score := range scoreSlice {
			s, _ := strconv.Atoi(score)
			sum += s
		}
		scores = append(scores, float64(sum/len(scoreSlice)))
	}

	for i := 0; i < len(scores); i++ {
		v := scores[i]

		if v >= 80 {
			fmt.Printf("%s Excellent\n", teacherNames[i])
		} else if v >= 60 {
			fmt.Printf("%s Very Good\n", teacherNames[i])
		} else if v >= 40 {
			fmt.Printf("%s Good\n", teacherNames[i])
		} else {
			fmt.Printf("%s Fair\n", teacherNames[i])
		}
	}
}
