package product

import (
	"github.com/YogiTan00/Reseller/pkg/logger"
	productPb "github.com/YogiTan00/Reseller/proto/_generated/product"
	"github.com/YogiTan00/Reseller/services/transactions/domain/usecase"
	"google.golang.org/grpc"
)

type ProductService struct {
	l          logger.Logger
	prdService productPb.ProductServiceClient
}

type ProductServiceFactory struct {
	L    logger.Logger
	Conn *grpc.ClientConn
}

func (prod *ProductServiceFactory) Create() usecase.ProductServiceInterface {
	prd := productPb.NewProductServiceClient(prod.Conn)
	return &ProductService{
		l:          prod.L,
		prdService: prd,
	}
}
