package firestore

import "strconv"

type Query struct {
	Field string `json:"field"`
	Operator QueryOperator `json:"operator"`
	Value interface{} `json:"value"`
}

type QueryOperator string

const (
	_Equal QueryOperator = "=="
	_NotEqual QueryOperator = "!="
	_GreaterThan QueryOperator = ">"
	_LessThan QueryOperator = "<"
	_GreaterThanOrEqual QueryOperator = ">="
	_LessThanOrEqual QueryOperator = "<="
	_In QueryOperator = "in"
	_ArrayContains QueryOperator = "array-contains"
	_ArrayContainsAny QueryOperator = "array-contains-any"
)

type Pagination struct {
	PageSize int `json:"page_size"`
	PageNumber int `json:"page_number"`
	OrderBy string `json:"order_by"`
	Ascending bool `json:"asc"`
	Count int `json:"count"`
	TotalPages int `json:"total_pages"`
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
		return q.Field + " " + string(q.Operator) + " \"" + q.Value.(string) + "\""
	case int64:
		return q.Field + " " + string(q.Operator) + " " + strconv.FormatInt(q.Value.(int64), 10)
	case int:
		return q.Field + " " + string(q.Operator) + " " + strconv.Itoa(q.Value.(int))
	case float64:
		return q.Field + " " + string(q.Operator) + " " + strconv.FormatFloat(q.Value.(float64), 'f', -1, 64)
	case bool:
		return q.Field + " " + string(q.Operator) + " " + strconv.FormatBool(q.Value.(bool))
	default:
		return q.Field + " " + string(q.Operator) + " INVALID_TYPE"
	}
}