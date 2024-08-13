package repositories

import (
	"github.com/anandu86130/Microservice_gRPC_product/internal/model"
	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}

func (u *ProductRepo) CreateProduct(product *model.Product) error {
	if err := u.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (u *ProductRepo) FetchByProductID(id uint) (*model.Product, error) {
	var product model.Product
	err := u.DB.First(&product, id).Error
	return &product, err
}

func (u *ProductRepo) FetchByName(name string) (*model.Product, error) {
	var product model.Product
	err := u.DB.Where("name=?", name).First(&product).Error
	return &product, err
}

func (u *ProductRepo) EditProduct(product *model.Product) error {
	err := u.DB.Save(product).Error
	return err
}

func (u *ProductRepo) DeleteProduct(id uint) error {
	err := u.DB.Delete(&model.Product{}, id).Error
	return err
}

func (u *ProductRepo) FetchProducts() ([]*model.Product, error) {
	var product []*model.Product
	err := u.DB.Find(&product).Error
	return product, err
}
