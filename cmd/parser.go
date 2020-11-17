package cmd

import (
    "flag"
    "fmt"
    "github.com/taiar/stock_ballance/migrations"
)

var (
    Command *string

    // Populate param
    Entity *string

    // Populate Stock params
    Code *string
    Name *string

    // Populate Wallet params
    Description *string

    // Populate Asset params
    Amount *uint
    StockId *uint
    WalletID *uint
    Value *uint

    // Import params
    FileInput *string
)
const Migrate   = "migrate"
const Server    = "server"
const Populate  = "populate"
const Import    = "import"

func Init() {
    Command     = flag.String("mode", "", "mode selection [migrate|populate|server|import]")
    Entity      = flag.String("entity", "", "set entity name for creation")
    Code        = flag.String("code", "", "set stock code")
    Name        = flag.String("name", "", "set stock name")
    Description = flag.String("description", "", "set wallet description")
    Amount      = flag.Uint("amount", 0, "set asset amount")
    StockId     = flag.Uint("stock_id", 0, "set asset stock id")
    WalletID    = flag.Uint("wallet_id", 0, "set asset wallet id")
    Value       = flag.Uint("value", 0, "set asset current value")
    FileInput   = flag.String("file_input", "", "file path to be imported")

    flag.Parse()
}

func Run() {
    switch *Command {
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
