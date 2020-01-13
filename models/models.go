package models

import (
    "github.com/jinzhu/gorm"
    "time"
)

type Stock struct {
    gorm.Model
    Name string `gorm:"required"`
    Code string `gorm:"required"`
}

type Asset struct {
    gorm.Model
    Amount      int         `gorm:"required"`
    Stock       Stock       `gorm:"required"`
    Time        time.Time   `gorm:"required"`
    TimeZone    string      `gorm:"required"`
}

type Wallet struct {
    gorm.Model
    Description string
    Assets      []Asset
}

type Quotation struct {
    gorm.Model
    Stock       Stock
    Value       int         `gorm:"required"`
    Time        time.Time   `gorm:"required"`
    TimeZone    string      `gorm:"required"`
}
