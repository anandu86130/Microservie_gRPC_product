package services

import (
	"github.com/anandu86130/Microservice_gRPC_product/internal/model"
	pb "github.com/anandu86130/Microservice_gRPC_product/internal/pb"
	"github.com/anandu86130/Microservice_gRPC_product/internal/repositories/interfaces"
	services_interfaces "github.com/anandu86130/Microservice_gRPC_product/internal/services/interfaces"
)

type ProductService struct {
	repo interfaces.ProductRepoInter
}

func (m *ProductService) CreateProductService(p *pb.ProductDetails) (*pb.ProductResponse, error) {
	product := &model.Product{
		Category:    p.Category,
		Name:        p.Name,
		Price:       p.Price,
		Imagepath:   p.Imagepath,
		Description: p.Description,
		Size:        p.Size,
		Quantity:    p.Quantity,
	}

	err := m.repo.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		Status:  "Success",
		Error:   "",
		Message: "Product added successfully",
	}, nil
}

func (m *ProductService) FetchByProductID(p *pb.ProductID) (*pb.ProductDetails, error) {
	product, err := m.repo.FetchByProductID(uint(p.Id))
	if err != nil {
		return nil, err
	}

	return &pb.ProductDetails{
		Id:          uint32(product.ID),
		Category:    product.Category,
		Name:        product.Name,
		Price:       product.Price,
		Imagepath:   product.Imagepath,
		Description: product.Description,
		Size:        product.Size,
		Quantity:    product.Quantity,
	}, nil
}

func (m *ProductService) FetchByName(p *pb.ProductByName) (*pb.ProductDetails, error) {
	product, err := m.repo.FetchByName(p.Name)
	if err != nil {
		return nil, err
	}

	return &pb.ProductDetails{
		Id:          uint32(product.ID),
		Category:    product.Category,
		Name:        product.Name,
		Price:       product.Price,
		Imagepath:   product.Imagepath,
		Description: product.Description,
		Size:        product.Size,
		Quantity:    product.Quantity,
	}, nil
}

func (m *ProductService) FetchProducts(p *pb.NoParam) (*pb.ProductList, error) {
	result, err := m.repo.FetchProducts()
	if err != nil {
		return nil, err
	}

	var products []*pb.ProductDetails
	for _, product := range result {
		products = append(products, &pb.ProductDetails{
			Id:          uint32(product.ID),
			Category:    product.Category,
			Name:        product.Name,
			Price:       product.Price,
			Imagepath:   product.Imagepath,
			Description: product.Description,
			Size:        product.Size,
			Quantity:    product.Quantity,
		})
	}
	return &pb.ProductList{
		Item: products,
	}, nil
}

func NewUserService(repo interfaces.ProductRepoInter) services_interfaces.ProductServiceInter {
	return &ProductService{
		repo: repo,
	}
}
