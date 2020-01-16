package cmd

import (
    "flag"
)

var (
    Command *string
    Entity *string

    Code *string
    Name *string
    Description *string
    Amount *uint
    StockId *uint
    WalletID *uint
    Value *uint
)
const Migrate = "migrate"
const Server = "server"
const Populate = "populate"

func Init() {
    Command = flag.String("mode", "", "mode selection [migrate|populate|server|import]")
    Entity = flag.String("entity", "", "set entity name for creation")
    Code = flag.String("code", "", "set stock code")
    Name = flag.String("name", "", "set stock name")
    Description = flag.String("description", "", "set wallet description")
    Amount = flag.Uint("amount", 0, "set asset amount")
    StockId = flag.Uint("stock_id", 0, "set asset stock id")
    WalletID = flag.Uint("wallet_id", 0, "set asset wallet id")
    Value = flag.Uint("value", 0, "set asset current value")
    flag.Parse()
}

func Mode() string {
    return *Command
}

func Run() {
    switch Mode() {
    case Populate:
        fmt.Println("Create a new", *Entity)
        PopulateHandler()
    case Server:
        fmt.Println("Run web server")
    case Migrate:
        migrations.Migrate()
    case Import:
        fmt.Println("Importing file", *FileInput)
        ImportHandler()
    default:
        fmt.Println("Please see the usage help below:")
        flag.PrintDefaults()
    }
}
