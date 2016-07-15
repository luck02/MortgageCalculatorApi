package mortgageInsurance

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testStruct struct {
	AskingPrice                   int64
	Downpayment                   int64
	ExpectedMortgageInsuranceCost int64
}

func TestSpec(t *testing.T) {
	Convey("mortgageInsurance", t, func() {
		Convey("mortgage insurance rates", func() {

			requests := []testStruct{
				{AskingPrice: 100000, Downpayment: 5000, ExpectedMortgageInsuranceCost: 2993},
				{AskingPrice: 100000, Downpayment: 10000, ExpectedMortgageInsuranceCost: 2160},
				{AskingPrice: 100000, Downpayment: 15000, ExpectedMortgageInsuranceCost: 1530},
				{AskingPrice: 100000, Downpayment: 20000, ExpectedMortgageInsuranceCost: 0},
			}

			for _, testValues := range requests {
				insuranceAmount := CalculateMortgageInsurance(testValues.AskingPrice, testValues.Downpayment)
				So(insuranceAmount, ShouldEqual, testValues.ExpectedMortgageInsuranceCost)
			}
		})

	})
}
