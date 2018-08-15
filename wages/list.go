package wages

type List map[int]float32

func MostRecentYear(list List) (int, float32) {
	var mostRecentYear int = 0
	var mostRecentWage float32 = 0.0

	for year, wage := range list {
		if year > mostRecentYear {
			mostRecentYear = year
			mostRecentWage = wage 
		}
	}
	return mostRecentYear, mostRecentWage
}

func Add(year int, wage float32, list List) {
	var dup float32 = list[year]

	if dup == 0 {
		list[year] =  wage
	}
}

