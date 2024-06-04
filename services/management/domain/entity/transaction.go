package entity

import "time"

type Transactin struct {
	Id                string
	TransactionNumber string
	SalesDate         time.Time
	IdName            string
	Unit              int
	Discount          float32
	AdminFee          float32
	AddAdminFee       float32
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type TransactinDto struct {
	id                string
	transactionNumber string
	salesDate         time.Time
	idName            string
	unit              int
	discount          float32
	adminFee          float32
	addAdminFee       float32
	createdAt         time.Time
	updatedAt         time.Time
}
