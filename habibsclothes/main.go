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
			clothes.coat = append(clothes.coat, parts[1:]...)

		case firstfield == "SHIRT:":
			clothes.shirt = append(clothes.shirt, parts[1:]...)

		case firstfield == "PANTS:":
			clothes.pants = append(clothes.pants, parts[1:]...)

		case firstfield == "CAP:":
			clothes.cap = append(clothes.cap, parts[1:]...)

		case firstfield == "JACKET:":
			clothes.jacket = append(clothes.jacket, parts[1:]...)
		}
	}

	clothes.coat = append(clothes.coat, "0")
	clothes.jacket = append(clothes.jacket, "5")
	myscanner.Scan()
	Season := myscanner.Text()
	switch Season {
	case "SPRING":
		for _, i := range clothes.shirt {
			for _, j := range clothes.pants {
				if len(clothes.cap) > 0 {
					for _, k := range clothes.cap {
						if len(clothes.coat) > 0 {
							for _, q := range clothes.coat {
								if q != "0" {
									fmt.Printf("COAT: %s SHIRT: %s PANTS: %s CAP: %s\n", q, i, j, k)
								}
							}
						} else {
							fmt.Printf("SHIRT: %s PANTS: %s CAP: %s\n", i, j, k)
						}
					}
				} else {
					fmt.Printf("SHIRT: %s PANTS: %s\n", i, j)
				}

			}
		}

	case "SUMMER":
		for i := 0; i < len(clothes.shirt); i++ {
			for j := 0; j < len(clothes.pants); j++ {
				for k := 0; k < len(clothes.cap); k++ {
					fmt.Printf("SHIRT: %s PANTS: %s CAP: %s\n", clothes.shirt[i], clothes.pants[j], clothes.cap[k])
				}
			}
		}

	case "FALL":
		for _, i := range clothes.shirt {
			for _, j := range clothes.pants {
				for _, k := range clothes.cap {
					if len(clothes.coat) > 0 {
						for _, q := range clothes.coat {
							if q != "0" && q != "orange" && q != "yellow" {
								fmt.Printf("COAT: %s SHIRT: %s PANTS: %s CAP: %s\n", q, i, j, k)
							}
						}
					} else {
						fmt.Printf("SHIRT: %s PANTS: %s CAP: %s\n", i, j, k)
					}
				}
			}
		}

	case "WINTER":
		for i := 0; i < len(clothes.coat); i++ {
			for j := 0; j < len(clothes.shirt); j++ {
				for k := 0; k < len(clothes.pants); k++ {
					for q := 0; q < len(clothes.jacket); q++ {
						if clothes.coat[i] == "0" && clothes.jacket[q] != "5" {
							fmt.Printf("SHIRT: %s PANTS: %s JACKET: %s\n", clothes.shirt[j], clothes.pants[k], clothes.jacket[q])
						} else if clothes.jacket[q] == "5" && clothes.coat[i] != "0" {
							fmt.Printf("COAT: %s SHIRT: %s PANTS: %s\n", clothes.coat[i], clothes.shirt[j], clothes.pants[k])
						}
					}

				}
			}
		}

	}

}
