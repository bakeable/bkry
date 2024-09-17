package delivery_entry

var ExampleJSON = `{
    customerId: string, // Customer ID associated with the delivery entry
    customerName: string, // Name of the customer associated with the delivery entry
    date: string, // Date of the delivery entry
    description: string, // Description of the delivery entry
    description_2: string, // Additional description for the delivery entry
    examination: {
      priority: false, // Whether the examination is a priority 
      priorityScore: 0, // The priority score of the examination 
      properties: {} satisfies Record<string, any>, // External properties of the examination 
      timeout: 0, // The timeout of the examination 
    }, // The product's examination
    extraDescription: string, // Extra description for the delivery entry
    guid: string, // Globally unique identifier for the delivery entry
    orderNumber: string, // Order number for the delivery entry
    orderType: string, // Type of order for the delivery entry
    serialNumber: number, // Serial number for the delivery entry
    sku: string, // SKU (Stock Keeping Unit) for the delivery entry
    unitType: string, // Type of unit for the delivery entry
    units: number, // Number of units in the delivery entry
    unitsPerPallet: number, // Number of units per pallet for the delivery entry
    warehouse: string, // Warehouse associated with the delivery entry
    _kind: string, // The kind of the object
}`