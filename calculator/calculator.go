package calculator

import (
	"github.com/luck02/MortgageCalculatorApi/calculator/mortgageInsurance"
	"github.com/luck02/MortgageCalculatorApi/calculator/validator"
	"github.com/luck02/MortgageCalculatorApi/models"
)

const weekly = "Weekly"
const biweekly = "BiWeekly"
const monthly = "Monthly"

func CalculatePayment(model models.MortgagePaymentRequest) (int64, error) {
	ok, err := validator.Validate(model)
	if err != nil {
		return 0, err
	}

	loanInsurance := mortgageInsurance.CalculateMortgageInsurance(model.AskingPrice, model.DownPayment)
	loanPrincipal := model.AskingPrice - model.DownPayment
	numberOfPayments := calculateNumberPayments(model.PaymentSchedule, model.AmortizationPeriod)

	return 0, nil
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

	return 12
}
