package matrix

import (
    "fmt"
    "strconv"
    "strings"
)

type Matrix struct {
    rows [][]int
}

func New(s string) (*Matrix, error) {
    rows := strings.Split(s, "\n")
    data := make([][]int, len(rows))
    res := &Matrix{data}
    for i, r := range rows {
        row := strings.FieldsFunc(r, func(c rune) bool { return c == ' ' })
        data[i] = make([]int, len(row))
        if len(data[i]) != res.colsNum() {
            return nil, fmt.Errorf("row %d length (%d) is diffrent than first row length (%d)", i, len(data[i]), res.colsNum())
        }
        for j, val := range row {
            num, err := strconv.Atoi(val)
            if err != nil {
                return nil, err
            }
            if !res.Set(i, j, num) {
                return nil, fmt.Errorf("set value %d on position (%d,%d) failed", num, i, j)
            }
        }
    }
    return res, nil
}

func (m *Matrix) Rows() [][]int {
    res := make([][]int, m.rowsNum())
    for i := range m.rows {
        res[i] = make([]int, m.colsNum())
        copy(res[i], m.rows[i])
    }
    return res
}

func (m *Matrix) Cols() [][]int {
    res := make([][]int, m.colsNum())
    for i := range res {
        res[i] = make([]int, m.rowsNum())
    }
    for i, r := range m.rows {
        for j := range r {
            res[j][i] = r[j]
        }
    }
    return res
}

func (m *Matrix) Set(row, col, val int) bool {
    if row < 0 || row >= m.rowsNum() || col < 0 || col >= m.colsNum() {
        return false
    }
    m.rows[row][col] = val
    return true
}

func (m *Matrix) rowsNum() int {
    return len(m.rows)
}

func (m *Matrix) colsNum() int {
    return len(m.rows[0])
}
