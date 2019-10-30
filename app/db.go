package app

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db = &gorm.DB{}
var err error

// Initialises the database
func DBInit() {
	env := godotenv.Load()
	if env != nil {
		fmt.Print(env)
	}
	appEnv := os.Getenv("APP_ENV")

	// Deploying on heroku
	if appEnv == "production" {
		databaseURL := os.Getenv("DATABASE_URL")
		db, err = gorm.Open("postgres", databaseURL)
	} else {
		username := os.Getenv("db_user")
		password := os.Getenv("db_pass")
		databaseName := os.Getenv("db_name")
		host := os.Getenv("db_host")
		dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, username, databaseName, password)
		fmt.Println(dbURI)
		db, err = gorm.Open("postgres", dbURI)
	}

	if err != nil {
		fmt.Print(err)
	}

	db.Debug().AutoMigrate(&User{})

}

// Return *gorm.DB image
func GetDB() *gorm.DB {
	return db
}
