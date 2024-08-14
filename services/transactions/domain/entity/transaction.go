package entity

import (
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	id          string
	idName      string
	name        string
	returnTrans *bool
	salesDate   *time.Time
	unit        int
	description string
	createdAt   time.Time
	updatedAt   time.Time
}

type TransactionDto struct {
	Id          string
	IdName      string
	Name        string
	ReturnTrans *bool
	SalesDate   *time.Time
	Unit        int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Transaction) GetId() string {
	return t.id
}
func (t *Transaction) GetIdName() string {
	return t.idName
}
func (t *Transaction) GetName() string {
	return t.name
}
func (t *Transaction) GetReturnTrans() *bool {
	return t.returnTrans
}
func (t *Transaction) GetSalesDate() *time.Time {
	return t.salesDate
}
func (t *Transaction) GetUnit() int {
	return t.unit
}
func (t *Transaction) GetDescription() string {
	return t.description
}
func (t *Transaction) GetCreatedAt() time.Time {
	return t.createdAt
}
func (t *Transaction) GetUpdatedAt() time.Time {
	return t.updatedAt
}

func (dto *TransactionDto) New() *Transaction {
	timeNow := time.Now()
	return &Transaction{
		id:          uuid.New().String(),
		idName:      dto.IdName,
		returnTrans: dto.ReturnTrans,
		salesDate:   dto.SalesDate,
		unit:        dto.Unit,
		description: dto.Description,
		createdAt:   timeNow,
		updatedAt:   timeNow,
	}
}

func (dto *TransactionDto) Update(data *TransactionDto) *Transaction {
	timeNow := time.Now()
	if data.IdName != "" {
		dto.IdName = data.IdName
	}

	if data.ReturnTrans != nil {
		dto.ReturnTrans = data.ReturnTrans
	}

	if !data.SalesDate.IsZero() {
		dto.SalesDate = data.SalesDate
	}

	if data.Unit != 0 {
		dto.Unit = data.Unit
	}

	if data.Description != "" {
		dto.Description = data.Description
	}

	dto.UpdatedAt = timeNow
	return dto.Create()
}

func (dto *TransactionDto) Create() *Transaction {
	return &Transaction{
		id:          dto.Id,
		idName:      dto.IdName,
		returnTrans: dto.ReturnTrans,
		salesDate:   dto.SalesDate,
		unit:        dto.Unit,
		description: dto.Description,
		createdAt:   dto.CreatedAt,
		updatedAt:   dto.UpdatedAt,
	}
}

func (data *TransactionDto) Validate() error {
	if data.IdName == "" {
		return exceptions.ErrRequired("idName")
	}

	if data.SalesDate.IsZero() {
		return exceptions.ErrRequired("salesDate")
	}
	return nil
}
