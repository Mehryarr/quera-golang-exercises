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

	myscanner.Scan()
	Season := myscanner.Text()
	switch Season {
	case "SPRING":
		for _, coat := range clothes.coat {
			for _, shirt := range clothes.shirt {
				for _, pants := range clothes.pants {
					fmt.Printf("COAT: %s SHIRT: %s PANTS: %s\n", coat, shirt, pants)
				}
			}
		}
		for _, shirt := range clothes.shirt {
			for _, pants := range clothes.pants {
				for _, caps := range clothes.cap {
					fmt.Printf("SHIRT: %s PANTS: %s CAP: %s\n", shirt, pants, caps)
				}
			}
		}
		for _, coat := range clothes.coat {
			for _, shirt := range clothes.shirt {
				for _, pants := range clothes.pants {
					for _, caps := range clothes.cap {
						fmt.Printf("COAT: %s SHIRT: %s PANTS: %s CAP: %s\n", coat, shirt, pants, caps)
					}
				}
			}
		}
		for _, shirt := range clothes.shirt {
			for _, pants := range clothes.pants {
				fmt.Printf("SHIRT: %s PANTS: %s\n", shirt, pants)
			}
		}

	case "SUMMER":
		for _, shirt := range clothes.shirt {
			for _, pants := range clothes.pants {
				for _, caps := range clothes.cap {
					fmt.Printf("SHIRT: %s PANTS: %s CAP: %s\n", shirt, pants, caps)
				}
			}
		}

	case "FALL":
		for _, coat := range clothes.coat {
			if coat != "orange" && coat != "yellow" {
				for _, shirt := range clothes.shirt {
					for _, pants := range clothes.pants {
						fmt.Printf("COAT: %s SHIRT: %s PANTS: %s\n", coat, shirt, pants)
					}
				}
			}
		}
		for _, shirt := range clothes.shirt {
			for _, pants := range clothes.pants {
				for _, caps := range clothes.cap {
					fmt.Printf("SHIRT: %s PANTS: %s CAP: %s\n", shirt, pants, caps)
				}
			}
		}
		for _, coat := range clothes.coat {
			if coat != "orange" && coat != "yellow" {
				for _, shirt := range clothes.shirt {
					for _, pants := range clothes.pants {
						for _, caps := range clothes.cap {
							fmt.Printf("COAT: %s SHIRT: %s PANTS: %s CAP: %s\n", coat, shirt, pants, caps)
						}
					}
				}
			}
		}
		for _, shirt := range clothes.shirt {
			for _, pants := range clothes.pants {
				fmt.Printf("SHIRT: %s PANTS: %s\n", shirt, pants)
			}
		}
	case "WINTER":
		for _, shirt := range clothes.shirt {
			for _, pants := range clothes.pants {
				for _, jacket := range clothes.jacket {
					fmt.Printf("SHIRT: %s PANTS: %s JACKET: %s\n", shirt, pants, jacket)
				}
			}
		}
		for _, coat := range clothes.coat {
			for _, shirt := range clothes.shirt {
				for _, pants := range clothes.pants {
					fmt.Printf("COAT: %s SHIRT: %s PANTS: %s\n", coat, shirt, pants)
				}
			}
		}
	}

}
