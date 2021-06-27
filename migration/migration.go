package migration

import (
	"log"
	"myapp/config"
	"myapp/graph/model"
)

func MigrateTable() {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.AutoMigrate(&model.User{}, &model.Item{}); err != nil {
		log.Println(err)
	}
}
