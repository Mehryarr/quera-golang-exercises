package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var clothes struct {
	coat   []string
	shirt  []string
	pants  []string
	cap    []string
	jacket []string
}

func main() {
	myscanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 5; i++ {
		myscanner.Scan()
		parts := strings.Split(myscanner.Text(), " ")
		firstfield := parts[0]

		switch {
		case firstfield == "COAT:":
			for i := 0; i < len(parts)-1; i++ {
				clothes.coat = append(clothes.coat, parts[i+1])
			}
		case firstfield == "SHIRT:":
			for i := 0; i < len(parts)-1; i++ {
				clothes.coat = append(clothes.coat, parts[i+1])
			}
		case firstfield == "PANTS:":
			for i := 0; i < len(parts)-1; i++ {
				clothes.coat = append(clothes.coat, parts[i+1])
			}
		case firstfield == "CAP:":
			for i := 0; i < len(parts)-1; i++ {
				clothes.coat = append(clothes.coat, parts[i+1])
			}
		case firstfield == "JACKET:":
			for i := 0; i < len(parts)-1; i++ {
				clothes.coat = append(clothes.coat, parts[i+1])
			}
		default:
			break
		}
	}
	myscanner.Scan()
	Season := myscanner.Text()
	switch Season {
	case "SPRING":
		fmt.Println("spring")
	case "SUMMER":
		fmt.Println("summer")
	case "FALL":
		fmt.Println("fall")
	case "WINTER":
		fmt.Println("winter")
	}

	fmt.Println(clothes.coat)
	fmt.Println(clothes.shirt)
	fmt.Println(clothes.pants)
	fmt.Println(clothes.cap)
	fmt.Println(clothes.jacket)

}
