package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/brewinski/grcp-learning/go-grpc/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:9000", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)

	limit := make(chan struct{}, 50)

	for i := 0; i < 1000000000; i++ {
		limit <- struct{}{}
		go func() {
			start := time.Now()
			// Contact the server and print out its response.
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			defer cancel()

			r, err := c.ReadChatByID(ctx, &chat.GetMessageRequest{Id: 1})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}

			log.Printf("Greeting: %s", r.GetBody())

			elapsed := time.Since(start)
			fmt.Printf("Request took %s\n", elapsed)
			<-limit
		}()
	}
}
