package main

import (
    "fmt"
    "github.com/joho/godotenv"
    "github.com/taiar/stock_ballance/cmd"
    "github.com/taiar/stock_ballance/database"
    "github.com/taiar/stock_ballance/models"
)

func main() {
    _ = godotenv.Load()
    database.Init()

    test()

    cmd.Init()
    cmd.Run()

    //client := av.NewClientConnection(os.Getenv("ALPHA_VANTAGE_API_KEY"), av.NewConnection())
    //data, _ := client.StockTimeSeries(av.TimeSeriesDaily, "BIDI4.SAO")
}

func test() {
    var wallet models.Wallet
    database.DBCon.Preload("Assets").First(&wallet, 1)
    fmt.Println(wallet.FormattedTotal())
}
