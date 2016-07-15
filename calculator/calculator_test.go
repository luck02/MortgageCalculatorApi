package calculator

import (
	"testing"

	"github.com/luck02/MortgageCalculatorApi/models"
	. "github.com/smartystreets/goconvey/convey"
)

var sampleRequest = models.MortgagePaymentRequest{
	AskingPrice:        100000,
	DownPayment:        5000,
	PaymentSchedule:    "Monthly",
	AmortizationPeriod: 25,
}

func TestSpec(t *testing.T) {
	Convey("calculator", t, func() {
		Convey("mortgage insurance rates", func() {
			req := sampleRequest

			req.AskingPrice = 0

			payment, err := CalculatePayment(req)

			So(payment, ShouldEqual, 0)
			So(err, ShouldBeNil)

		})

	})
}
