package config

import (
	"fmt"

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

	dsn := fmt.Sprintf("mysql://root:2KDqdTYvyVpV42Z22uO4@containers-us-west-105.railway.app:7596/railway")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

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
