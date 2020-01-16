package models

import (
    "github.com/jinzhu/gorm"
    "time"
)

type Asset struct {
    gorm.Model

    Amount      uint
    Value       uint
    Time        time.Time
    TimeZone    string

    StockID     uint
    Stock       Stock

    WalletID    uint
    Wallet      Wallet
}
