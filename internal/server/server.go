package server

import (
	"log"
	"net"

	"github.com/anandu86130/Microservice_gRPC_product/internal/handlers"
	pb "github.com/anandu86130/Microservice_gRPC_product/internal/pb"
	"google.golang.org/grpc"
)

func NewGrpcServer(handlr *handlers.ProductHandlers) {
	log.Println("connecting to gRPC server")
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal("error when creating listener on port 8083")
	}

	grp := grpc.NewServer()
	pb.RegisterProductServiceServer(grp, handlr)

	log.Printf("listening on gRPC server 8083")
	if err := grp.Serve(lis); err != nil {
		log.Fatal("error connecting to gRPC server")
	}
}
