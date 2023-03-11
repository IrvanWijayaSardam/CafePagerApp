package config

import (
	// "fmt"
	// "os"

	"github.com/IrvanWijayaSardam/mynotesapp/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetupDatabaseConnection is creating a new connection to our database
func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	// dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASS")
	// dbHost := os.Getenv("DB_HOST")
	// dbName := os.Getenv("DB_NAME")

	//dsn := fmt.Sprintf("%s:%s@tcp(%s:7596)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open("mysql -hcontainers-us-west-105.railway.app -uroot -pEGeAWNG2In2EWfptYa7r --port 7596 --protocol=TCP railway"), &gorm.Config{})

	if err != nil {
		panic("Failed to create connection to database")
	}

	//Isi model / table disini
	db.AutoMigrate(&entity.Notes{}, &entity.User{}, &entity.Pager{})
	return db
}

// CloseDatabaseConnection Close database connection
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
