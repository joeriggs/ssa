package wages

type List map[int]float32

func Add(year int, wage float32, list List) {
	var dup float32 = list[year]

	if dup == 0 {
		list[year] =  wage
	}
}

