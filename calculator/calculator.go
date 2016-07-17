package calculator

import (
	"math"

	"github.com/luck02/MortgageCalculatorApi/calculator/mortgageInsurance"
	"github.com/luck02/MortgageCalculatorApi/calculator/validator"
	"github.com/luck02/MortgageCalculatorApi/models"
)

const weekly = "Weekly"
const biweekly = "BiWeekly"
const monthly = "Monthly"
const interestRate = 0.025 // Another distributed component to serve this
func CalculatePayment(model models.MortgagePaymentRequest) (int64, error) {
	err := validator.Validate(model)
	if err != nil {
		return 0, err
	}

	loanInsurance := float64(mortgageInsurance.CalculateMortgageInsurance(model.AskingPrice, model.DownPayment))
	loanPrincipal := float64(model.AskingPrice) - float64(model.DownPayment) + loanInsurance
	numberOfPayments := float64(calculateNumberPayments(model.PaymentSchedule, model.AmortizationPeriod))

	payment := loanPrincipal * (interestRate * math.Pow((1+interestRate), numberOfPayments)) /
		(math.Pow(1+interestRate, numberOfPayments) - 1)

	return int64(math.Floor(payment + .5)), nil
}

func calculateNumberPayments(paymentSchedule string, amortizationPeriod int16) int {
	switch paymentSchedule {
	case weekly:
		return int(amortizationPeriod * int16(52))
	case biweekly:
		return int(amortizationPeriod * int16(26))
	case monthly:
		return int(amortizationPeriod * int16(12))
	}
}
