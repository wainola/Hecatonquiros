package main

import (
	"fmt"
	"github/wainola/proglog/grpcV3/cmd/client/router"
	list "github/wainola/proglog/grpcV3/proto"
	"net/http"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	client := list.NewListServiceClient(conn)

	r := router.BuildRouter(client)

	fmt.Println("Server runing on port 3500")

	http.ListenAndServe(":3500", r)

}
