package calculator

import (
	"testing"

	"github.com/luck02/MortgageCalculatorApi/models"
	. "github.com/smartystreets/goconvey/convey"
)

/*
confirmed calculation at:
http://www.vertex42.com/ExcelArticles/amortization-calculation.html

This is pretty shallow testing but gets the point across.  I would probably
require a set of test cases from business or QA or someone of that ilk rather
than rolling a whole set of my own.
*/

var sampleRequest = models.MortgagePaymentRequest{
	AskingPrice:        100000,
	DownPayment:        5000,
	PaymentSchedule:    "Monthly",
	AmortizationPeriod: 25,
}

func TestSpec(t *testing.T) {
	Convey("calculator", t, func() {
		Convey("payment is correct", func() {
			req := sampleRequest

			payment, err := CalculatePayment(req)

			So(payment, ShouldEqual, 2451)
			So(err, ShouldBeNil)
		})
	})
}
