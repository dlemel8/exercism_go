package pythagorean

import "sort"

const testVersion = 1

type Triplet [3]int

func Range(min, max int) []Triplet {
	size := max - min + 1
	squares := make([]int, size)
	for i := 0; i < size; i++ {
		squares[i] = (min + i) * (min + i)
	}

	res := []Triplet{}
	for i := 0; i < size; i++ {
		for j := 1; j <= i; j++ {
			for k := j; k >= 1; k-- {
				if squares[i] == squares[j]+squares[k] {
					res = append(res, Triplet{min + k, min + j, min + i})
				}
			}
		}
	}
	sort.Slice(res, func(i, j int) bool {
		if res[i][0] <= res[j][0] {
			return true
		}
		return false
	})
	return res
}

func Sum(p int) []Triplet {
	res := []Triplet{}
	options := Range(1, p)
	for _, o := range options {
		if o[0]+o[1]+o[2] == p {
			res = append(res, o)
		}
	}
	return res
}
