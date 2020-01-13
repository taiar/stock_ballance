package migrations

import (
    "fmt"
    "github.com/taiar/stock_ballance/database"
    "github.com/taiar/stock_ballance/models"
)

func Migrate() {
    fmt.Print("Migrating...")
    database.DBCon.AutoMigrate(
        models.Asset{},
        models.Quotation{},
        models.Stock{},
        models.Wallet{})
    fmt.Println(" done!")
}
