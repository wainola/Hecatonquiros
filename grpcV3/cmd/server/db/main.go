package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "grpcv3"

	session, _ := cluster.CreateSession()
	defer session.Close()

	if err := session.Query("create table lists (id uuid primary key, context text, userId uuid);").Exec(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table lists created!")
}
