package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to find all arithmetic sequence subarrays
func findArithmeticSubarrays(arr []int) [][]int {
	n := len(arr)
	if n < 3 {
		return [][]int{}
	}

	result := [][]int{}

	for i := 0; i < n-2; {
		start := i
		diff := arr[i+1] - arr[i]
		end := i + 2

		// Extend the subarray as long as the difference is constant
		for end < n && arr[end]-arr[end-1] == diff {
			end++
		}

		// If we have a valid arithmetic subarray with at least 3 elements
		if end-start >= 3 {
			for k := start + 2; k < end; k++ {
				result = append(result, arr[start:k+1])
			}
		}

		// Move to the next potential starting point
		i = start + 1
	}

	return result
}

func main() {
	myscanner := bufio.NewScanner(os.Stdin)
	var spaceship []string

	var resultcount []int

	myscanner.Scan()
	n, _ := strconv.Atoi(myscanner.Text())

	for i := 0; i < n; i++ {
		myscanner.Scan()
		fields := strings.Fields(myscanner.Text())
		firstfield := []string{fields[0]}
		spaceship = append(spaceship, firstfield...)
		var gas []int
		for _, field := range fields[1:] {
			num, _ := strconv.Atoi(field)
			gas = append(gas, int(num))
		}
		result := len(findArithmeticSubarrays(gas))
		resultcount = append(resultcount, result)
		//fmt.Println(gas)
	}

	//fmt.Println(spaceship)

	for i := 0; i < n; i++ {
		fmt.Printf("%s %d", spaceship[i], resultcount[i])
		fmt.Println()
	}

}

//****************************************************************
/*
func countArithmeticSequences(fuelData []int) int {
    count := 0
    n := len(fuelData)

    for i := 0; i < n; i++ {
        for j := i + 2; j < n; j++ {
            d := fuelData[i+1] - fuelData[i]
            isArithmetic := true
            for k := i + 1; k <= j; k++ {
                if fuelData[k]-fuelData[k-1] != d {
                    isArithmetic = false
                    break
                }
            }
            if isArithmetic {
                count++
            }
        }
    }

    return count
}
*/
//***************************************************************
/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to count all arithmetic sequence subarrays
func countArithmeticSubarrays(arr []int) int {
	n := len(arr)
	if n < 3 {
		return 0
	}

	count := 0

	for i := 0; i < n-2; {
		start := i
		diff := arr[i+1] - arr[i]
		end := i + 2

		// Extend the subarray as long as the difference is constant
		for end < n && arr[end]-arr[end-1] == diff {
			end++
		}

		// If we have a valid arithmetic subarray with at least 3 elements
		if end-start >= 3 {
			count += (end - start - 1) * (end - start - 2) / 2
		}

		// Move to the next potential starting point
		i = end - 1
	}

	return count
}

func main() {
	myscanner := bufio.NewScanner(os.Stdin)
	var spaceships []string
	var resultcount []int

	// Read the number of input lines
	myscanner.Scan()
	n, _ := strconv.Atoi(myscanner.Text())

	for i := 0; i < n; i++ {
		myscanner.Scan()
		fields := strings.Fields(myscanner.Text())
		if len(fields) < 2 {
			continue // Skip invalid lines
		}

		// First field for the spaceship name
		spaceship := fields[0]
		spaceships = append(spaceships, spaceship)

		// Reinitialize gas for each line
		gas := []int{}
		for _, field := range fields[1:] {
			num, _ := strconv.Atoi(field)
			gas = append(gas, num)
		}

		// Count the arithmetic subarrays and store the count
		count := countArithmeticSubarrays(gas)
		resultcount = append(resultcount, count)
	}

	// Print the results
	for i := 0; i < n; i++ {
		fmt.Printf("%s %d\n", spaceships[i], resultcount[i])
	}
}
*/
