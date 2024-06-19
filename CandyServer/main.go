package main

import (
	"GolangBootcamp/CandyServer/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/buy_candy", handler.BuyCandy)
	fmt.Println("Server started on :3333")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
