package main

import (
	"fmt"
	"net/http"
)

// go run server.go
func main() {
	err := http.ListenAndServe(":9090", http.FileServer(http.Dir("../public")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
