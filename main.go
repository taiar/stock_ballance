package main

import (
    "fmt"
    av "github.com/cmckee-dev/go-alpha-vantage"
    "github.com/joho/godotenv"
    "github.com/taiar/stock_ballance/cmd"
    "os"
)

func main() {
    cmd.Init()
    _ = godotenv.Load()

    client := av.NewClientConnection(os.Getenv("ALPHA_VANTAGE_API_KEY"), av.NewConnection())
    data, _ := client.StockTimeSeries(av.TimeSeriesDaily, "BIDI4.SAO")

    fmt.Println(data)

    //fmt.Println(*cmd.Mode)
}
