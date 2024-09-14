import { type JSONData } from '../../base_classes/dto.d'
import { Dto } from '../../base_classes/dto'
import type { IAirlinePricingDto, IMargin, IValues,  } from './dto.gen.d'

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


// Types

export class AirlinePricingDto extends Dto implements IAirlinePricingDto {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  _path = '/airline_pricing/{id}'
  _reference = '/AirlinePricing/{id}'

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
    
    colors = [] as number[]
    
    
    description = ""
    
    
    dividers = 0
    
    
    kind = "AirlinePricing"
    
    
    margin: IMargin = {
      amount: 0,
      percentage: 0,
      percentageType: "profit-margin",
      staggered: false,
      type: "percentage",
      values: [] as any[],
    }

    
    
    units = [] as number[]
    
    
    values: IValues[] = []

    
    
    width = 0.0
    
    
    widthType = "MILLIMETERS"
    

  /// ///////////////////////////////////////
  /// ///////////// METHODS /////////////////
  /// ///////////////////////////////////////

  static FromJSON(data: unknown): AirlinePricingDto {
    // Create a new instance of the class
    const obj = new AirlinePricingDto()
    obj.set(data as JSONData)
    return obj
  }

  static FromJSONArray(data: unknown[]): AirlinePricingDto[] {
    return data.map((item) => AirlinePricingDto.FromJSON(item))
  }

  override set(data: JSONData): void {
    super.set(data)
 
    // Implement variables here
    if (data.colors !== undefined && data.colors !== null ) {
      this.colors = data.colors as number[] ?? this.colors
      
    }
    if (data.description !== undefined && data.description !== null ) {
      this.description = data.description as string ?? this.description
      
    }
    if (data.dividers !== undefined && data.dividers !== null ) {
      this.dividers = data.dividers as number ?? this.dividers
      
    }
    if (data.kind !== undefined && data.kind !== null ) {
      this.kind = data.kind as string ?? this.kind
      
    }
    if (data.margin !== undefined && data.margin !== null ) {
      if (data.margin) {
        this.margin = Margin.FromJSON(data.margin)
      }
    }
    if (data.units !== undefined && data.units !== null ) {
      this.units = data.units as number[] ?? this.units
      
    }
    if (data.values !== undefined && data.values !== null && Array.isArray(data.values)) {
      if (data.values) {
        this.values = Values.FromJSONArray(data.values as JSONData[])
      }
    }
    if (data.width !== undefined && data.width !== null ) {
      this.width = data.width as number ?? this.width
      
    }
    if (data.widthType !== undefined && data.widthType !== null ) {
      this.widthType = data.widthType as string ?? this.widthType
      
    }
  }

}

export class Margin implements IMargin {
  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
    
      amount = 0
    
    
      percentage = 0
    
    
      percentageType = "profit-margin"
    
    
      staggered = false
    
    
      type = "percentage"
    
    
      values = [] as any[]
    

  /// ///////////////////////////////////////
  /// ///////////// METHODS /////////////////
  /// ///////////////////////////////////////

  static FromJSON(data: unknown): Margin {
    // Create a new instance of the class
    const obj = new Margin()
    obj.set(data as JSONData)
    return obj
  }

  static FromJSONArray(data: unknown[]): Margin[] {
    return data.map((item) => Margin.FromJSON(item))
  }


  set(data: JSONData): void {
    // Implement variables here
    if (data.amount !== undefined && data.amount !== null ) {
      this.amount = data.amount as number ?? this.amount
      
    }
    if (data.percentage !== undefined && data.percentage !== null ) {
      this.percentage = data.percentage as number ?? this.percentage
      
    }
    if (data.percentageType !== undefined && data.percentageType !== null ) {
      this.percentageType = data.percentageType as string ?? this.percentageType
      
    }
    if (data.staggered !== undefined && data.staggered !== null ) {
      this.staggered = data.staggered as boolean ?? this.staggered
      
    }
    if (data.type !== undefined && data.type !== null ) {
      this.type = data.type as string ?? this.type
      
    }
    if (data.values !== undefined && data.values !== null ) {
      this.values = data.values as any[] ?? this.values
      
    }
  }
}

export class Values implements IValues {
  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
    
      attributes = new PricingAttributes()
    
    
      value = new PricingValue()
    

  /// ///////////////////////////////////////
  /// ///////////// METHODS /////////////////
  /// ///////////////////////////////////////

  static FromJSON(data: unknown): Values {
    // Create a new instance of the class
    const obj = new Values()
    obj.set(data as JSONData)
    return obj
  }

  static FromJSONArray(data: unknown[]): Values[] {
    return data.map((item) => Values.FromJSON(item))
  }


  set(data: JSONData): void {
    // Implement variables here
    if (data.attributes !== undefined && data.attributes !== null ) {
      if (data.attributes) {
        this.attributes = PricingAttributes.FromJSON(data.attributes)
      }
    }
    if (data.value !== undefined && data.value !== null ) {
      if (data.value) {
        this.value = PricingValue.FromJSON(data.value)
      }
    }
  }
}

export class PricingAttributes {
  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
    
      colors = 0
    
    
      units = 0
    

  /// ///////////////////////////////////////
  /// ///////////// METHODS /////////////////
  /// ///////////////////////////////////////

  static FromJSON(data: unknown): PricingAttributes {
    // Create a new instance of the class
    const obj = new PricingAttributes()
    obj.set(data as JSONData)
    return obj
  }

  static FromJSONArray(data: unknown[]): PricingAttributes[] {
    return data.map((item) => PricingAttributes.FromJSON(item))
  }


  set(data: JSONData): void {
    // Implement variables here
    if (data.colors !== undefined && data.colors !== null ) {
      this.colors = data.colors as number ?? this.colors
      
    }
    if (data.units !== undefined && data.units !== null ) {
      this.units = data.units as number ?? this.units
      
    }
  }
}

export class PricingValue {
  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
    
      decimalValue = 0
    
    
      floatValue = 0
    
    
      intValue = 0
    
    
      textValue = ""
    

  /// ///////////////////////////////////////
  /// ///////////// METHODS /////////////////
  /// ///////////////////////////////////////

  static FromJSON(data: unknown): PricingValue {
    // Create a new instance of the class
    const obj = new PricingValue()
    obj.set(data as JSONData)
    return obj
  }

  static FromJSONArray(data: unknown[]): PricingValue[] {
    return data.map((item) => PricingValue.FromJSON(item))
  }


  set(data: JSONData): void {
    // Implement variables here
    if (data.decimalValue !== undefined && data.decimalValue !== null ) {
      this.decimalValue = data.decimalValue as number ?? this.decimalValue
      
    }
    if (data.floatValue !== undefined && data.floatValue !== null ) {
      this.floatValue = data.floatValue as number ?? this.floatValue
      
    }
    if (data.intValue !== undefined && data.intValue !== null ) {
      this.intValue = data.intValue as number ?? this.intValue
      
    }
    if (data.textValue !== undefined && data.textValue !== null ) {
      this.textValue = data.textValue as string ?? this.textValue
      
    }
  }
}
