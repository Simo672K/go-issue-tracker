package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Simo672K/issue-tracker/api/routes"
)

func Run() {
	mux := http.NewServeMux()
	ctx := context.Background()

	fmt.Println("Running app on port :3000")
	if err := routes.AppRoutes(&ctx, mux); err != nil {
		log.Fatalf("An error accured: %+v", err)
	}

	http.ListenAndServe(":3000", mux)
}
