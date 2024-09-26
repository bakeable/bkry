package examination_settings

var ExampleJSON = `{
    exclusionTime: number, // Standard timeframe from which to select tasks (in milliseconds)
    exclusionTimeFrame: {
      days: 0, // Number of days in the exclusion time frame 
      hours: 0, // Number of hours in the exclusion time frame 
    }, // Exclusion time frame
    kind: string, // The kind of the object
    maxOpenTasks: number, // Maximum number of open tasks
    maxProducts: number, // Maximum number of delivery entries to examine
    maxProperties: number, // Maximum number of properties to examine
    minProducts: number, // Minimum number of delivery entries to examine
    minProperties: number, // Minimum number of properties to examine
    properties: string[], // Properties to examine
    selectionTime: number, // Timeframe from which to select delivery entries (in milliseconds)
    selectionTimeFrame: {
      days: 0, // Number of days in the selection time frame 
      hours: 0, // Number of hours in the selection time frame 
    }, // Selection time frame
}`