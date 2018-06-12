package social_security

import ( "fmt"
         "github.com/joeriggs/ssa/wages"
)

type WageListEntry struct {
	year int
	wage float32
}
type WageList [] WageListEntry

/* This function calculates the "indexing factor" for the specified year. */
func indexingFactor(dob, year int) float32 {
	var indexing_factor float32 = 0.0

	/* If the person is age 60+, then their indexing factor is always 1. */
	var age_60 int = dob + 60
	if year >= age_60 {
		indexing_factor = 1.0
	} else {
		/* If the person is less than age 60, calculate their indexing
		 * factor. */
		var awi_60 float32 = AverageWageIndex(age_60)
		var awi    float32 = AverageWageIndex(year)
		indexing_factor = awi_60 / awi
		//fmt.Printf("indexingFactor(): dob %d : year %d : awi_60 %f : awi %f : indexing_factor %f\n", dob, year, awi_60, awi, indexing_factor)
	}

	return indexing_factor
}

/* Add all of the top 35 wage years to produce the "total indexed earnings". */
func WageListTotalIndexedEarnings(highestEarnings WageList) float32 {
	var total float32 = 0

	for _, wage := range highestEarnings {
		total += wage.wage
	}

	return total
}

func WageListCreate() WageList {
	var list = make(WageList, 35)
	return list
}

func WageListHighestIndexedEarnings(dob int, earnings_list wages.List, highest_earnings WageList) {
	for wageYear, wageWage := range earnings_list {
		var maxEarnings int = MaxEarnings(wageYear)
		var allowedWage float32
		if wageWage < float32(maxEarnings) {
			allowedWage = wageWage
		} else {
			allowedWage = float32(maxEarnings)
		}

		var indexingFactor float32 = indexingFactor(dob, wageYear)
		var indexedEarnings float32 = allowedWage * indexingFactor
		//fmt.Printf("WageListHighestIndexedEarnings(): %d: %6.2f: %6.2f * %f = %6.2f\n",
		//             wageYear, wageWage, allowedWage, indexingFactor, indexedEarnings)

		/* Do we need to round indexedEarnings up? */
		var ie float32 = float32(allowedWage) * indexingFactor
		var a float32 = float32(indexedEarnings)
		a += 0.5
		if(ie >= 1) {
			//fmt.Printf("Bumping indexedEarnings\n")
			indexedEarnings++
		}

		//fmt.Printf("WageListHighestIndexedEarnings(): %d: %6.2f: %6d: %6.2f * %f = %6.2f\n",
		//             wageYear, wageWage, maxEarnings, allowedWage, indexingFactor, indexedEarnings)

		var y int
		for i := 0; i < len(highest_earnings); i++ {
			if highest_earnings[i].wage == 0.0 {
				y = i
				break
			} else if highest_earnings[i].wage < highest_earnings[y].wage {
				y = i
			}
		}

		if highest_earnings[y].year != 0 {
			//fmt.Printf("Lowest index is %d (%d %f).\n", y, highest_earnings[y].year, highest_earnings[y].wage)
		}
		if highest_earnings[y].wage > float32(indexedEarnings) {
			//fmt.Printf("Throwing away %d : %f.\n", wageYear, wageWage)
			continue
		}

		//fmt.Printf("Replacing %d (%f) with %d (%f).\n", highest_earnings[y].year, highest_earnings[y].wage, wageYear, indexedEarnings)
		highest_earnings[y].year = wageYear
		highest_earnings[y].wage = float32(indexedEarnings)
	}

	/* Take a look at the result. */
	fmt.Printf("List: Begin\n");
	for i := 0; i < len(highest_earnings); i++ {
		if highest_earnings[i].year > 0 {
			fmt.Printf("List: %d: %11.2f\n", highest_earnings[i].year, highest_earnings[i].wage)
		}
	}
	fmt.Printf("List: End\n");
}

