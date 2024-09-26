package printing_order_supplier

var ExampleJSON = `{
    code: string, // The supplier code
    contactName: string, // Name of the supplier
    deliveryType: PrintingOrderSupplierDeliveryType, // The supplier delivery type
    email: string, // The main email of the supplier
    emails: string[], // The other email addresses of the supplier
    _kind: string, // The kind of the entity
    name: string, // The name of the supplier
    skus: string[], // The sku's associated with the supplier
}`