import { type JSONData } from '../../base_classes/dto.d'
import type IDto from '../../base_classes/dto.d'

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


// Types

// Interfaces
export interface IExclusionTimeFrame {
    // Number of days in the exclusion time frame
      days: number
    // Number of hours in the exclusion time frame
      hours: number
}
export interface ISelectionTimeFrame {
    // Number of days in the selection time frame
      days: number
    // Number of hours in the selection time frame
      hours: number
}

// Data Transfer Object Interface
export interface IExaminationSettingsDto extends IDto {
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

    // Standard timeframe from which to select tasks (in milliseconds)
    exclusionTime: number

    // Exclusion time frame
    exclusionTimeFrame: IExclusionTimeFrame

    // The kind of the object
    kind: string

    // Maximum number of open tasks
    maxOpenTasks: number

    // Maximum number of delivery entries to examine
    maxProducts: number

    // Maximum number of properties to examine
    maxProperties: number

    // Minimum number of delivery entries to examine
    minProducts: number

    // Minimum number of properties to examine
    minProperties: number

    // Properties to examine
    properties: string[]

    // Timeframe from which to select delivery entries (in milliseconds)
    selectionTime: number

    // Selection time frame
    selectionTimeFrame: ISelectionTimeFrame

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /**
   * Set the ExaminationSettings's properties
   */
  set: (data: JSONData) => void
}
