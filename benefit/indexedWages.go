package benefit

import ( "github.com/joeriggs/ssa/wages" )

type IndexedWagesEntry struct {
	year int
	wage float32
}
type IndexedWages [] IndexedWagesEntry

/* Calculate the "indexing factor" for the specified year.  We calculate the
 * indexing factor by dividing the "average wage index" for the year they turn
 * 60 by the "average wage index" for the requested year.
 */
func indexingFactor(dob, year int) float32 {
	var indexing_factor float32 = 0.0

	/* If the person is age 60+, then their indexing factor is always 1. */
	var age60Year int = dob + 60
	if year >= age60Year {
		indexing_factor = 1.0
	} else {
		/* If the person is less than age 60, calculate their indexing
		 * factor. */
		var awiAge60 float32 = AverageWageIndex(age60Year)
		var awi    float32 = AverageWageIndex(year)
		indexing_factor = awiAge60 / awi
	}

	return indexing_factor
}

/*******************************************************************************
 * Beginning of public API.
 ******************************************************************************/

func IndexedWagesCreate() IndexedWages {
	var list = make(IndexedWages, 35)
	return list
}

/* Add all of the top 35 wage years to produce the "total indexed earnings". */
func IndexedWagesTotalIndexedEarnings(highestEarnings IndexedWages) float32 {
	var total float32 = 0

	for _, wage := range highestEarnings {
		total += wage.wage
	}

	return total
}

func IndexedWagesHighestIndexedEarnings(dob int, earningsList wages.List, highestEarnings IndexedWages) {
	for wageYear, wageAmount := range earningsList {
		var maxEarnings int = MaxEarnings(wageYear)
		var allowedWage float32
		if wageAmount < float32(maxEarnings) {
			allowedWage = wageAmount
		} else {
			allowedWage = float32(maxEarnings)
		}

		var indexingFactor float32 = indexingFactor(dob, wageYear)
		var indexedEarnings float32 = allowedWage * indexingFactor
		//fmt.Printf("IndexedWagesHighestIndexedEarnings(): %d: %6.2f: %6.2f * %f = %6.2f\n",
		//             wageYear, wageAmount, allowedWage, indexingFactor, indexedEarnings)

		/* Do we need to round indexedEarnings up? */
		var ie float32 = float32(allowedWage) * indexingFactor
		var a float32 = float32(indexedEarnings)
		a += 0.5
		if(ie >= 1) {
			//fmt.Printf("Bumping indexedEarnings\n")
			indexedEarnings++
		}

		//fmt.Printf("IndexedWagesHighestIndexedEarnings(): %d: %6.2f: %6d: %6.2f * %f = %6.2f\n",
		//             wageYear, wageAmount, maxEarnings, allowedWage, indexingFactor, indexedEarnings)

		var y int
		for i := 0; i < len(highestEarnings); i++ {
			if highestEarnings[i].wage == 0.0 {
				y = i
				break
			} else if highestEarnings[i].wage < highestEarnings[y].wage {
				y = i
			}
		}

		if highestEarnings[y].year != 0 {
			//fmt.Printf("Lowest index is %d (%d %f).\n", y, highestEarnings[y].year, highestEarnings[y].wage)
		}
		if highestEarnings[y].wage > float32(indexedEarnings) {
			//fmt.Printf("Throwing away %d : %f.\n", wageYear, wageAmount)
			continue
		}

		highestEarnings[y].year = wageYear
		highestEarnings[y].wage = float32(indexedEarnings)
	}
}

