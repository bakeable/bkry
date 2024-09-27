package airline_order

var ExampleJSON = `{
    afasOrderNumber: string, // The AFAS order number of the airline order
    colors: number, // The number of colors of the airline order
    customerId: string, // The ID of the customer associated with the airline order
    customerName: string, // The name of the customer associated with the airline order
    date: string, // The date of the airline order
    deliveryAddress: string, // The delivery address of the order
    description: string, // The description of the airline order
    dividers: number, // The number of dividers of the airline order
    guid: string, // The unique identifier of the airline order
    supplierId: string, // The ID of the supplier associated with the airline order
    isGrouped: boolean, // Whether the airline order has already been grouped
    _kind: string, // The kind of the object
    margin: {
      amount: 0, // The amount of margin 
      percentage: 0, // The percentage of margin 
      percentageType: "profit-margin", // The type of percentage 
      staggered: false, // Whether the margin is staggered 
      type: "percentage", // The type of margin 
      values: [] as any[], // The values of the margin 
    }, // The order margin
    number: string, // The number of the airline order
    orderNumber: string, // The order number of the airline order
    originalTotalCosts: number, // The original total costs (when separately ordered)
    originalUnitCosts: number, // The original unit costs (when separately ordered)
    serialNumber: number, // The serial number of the airline order
    sku: string, // The SKU (Stock Keeping Unit) of the airline order
    status: string, // The status of the airline order
    totalCosts: number, // The total costs of the airline order
    unitCosts: number, // The unit costs of the airline order
    unitType: string, // The unit type of the airline order
    units: number, // The number of units in the airline order
    warehouse: string, // The warehouse of the airline order
    width: number, // The width of the ordered airline foil
    widthUnitType: number, // The width unit type of the ordered airline foil
}`