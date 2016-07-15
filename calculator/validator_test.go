package validator

import (
	"testing"

	"github.com/luck02/MortgageCalculatorApi/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("A minimum downpayment", t, func() {
		Convey("Must be at least 5\\% of the first 500k", func() {
			mortgagePaymentRequest := models.MortgagePaymentRequest{
				AskingPrice:        10000,
				DownPayment:        500,
				PaymentSchedule:    "Weekly",
				AmortizationPeriod: 5,
			}

			ok, err := Validate(mortgagePaymentRequest)

			So(ok, ShouldBeFalse)
			So(err, ShouldEqual, "validation error, minimum downpayment on $100.00 should $20.00")

		})

	})
}
