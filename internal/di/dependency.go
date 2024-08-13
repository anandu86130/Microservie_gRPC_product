package di

import (
	"github.com/anandu86130/Microservice_gRPC_product/config"
	"github.com/anandu86130/Microservice_gRPC_product/internal/db"
	"github.com/anandu86130/Microservice_gRPC_product/internal/handlers"
	"github.com/anandu86130/Microservice_gRPC_product/internal/repositories"
	"github.com/anandu86130/Microservice_gRPC_product/internal/server"
	"github.com/anandu86130/Microservice_gRPC_product/internal/services"
)

func Init() {
	config.LoadConfig()
	db := db.ConnectDB()

	productRepo := repositories.NewProductRepo(db)
	productsvc := services.NewUserService(productRepo)
	ProductHandle := handlers.NewProductHandler(productsvc)
	server.NewGrpcServer(ProductHandle)
}
