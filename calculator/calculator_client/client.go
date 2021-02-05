package main

import (
	"context"
	"github.com/RyanYang1221/grpc-api/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cc, err := grpc.Dial("0.0.0.0:50551", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)
	doUnary(c)

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		FirstNumber: 1,
		SecondNumber: 2,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet: %v", err)
	}
	log.Printf("Response: %v", res.SumResult)
}