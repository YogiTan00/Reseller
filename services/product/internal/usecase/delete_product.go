package usecase

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/google/uuid"
)

func (p ProductUsecase) DeleteProduct(ctx context.Context, productId string) error {
	err := uuid.Validate(productId)
	if err != nil {
		return err
	}

	err = p.repoProduct.DeleteProduct(ctx, productId)
	if err != nil {
		p.l.Error(err)
		return exceptions.ErrInternalServer
	}

	return nil
}
