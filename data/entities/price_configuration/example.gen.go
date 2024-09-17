package price_configuration

var ExampleJSON = `{
    additional_costs: number, // The additional costs
    attributes: Record<string, string>, // The attribute configuration for this price configuration
    currency: string, // The currency used in the response
    final_price: number, // The final price
    fixed_costs: number, // The fixed costs
    _kind: string, // The kind of the object
    margin_amount: number, // The margin amount
    margin_percentage: number, // The margin percentage
    product_id: string, // The priceConfiguration's corresponding product ID
    sku: string, // The priceConfiguration's corresponding product SKU
    total_costs: number, // The total costs
    unit_costs: number, // The costs associated with a single unit
    unit_margin: number, // The costs associated with a single unit
    units: number, // The number of units
    variable_costs: number, // The costs associated with a single unit
}`