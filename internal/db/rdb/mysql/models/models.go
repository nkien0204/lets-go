package models

import (
	"time"
)

const PaymentsTable string = "payments"

type Payment struct {
	CustomerNumber int       `gorm:"column:customerNumber"`
	CheckNumber    string    `gorm:"column:checkNumber"`
	PaymentDate    time.Time `gorm:"column:paymentDate"`
	Amount         float32   `gorm:"type:decimal(10,2);"`
}
