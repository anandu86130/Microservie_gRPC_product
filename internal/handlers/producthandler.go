package handlers

import (
	"context"

	pb "github.com/anandu86130/Microservice_gRPC_product/internal/pb"
	services_interfaces "github.com/anandu86130/Microservice_gRPC_product/internal/services/interfaces"
)

type ProductHandlers struct{
	pb.ProductServiceServer
	svc services_interfaces.ProductServiceInter
}

func NewProductHandler(svc services_interfaces.ProductServiceInter) *ProductHandlers{
	return &ProductHandlers{
		svc: svc,
	}
}

func (m *ProductHandlers) CreateProduct(ctx context.Context, p *pb.ProductDetails) (*pb.ProductResponse, error){
	result, err := m.svc.CreateProductService(p)
	if err != nil{
		return nil, err
	}
	return result,nil
}

func (m *ProductHandlers) FetchByProductID(ctx context.Context, p *pb.ProductID) (*pb.ProductDetails,error){
	result, err := m.svc.FetchByProductID(p)
	if err != nil{
		return nil, err
	}
	return result,nil
}

func (m *ProductHandlers) FetchByName(ctx context.Context, p *pb.ProductByName) (*pb.ProductDetails,error){
	result, err := m.svc.FetchByName(p)
	if err != nil{
		return nil,err
	}
	return result, nil
}

func (m *ProductHandlers) FetchProducts(ctx context.Context, p *pb.NoParam) (*pb.ProductList,error){
	result, err := m.svc.FetchProducts(p)
	if err != nil{
		return nil, err
	}
	return result,nil
}