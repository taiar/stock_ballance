package models

import (
    "github.com/jinzhu/gorm"
)

type Stock struct {
    gorm.Model

    Name string
    Code string
}
