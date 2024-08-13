package entity

import (
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/google/uuid"
	"time"
)

type Product struct {
	id           string
	name         string
	typeSize     string
	image        string
	defaultPrice int64
	price        int64
	createdAt    time.Time
	updatedAt    time.Time
	deletedAt    *time.Time
}

type ProductDto struct {
	Id           string
	Name         string
	TypeSize     string
	Image        string
	DefaultPrice int64
	Price        int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

func (data *Product) GetId() string {
	return data.id
}
func (data *Product) GetName() string {
	return data.name
}
func (data *Product) GetTypeSize() string {
	return data.typeSize
}
func (data *Product) GetImage() string {
	return data.image
}
func (data *Product) GetDefaultPrice() int64 {
	return data.defaultPrice
}
func (data *Product) GetPrice() int64 {
	return data.price
}
func (data *Product) GetCreatedAt() time.Time {
	return data.createdAt
}
func (data *Product) GetUpdatedAt() time.Time {
	return data.updatedAt
}

func (dto *ProductDto) New() *Product {
	timeNow := time.Now()
	return &Product{
		id:           uuid.New().String(),
		name:         dto.Name,
		typeSize:     dto.TypeSize,
		image:        dto.Image,
		defaultPrice: dto.DefaultPrice,
		price:        dto.Price,
		createdAt:    timeNow,
		updatedAt:    timeNow,
	}
}

func (dto *ProductDto) Update(data *ProductDto) *Product {
	timeNow := time.Now()
	if data.Name != "" {
		dto.Name = data.Name
	}

	if data.TypeSize != "" {
		dto.TypeSize = data.TypeSize
	}

	if data.Image != "" {
		dto.Image = data.Image
	}

	if data.DefaultPrice != 0 {
		dto.DefaultPrice = data.DefaultPrice
	}

	if data.Price != 0 {
		dto.Price = data.Price
	}

	dto.UpdatedAt = timeNow

	return dto.Create()
}

func (dto *ProductDto) Create() *Product {
	return &Product{
		id:           dto.Id,
		name:         dto.Name,
		typeSize:     dto.TypeSize,
		image:        dto.Image,
		defaultPrice: dto.DefaultPrice,
		price:        dto.Price,
		createdAt:    dto.CreatedAt,
		updatedAt:    dto.UpdatedAt,
	}
}

func (data *ProductDto) Validate() error {
	if data.Name == "" {
		return exceptions.ErrRequired("name")
	}

	if data.TypeSize == "" {
		return exceptions.ErrRequired("typeSize")
	}

	if data.DefaultPrice == 0 {
		return exceptions.ErrRequired("defaultPrice")
	}

	if data.Price == 0 {
		return exceptions.ErrRequired("price")
	}

	return nil
}
