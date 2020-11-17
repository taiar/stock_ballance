package cmd

import (
    "fmt"
    "github.com/taiar/stock_ballance/database"
    "github.com/taiar/stock_ballance/models"
    "strings"
)

func PopulateHandler() {
    switch *Entity {
    case "stock":
        stock := models.Stock{Code: strings.ToUpper(*Code), Name: *Name}
        createAndPrintEntity(&stock)
        fmt.Println(stock.ID)
    case "wallet":
        wallet := models.Wallet{Description: *Description}
        createAndPrintEntity(&wallet)
        fmt.Println(wallet.ID)
    case "asset":
        var wallet models.Wallet
        var stock models.Stock

        database.DBCon.Find(&wallet, *WalletId)
        database.DBCon.Find(&stock, *StockId)

        asset := models.Asset{ Amount: *Amount, Value:  *Value, Stock: stock, Wallet: wallet }
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
