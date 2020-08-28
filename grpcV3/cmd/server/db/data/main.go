package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Quote string `json:"quote"`
}

type ResponseCollection []string

func main() {
	res, err := http.Get("https://api.kanye.rest")

	if err != nil {
		log.Fatal(err)
	}

	r := Response{}
	decodeError := json.NewDecoder(res.Body).Decode(&r)

	if decodeError != nil {
		log.Fatal(err)
	}

	fmt.Println("r::", r)

	// ticker := time.NewTicker(time.Second)
	// defer ticker.Stop()

	// done := make(chan bool)
	// go func(){
	// 	time.Sleep(10 * time.Second)
	// 	done <- true
	// }()

	// for {
	// 	select {
	// 	case <- done:

	// 	}
	// }
}
