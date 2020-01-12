package migrations

import (
    "github.com/taiar/stock_ballance/database"
    "github.com/taiar/stock_ballance/models"
)

func Migrate() {
    database.DBCon.AutoMigrate(
        models.Asset{},
        models.Quotation{},
        models.Stock{},
        models.Wallet{})
}
