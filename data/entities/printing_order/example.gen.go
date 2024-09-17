package printing_order

var ExampleJSON = `{
    approved: boolean, // Flag indicating if the printing order has been approved
    approvedTimestamp: number, // The timestamp of the approved
    artworkApproved: boolean, // Flag indicating whether the artwork is approved
    artworkIsUrl: boolean, // Flag indicating whether the artwork is a URL or a file
    autoCorrespondence: boolean, // Flag indicating whether to automatically send correspondence
    companyName: string, // The name of the company associated with the printing order
    correspondence: {
      id: "", // ID of the correspondence 
      key: PrintingOrderCorrespondenceTemplateKeyValues.CUSTOMER_APPROVE_SAMPLE, // Key of the correspondence (an enum with possible values: [CUSTOMER_APPROVE_SAMPLE CUSTOMER_PRINTING_ORDER_RECEIVED CUSTOMER_TRACK_AND_TRACE SUPPLIER_PRINTING_ORDER_RECEIVED SUPPLIER_SAMPLE_ACCEPTED SUPPLIER_SAMPLE_REJECTED])
      sent: false, // Flag indicating if the correspondence was sent 
      sentAt: 0, // Timestamp when the correspondence was sent 
      statusLogIndex: 0, // The index of the status log change that triggered the correspondence 
      timestamp: 0, // Timestamp of the correspondence 
    }[], // The correspondence tracking
    customerAddress: {
      city: "", // City of the address 
      country: "NL", // Country of the address 
      line: "", // Address line 
      postalCode: "", // Postal code of the address 
    }, // The address of the customer
    customerEmail: string, // Email of the customer
    customerName: string, // Name of the customer
    description: string, // Description of the printing order
    emailCode: string, // The 10 digit e-mail code
    files: {
      artwork: , // The artwork file 
      media: , // The media file 
      orderConfirmation: , // The order confirmation file 
    }, // The files associated with the printing order
    isAirtableOrder: boolean, // Whether the printing order is from Airtable
    itemIndex: number, // The item index of the printing order
    _kind: string, // Type of the printing order
    notes: string, // Notes of the printing order
    numberOfUnits: number, // Number of units ordered per quantity
    orderNumber: string, // The order number of the printing order
    orderStatus: string, // The status of the order in Magento
    productOptions: {
      id: "", // ID of the product option 
      key: "", // Key of the product option 
      value: "", // Value of the product option 
    }[], // The options of the product
    quantity: number, // Quantity of the printing order
    quantityOrdered: number, // Quantity ordered of the printing order
    rejected: boolean, // Whether the sample was rejected
    rejectionReason: string, // The reason for rejection
    sku: string, // SKU of the printing order
    status: PrintingOrderStatus, // Status of the printing order
    statusIndex: number, // The ordering index of the status
    statusLog: {
      prevStatus: PrintingOrderStatusValues.new, // Previous status of the printing order (an enum with possible values: [])
      status: PrintingOrderStatusValues.new, // Status of the printing order (an enum with possible values: [])
      timestamp: 0, // Timestamp of the status change 
    }[], // The log of previous statuses
    supplierContactName: string, // Name of the supplier contact associated with the printing order
    supplierId: string, // ID of the supplier associated with the printing order
    supplierName: string, // Name of the supplier associated with the printing order
    supplierReference: string, // The supplier reference
    trackingUrl: string, // Tracking URL of the printing order
    trackingUrlApproved: boolean, // Whether the tracking URL is approved
}`