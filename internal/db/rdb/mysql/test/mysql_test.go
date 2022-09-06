package test

import (
	"testing"

	"github.com/nkien0204/projectTemplate/internal/db/rdb/mysql"
	"github.com/nkien0204/projectTemplate/internal/db/rdb/mysql/models"
)

func TestInitConnection(t *testing.T) {
	mysqlService, err := mysql.InitMysqlConnection("ntkien11:1@tcp(127.0.0.1:3306)/classicmodels")
	if err != nil {
		t.Errorf("mysql.InitMysqlConnection: %v", err.Error())
		return
	}
	t.Logf("init connection successfully: %v", mysqlService.Address)
}

func TestGetPayments(t *testing.T) {
	mysqlService, err := mysql.InitMysqlConnection("ntkien11:1@tcp(127.0.0.1:3306)/classicmodels")
	if err != nil {
		t.Errorf("mysql.InitMysqlConnection: %v", err.Error())
		return
	}

	var payment []models.Payment
	if result := mysqlService.Db.Table(models.PaymentsTable).Limit(10).Offset(10).Find(&payment); result.Error != nil {
		t.Errorf("mysqlService.Db.Table(models.PaymentsTable).Find: %v", result.Error)
		return
	}
	t.Logf("TestGetPayments OK: %v", payment[0].PaymentDate.Unix())
}
