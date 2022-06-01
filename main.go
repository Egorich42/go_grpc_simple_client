package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	proto "github.com/Egorich42/app_proto"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &proto.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
