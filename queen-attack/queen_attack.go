package queenattack

import "errors"

const testVersion = 2

type position struct {
	file, rank int
}

func (p position) isOffBoard() bool {
	return p.file > 8 || p.file < 1 || p.rank > 8 || p.rank < 1
}

func (p position) getDiagonals() []position {
	res := make([]position, 0, 32)
	for i := 1; i <= 8; i++ {
		positions := []position{
			{p.file + i, p.rank + i},
			{p.file + i, p.rank - i},
			{p.file - i, p.rank + i},
			{p.file - i, p.rank - i},
		}
		for _, pos := range positions {
			if !pos.isOffBoard() {
				res = append(res, pos)
			}
		}
	}
	return res
}

func (p position) canAttack(another position) bool {
	if p.file == another.file || p.rank == another.rank {
		return true
	}
	for _, d := range p.getDiagonals() {
		if another == d {
			return true
		}
	}
	return false
}

func getPosition(s string) (position, error) {
	if len(s) != 2 {
		return position{}, errors.New("invalid")
	}

	file := int(s[0] - 'a' + 1)
	rank := int(s[1] - '0')
	res := position{file, rank}
	if res.isOffBoard() {
		return position{}, errors.New("off board")
	}
	return res, nil
}

func CanQueenAttack(w string, b string) (bool, error) {
	if w == b {
		return false, errors.New("same square")
	}
	wPosition, err := getPosition(w)
	if err != nil {
		return false, err
	}
	bPosition, err := getPosition(b)
	if err != nil {
		return false, err
	}
	return wPosition.canAttack(bPosition), nil
}
