package database

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "os"
)

var DBCon *gorm.DB

func Init() {
    var err error

    connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
        os.Getenv("DATABASE_HOST"),
        os.Getenv("DATABASE_PORT"),
        os.Getenv("DATABASE_USER"),
        os.Getenv("DATABASE_NAME"),
        os.Getenv("DATABASE_PASSWORD"))

    DBCon, err = gorm.Open("postgres", connectionString)

    if err != nil {
        panic(fmt.Sprintf("failed to connect database: %s", err))
    }
}
