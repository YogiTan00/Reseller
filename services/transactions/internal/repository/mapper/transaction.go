package mapper

import (
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
	"github.com/YogiTan00/Reseller/services/transactions/internal/repository/models"
)

func ToModelTransaction(p *entity.Transaction) *models.Transaction {
	return &models.Transaction{
		Id:          p.GetId(),
		IdName:      p.GetIdName(),
		Return:      p.GetReturnItem(),
		SalesDate:   p.GetSalesDate(),
		Unit:        p.GetUnit(),
		Description: p.GetDescription(),
	}

}

func ToEntityTransaction(p *models.Transaction) *entity.TransactionDto {
	return &entity.TransactionDto{
		Id:          p.Id,
		IdName:      p.IdName,
		ReturnItem:  p.Return,
		SalesDate:   p.SalesDate,
		Unit:        p.Unit,
		Description: p.Description,
	}
}

func ToListEntityProduct(data []*models.Transaction) []*entity.TransactionDto {
	listProduct := make([]*entity.TransactionDto, 0)
	for _, v := range data {
		listProduct = append(listProduct, ToEntityTransaction(v))
	}
	return listProduct
}
