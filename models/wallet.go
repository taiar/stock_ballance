package models

import (
    "github.com/Rhymond/go-money"
    "github.com/jinzhu/gorm"
)

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
