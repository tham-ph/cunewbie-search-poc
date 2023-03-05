package main

import (
	"context"
	"github.com/tham-ph/cunewbie-search-poc/service1/src/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func callSayHello(client pb.SearchServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.SayHelloRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)

	callSayHello(client)

}
