import { type JSONData } from '../../base_classes/dto.d'
import { Dto } from '../../base_classes/dto'
import type { ICloudFunctionDto,  } from './dto.gen.d'

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


// Types

export class CloudFunctionDto extends Dto implements ICloudFunctionDto {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  _path = '/cloud_function/{id}'
  _reference = '/CloudFunction/{id}'

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
    
    description = ""
    
    
    _kind = "CloudFunction"
    
    
    name = ""
    
    
    region = ""
    
    
    runtime = ""
    
    
    triggerType = ""
    
    
    triggerValue = ""
    

  /// ///////////////////////////////////////
  /// ///////////// METHODS /////////////////
  /// ///////////////////////////////////////

  static FromJSON(data: unknown): CloudFunctionDto {
    // Create a new instance of the class
    const obj = new CloudFunctionDto()
    obj.set(data as JSONData)
    return obj
  }

  static FromJSONArray(data: unknown[]): CloudFunctionDto[] {
    return data.map((item) => CloudFunctionDto.FromJSON(item))
  }

  override set(data: JSONData): void {
    super.set(data)
 
    // Implement variables here
    if (data.description !== undefined && data.description !== null ) {
      this.description = data.description as string ?? this.description
      
    }
    if (data._kind !== undefined && data._kind !== null ) {
      this._kind = data._kind as string ?? this._kind
      
    }
    if (data.name !== undefined && data.name !== null ) {
      this.name = data.name as string ?? this.name
      
    }
    if (data.region !== undefined && data.region !== null ) {
      this.region = data.region as string ?? this.region
      
    }
    if (data.runtime !== undefined && data.runtime !== null ) {
      this.runtime = data.runtime as string ?? this.runtime
      
    }
    if (data.triggerType !== undefined && data.triggerType !== null ) {
      this.triggerType = data.triggerType as string ?? this.triggerType
      
    }
    if (data.triggerValue !== undefined && data.triggerValue !== null ) {
      this.triggerValue = data.triggerValue as string ?? this.triggerValue
      
    }
  }

}
