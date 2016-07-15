package validator

import (
	"testing"

	"github.com/luck02/MortgageCalculatorApi/models"
	. "github.com/smartystreets/goconvey/convey"
)

var sampleRequest = models.MortgagePaymentRequest{
	AskingPrice:        1,
	DownPayment:        1,
	PaymentSchedule:    "Weekly",
	AmortizationPeriod: 5,
}

func TestSpec(t *testing.T) {
	Convey("request validation", t, func() {
		Convey("The basic request must be sane", func() {

		})

		Convey("A minimum downpayment", func() {
			Convey("Must be at least 5\\% of the first 500k", func() {
				req := sampleRequest
				req.AskingPrice = 10000
				req.DownPayment = 499

				ok, err := Validate(req)

				So(ok, ShouldBeFalse)
				So(err.Error(), ShouldEqual, "validation error, minimum downpayment on $100.00 should $5.00")

			})

		})
	})

}
