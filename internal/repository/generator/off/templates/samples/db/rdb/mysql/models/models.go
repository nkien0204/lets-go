package models

import (
	"time"
)

const PaymentsTable string = "payments"
const UsersTable string = "users"

type Payment struct {
	CustomerNumber int       `gorm:"column:customerNumber"`
	CheckNumber    string    `gorm:"column:checkNumber"`
	PaymentDate    time.Time `gorm:"column:paymentDate"`
	Amount         float32   `gorm:"type:decimal(10,2);"`
}

type User struct {
	Id       string `gorm:"column:id"`
	FullName string `gorm:"column:full_name"`
	UserName string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Role     int    `gorm:"column:role"`
}
