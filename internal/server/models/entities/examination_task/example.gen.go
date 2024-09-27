package examination_task

var ExampleJSON = `{
    approved: boolean, // Whether the task is approved
    assignedTo: ExaminationTaskRole, // The person assigned to the task
    changeLog: {
      action: ExaminationTaskActionValues.created, // The action of the change log entry (an enum with possible values: [created updated observed solved approved archived])
      status: "", // The status of the change log entry 
      timestamp: "", // The timestamp of the change log entry 
    user: {
      email: "", // The email of the user 
      id: "", // The ID of the user 
    }, // The user of the change log entry
    }[], // The change log entries for the task
    kind: string, // The kind of the object
    observations: Record<string, any>, // The observations for the task
    productId: string, // The ID of the product
    productName: string, // The name of the product
    properties: {
    acceptanceRegion: {
      max: 0, // The maximum value of the acceptance region 
      min: 0, // The minimum value of the acceptance region 
      symmetric: false, // Whether the acceptance region is symmetric 
      type: AcceptanceRegionTypeValues.percentage, // The type of the acceptance region (an enum with possible values: [range percentage])
      value: 0, // The value of the acceptance region 
    }, // The acceptance region of the property
      description: "", // The description of the property 
      expectedValue: "", // The expected value of the property 
      inputType: "", // The input type of the property 
      instructions: "", // The instructions of the property 
      key: "", // The key of the property 
      lowerBound: 0, // The lower bound of the property 
      name: "", // The name of the property 
      productSpecific: false, // Whether the property is product specific 
      required: false, // Whether the property is required 
      type: "", // The type of the property 
      unitType: "", // The unit type of the property 
      upperBound: 0, // The upper bound of the property 
    }[], // The properties of the task
    sku: string, // The SKU of the product
    status: ExaminationTaskStatus, // The status of the task
    statusIndex: number, // The ordering of the status
}`