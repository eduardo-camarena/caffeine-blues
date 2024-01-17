package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eduardo-camarena/caffeineBlues/internal/routes"
)

func main() {
	router := routes.NewRouter()

	addr := fmt.Sprintf(":%d", 8080)
	fmt.Printf("Server listening on http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
