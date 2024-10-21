package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go server!")

	c := polygon.New("API_KEY")

	params := models.GetTickerDetailsParams{
		Ticker: "AAPL",
	}

	res, err := c.GetTickerDetails(context.Background(), &params)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(res)
}

func main() {
	http.HandleFunc("/stocks/appl", handler)

	fmt.Println("Server is starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
