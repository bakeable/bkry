package quotation

var ExampleJSON = `{
    customer: {
      address: "", // The customer's address 
      email: "", // The customer's email 
      name: "", // The customer's name 
      phone: "", // The customer's phone number 
    }, // The customer's information
    _kind: string, // The kind of the object
    lineItems: {
      attributes: {} satisfies Record<string, string>, // The quotation's attributes 
      finalPrice: 0.0, // The final price of the quotation 
      fixedCosts: 0.0, // The fixed costs of the quotation 
      marginAmount: 0.0, // The margin amount of the quotation 
      marginPercentage: 0.0, // The margin percentage of the quotation 
    product: {
      id: "", // The product's id 
      name: "", // The product's name 
      sku: "", // The product's name 
    }, // The products information
      totalCosts: 0.0, // The total costs of the quotation 
      units: 0, // The number of units in the quotation 
      variableCosts: 0.0, // The variable costs of the quotation 
    }[], // The quotation's added products
    name: string, // The quotation's name
    notes: string, // Any notes made on the quotation
    referralCode: string, // A supplied referral code
    status: QuotationStatus, // The quotation's status
    statusIndex: number, // The ordering index of the status
    totals: {
      finalPrice: 0.0, // The final price of the quotation 
      marginAmount: 0.0, // The margin amount of the quotation 
      marginPercentage: 0.0, // The margin percentage of the quotation 
      totalCosts: 0.0, // The total costs of the quotation 
    }, // The quotation totals
}`