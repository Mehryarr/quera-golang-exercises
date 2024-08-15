package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Server struct {
	p string
}

func NewServer(port string) *Server {
	return &Server{p: port}
}

func (s *Server) Start() {
	http.HandleFunc("/add", addition)
	http.HandleFunc("/sub", subtraction)
	err := http.ListenAndServe(":"+s.p, nil)
	if err != nil {
		panic("server didn't start")
	}
}

func addition(w http.ResponseWriter, r *http.Request) {
	numbersParam := r.URL.Query().Get("numbers")
	if numbersParam == "" {
		respondWithError(w, "'numbers' parameter missing")
		return
	}

	numbers, err := parseNumbers(numbersParam)
	if err != nil {
		respondWithError(w, err.Error())
		return
	}

	result, err := addNumbers(numbers)
	if err != nil {
		respondWithError(w, err.Error())
		return
	}

	respondWithResult(w, result)
}

func subtraction(w http.ResponseWriter, r *http.Request) {
	numbersParam := r.URL.Query().Get("numbers")
	if numbersParam == "" {
		respondWithError(w, "'numbers' parameter missing")
		return
	}

	numbers, err := parseNumbers(numbersParam)
	if err != nil {
		respondWithError(w, err.Error())
		return
	}

	result, err := subtractNumbers(numbers)
	if err != nil {
		respondWithError(w, err.Error())
		return
	}

	respondWithResult(w, result)
}

func parseNumbers(numbersParam string) ([]int, error) {
	numberStrings := strings.Split(numbersParam, ",")
	numbers := make([]int, len(numberStrings))

	for i, numStr := range numberStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", numStr)
		}
		numbers[i] = num
	}

	return numbers, nil
}

func addNumbers(numbers []int) (int, error) {
	sum := 0
	for _, num := range numbers {
		if sum > 0 && num > (int(^uint(0)>>1)-sum) {
			return 0, fmt.Errorf("Overflow")
		}
		sum += num
	}
	return sum, nil
}

func subtractNumbers(numbers []int) (int, error) {
	result := numbers[0]
	for _, num := range numbers[1:] {
		if (num > 0 && result < num-int(^uint(0)>>1)) || (num < 0 && result > num+int(^uint(0)>>1)) {
			return 0, fmt.Errorf("Overflow")
		}
		result -= num
	}
	return result, nil
}

func respondWithResult(w http.ResponseWriter, result int) {
	response := map[string]string{
		"result": fmt.Sprintf("The result of your query is: %d", result),
		"error":  "",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func respondWithError(w http.ResponseWriter, errorMessage string) {
	response := map[string]string{
		"result": "",
		"error":  errorMessage,
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}
