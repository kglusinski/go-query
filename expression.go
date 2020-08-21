package query

import (
	"fmt"
	"strings"
)

const (
	ExpIn Exp = "IN"
	ExpAnd Exp = "AND"
	ExpOr Exp = "OR"
	ExpEq Exp = "="
)

type Exp string

func In(list ...string) Exp {
	return Exp(fmt.Sprintf("IN (%v)", strings.Join(list, ",")))
}

func Eq() Exp {
	return ExpEq
}