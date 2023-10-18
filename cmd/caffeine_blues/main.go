package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eduardo-camarena/caffeineBlues/internal/routes"
)

func main() {
	router := routes.NewRouter()

	port := 8000
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
