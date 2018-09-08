package benefit

/* These are the "bend points" for the Primary Insurance Amount (PIA)
 * calculation.  They were lifted from:
 *   https://www.ssa.gov/oact/COLA/bendpoints.html
 */
type bendPoint struct {
  pia_bend1 int
  pia_bend2 int
  family1 int
  family2 int
  family3 int
}
var bendPoints = map[int]bendPoint{
  1979: { 180, 1085,  230,  332,  433 },
  1980: { 194, 1171,  248,  358,  467 },
  1981: { 211, 1274,  270,  390,  508 },
  1982: { 230, 1388,  294,  425,  554 },
  1983: { 254, 1528,  324,  468,  610 },
  1984: { 267, 1612,  342,  493,  643 },
  1985: { 280, 1691,  358,  517,  675 },
  1986: { 297, 1790,  379,  548,  714 },
  1987: { 310, 1866,  396,  571,  745 },
  1988: { 319, 1922,  407,  588,  767 },
  1989: { 339, 2044,  433,  626,  816 },
  1990: { 356, 2145,  455,  656,  856 },
  1991: { 370, 2230,  473,  682,  890 },
  1992: { 387, 2333,  495,  714,  931 },
  1993: { 401, 2420,  513,  740,  966 },
  1994: { 422, 2545,  539,  779, 1016 },
  1995: { 426, 2567,  544,  785, 1024 },
  1996: { 437, 2635,  559,  806, 1052 },
  1997: { 455, 2741,  581,  839, 1094 },
  1998: { 477, 2875,  609,  880, 1147 },
  1999: { 505, 3043,  645,  931, 1214 },
  2000: { 531, 3202,  679,  980, 1278 },
  2001: { 561, 3381,  717, 1034, 1349 },
  2002: { 592, 3567,  756, 1092, 1424 },
  2003: { 606, 3653,  774, 1118, 1458 },
  2004: { 612, 3689,  782, 1129, 1472 },
  2005: { 627, 3779,  801, 1156, 1508 },
  2006: { 656, 3955,  838, 1210, 1578 },
  2007: { 680, 4100,  869, 1255, 1636 },
  2008: { 711, 4288,  909, 1312, 1711 },
  2009: { 744, 4483,  950, 1372, 1789 },
  2010: { 761, 4586,  972, 1403, 1830 },
  2011: { 749, 4517,  957, 1382, 1803 },
  2012: { 767, 4624,  980, 1415, 1845 },
  2013: { 791, 4768, 1011, 1459, 1903 },
  2014: { 816, 4917, 1042, 1505, 1962 },
  2015: { 826, 4980, 1056, 1524, 1987 },
  2016: { 856, 5157, 1093, 1578, 2058 },
  2017: { 885, 5336, 1131, 1633, 2130 },
  2018: { 895, 5397, 1144, 1651, 2154 },
}

var bpMostRecentYear int = 0
func mostRecentBendPointsYear() int {
	if bpMostRecentYear == 0 {
		for key, _ := range maxWages {
			if key > bpMostRecentYear {
				bpMostRecentYear = key
			}
		}
	}
	return bpMostRecentYear
}

func BendPoints(dob int) (int, int, int, int, int) {
	var age_62 int = dob + 62
	yr := mostRecentBendPointsYear()
	if age_62 < yr {
		yr = age_62
	}

	var point bendPoint = bendPoints[yr]

	return point.pia_bend1, point.pia_bend2, point.family1, point.family2, point.family3
}

