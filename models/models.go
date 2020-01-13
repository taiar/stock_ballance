package models

import (
    "github.com/jinzhu/gorm"
    "time"
)

type Stock struct {
    gorm.Model

    Name string
    Code string
}

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

type Wallet struct {
    gorm.Model

    Description string
    Assets      []Asset
}

type Quotation struct {
    gorm.Model

    Value       int
    Time        time.Time
    TimeZone    string

    StockID     uint
    Stock       Stock
}
