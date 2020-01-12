package main

import (
    "github.com/joho/godotenv"
    "github.com/taiar/stock_ballance/cmd"
    "github.com/taiar/stock_ballance/database"
    "github.com/taiar/stock_ballance/migrations"
)

func main() {
    cmd.Init()
    _ = godotenv.Load()
    database.Init()
    migrations.Migrate()


    //client := av.NewClientConnection(os.Getenv("ALPHA_VANTAGE_API_KEY"), av.NewConnection())
    //data, _ := client.StockTimeSeries(av.TimeSeriesDaily, "BIDI4.SAO")

    //fmt.Println(data)

    //fmt.Println(*cmd.Mode)
}
