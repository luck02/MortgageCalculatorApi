package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/luck02/MortgageCalculatorApi/calculator"
	"github.com/luck02/MortgageCalculatorApi/models"
)

const port = 8081

/*
Missing pieces:
* Logging, I'm just going to print to console for now.
* if it blows up there's no infra around it to report / restart / handle

*/

func main() {
	fmt.Println(fmt.Sprintf("Server starting on port:%d", port))
	http.HandleFunc("/", paymentEstimatorHandler)
	fmt.Println(http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil))
}

func paymentEstimatorHandler(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	model := models.MortgagePaymentRequest{}
	
	err := json.Unmarshal(buf.Bytes(), &model)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	payment, err := calculator.CalculatePayment(model)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}

	mortgagePaymentResponse := models.MortgagePaymentResponse{
		Payment: payment,
	}
	
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mortgagePaymentResponse)
}
