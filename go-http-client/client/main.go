package main

import (
	"context"
	"fmt"
	"log"

	gohttpclient "github.com/bozd4g/go-http-client"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	client := gohttpclient.New("http://localhost:8080/sayHelloWorld")

	response, err := client.Get(context.Background(), "")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var res Response
	if err := response.Unmarshal(&res); err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("%+v\n", res)
}
