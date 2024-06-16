package main

import "fmt"

func ConvertToDigitalFormat(hour, minute, second int) string {
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

func ExtractTimeUnits(seconds int) (int, int, int) {
	//example:
	// 4510 / 3600 = 1.252311 = 1
	// 4510 mod 3600 = 910 | 910 / 60 =15.1666 = 15
	// 910 mod 60 = 10
	hour := seconds / 3600
	minute := (seconds % 3600) / 60
	second := seconds % 60
	return hour, minute, second
}

func main() {
	// Example usage
	digitalTime := ConvertToDigitalFormat(2, 23, 4)
	fmt.Println(digitalTime) // Outputs: 02:23:04
}
