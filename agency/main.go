package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	myscanner := bufio.NewScanner(os.Stdin)

	var countnryName []string
	var countryCode []string
	var result []string
	m := make(map[string]string)
	myscanner.Scan()
	n, _ := strconv.Atoi(myscanner.Text())

	for i := 0; i < n; i++ {
		myscanner.Scan()
		fields := strings.Fields(myscanner.Text())
		firstfield := []string{fields[0]}
		countnryName = append(countnryName, firstfield...)
		secondfield := []string{fields[1]}
		countryCode = append(countryCode, secondfield...)
		//m[countnryCode[i]] = countnryName[i]
		m[countryCode[i]] = countnryName[i]

	}

	myscanner.Scan()
	q, _ := strconv.Atoi(myscanner.Text())

	for i := 0; i < q; i++ {
		var index int
		var test bool
		myscanner.Scan()
		str := myscanner.Text()
		test = false
		for i := 0; i < n; i++ {
			index = strings.Index(str, countryCode[i])
			if index != -1 {
				result = append(result, m[countryCode[i]])
				test = true
			}
		}
		if test == false {
			result = append(result, "Invalid Number")
		}

	}

	//fmt.Println(countnryName)
	//fmt.Println(countryCode)
	//fmt.Println(m)

	for _, v := range result {
		fmt.Println(v)
	}

}
