package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/tham-ph/cunewbie-search-poc/search/src/database"
	"github.com/tham-ph/cunewbie-search-poc/search/src/pb"
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
	lis, err := net.Listen("tcp", ":8080")
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

	fmt.Println("server started listening on :8080")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}