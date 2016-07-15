package validator

import (
	"errors"

	"github.com/luck02/MortgageCalculatorApi/models"
)

// Validate : Takes a request and confirms it matches validation rules
// as given in ../requirements.txt
func Validate(mortgagePaymentRequest models.MortgagePaymentRequest) (bool, error) {
	return false, errors.New("derp")
}
