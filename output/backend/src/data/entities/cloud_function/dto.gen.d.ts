import { type JSONData } from '../../base_classes/dto.d'
import type IDto from '../../base_classes/dto.d'

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


// Types

// Interfaces

// Data Transfer Object Interface
export interface ICloudFunctionDto extends IDto {
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

    // The cloud function's description
    description: string

    // The kind of the object
    _kind: string

    // The cloud function's name
    name: string

    // The cloud function's region
    region: string

    // The cloud function's runtime
    runtime: string

    // The cloud function's trigger type
    triggerType: string

    // The cloud function's trigger value
    triggerValue: string

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /**
   * Set the CloudFunction's properties
   */
  set: (data: JSONData) => void
}
