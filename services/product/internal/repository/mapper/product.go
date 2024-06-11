package mapper

import (
	"Reseller/services/product/domain/entity"
	"Reseller/services/product/internal/repositroy/models"
)

func ToModelProduct(data entity.Product) *models.Product {
	return &models.Product{
		Id:           data.GetId(),
		Name:         data.GetName(),
		TypeSize:     data.GetTypeSize(),
		MarketType:   data.GetMarketType(),
		Image:        data.GetImage(),
		DefaultPrice: data.GetDefaultPrice(),
		Price:        data.GetPrice(),
		CreatedAt:    data.GetCreatedAt(),
		UpdatedAt:    data.GetUpdatedAt(),
	}
}

func ToEntityProduct(data *models.Product) *entity.ProductDto {
	return &entity.ProductDto{
		Id:           data.Id,
		Name:         data.Name,
		TypeSize:     data.TypeSize,
		MarketType:   data.MarketType,
		Image:        data.Image,
		DefaultPrice: data.DefaultPrice,
		Price:        data.Price,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func ToListEntityProduct(data []*models.Product) []*entity.ProductDto {
	listProduct := make([]*entity.ProductDto, 0)
	for _, v := range data {
		listProduct = append(listProduct, ToEntityProduct(v))
	}
	return listProduct
}
