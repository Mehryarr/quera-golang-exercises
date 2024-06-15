package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	myscanner := bufio.NewScanner(os.Stdin)

	var shabak []string
	var bookname []string
	library := make(map[string]string)

	myscanner.Scan()
	n, _ := strconv.Atoi(myscanner.Text())

	for i := 0; i < n; i++ {
		myscanner.Scan()
		parts := strings.SplitN(myscanner.Text(), " ", 3)
		firstfield := parts[0]
		secondfield := parts[1]

		if firstfield == "ADD" {
			thirdfield := parts[2]
			updated := false
			for i := 0; i < len(shabak); i++ {
				if secondfield == shabak[i] {
					library[shabak[i]] = thirdfield
					updated = true
					break
				}
			}
			if !updated {
				shabak = append(shabak, secondfield)
				bookname = append(bookname, thirdfield)
				library[secondfield] = thirdfield
			}

		} else if firstfield == "REMOVE" {
			delete(library, parts[1])
		}
	}

	keys := make([]string, 0, len(library))
	for key := range library {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		if library[keys[i]] == library[keys[j]] {
			numi, _ := strconv.Atoi(keys[i])
			numj, _ := strconv.Atoi(keys[j])
			return numi < numj
		}
		return library[keys[i]] < library[keys[j]]
	})

	for _, key := range keys {
		fmt.Println(key)
	}

}
