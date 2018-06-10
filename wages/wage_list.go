package wages

type WageList map[int]float32

func WageListAdd(year int, wage float32, wages WageList) {
	var dup float32 = wages[year]

	if dup == 0 {
		wages[year] =  wage
	}
}

