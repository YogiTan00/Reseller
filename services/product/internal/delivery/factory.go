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
	l       logger.Logger
	srv     *grpc.Server
	product usecase.ProductUsecaseInterface
}

func (prod *ProductHandlerFactory) Create() {
	productPb.RegisterProductServiceServer(prod.srv, &ProductHandlerFactory{
		l:       prod.l,
		srv:     prod.srv,
		product: prod.product,
	})
}
