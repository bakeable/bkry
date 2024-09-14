import { type JSONData } from '../../base_classes/dto.d'
import type IDto from '../../base_classes/dto.d'

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


// Types
export type AttributeIncrementalPricing = "staggered" | "standard"
export type PriceType = "variable" | "fixed"
export type RoundingUnits = "up" | "ratio" | "down"

// Interfaces
export interface IVariablePriceValues {
    // The attributes associated with the variable price
      attributes: Record<string, any>
    // The value of the variable price
    value: {
    // The decimal value of the variable price
      decimalValue: number
    // The float value of the variable price
      floatValue: number
    // The integer value of the variable price
      intValue: number
    // The text value of the variable price
      textValue: string
    }
}

// Data Transfer Object Interface
export interface IPriceLayerDto extends IDto {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  /**
   * The path to the corresponding API endpoint with variables marked as {?}
   */
  _path: string

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

    // The attributes associated with the variable price
    attributes: string[]

    // The description of the price layer
    description: string

    // Whether the price layer is external
    external: boolean

    // Whether the price layer should be included in the margin calculation
    includeInMarginCalculation: boolean

    // Whether the price layer should be included in the purchase price
    includeInPurchasePrice: boolean

    // The type of incremental pricing of the price layer
    incrementalPricing: AttributeIncrementalPricing

    // The kind of the object
    _kind: string

    // The name of the price layer
    name: string

    // The options associated with the price layer
    options: string[]

    // The unit size of the variable price
    perUnits: number

    // The pricing type of the price layer
    priceType: PriceType

    // How to round the unit sizes if they do not correspond to the unit size
    roundingUnits: RoundingUnits

    // Whether to save the price-layer to be used later
    saveAsNew: boolean

    // The template of the price layer
    template: string

    // The values of the variable price
    values: IVariablePriceValues[]

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /**
   * Set the PriceLayer's properties
   */
  set: (data: JSONData) => void
}
