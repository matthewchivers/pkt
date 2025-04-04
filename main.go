package main

import (
	"log"

	"github.com/alecthomas/kong"
	"github.com/matthewchivers/pkt/cmd"
)

type CLI struct {
	Trace cmd.TraceCmd `cmd:"" help:"Trace the route to a host using ICMP."`
}

func main() {
	var cli CLI
	kctx := kong.Parse(&cli,
		kong.Name("pkt"),
		kong.Description("A minimal network diagnostics toolkit."),
		kong.UsageOnError(),
	)

	err := kctx.Run()
	if err != nil {
		log.Fatal(err)
	}
}
