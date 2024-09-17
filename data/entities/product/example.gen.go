package product

var ExampleJSON = `{
    attributeMap: Record<string, any>, // The Products default attributes
    attributes: string[], // The Products default attributes
    categories: string[], // The product's categories
    configurablePrice: boolean, // Whether the product has a configurable price
    examination: {
      priority: false, // Whether the examination is a priority 
      priorityScore: 0, // The priority score of the examination 
      properties: {} satisfies Record<string, any>, // External properties of the examination 
      timeout: 0, // The timeout of the examination 
    }, // The product's examination
    examine: boolean, // Whether the product needs to be examined
    family: string, // The product's family
    _kind: string, // The kind of the object
    margins: {
      direct: , // The direct margin 
      premium: , // The premium margin 
      standard: , // The standard margin 
    }, // The product's margins
    name: string, // The product's name
    options: string[], // The products attribute options
    priceConfigurationsGenerated: boolean, // Whether all price configurations have been generated
    priceConfigurationGeneratedTimestamp: number, // Whether all price configurations have been generated
    priceLayerIds: string[], // The price layers ID's
    properties: Record<string, any>, // External properties of the product
    sku: string, // The product's SKU
    status: ProductStatus, // The product's status
    statusIndex: number, // The ordering index of the status
}`