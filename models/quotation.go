package models

import (
    "github.com/jinzhu/gorm"
    "time"
)

type Quotation struct {
    gorm.Model

    Value       int
    Time        time.Time
    TimeZone    string

    StockID     uint
    Stock       Stock
}
