package packing_slip

var ExampleJSON = `{
    addedBy: string, // Name of the person who added this packing slip
    address: {
      city: "", // City of the address 
      country: "", // Country of the address 
      description: "", // Description of the address 
      formatted: "", // Formatted address string 
      houseNumber: "", // House number of the address 
      postalCode: "", // Postal code of the address 
      street: "", // Street of the address 
    }, // Recipient's address
    administration: string, // The administration related to this packing slip
    awb: string, // Airway bill number
    carrierCode: CarrierCode, // Carrier code for the shipment
    combinedOrderNumbers: string[], // List of combined order numbers
    combinedPackingSlipIds: string[], // List of combined packing slip IDs
    companyName: string, // Name of the company
    customer: {
      email: "", // Email of the customer 
      name: "", // Name of the customer 
      phoneNumber: "", // Phone number of the customer 
    }, // Customer details
    date: string, // Date of the packing slip
    delivery: {
    afterTime: {
      hour: 9, // Hour of the delivery time 
      minute: 0, // Minute of the delivery time 
    }, // Time after which the delivery is expected
    beforeTime: {
      hour: 17, // Hour of the delivery time 
      minute: 0, // Minute of the delivery time 
    }, // Time before which the delivery is expected
      date: "", // Date of the delivery 
      notes: "", // Notes about the delivery 
    }, // Delivery details
    description: string, // Description of the packing slip
    incotermCode: string, // Incoterm code for the shipment
    _kind: string, // Type of the packing slip
    mainPackingSlipId: string, // ID of the main packing slip
    markedAsCompleted: boolean, // Flag indicating if the packing slip is marked as completed
    markedAsPrinted: boolean, // Flag indicating if the packing slip is marked as printed
    markedAsPushed: boolean, // Flag indicating if the packing slip is marked as pushed
    note: string, // Additional note for the packing slip
    orderCreatedBy: string, // Name of the person who created the order
    orderNumber: string, // Order number associated with the packing slip
    orderProcessing: string, // Order processing details
    originalTotalGrossWeight: number, // Original total gross weight of the shipment
    packages: {
      amount: 0, // Number of packages 
    dimensions: {
      height: [] as Dimension, // Height of the package 
      length: [] as Dimension, // Length of the package 
      width: [] as Dimension, // Width of the package 
    }, // Dimensions of the package
      transsmartCode: "", // Transsmart code for the package 
      type: "", // Type of the package 
      weight: [] as Dimension, // Weight of the package 
    }[], // List of packages in the shipment
    promisedDeliveryDate: string, // Promised delivery date
    reference: string, // Reference for the packing slip
    service: string, // Service details
    slipNumber: string, // Packing slip number
    status: PackingSlipStatus, // Status of the packing slip
    statusIndex: number, // The ordering index of the status
    supplierContactName: string, // Name of the supplier contact
    supplierId: string, // ID of the supplier
    supplierName: string, // Name of the supplier
    totalAmountExcludingVAT: number, // Total amount excluding VAT
    totalGrossWeight: number, // Total gross weight
    totalQuantity: number, // Total quantity
    trackingUrl: string, // Tracking URL
}`