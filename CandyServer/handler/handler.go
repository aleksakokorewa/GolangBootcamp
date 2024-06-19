package handler

import (
	"GolangBootcamp/CandyServer/config"
	"encoding/json"
	"fmt"
	"net/http"
)

func BuyCandy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var order config.Order

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	candyPrice, ok := config.CandyPrice[order.CandyType]
	totalPrice := candyPrice * order.CandyCount
	differenceMoney := totalPrice - order.Money

	if order.Money < totalPrice {
		w.WriteHeader(http.StatusPaymentRequired)
		err = json.NewEncoder(w).Encode(config.Response{ErrorMsg: fmt.Sprintf("You need %d more money!", differenceMoney)})
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	if !ok {
		http.Error(w, "Unknown candy type", http.StatusNotFound)
		return
	}
	if order.CandyCount <= 0 {
		http.Error(w, "Invalid candy count", http.StatusBadRequest)
		return
	} else {
		change := order.Money - totalPrice
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(config.Response{Change: change, Thanks: "Thank you!"})
	}
}
