package validator

import (
	"fmt"

	"github.com/luck02/MortgageCalculatorApi/models"
)

const minimumDownpaymentPercent = .05

func Validate(mortgagePaymentRequest models.MortgagePaymentRequest) (bool, error) {
	err := validateDownpaymentMinimum(
		mortgagePaymentRequest.AskingPrice,
		mortgagePaymentRequest.DownPayment,
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

func formatAmountForError(amount int64) string {
	return fmt.Sprintf("%.2f", float64(amount)/100)
}

func validateDownpaymentMinimum(askingPrice, downpayment int64) error {
	var minimumRequiredDownpayment = int64(float64(askingPrice) * minimumDownpaymentPercent)

	if downpayment < minimumRequiredDownpayment {
		return fmt.Errorf("validation error, minimum downpayment on $%s should $%s",
			formatAmountForError(askingPrice), formatAmountForError(minimumRequiredDownpayment))
	}

	return nil
}
