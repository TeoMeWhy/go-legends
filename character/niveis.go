package character

import "sort"

var xpNivel = map[int]int{
	0:      1,
	300:    2,
	900:    3,
	2700:   4,
	6500:   5,
	14000:  6,
	23000:  7,
	34000:  8,
	48000:  9,
	64000:  10,
	85000:  11,
	100000: 12,
	120000: 13,
	140000: 14,
	165000: 15,
	195000: 16,
	225000: 17,
	265000: 18,
	305000: 19,
	355000: 20,
}

var nivelPonto = map[int]int{
	1:  0,
	2:  1,
	3:  1,
	4:  1,
	5:  2,
	6:  2,
	7:  2,
	8:  2,
	9:  3,
	10: 3,
	11: 3,
	12: 3,
	13: 4,
	14: 4,
	15: 4,
	16: 4,
	17: 5,
	18: 5,
	19: 5,
	20: 5,
}

func GetNivel(exp int) int {

	keys := make([]int, 0, len(xpNivel))
	for i := range xpNivel {
		keys = append(keys, i)
	}

	sort.Ints(keys)

	for i := range keys {
		if keys[i] > exp {
			return xpNivel[keys[i-1]]
		}
	}
	return 20
}
