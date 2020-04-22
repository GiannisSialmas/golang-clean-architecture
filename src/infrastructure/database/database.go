package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"application/models"

	"github.com/jinzhu/gorm"
	// Postgres driver to be used by gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type databaseConfigType struct {
	port     int
	host     string
	user     string
	password string
	dbname   string
}

// NewDB creates a new client to a postgres database and returns it for use by other components of the application
func NewDB() *gorm.DB {

	dbConfig := config()
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", dbConfig.host, dbConfig.port, dbConfig.user, dbConfig.password, dbConfig.dbname)

	var db *gorm.DB
	for i := 0; i < 10; i++ {
		dbTemp, err := gorm.Open("postgres", connStr)
		if err != nil {
			log.Println("Error connecting to the database: ", err)
			log.Println("Retrying in 3 seconds.")
			if i == 10 {
				log.Fatalln("Error connecting to the database: ", err)
			}
			time.Sleep(3 * time.Second)
		} else {
			log.Println("Succesfully connected to the database")
			db = dbTemp
			break
		}
	}
	db.SingularTable(true)
	db.AutoMigrate(&models.User{})
	return db

}

// Creates a configuration struct to be used to connect to the postgres instance.
func config() databaseConfigType {
	var databaseConfig databaseConfigType

	// db configuration
	if db, ok := os.LookupEnv("DB_NAME"); ok {
		databaseConfig.dbname = db
	} else {
		log.Fatalf("No database name provided")
	}

	// password configuration
	if password, ok := os.LookupEnv("DB_PASSWORD"); ok {
		databaseConfig.password = password
	} else {
		log.Fatalf("No user password provided")
	}

	// Port configuration
	if portString, ok := os.LookupEnv("DB_PORT"); ok {
		port, err := strconv.Atoi(portString)
		if err != nil {
			log.Fatal("Invalid port", port)
		}
		databaseConfig.port = port
	} else {
		databaseConfig.port = 5432
	}

	// host configuration
	if host, ok := os.LookupEnv("DB_HOST"); ok {
		databaseConfig.host = host
	} else {
		databaseConfig.host = "127.0.0.1"
	}

	// user configuration
	if user, ok := os.LookupEnv("DB_USER"); ok {
		databaseConfig.user = user
	} else {
		databaseConfig.user = "postgres"
	}

	log.Println(fmt.Sprintf("Connection string is: host=%s port=%d user=%s "+"password=***** dbname=%s sslmode=disable", databaseConfig.host, databaseConfig.port, databaseConfig.user, databaseConfig.dbname))

	return databaseConfig
}
