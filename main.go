package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	fmt.Printf(fmt.Sprintf("Starting server on: localhost:%s", port))

	routes()
	_ = http.ListenAndServe(port, nil)
}
