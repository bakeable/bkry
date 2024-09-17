package dimension

import utils "github.com/bakeable/bkry/tools"

// Decode a database map to an Dimension struct
func Decode(data interface{}) Dimension {
	if data == nil {
		return Dimension{}
	}
	if val, ok := data.(map[string]interface{}); ok {
		return Dimension{
			Amount: float64(utils.DecodeFloat64(val["amount"], 0)),
			Unit:    utils.DecodeString(val["unit"], ""),
		}
	}
	return Dimension{}
}