package product
import (
	"github.com/bakeable/bkry/data/general/margin"
	"github.com/bakeable/bkry/data/entities/price_layer"
   "github.com/bakeable/bkry/data/general/auditinfo"
)

type Product struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The Products default attributes
	AttributeMap map[string]interface{} `json:"attributeMap"`
	// The Products default attributes
	Attributes []string `json:"attributes"`
	// The product's categories
	Categories []string `json:"categories"`
	// Whether the product has a configurable price
	ConfigurablePrice bool `json:"configurablePrice"`
	// The product's examination
	Examination Examination `json:"examination"`
	// Whether the product needs to be examined
	Examine bool `json:"examine"`
	// The product's family
	Family string `json:"family"`
	// The kind of the object
	Kind string `json:"_kind"`
	// The product's defalt margin
	Margin margin.Margin `json:"margin"`
	// The product's margins
	Margins Margins `json:"margins"`
	// The product's name
	Name string `json:"name"`
	// The products attribute options
	Options []string `json:"options"`
	// Whether all price configurations have been generated
	PriceConfigurationGenerated bool `json:"priceConfigurationsGenerated"`
	// Whether all price configurations have been generated
	PriceConfigurationGeneratedTimestamp int `json:"priceConfigurationGeneratedTimestamp"`
	// The price layers ID's
	PriceLayerIDs []string `json:"priceLayerIds"`
	// The price layers determining the product price
	PriceLayers []price_layer.PriceLayer `json:"priceLayers"`
	// External properties of the product
	Properties map[string]interface{} `json:"properties"`
	// The product's SKU
	Sku string `json:"sku"`
	// The product's status
	Status ProductStatus `json:"status"`
	// The ordering index of the status
	StatusIndex int `json:"statusIndex"`
}



type Examination struct {
	// Whether the examination is a priority
	Priority bool `json:"priority"`
	// The priority score of the examination
	PriorityScore int `json:"priorityScore"`
	// External properties of the examination
	Properties map[string]interface{} `json:"properties"`
	// The timeout of the examination
	Timeout int `json:"timeout"`
}


type Margins struct {
	// The direct margin
	Direct margin.Margin `json:"direct"`
	// The premium margin
	Premium margin.Margin `json:"premium"`
	// The standard margin
	Standard margin.Margin `json:"standard"`
}



type ProductStatus string




