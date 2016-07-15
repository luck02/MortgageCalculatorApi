package calculator_test

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
			req := sampleRequest

			req.AskingPrice = 0

			ok, err := Validate(req)

			So(ok, ShouldBeFalse)
			So(err.Error(), ShouldEqual,
				"validation error, asking price must be > 0")
		})

	})
}
