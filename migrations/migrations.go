package migrations

import (
    "fmt"
    "github.com/taiar/stock_ballance/database"
    "github.com/taiar/stock_ballance/models"
)

func Migrate() {
    fmt.Print("Migrating...")
    database.DBCon.AutoMigrate(models.Stock{})
    database.DBCon.AutoMigrate(models.Quotation{}).AddForeignKey("stock_id", "stocks(id)", "CASCADE", "RESTRICT")
    database.DBCon.AutoMigrate(models.Wallet{})
    database.DBCon.AutoMigrate(models.Asset{}).AddForeignKey("stock_id", "stocks(id)", "CASCADE", "RESTRICT").AddForeignKey("wallet_id", "wallets(id)", "RESTRICT", "RESTRICT")
    fmt.Println(" done!")
}
