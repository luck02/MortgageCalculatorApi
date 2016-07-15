package mortgageInsurance

import "math"

/*
Reviwer notes:  This would definitely be some sort of independant service.

I cant imagine this simplistic of an implementation that didn't change
constantly.

Certainly at the least the lookup table would be dynamically queried, possibly
cached in case of massive access but would need to be queried in real time
before approving a mortgage.  In fact I don't think this should be part of this
calculator at all, it should be a service that lives elsewhere. Realistically
we would want to model the interface and then mock that out and not test the
implementation here at all.

One additional point is that
*/

type rate struct {
	floor float64
	cost  float64
}

var rateTable = []rate{
	{.05, 0.0315},
	{.10, 0.024},
	{.15, 0.018},
	{.20, 0},
}

func CalculateMortgageInsurance(askingPrice, downpayment int64) int64 {
	var rate = getRate(askingPrice, downpayment)
	var insuranceCost = rate * float64(askingPrice-downpayment)
	return int64(math.Floor(insuranceCost + .5))
}

func getRate(askingPrice, downpayment int64) float64 {
	var downpaymentRatio = float64(downpayment) / float64(askingPrice)
	var assessedRate rate

	for _, rate := range rateTable {
		if downpaymentRatio >= rate.floor {
			assessedRate = rate
		}
	}

	return assessedRate.cost
}
