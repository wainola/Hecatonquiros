package router

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	list "github/wainola/proglog/grpcV3/proto"
)

func BuildRouter(client list.ListServiceClient) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/items", itemsHandler(client)).Methods("GET")

	return r
}

func itemsHandler(client list.ListServiceClient) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		printList(client)
	}
}

func printList(client list.ListServiceClient) {
	log.Print("Getting the fucking list item!")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rr := &list.GetListReq{}
	listItem, err := client.GetList(ctx, rr)

	if err != nil {
		log.Fatalf("Error on getting item list")
	}
	log.Println("List item:::", listItem)
}
