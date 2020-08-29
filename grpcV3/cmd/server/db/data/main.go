package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gocql/gocql"
)

type Response struct {
	Quote string `json:"quote"`
}

type ResponseCollection []string

func makeRequest() {
	res, err := http.Get("https://api.kanye.rest")

	if err != nil {
		log.Fatal(err)
	}

	r := Response{}
	decodeError := json.NewDecoder(res.Body).Decode(&r)

	if decodeError != nil {
		log.Fatal(err)
	}

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "grpcv3"

	session, _ := cluster.CreateSession()

	defer session.Close()

	if err := session.Query("insert into lists (id, content, userid) values (?, ?, ?);", gocql.TimeUUID(), r.Quote, gocql.TimeUUID()).Exec(); err != nil {
		log.Fatal("Error inserting data", err)
	}

	fmt.Printf("Success inserting quote: %s\n", r.Quote)
}

func main() {

	for i := 1; i < 20; i++ {
		time.Sleep(1 * time.Second)
		makeRequest()
	}
}
