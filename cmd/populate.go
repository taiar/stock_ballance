package cmd

import (
    "fmt"
    "github.com/taiar/stock_ballance/database"
    "github.com/taiar/stock_ballance/models"
)

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

func createAndPrintEntity(entity interface{}) {
    database.DBCon.Create(entity)
    fmt.Printf("%+v\n", entity)
}
