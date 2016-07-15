package validator

import (
	"errors"
	"fmt"

	"github.com/forestgiant/sliceutil"
	"github.com/luck02/MortgageCalculatorApi/models"
)

var paymentScheduleValues = []string{"Weekly", "BiWeekly", "Monthly"}

const lowMortgageMinimumDownpaymentAmount = .05
const highMortgageAmount = 50000000
const highMortgageMinimumDownpaymentPercent = .10

func formatAmountForError(amount int64) string {
	return fmt.Sprintf("%.2f", float64(amount)/100)
}

func Validate(mortgagePaymentRequest models.MortgagePaymentRequest) (bool, error) {
	err := sanityCheck(mortgagePaymentRequest)
	if err != nil {
		return false, err
	}

	err = validateDownpaymentMinimum(
		mortgagePaymentRequest.AskingPrice,
		mortgagePaymentRequest.DownPayment,
	)

	if err != nil {
		return false, err
	}

	err = validateAmortizationPeriod(mortgagePaymentRequest.AmortizationPeriod)

	if err != nil {
		return false, err
	}

	return true, nil
}

func sanityCheck(mortgagePaymentRequest models.MortgagePaymentRequest) error {
	if mortgagePaymentRequest.AskingPrice <= 0 {
		return errors.New("validation error, asking price must be > 0")
	}

	if !sliceutil.Contains(paymentScheduleValues, mortgagePaymentRequest.PaymentSchedule) {
		return errors.New("validation error, PaymentSchedule must be one of Weekly, biweekly or monthly")
	}

	return nil
}

func validateDownpaymentMinimum(askingPrice, downpayment int64) error {
	var minimumRequiredDownpayment = calculateMinimumDownpaymentAmount(askingPrice, downpayment)

	if downpayment < minimumRequiredDownpayment {
		return fmt.Errorf("validation error, minimum downpayment on $%s is $%s",
			formatAmountForError(askingPrice), formatAmountForError(minimumRequiredDownpayment))
	}

	return nil
}

func calculateMinimumDownpaymentAmount(askingPrice, downpayment int64) int64 {
	var minimumDownpayment = int64(0)

	if askingPrice > highMortgageAmount {
		minimumDownpayment += int64(float64(highMortgageAmount) * lowMortgageMinimumDownpaymentAmount)
		minimumDownpayment += int64(float64(askingPrice-highMortgageAmount) * highMortgageMinimumDownpaymentPercent)
	} else {
		minimumDownpayment += int64(float64(askingPrice) * lowMortgageMinimumDownpaymentAmount)
	}

	return minimumDownpayment
}

func validateAmortizationPeriod(amortizationPeriod int16) error {
	return nil
}
