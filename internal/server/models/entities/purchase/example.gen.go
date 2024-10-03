package purchase

var ExampleJSON = `{
    billing_address: {
      city: "", // The city of the billing address 
      country_code: "", // The country code of the billing address 
      line1: "", // The first line of the billing address 
      line2: "", // The second line of the billing address 
      postal_code: "", // The postal code of the billing address 
      state: "", // The state of the billing address 
      street: "", // The street of the billing address 
    }, // The billing address for the purchase
    package_id: string, // The Id of the package being purchased
    payment_method: string, // The method of payment
    price: {
      currency: "", // The currency of the purchase 
      discount_amount: 0, // The discount amount 
      final_amount: 0, // The final amount after tax and discounts 
      initial_amount: 0, // The initial amount before tax and discounts 
    tax: [], // The tax details of the purchase
      tax_amount: 0, // The total tax amount 
    }, // The price details of the purchase
    status: Status, // The status of the purchase
    status_log: {
      new_status: "", // The new status of the purchase 
      prev_status: "", // The previous status of the purchase 
      timestamp: 0, // The timestamp when the status was changed 
    }[], // The log of status changes for the purchase
    transaction_id: string, // The Id of the transaction
    user_id: string, // The Id of the user making the purchase
}`