package docref

import "github.com/bakeable/bkry/utils"

func Decode(m interface{}) DocRef {
	if m == nil {
		return ""
	}
	
	// Try decode as a map
	if vMap, ok := m.(map[string]interface{}); ok {
		str := utils.DefaultDecode(vMap["documentPath"], "").(string)
		return DocRef(str)
	}

	// Try decode as a string
	str := utils.DefaultDecode(m, "").(string)
	return DocRef(str)
}