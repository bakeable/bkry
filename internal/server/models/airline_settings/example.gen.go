package airline_settings

var ExampleJSON = `{
    clicheCosts: number, // The cliche costs for the airline settings
    _kind: string, // The kind of the object
    metersPerUnit: number, // The per units value for the airline settings
    rollLength: number, // The roll length for the airline settings
    rollWidthDividerMap: Record<number, number>, // The roll width to divider map
    rollWidthOptions: number[], // The roll width options
    sleeveCosts: number, // The sleeve costs for the airline settings
    switchCosts: number, // The switch costs for the airline settings
    unitOptionSets: number[][], // The unit options for the airline pricing
}`