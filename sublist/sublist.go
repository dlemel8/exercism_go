package sublist

import (
    "fmt"
    "strings"
)

type Relation string

func Sublist(first, second []int) Relation {
    firstString := sliceToString(first)
    secondString := sliceToString(second)
    firstContainsSecond := strings.Contains(firstString, secondString)
    secondContainsFirst := strings.Contains(secondString, firstString)
    if firstContainsSecond {
        if secondContainsFirst {
            return "equal"
        } else {
            return "superlist"
        }
    }
    if secondContainsFirst {
        return "sublist"
    } else {
        return "unequal"
    }
}

func sliceToString(str []int) string {
    return strings.Trim(fmt.Sprint(str), "[]")
}
