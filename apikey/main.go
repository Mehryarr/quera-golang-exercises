package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if apiKey == "" {
		fmt.Println("OPENWEATHERMAP_API_KEY is not set in the .env file")
		return
	}
	apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=Tehran&appid=%s", apiKey)

	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error making the API request:", err)
		return
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	fmt.Println(string(responseBody))
}
