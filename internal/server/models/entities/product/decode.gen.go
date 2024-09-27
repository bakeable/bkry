package product

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/internal/server/models/general/auditinfo"
)

// Decode converts a map to Product struct
func Decode(m interface{}) Product {
	if m, ok := m.(map[string]interface{}); ok {
		return Product{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			AttributeMap: utils.DecodeMap(m["attributeMap"], nil),
			Attributes: utils.DecodeStringArray(m["attributes"], []string{}),
			Categories: utils.DecodeStringArray(m["categories"], []string{}),
			ConfigurablePrice: utils.DecodeBool(m["configurablePrice"], false),
			Examination: DecodeExamination(m["examination"]),
			Examine: utils.DecodeBool(m["examine"], false),
			Family: utils.DecodeString(m["family"], ""),
			Kind: utils.DecodeString(m["_kind"], "Product"),
			Margin: (m["margin"]),
			Margins: DecodeMargins(m["margins"]),
			Name: utils.DecodeString(m["name"], ""),
			Options: utils.DecodeStringArray(m["options"], []string{}),
			PriceConfigurationGenerated: utils.DecodeBool(m["priceConfigurationsGenerated"], false),
			PriceConfigurationGeneratedTimestamp: utils.DecodeInt(m["priceConfigurationGeneratedTimestamp"], 0),
			PriceLayerIDs: utils.DecodeStringArray(m["priceLayerIds"], []string{}),
			PriceLayers: All(m["priceLayers"]),
			Properties: utils.DecodeMap(m["properties"], nil),
			Sku: utils.DecodeString(m["sku"], ""),
			Status: ProductStatus(utils.DecodeString(m["status"], "concept")),
			StatusIndex: utils.DecodeInt(m["statusIndex"], 0),
		}
	}

	return Product{}
}

// DecodeAll converts a slice of maps to a slice of Product struct
func DecodeAll(ms interface{}) []Product {
	var entities []Product

	if arr, ok := ms.([]map[string]interface{}); ok {
		for _, m := range arr {
			entities = append(entities, Decode(m))
		}
		return entities
	}

	if arr, ok := ms.([]interface{}); ok {
		for _, m := range arr {
			entities = append(entities, Decode(m))
		}
		return entities
	}

	return entities
}

// DecodeExamination converts a map to Examination struct
func DecodeExamination(m interface{}) Examination {
	if m == nil {
		return Examination{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return Examination{
			Priority: utils.DecodeBool(val["priority"], false),
			PriorityScore: utils.DecodeInt(val["priorityScore"], 0),
			Properties: utils.DecodeMap(val["properties"], nil),
			Timeout: utils.DecodeInt(val["timeout"], 0),
		}
	}
	return Examination{}
}


// DecodeMargins converts a map to Margins struct
func DecodeMargins(m interface{}) Margins {
	if m == nil {
		return Margins{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return Margins{
			Direct: (val["direct"]),
			Premium: (val["premium"]),
			Standard: (val["standard"]),
		}
	}
	return Margins{}
}


