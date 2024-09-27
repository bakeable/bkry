package packing_slip_package

var ExampleJSON = `{
    dimensions: {
      height: [] as Dimension, // The height of the package 
      length: [] as Dimension, // The length of the package 
      width: [] as Dimension, // The width of the package 
    }, // The dimensions of the package
    _kind: string, // The kind of the datastore entity
    name: string, // The name of the package
    transsmartCode: TranssmartCode, // The code of the package in Transsmart
    type: string, // The type of the package
}`