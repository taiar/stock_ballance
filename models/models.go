package models

import (
    "github.com/Rhymond/go-money"
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

func (wallet Wallet) Total() int64 {
    var total int64
    for _, asset := range wallet.Assets {
        total = total + (int64(asset.Value * asset.Amount))
    }
    return total
}

func (wallet Wallet) FormattedTotal() string {
    return money.New(wallet.Total(), "BRL").Display()
}

type Quotation struct {
    gorm.Model

    Value       int
    Time        time.Time
    TimeZone    string

    StockID     uint
    Stock       Stock
}
