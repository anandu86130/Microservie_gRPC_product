package interfaces

import "github.com/anandu86130/Microservice_gRPC_product/internal/model"

type ProductRepoInter interface {
	CreateProduct(product *model.Product) error
	FetchByProductID(id uint) (*model.Product, error)
	EditProduct(product *model.Product) error
	FetchByName(name string) (*model.Product, error)
	FetchProducts() ([]*model.Product, error)
}
