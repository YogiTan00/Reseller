package delivery

import (
	"Reseller/pkg/logger"
	productPb "Reseller/proto/_generated/product"
	"Reseller/services/product/domain/usecase"
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
