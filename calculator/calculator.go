package calculator

import (
	"github.com/luck02/MortgageCalculatorApi/calculator/mortgageInsurance"
	"github.com/luck02/MortgageCalculatorApi/models"
)

const weekly = "Weekly"
const biweekly = "BiWeekly"
const monthly = "Monthly"

func CalculatePayment(model models.MortgagePaymentRequest) (int64, error) {
	var loanInsurance = mortgageInsurance.CalculateMortgageInsurance(model.AskingPrice, model.DownPayment)
	var loanPrincipal = model.AskingPrice - model.DownPayment
	var numberOfPayments = calculateNumberPayments(model.PaymentSchedule, model.AmortizationPeriod)

	return 0, nil
}

func calculateNumberPayments(paymentSchedule string, amortizationPeriod int16) int {
	switch paymentSchedule {
	case weekly:
		return int(amortizationPeriod * int16(52))

	case monthly:
		return int(amortizatonPeriod * int16(12))
	}

	return 12
}
