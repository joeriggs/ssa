package social_security

/*******************************************************************************
 * The ssa package an implementation of the algorithms that are used to
 * calculate or estimate the Social Security benefits for a retiree.
 *
 * Here's a simple description of how it all works:
 * - The Average Wage Index (AWI) Table contains the AWI for each year, going
 *   back to 1951.
 * - The AWI is used to calculate the "indexing factor" for each year that a
 *   person worked and earned wages.
 * - The "indexing factor" for year xxxx is calculated by dividing the AWI for
 *   the year the worker turned 60 by the AWI for year xxxx.
 * - There is a maximum "Nominal Earnings" each year.  You can see the maximum
 *   values in the maxWages table (max_earnings.go).  If the person earned more
 *   than the max in a given year, then the max is used.
 * - The "indexing factor" is multiplied by their nominal earnings for year xxxx
 *   to calculate their "indexed earnings" for that year.
 * - The highest 35 "indexed earnings" are added together to produce the
 *   "highest-35 total".
 * - The "highest-35 total" is divided by 420 to produce their "Average
 *   Indexed Monthly Earnings" (AIME).  FYI, there are 420 months in 35 years.
 *
 * Here is a simple example:
 * - The worker was born in 1954.  We need to calculate the "indexing factor"
 *   for this person for the year 1976.
 * - They were born in 1954, which means they turn 60 in 2014.
 * - AWI(2014) = 46,481.52
 * - AWI(1976) =  9,226.48
 * - Indexing Factor = 46,481.52 / 9,226.48 = 5.0378.
 * - Indexed Earnings for 1976 for this worker = Nominal_Earnings * 5.0378
 *
 ******************************************************************************/

import ( "github.com/joeriggs/ssa/wages" )

func min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	} else {
		return val2
	}
}

func Benefit(dob int, career_earnings wages.List) int {
	var highest_earnings WageList = WageListCreate()

	WageListHighestIndexedEarnings(dob, career_earnings, highest_earnings)

	var total float32 = WageListTotalIndexedEarnings(highest_earnings)

	var AIME int = int(total) / (35 * 12)

	//fmt.Printf("total %9.2f : AIME %d\n", total, AIME)

	var bend1, bend2 int
	bend1, bend2, _, _, _ = BendPoints(dob)
	//fmt.Printf("bend1 = %d : bend2 = %d\n", bend1, bend2)

	var tempAIME int = AIME

	var bend1Amt int = min(bend1, tempAIME)
	var bend1BenefitFloat float32 = float32(bend1Amt) * 0.90
	var bend1Benefit int = int(bend1BenefitFloat)
	tempAIME -= bend1Amt
	//fmt.Printf("bend1Amt = %d : bend1Benefit = %d\n", bend1Amt, bend1Benefit)

	var bend2Amt int = min((bend2 - bend1), tempAIME)
	var bend2BenefitFloat float32 = float32(bend2Amt) * 0.32
	var bend2Benefit int = int(bend2BenefitFloat)
	tempAIME -= bend2Amt
	//fmt.Printf("bend2Amt = %d : bend2Benefit = %d\n", bend2Amt, bend2Benefit)

	var moreAmt int = tempAIME
	var moreBenefitFloat float32 = float32(moreAmt) * 0.15
	var moreBenefit int = int(moreBenefitFloat)
	//fmt.Printf("moreAmt = %d : moreBenefit = %d\n", moreAmt, moreBenefit)

	var PIA int = bend1Benefit + bend2Benefit + moreBenefit
	//fmt.Printf("PIA = %d\n", PIA)

	return PIA
}

