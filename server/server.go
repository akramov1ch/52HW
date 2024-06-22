package main

import (
	"io"
	"log"
	"net"

	pb "52HW/gen"
	db "52HW/server/db"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedSalesServiceServer
}

func (s *server) StreamSalesTransactions(stream pb.SalesService_StreamSalesTransactionsServer) error {
	var totalAmount float32
	var totalTransactions int32

	for {
		transaction, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.SalesSummary{
				TotalAmount:       totalAmount,
				TotalTransactions: totalTransactions,
			})
		}
		if err != nil {
			return err
		}

		err = db.SaveTransaction(transaction)
		if err != nil {
			return err
		}

		totalAmount += transaction.Price * float32(transaction.Quantity)
		totalTransactions++
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSalesServiceServer(s, &server{})
	reflection.Register(s)

	log.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
