package product

import (
	"github.com/YogiTan00/Reseller/services/product/internal/delivery"
	"github.com/YogiTan00/Reseller/services/product/internal/repository"
	"github.com/YogiTan00/Reseller/services/product/internal/usecase"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type ProductHandler struct {
	db  *gorm.DB
	srv *grpc.Server
}

type ProductHandlerFactory struct {
	Db  *gorm.DB
	Srv *grpc.Server
}

func (prod *ProductHandlerFactory) Create() {
	repoProduct := &repository.ProductRepositoryFactory{
		Db: *prod.Db,
	}
	ucProduct := &usecase.ProductUsecaseFactory{
		RepoProduct: repoProduct.Create(),
	}
	srvProduct := &delivery.ProductHandlerFactory{
		Srv:     prod.Srv,
		Product: ucProduct.Create(),
	}
	srvProduct.Create()

}
