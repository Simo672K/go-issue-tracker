package main

import (
	"context"

	"github.com/Simo672K/issue-tracker/api/routes"
	"github.com/Simo672K/issue-tracker/cmd"
)

func main() {
	ctx := context.Background()
	config := cmd.Config{
		Addr: ":3000",
	}
	app := cmd.Application{
		Config: config,
	}

	mux, err := routes.MuxRouter(&ctx)
	if err != nil {
	}

	app.Run(mux)
}
