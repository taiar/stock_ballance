package main

import (
    "fmt"
    "github.com/joho/godotenv"
    "github.com/taiar/stock_ballance/cmd"
    "github.com/taiar/stock_ballance/database"
    "github.com/taiar/stock_ballance/migrations"
)

func main() {
    _ = godotenv.Load()
    database.Init()

    cmd.Init()
    switch cmd.Mode() {
    case cmd.Populate:
        fmt.Println("Create a new", *cmd.Entity)
        cmd.PopulateHandler()
    case cmd.Server:
        // run web application
        fmt.Println("Run web server")
    case cmd.Migrate:
        migrations.Migrate()
    default:
        fmt.Println(cmd.Mode())
    }

    //client := av.NewClientConnection(os.Getenv("ALPHA_VANTAGE_API_KEY"), av.NewConnection())
    //data, _ := client.StockTimeSeries(av.TimeSeriesDaily, "BIDI4.SAO")

    //fmt.Println(data)

    //fmt.Println(*cmd.Mode)
}
