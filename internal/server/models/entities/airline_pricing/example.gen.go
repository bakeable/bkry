package airline_pricing

var ExampleJSON = `{
    colors: number[], // The number of colors for the airline pricing
    description: string, // The description of the airline pricing
    dividers: number, // The number of dividers for the airline pricing
    kind: string, // The kind of the object
    margin: {
      amount: 0, // The amount of margin 
      percentage: 0, // The percentage of margin 
      percentageType: "profit-margin", // The type of percentage 
      staggered: false, // Whether the margin is staggered 
      type: "percentage", // The type of margin 
      values: [] as any[], // The values of the margin 
    }, // The pricing table's margin
    units: number[], // An array of units for the airline pricing
    values: {
    attributes: {
      colors: 0, // The colors for the airline pricing 
      units: 0, // The units for the airline pricing 
    }, // The attributes for the airline pricing
    value: {
      decimalValue: 0, // The decimal value for the airline pricing 
      floatValue: 0, // The float value for the airline pricing 
      intValue: 0, // The int value for the airline pricing 
      textValue: "", // The text value for the airline pricing 
    }, // The value for the airline pricing
    }[], // An array of values for the airline pricing
    width: number, // The width of the airline pricing
    widthType: string, // The type of width for the airline pricing
}`