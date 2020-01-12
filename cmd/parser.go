package cmd

import (
    "flag"
)

var (
    Mode *string
)

func Init() {
    Mode = flag.String("mode", "server", "mode selection [server|cli]")
    flag.Parse()
}
