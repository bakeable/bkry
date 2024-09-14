import { type JSONData } from '../../base_classes/dto.d'
import { Dto } from '../../base_classes/dto'
import type { IPriceLayerDto, IVariablePriceValues, AttributeIncrementalPricing, PriceType, RoundingUnits,  } from './dto.gen.d'

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


// Types
export const AttributeIncrementalPricingValues: Record<string,AttributeIncrementalPricing> = Object.freeze({
  STAGGERED: 'staggered',
  STANDARD: 'standard',
})
export const PriceTypeValues: Record<string,PriceType> = Object.freeze({
  VARIABLE: 'variable',
  FIXED: 'fixed',
})
export const RoundingUnitsValues: Record<string,RoundingUnits> = Object.freeze({
  UP: 'up',
  RATIO: 'ratio',
  DOWN: 'down',
})

export class PriceLayerDto extends Dto implements IPriceLayerDto {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  _path = '/price_layer/{id}'
  _reference = '/PriceLayer/{id}'

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
    
    attributes = [] as string[]
    
    
    description = ""
    
    
    external = false
    
    
    includeInMarginCalculation = false
    
    
    includeInPurchasePrice = false
    
    
    incrementalPricing = AttributeIncrementalPricingValues.standard
    
    
    _kind = "PriceLayer"
    
    
    name = ""
    
    
    options = [] as string[]
    
    
    perUnits = 1
    
    
    priceType = PriceTypeValues.variable
    
    
    roundingUnits = RoundingUnitsValues.ratio
    
    
    saveAsNew = false
    
    
    template = ""
    
    
    values: IVariablePriceValues[] = []

    

  /// ///////////////////////////////////////
  /// ///////////// METHODS /////////////////
  /// ///////////////////////////////////////

  static FromJSON(data: unknown): PriceLayerDto {
    // Create a new instance of the class
    const obj = new PriceLayerDto()
    obj.set(data as JSONData)
    return obj
  }

  static FromJSONArray(data: unknown[]): PriceLayerDto[] {
    return data.map((item) => PriceLayerDto.FromJSON(item))
  }

  override set(data: JSONData): void {
    super.set(data)
 
    // Implement variables here
    if (data.attributes !== undefined && data.attributes !== null ) {
      this.attributes = data.attributes as string[] ?? this.attributes
      
    }
    if (data.description !== undefined && data.description !== null ) {
      this.description = data.description as string ?? this.description
      
    }
    if (data.external !== undefined && data.external !== null ) {
      this.external = data.external as boolean ?? this.external
      
    }
    if (data.includeInMarginCalculation !== undefined && data.includeInMarginCalculation !== null ) {
      this.includeInMarginCalculation = data.includeInMarginCalculation as boolean ?? this.includeInMarginCalculation
      
    }
    if (data.includeInPurchasePrice !== undefined && data.includeInPurchasePrice !== null ) {
      this.includeInPurchasePrice = data.includeInPurchasePrice as boolean ?? this.includeInPurchasePrice
      
    }
    if (data.incrementalPricing !== undefined && data.incrementalPricing !== null ) {
        this.incrementalPricing = AttributeIncrementalPricingValues[(data?.incrementalPricing as string)?.toUpperCase() ?? this.incrementalPricing.toUpperCase()]
    }
    if (data._kind !== undefined && data._kind !== null ) {
      this._kind = data._kind as string ?? this._kind
      
    }
    if (data.name !== undefined && data.name !== null ) {
      this.name = data.name as string ?? this.name
      
    }
    if (data.options !== undefined && data.options !== null ) {
      this.options = data.options as string[] ?? this.options
      
    }
    if (data.perUnits !== undefined && data.perUnits !== null ) {
      this.perUnits = data.perUnits as number ?? this.perUnits
      
    }
    if (data.priceType !== undefined && data.priceType !== null ) {
        this.priceType = PriceTypeValues[(data?.priceType as string)?.toUpperCase() ?? this.priceType.toUpperCase()]
    }
    if (data.roundingUnits !== undefined && data.roundingUnits !== null ) {
        this.roundingUnits = RoundingUnitsValues[(data?.roundingUnits as string)?.toUpperCase() ?? this.roundingUnits.toUpperCase()]
    }
    if (data.saveAsNew !== undefined && data.saveAsNew !== null ) {
      this.saveAsNew = data.saveAsNew as boolean ?? this.saveAsNew
      
    }
    if (data.template !== undefined && data.template !== null ) {
      this.template = data.template as string ?? this.template
      
    }
    if (data.values !== undefined && data.values !== null && Array.isArray(data.values)) {
      if (data.values) {
        this.values = VariablePriceValues.FromJSONArray(data.values as JSONData[])
      }
    }
  }

}

export class VariablePriceValues implements IVariablePriceValues {
  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
    
      attributes = {} satisfies Record<string, any>
    
    
      value = new VariablePriceValue()
    

  /// ///////////////////////////////////////
  /// ///////////// METHODS /////////////////
  /// ///////////////////////////////////////

  static FromJSON(data: unknown): VariablePriceValues {
    // Create a new instance of the class
    const obj = new VariablePriceValues()
    obj.set(data as JSONData)
    return obj
  }

  static FromJSONArray(data: unknown[]): VariablePriceValues[] {
    return data.map((item) => VariablePriceValues.FromJSON(item))
  }


  set(data: JSONData): void {
    // Implement variables here
    if (data.attributes !== undefined && data.attributes !== null ) {
      this.attributes = data.attributes as Record<string, any> ?? this.attributes
      
    }
    if (data.value !== undefined && data.value !== null ) {
      if (data.value) {
        this.value = VariablePriceValue.FromJSON(data.value)
      }
    }
  }
}

export class VariablePriceValue {
  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
    
      decimalValue = 0.0
    
    
      floatValue = 0.0
    
    
      intValue = 0
    
    
      textValue = ""
    

  /// ///////////////////////////////////////
  /// ///////////// METHODS /////////////////
  /// ///////////////////////////////////////

  static FromJSON(data: unknown): VariablePriceValue {
    // Create a new instance of the class
    const obj = new VariablePriceValue()
    obj.set(data as JSONData)
    return obj
  }

  static FromJSONArray(data: unknown[]): VariablePriceValue[] {
    return data.map((item) => VariablePriceValue.FromJSON(item))
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
