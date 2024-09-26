package db

import "strconv"

type Query struct {
	Field string
	Operator string
	Value interface{}
}


func (q Query) ToString() string {
	// Infer on the type of the value
	// and convert it to a string
	// to be used in the query
	// e.g. "Field == Value"
	// e.g. "Field > Value"
	// e.g. "Field < Value"

	switch q.Value.(type) {
	case string:
		return q.Field + " " + q.Operator + " \"" + q.Value.(string) + "\""
	case int64:
		return q.Field + " " + q.Operator + " " + strconv.FormatInt(q.Value.(int64), 10)
	case int:
		return q.Field + " " + q.Operator + " " + strconv.Itoa(q.Value.(int))
	case float64:
		return q.Field + " " + q.Operator + " " + strconv.FormatFloat(q.Value.(float64), 'f', -1, 64)
	case bool:
		return q.Field + " " + q.Operator + " " + strconv.FormatBool(q.Value.(bool))
	default:
		return q.Field + " " + q.Operator + " INVALID_TYPE"
	}
}