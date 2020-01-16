package models

import (
    "github.com/jinzhu/gorm"
    "time"
)

type Quotation struct {
    gorm.Model

    Opening     int
    Max         int
    Min         int
    Med         int
    Closing     int
    BestBuy     int
    BestSell    int
    Volume      int

    Time        time.Time
    TimeZone    string

    StockID     uint
    Stock       Stock
}
