package test

import (
	"testing"

	"github.com/nkien0204/lets-go/internal/infrastructure/db/rdb/mysql"
	"github.com/nkien0204/lets-go/internal/infrastructure/db/rdb/mysql/models"
)

func TestGetPayments(t *testing.T) {
	mysqlService, err := mysql.NewMysqlConnection("user:pass@tcp(127.0.0.1:3306)/classicmodels")
	if err != nil {
		t.Errorf("create new connection failed: %v", err)
		return
	}

	var payment []models.Payment
	if result := mysqlService.Table(models.PaymentsTable).Limit(10).Offset(10).Find(&payment); result.Error != nil {
		t.Errorf("mysqlService.Db.Table(models.PaymentsTable).Find: %v", result.Error)
		return
	}
	t.Logf("TestGetPayments OK: %v", payment[0].PaymentDate.Unix())
}
