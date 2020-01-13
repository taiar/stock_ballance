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
    Amount *uint
    StockId *uint
    WalletID *uint
    Value *uint
)
const Migrate = "migrate"
const Server = "server"
const Populate = "populate"

func Init() {
    Command = flag.String("mode", "server", "mode selection [migrate|populate|server]")
    Entity = flag.String("entity", "", "set entity name for creation")
    Code = flag.String("code", "", "set stock code")
    Name = flag.String("name", "", "set stock name")
    Description = flag.String("description", "", "set wallet description")
    Amount = flag.Uint("amount", 0, "set asset amount")
    StockId = flag.Uint("stock_id", 0, "set asset stock id")
    WalletID = flag.Uint("wallet_id", 0, "set asset wallet id")
    Value = flag.Uint("value", 0, "set asset current value")
    flag.Parse()
}

func PopulateHandler() {
    switch *Entity {
    case "stock":
        stock := models.Stock{Code: *Code, Name: *Name}
        createAndPrintEntity(&stock)
        fmt.Println(stock.ID)
    case "wallet":
        wallet := models.Wallet{Description: *Description}
        createAndPrintEntity(&wallet)
        fmt.Println(wallet.ID)
    case "asset":
        asset := models.Asset{ Amount: *Amount, Value:  *Value, StockID: *StockId, WalletID: *WalletID }
        createAndPrintEntity(&asset)
        fmt.Println(asset.ID)
    default:
        fmt.Println("Unknown entity", *Entity)
    }
}

func Mode() string {
    return *Command
}

func createAndPrintEntity(entity interface{}) {
    database.DBCon.Create(entity)
    fmt.Printf("%+v\n", entity)
}
