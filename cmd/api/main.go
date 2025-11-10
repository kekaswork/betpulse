package main

import (
	"fmt"
	"net/http"

	"github.com/kekaswork/betpulse/pkg/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	fmt.Println("Starting BetPulse API...")

	cfg := config.MustLoad()

	fmt.Printf("Config: %v", cfg)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
