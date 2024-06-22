package main

import (
	"context"
	"log"
	"time"

	pb "52HW/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost: 9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSalesServiceClient(conn)

	stream, err := c.StreamSalesTransactions(context.Background())
	if err != nil {
		log.Fatalf("could not stream: %v", err)
	}

	transactions := []*pb.SalesTransaction{
		{TransactionId: "1", ProductId: "P1", Quantity: 1, Price: 10.0, Timestamp: time.Now().Unix()},
		{TransactionId: "2", ProductId: "P2", Quantity: 2, Price: 20.0, Timestamp: time.Now().Unix()},
	}

	for _, transaction := range transactions {
		if err := stream.Send(transaction); err != nil {
			log.Fatalf("could not send transaction: %v", err)
		}
	}

	summary, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not receive summary: %v", err)
	}

	log.Printf("Summary: %v transactions, total amount: $%v", summary.TotalTransactions, summary.TotalAmount)
}
