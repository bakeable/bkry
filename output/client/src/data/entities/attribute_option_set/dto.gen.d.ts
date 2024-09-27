import { type JSONData } from '../../base_classes/dto.d'
import type IDto from '../../base_classes/dto.d'

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


// Types

// Interfaces

// Data Transfer Object Interface
export interface IAttributeOptionSetDto extends IDto {
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

    // The optionset's label
    label: string

    // The optionset's options
    options: attribute_option.AttributeOption[]

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /**
   * Set the AttributeOptionSet's properties
   */
  set: (data: JSONData) => void
}
