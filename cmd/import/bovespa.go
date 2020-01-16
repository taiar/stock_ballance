package _import

var HeaderRegType       = make([]byte, 2)
var HeaderFileName      = make([]byte, 13)
var HeaderOriginCode    = make([]byte, 8)
var HeaderGenDate       = make([]byte, 8)
var HeaderOffset        = make([]byte, 8)

var LineTypeReg        = make([]byte, 2)
var LineMarketDate     = make([]byte, 8)
var LineBdiCode        = make([]byte, 2)
var LineStockCode      = make([]byte, 12)
var LineMarketType     = make([]byte, 3)
var LineCompanyName    = make([]byte, 12)
var LineStockSpec      = make([]byte, 10)
var LineTermDays       = make([]byte, 10)
var LineCurrency       = make([]byte, 4)
var LineOpeningPrice   = make([]byte, 11)
var LineMaxPrice       = make([]byte, 11)
var LineMinPrice       = make([]byte, 11)
var LineMedPrice       = make([]byte, 11)
var LineLastPrice      = make([]byte, 11)
var LineBestBuyPrice   = make([]byte, 11)
var LineBestSellPrice  = make([]byte, 11)
var LineNegTotal       = make([]byte, 5)
var LineTitTotal       = make([]byte, 18)
var LineVolumeTotal    = make([]byte, 16)

var FooterTypeReg        = make([]byte, 2)

var LineSpec = map[string]int{
 "LineTypeReg":        2,
 "LineMarketDate":     8,
 "LineBdiCode":        2,
 "LineStockCode":      12,
 "LineMarketType":     3,
 "LineCompanyName":    12,
 "LineStockSpec":      10,
 "LineTermDays":       10,
 "LineCurrency":       4,
 "LineOpeningPrice":   11,
 "LineMaxPrice":       11,
 "LineMinPrice":       11,
 "LineMedPrice":       11,
 "LineLastPrice":      11,
 "LineBestBuyPrice":   11,
 "LineBestSellPrice":  11,
 "LineNegTotal":       5,
 "LineTitTotal":       18,
 "LineVolumeTotal":    16,
}

var HeaderSpec = map[string]int{
 "HeaderRegType":    2,
 "HeaderFileName":   13,
 "HeaderOriginCode": 8,
 "HeaderGenDate":    8,
 "HeaderOffset":     8,
}
