package main

import (
	"fmt"
	"goSumerGame/server/handlers"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	handlers.SetupHandlers(mux)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		_ = fmt.Errorf("error: %w, err\n", err)
		os.Exit(1)
	}
}
