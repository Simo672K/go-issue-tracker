package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Simo672K/issue-tracker/api/routes"
	"github.com/Simo672K/issue-tracker/cmd"
)

func main() {
	ctx := context.Background()
	mux := http.NewServeMux()

	config := cmd.Config{
		Addr: ":3000",
	}
	app := cmd.Application{
		Config: config,
	}

	err := routes.MuxRouter(ctx, mux)
	if err != nil {
		fmt.Println("Failed at mounting routes")
		return
	}
	app.Mount(mux)
	app.Run()
}

// func main() {
// 	migration.MigrateDB()
// }
