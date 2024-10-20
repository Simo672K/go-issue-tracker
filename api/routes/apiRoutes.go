package routes

import (
	"context"
	"net/http"
)

func MuxRouter(ctx context.Context, mux *http.ServeMux) error {
	mux.HandleFunc("/checkhealth", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("{\"message\":\"Hello world\"}")) })

	return nil
}
