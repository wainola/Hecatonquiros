package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gocql/gocql"

	"google.golang.org/grpc/reflection"

	list "github/wainola/proglog/grpcV3/proto"

	"google.golang.org/grpc"
)

type listServer struct {
	ClusterInstance *gocql.ClusterConfig
}

type ListItems struct {
	Items []*list.List
}

// LIST COLLECTION FOR LEARNING PURPOSES
var ItemsCollection ListItems = ListItems{Items: []*list.List{
	&list.List{Id: "123", Content: "Content number 1"},
	&list.List{Id: "111", Content: "Content number 2"},
	&list.List{Id: "456", Content: "Content number 3"},
}}

func (l *listServer) SetList(ctx context.Context, listItem *list.List) (*list.SetListsResponse, error) {
	fmt.Println("List received:", listItem)

	return &list.SetListsResponse{}, nil
}

func (l *listServer) GetAllLists(req *list.RequestTrack, stream list.ListService_GetAllListsServer) error {

	session, _ := l.ClusterInstance.CreateSession()
	defer session.Close()

	iter := session.Query("select id, content, userid from lists;").Iter()

	var id gocql.UUID
	var content string
	var userId gocql.UUID

	var lCollection []*list.List = []*list.List{}

	for iter.Scan(&id, &content, &userId) {
		lCollection = append(lCollection,
			&list.List{
				Id:      id.String(),
				Content: content,
				UserId:  userId.String(),
			})
	}

	var collectionToSend ListItems = ListItems{
		Items: lCollection,
	}

	if err := stream.Send(
		&list.ListResp{
			Items: collectionToSend.Items,
		},
	); err != nil {
		return err
	}

	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (l *listServer) GetList(ctx context.Context, req *list.GetListReq) (*list.ListResp, error) {

	return &list.ListResp{Items: ItemsCollection.Items}, nil
}

func getOneItemList(id string) (*list.List, error) {
	var itemToReturn map[string]string = make(map[string]string)
	for _, value := range ItemsCollection.Items {
		if id == value.Id {
			itemToReturn["id"] = value.Id
			itemToReturn["content"] = value.Content
		}
	}
	return &list.List{Id: itemToReturn["id"], Content: itemToReturn["content"]}, nil
}

func (l *listServer) GetOneListItem(ctx context.Context, req *list.GetOneListReq) (*list.ListItem, error) {

	item, err := getOneItemList(req.ListId)

	if err != nil {
		return nil, err
	}

	var itemToSend *list.ListItem = &list.ListItem{
		Item: &list.List{Id: item.Id, Content: item.Content},
	}

	return itemToSend, nil
}

func main() {
	var cluster *gocql.ClusterConfig = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "grpcv3"
	var dbInstance DB = DB{ClusterInstance: cluster}

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	srv := grpc.NewServer()
	list.RegisterListServiceServer(srv, &listServer{ClusterInstance: dbInstance.ClusterInstance})
	reflection.Register(srv)

	fmt.Println("GRPC Server!")
	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
