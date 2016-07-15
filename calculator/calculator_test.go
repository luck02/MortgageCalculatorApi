package calculator

import (
	"testing"

	"github.com/luck02/MortgageCalculatorApi/models"
	. "github.com/smartystreets/goconvey/convey"
)

/*
confirmed calculation at:
http://www.vertex42.com/ExcelArticles/amortization-calculation.html

*/

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

			payment, err := CalculatePayment(req)

			So(payment, ShouldEqual, 2451)
			So(err, ShouldBeNil)

		})

	})
}
