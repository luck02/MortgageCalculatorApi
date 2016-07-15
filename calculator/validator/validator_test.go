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
			req := sampleRequest

			req.AskingPrice = 0

			ok, err := Validate(req)

			So(ok, ShouldBeFalse)
			So(err.Error(), ShouldEqual,
				"validation error, asking price must be > 0")
		})

		Convey("payment schedule must be recognized value", func() {
			req := sampleRequest

			req.PaymentSchedule = "someBadValue"

			ok, err := Validate(req)

			So(ok, ShouldBeFalse)
			So(err.Error(), ShouldEqual,
				"validation error, PaymentSchedule must be one of Weekly, biweekly or monthly")
		})

		Convey("A downpayment", func() {
			Convey("Must be at least 5\\% of the first 500k", func() {
				req := sampleRequest
				req.AskingPrice = 10000
				req.DownPayment = 499

				ok, err := Validate(req)

				So(ok, ShouldBeFalse)
				So(err.Error(), ShouldEqual,
					"validation error, minimum downpayment on $100.00 is $5.00")
			})

			Convey("Must be at least 10\\% of amount above 500k", func() {
				req := sampleRequest
				req.AskingPrice = 75000000
				req.DownPayment = 500

				ok, err := Validate(req)

				So(ok, ShouldBeFalse)
				So(err.Error(), ShouldEqual, "validation error, minimum downpayment on $750000.00 is $50000.00")
			})
		})

		Convey("The amortizationPeriod", func() {
			Convey("Must be between 5 and 25 years", func() {
				req := sampleRequest
				req.AmortizationPeriod = 4

				ok, err := Validate(req)

				So(ok, ShouldBeFalse)
				So(err.Error(), ShouldEqual, "validation error, the amortization period must be between 5 and 25 years")
			})
		})

		Convey("We should test to ensure the downpayment doesn't exceed the askingprice", func() {

		})
	})
}
