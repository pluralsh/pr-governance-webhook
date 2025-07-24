package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/pluralsh/pr-governance-webhook/pkg/handler"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP server address (e.g., ':8080' or '127.0.0.1:9090')")
	flag.Parse()

	http.HandleFunc("/v1/open", handler.OpenHandler)
	http.HandleFunc("/v1/confirm", handler.ConfirmHandler)
	http.HandleFunc("/v1/close", handler.CloseHandler)

	log.Printf("Server starting on %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
