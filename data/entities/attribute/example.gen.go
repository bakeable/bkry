package attribute

var ExampleJSON = `{
    default: boolean, // Whether an attribute is selected by default
    defaultOptions: {
      id: "", // Unique identifier for the option 
      key: "", // The option's key 
      max: 0, // The maximum value for the option 
      min: 0, // The minimum value for the option 
      value: "", // The option's value 
    }[], // The default options for this attribute
    description: string, // A description of the attribute
    included: boolean, // Whether the attribute is included
    incrementalPricing: AttributeIncrementalPricing, // Whether graduated pricing applies
    key: string, // The attribute's key
    _kind: string, // The kind of the object
    maxOptionId: number, // The maximum option ID
    name: string, // The attribute's name
    nid: number, // The attribute's numerical identifier
    optional: boolean, // Whether the attribute is optional
    order: number, // A numerical representation of the order in which the attribute should be displayed
    prefix: string, // The attribute's prefix
    priceTypes: AttributePriceType[], // The associated price types
    surcharges: Record<string, any>, // The attribute's surcharges
    type: AttributeType, // 
}`