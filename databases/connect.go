package databases

import (
	"database/sql"
    "fmt"
    "log"
    "os"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func loadEnv() {
	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
}

func connectDB() {
	// Capture connection properties.
	loadEnv()
	fmt.Println(os.Getenv("DBPASS"))
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   os.Getenv("DBHOST"),
        DBName: os.Getenv("DBNAME"),
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
	//return db
}

func GetDB() *sql.DB {
	connectDB()
	return db
}

