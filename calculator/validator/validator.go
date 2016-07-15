package validator

import (
	"errors"
	"fmt"

	"github.com/forestgiant/sliceutil"
	"github.com/luck02/MortgageCalculatorApi/models"
)

var paymentScheduleValues = []string{"Weekly", "BiWeekly", "Monthly"}

/*
Wanted to leave a note for reviewers.  Values below would probably originate in
a policy service of some sort.  This is not production worthy code, just
an example of problem decomposition.  In fact a quick spike to prove out
a point would be worthwhile, Mostly to ensure we'd caught all the unexpected
edge cases etc.
*/

const lowMortgageMinimumDownpaymentAmount = .05
const highMortgageAmount = 50000000
const highMortgageMinimumDownpaymentPercent = .10
const minAmortizationPeriod = 5
const maxAmortizationPeriod = 25

func formatAmountForError(amount int64) string {
	return fmt.Sprintf("%.2f", float64(amount)/100)
}

func Validate(mortgagePaymentRequest models.MortgagePaymentRequest) error {
	err := sanityCheck(mortgagePaymentRequest)
	if err != nil {
		return err
	}

	err = validateDownpaymentMinimum(
		mortgagePaymentRequest.AskingPrice,
		mortgagePaymentRequest.DownPayment,
	)

	if err != nil {
		return err
	}

	err = validateAmortizationPeriod(mortgagePaymentRequest.AmortizationPeriod)

	if err != nil {
		return err
	}

	return nil
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
	if amortizationPeriod < minAmortizationPeriod ||
		amortizationPeriod > maxAmortizationPeriod {
		return fmt.Errorf("validation error, the amortization period must be between %d and %d years", minAmortizationPeriod, maxAmortizationPeriod)
	}
	return nil
}
