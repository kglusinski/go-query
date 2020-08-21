package query

import (
	"fmt"
	"strings"
)

// Query is a representation of a MySQL query
type Query string
type Where struct {
	Field string
	Exp string
	Val interface{}
}

func (q *Query) Select(table, alias string, fields ...string) *Query {
	s := fmt.Sprintf("SELECT %s FROM %s %s", strings.Join(fields, ","), table, alias)

	res := Query(s)
	q = &res
	return q
}

func (q *Query) Where(w Where) *Query {
	s := fmt.Sprintf(q.ToString() +  " WHERE %s %s %v", w.Field, w.Exp, w.Val)

	res := Query(s)
	q = &res
	return q
}

func (q *Query) AndWhere(w Where) *Query {
	s := fmt.Sprintf(q.ToString() +  " AND %s %s %v", w.Field, w.Exp, w.Val)

	res := Query(s)
	q = &res
	return q
}

func (q *Query) OrWhere(w Where) *Query {
	s := fmt.Sprintf(q.ToString() +  " OR %s %s %v", w.Field, w.Exp, w.Val)

	res := Query(s)
	q = &res
	return q
}

func (q *Query) InnerJoin(table, alias, key1, key2 string) *Query {
	s := fmt.Sprintf(q.ToString() +  " INNER JOIN %s ON %s = %s", table, alias, key1, key2)

	res := Query(s)
	q = &res
	return q
}

func (q *Query) LeftJoin(table, alias, key1, key2 string) *Query {
	s := fmt.Sprintf(q.ToString() +  " LEFT JOIN %s %s ON %s = %s", table, alias, key1, key2)

	res := Query(s)
	q = &res
	return q
}

func (q *Query) Build() string {
	return q.ToString() + ";"
}

func (q *Query) ToString() string {
	return string(*q)
}
