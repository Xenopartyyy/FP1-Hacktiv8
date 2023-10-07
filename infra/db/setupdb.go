package setupdb

import (
	"FP1-Hacktiv8/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Todos = models.Todos{}

func ConnDB() {
	dsn := "host=localhost user=postgres password=dirham dbname=todos port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&Todos)
}
