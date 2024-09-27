package airline_order_grouping

var ExampleJSON = `{
    afasOrderNumbers: string[], // The returned order number from AFAS
    finalPrice: number, // The final price of this order grouping
    firstOrderDate: string, // The date the first order was placed
    _kind: string, // The kind of the object
    marginAmount: number, // The total amount of margin
    marginPercentage: number, // The total amount of margin
    name: string, // A reference name for the airline order grouping
    orderDate: string, // The date the order grouping needs to be ordered
    orderIds: string[], // The list of airline orders within the grouping
    originalTotalCosts: number, // The original costs when using separated orders
    profitAmount: number, // The total profit amount of this order grouping
    profitPercentage: number, // The total profit percentage of this order grouping
    savedCosts: number, // The costs saved when using this order grouping
    sentToAfas: boolean, // Whether the order grouping has been sent to AFAS
    sentToAfasTimestamp: number, // The timestamp of when the order grouping was sent to AFAS
    totalCosts: number, // The total costs of this order grouping
    units: number, // The total number of units of the airline order grouping
}`