package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"

	list "github/wainola/proglog/grpcV3/proto"
)

func BuildRouter(client list.ListServiceClient) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/items", itemsHandler(client)).Methods("GET")
	r.HandleFunc("/item/{id}", getItemHandler(client)).Methods("GET")

	return r
}

func itemsHandler(client list.ListServiceClient) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		printList(client)
	}
}

func printList(client list.ListServiceClient) {
	log.Print("Getting the fucking list item!")
	ctx, cancel := provideCtx()
	defer cancel()
	rr := &list.GetListReq{}
	listItem, err := client.GetList(ctx, rr)

	if err != nil {
		log.Fatalf("Error on getting item list")
	}
	log.Println("List item:::", listItem)
}

func getItemHandler(client list.ListServiceClient) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.Split(r.URL.Path, "/")

		var idValues []string
		for _, value := range id {
			fmt.Println("value::", value)
			if len(value) != 0 {
				idValues = append(idValues, value)
			}
		}

		var idToUse string = idValues[1]

		fmt.Println("id array:", idToUse)

		getOneListItem(client, idToUse)

	}
}

func getOneListItem(client list.ListServiceClient, id string) {
	ctx, cancel := provideCtx()

	defer cancel()

	rr := &list.GetOneListReq{ListId: id}

	listItem, err := client.GetOneListItem(ctx, rr)

	if err != nil {
		log.Fatal("Error on getting one item", err)
	}

	fmt.Println("List item", listItem)

}

func provideCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
