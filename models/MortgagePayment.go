package models

type MortgagePaymentRequest struct {
	AskingPrice        int64  `json:"askingPrice"`
	DownPayment        int64  `json:"downPayment"`
	PaymentSchedule    string `json:"paymentSchedule"`
	AmortizationPeriod int16  `json:"amortizationPeriod"`
}

type MortgagePaymentResponse struct {
	Payment int64 `json:"payment"`
}
