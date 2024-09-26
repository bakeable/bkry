package price_layer

var ExampleJSON = `{
    attributes: string[], // The attributes associated with the variable price
    description: string, // The description of the price layer
    external: boolean, // Whether the price layer is external
    includeInMarginCalculation: boolean, // Whether the price layer should be included in the margin calculation
    includeInPurchasePrice: boolean, // Whether the price layer should be included in the purchase price
    incrementalPricing: AttributeIncrementalPricing, // The type of incremental pricing of the price layer
    _kind: string, // The kind of the object
    name: string, // The name of the price layer
    options: string[], // The options associated with the price layer
    perUnits: number, // The unit size of the variable price
    priceType: PriceType, // The pricing type of the price layer
    roundingUnits: RoundingUnits, // How to round the unit sizes if they do not correspond to the unit size
    saveAsNew: boolean, // Whether to save the price-layer to be used later
    template: string, // The template of the price layer
    values: {
      attributes: {} satisfies Record<string, any>, // The attributes associated with the variable price 
    value: {
      decimalValue: 0.0, // The decimal value of the variable price 
      floatValue: 0.0, // The float value of the variable price 
      intValue: 0, // The integer value of the variable price 
      textValue: "", // The text value of the variable price 
    }, // The value of the variable price
    }[], // The values of the variable price
}`