package margin

import "github.com/bakeable/bkry/utils"

// Decode a database map to an Dimension struct
func Decode(data interface{}) Margin {
	if data == nil {
		return Margin{}
	}
	if val, ok := data.(map[string]interface{}); ok {
		return Margin{
			Amount: utils.DecodeFloat64(val["amount"], 0),
			Percentage: utils.DecodeFloat64(val["percentage"], 0),
			PercentageType: utils.DecodeString(val["percentageType"], ""),
			Staggered: utils.DecodeBool(val["staggered"], false),
			Type: utils.DecodeString(val["type"], ""),
			Values: decodeMarginValues(val["values"]),
		}
	}
	return Margin{}
}

func decodeMarginValues(data interface{}) []MarginValues {
	if data == nil {
		return []MarginValues{}
	}
	if val, ok := data.([]interface{}); ok {
		values := make([]MarginValues, 0)
		for _, v := range val {
			values = append(values, decodeMarginValue(v))
		}
		return values
	}
	return []MarginValues{}
}

func decodeMarginValue(data interface{}) MarginValues {
	if data == nil {
		return MarginValues{}
	}
	if val, ok := data.(map[string]interface{}); ok {
		return MarginValues{
			Attributes: utils.DecodeMap(val["attributes"], map[string]interface{}{}),
			Value: MarginValue{
				TextValue: utils.DecodeString(val["textValue"], ""),
				IntValue: utils.DecodeInt(val["intValue"], 0),
				FloatValue: utils.DecodeFloat64(val["floatValue"], 0),
				DecimalValue: utils.DecodeString(val["decimalValue"], ""),
			},
		}
	}
	return MarginValues{}
}