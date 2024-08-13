package services_interfaces

import (
	pb "github.com/anandu86130/Microservice_gRPC_product/internal/pb"
)

type ProductServiceInter interface{
	CreateProductService(p *pb.ProductDetails) (*pb.ProductResponse,error)
	FetchByProductID(p *pb.ProductID) (*pb.ProductDetails,error)
	FetchByName(p *pb.ProductByName) (*pb.ProductDetails,error)
	FetchProducts(p *pb.NoParam) (*pb.ProductList,error)
}