package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/tham-ph/cunewbie-search-poc/search-service/src/database"
	"github.com/tham-ph/cunewbie-search-poc/search-service/src/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type SearchServiceServer struct {
	pb.SearchServiceServer
}

func (s *SearchServiceServer) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{
		Name:   "Search Service",
		Number: 0,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSearchServiceServer(grpcServer, &SearchServiceServer{})

	es := database.ConnectElasticSearch()

	res, err := es.Create("students", "21", esutil.NewJSONReader(map[string]interface{}{"name": "Poom", "age": 33}))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)

	// query for search
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"name": "Poom",
			},
		},
	}

	res, err = es.Search(es.Search.WithIndex("students"), es.Search.WithBody(esutil.NewJSONReader(query)))
	log.Println(res)

	// rabbitmq
	rabbitmqConnection, err := database.ConnectRabbitMQ()
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitmqConnection.Close()

	rabbitmqChannel, err := rabbitmqConnection.Channel()
	if err != nil {
		log.Fatal(err)
	}

	q, err := rabbitmqChannel.QueueDeclare("queue1", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	messages, err := rabbitmqChannel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		fmt.Println("server started listening on :3001")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	forever := make(chan bool)
	go func() {
		for message := range messages {
			log.Println(string(message.Body))
		}
	}()
	<-forever
}
