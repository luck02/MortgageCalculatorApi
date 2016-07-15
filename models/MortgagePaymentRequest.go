package models

type MortgagePaymentRequest struct {
	AskingPrice        int64  `json:"askingPrice"`
	DownPayment        int64  `json:"downPayment"`
	PaymentSchedule    string `json:"schedule"`
	AmortizationPeriod int16  `json:"amorizationPeriod"`
}
