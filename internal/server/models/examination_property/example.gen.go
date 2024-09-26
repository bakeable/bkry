package examination_property

var ExampleJSON = `{
    acceptanceRegion: {
      max: 0.0, // The maximum value of the acceptance region 
      min: 0.0, // The minimum value of the acceptance region 
      symmetric: false, // Whether the acceptance region is symmetric 
      type: "", // The type of the acceptance region 
      value: 0.0, // The value of the acceptance region 
    }, // The acceptance region of the property
    description: string, // A description of the property to be examined
    examine: boolean, // Whether the property needs to be examined
    info: Record<string, any>, // All additional information about the property
    inputType: ExaminationInputType, // The input type of the property
    instructions: string, // A set of instructions for the examiner of the property
    key: string, // The key of the property to be examined
    _kind: string, // The kind of the object
    name: string, // The name of the property to be examined
    productId: string, // The ID of the product to which the property belongs
    productSpecific: boolean, // Whether the property is product specific
    required: boolean, // Whether it is required to observe the property
    testable: boolean, // Whether the property is testable
    type: string, // The type of the property to be examined
    unitType: string, // The unit type of the property to be examined
}`