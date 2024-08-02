package db

import (
	"log"

	"github.com/kchxng/ecom-api/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	//"host=127.0.0.1 user=postgres password=password dbname=db_name port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	dsn := configs.GetENV("CONNECT_STR")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("[GIN] Connect to database failed!\n")
		log.Fatal(err)
	}
	log.Printf("[GIN] Connected to database...\n")
	// db.AutoMigrate(&models.Users{})
	// db.AutoMigrate(&models.Roles{})
	// db.AutoMigrate(&models.Permissions{})
	// db.AutoMigrate(&models.RolesPermissions{})
	// db.AutoMigrate(&models.UsersRoles{})
	// db.AutoMigrate(&models.RolesPermissions{})
	return db
}
