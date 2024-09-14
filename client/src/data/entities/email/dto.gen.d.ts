import { type JSONData } from '../../base_classes/dto.d'
import type IDto from '../../base_classes/dto.d'

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


// Types

// Interfaces

// Data Transfer Object Interface
export interface IEmailDto extends IDto {
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

    // The entity kind
    _kind: string

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /**
   * Set the Email's properties
   */
  set: (data: JSONData) => void
}
