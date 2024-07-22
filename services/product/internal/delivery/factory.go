package delivery

import (
	"github.com/YogiTan00/Reseller/pkg/logger"
	productPb "github.com/YogiTan00/Reseller/proto/_generated/product"
	"github.com/YogiTan00/Reseller/services/product/domain/usecase"
	"google.golang.org/grpc"
)

type ProductHandler struct {
	l       logger.Logger
	srv     *grpc.Server
	product usecase.ProductUsecaseInterface
}

type ProductHandlerFactory struct {
	L       logger.Logger
	Srv     *grpc.Server
	Product usecase.ProductUsecaseInterface
}

func (prod *ProductHandlerFactory) Create() {
	productPb.RegisterProductServiceServer(prod.Srv, &ProductHandler{
		l:       prod.L,
		srv:     prod.Srv,
		product: prod.Product,
	})
}
