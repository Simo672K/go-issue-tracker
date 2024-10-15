package routes

import (
	"context"
	"net/http"
)

func MuxRouter(ctx *context.Context) (*http.ServeMux, error) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello\n")) })

	return mux, nil
}
