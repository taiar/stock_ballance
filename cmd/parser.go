package cmd

import (
    "flag"
)

var (
    Command *string
)
const Migrate = "migrate"
const Server = "server"
const Populate = "populate"

func Init() {
    Command = flag.String("mode", "server", "mode selection")
    flag.Parse()
}

func Mode() string {
    return *Command
}
