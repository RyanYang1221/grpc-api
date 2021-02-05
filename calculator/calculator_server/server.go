package main

import (
	"context"
	"fmt"
	"github.com/RyanYang1221/grpc-api/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumReponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber
	sum := firstNumber + secondNumber
	response := &calculatorpb.SumReponse{
		SumResult: sum,
	}
	return response, nil
}

func main() {
	fmt.Println("Hello!")
	lis, err := net.Listen("tcp", "0.0.0.0:50551")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
