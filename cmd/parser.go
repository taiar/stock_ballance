package cmd

import (
    "flag"
    "fmt"
    "github.com/taiar/stock_ballance/database"
    "github.com/taiar/stock_ballance/models"
)

var (
    Command *string
    Entity *string

    Code *string
    Name *string
    Description *string
)
const Migrate = "migrate"
const Server = "server"
const Populate = "populate"

func Init() {
    Command = flag.String("mode", "server", "mode selection")
    Entity = flag.String("entity", "", "entity for creation")
    Code = flag.String("code", "", "entity for creation")
    Name = flag.String("name", "", "entity for creation")
    Description = flag.String("description", "", "entity for creation")
    flag.Parse()
}

func PopulateHandler() {
    switch *Entity {
    case "stock":
        stock := models.Stock{Code: *Code, Name: *Name}
        database.DBCon.Create(&stock)
        printEntity(stock)
    case "wallet":
        wallet := models.Wallet{Description: *Description}
        database.DBCon.Create(&wallet)
        printEntity(wallet)
    default:
        fmt.Println("Unknown entity", *Entity)
    }
}

func Mode() string {
    return *Command
}

func printEntity(entity interface{}) {
    fmt.Printf("%+v\n", entity)
}
