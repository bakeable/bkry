package user_profile

var ExampleJSON = `{
    accessLevel: {
      admin: PermissionValues.none, // The user's administrative access level (an enum with possible values: [none read read_write])
      airlines: PermissionValues.none, // The user's access level to the airlines tool (an enum with possible values: [none read read_write])
      apiDocs: PermissionValues.none, // The user's access level to the API documentation tool (an enum with possible values: [none read read_write])
      packingSlips: PermissionValues.none, // The user's access level to the packing slip tool (an enum with possible values: [none read read_write])
      priceConfiguration: PermissionValues.none, // The user's access level to the price configuration tool (an enum with possible values: [none read read_write])
      printingOrders: PermissionValues.none, // The user's access level to the printing orders tool (an enum with possible values: [none read read_write])
      quotation: PermissionValues.none, // The user's access level to the quotation tool (an enum with possible values: [none read read_write])
      sampleTest: PermissionValues.none, // The user's access level to the sample test tool (an enum with possible values: [none read read_write])
      userManagement: PermissionValues.none, // The user's access level to the user management tool (an enum with possible values: [none read read_write])
    }, // Which tools the user has access to
    address: Record<string, any>, // A user's address
    avatarUrl: string, // The URL of a user's avatar
    email: string, // A user's email address
    firstName: string, // A user's first name
    infix: string, // A user's infix
    isAdmin: boolean, // Whether the user has administrative rights
    _kind: string, // The kind of the object
    language: Language, // A user's language
    lastName: string, // A user's last name
    prefix: string, // A user's prefix
}`