import { type JSONData } from '../../base_classes/dto.d'
import { Dto } from '../../base_classes/dto'
import type { IPackingSlipPackageDto, IDimensions, TranssmartCode,  } from './dto.gen.d'
import { Dimension } from '../dimension'

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


// Types
export const TranssmartCodeValues: Record<string,TranssmartCode> = Object.freeze({
  BOX: 'BOX',
  BLOKPALLET: 'BLOKPALLET',
  EUROGESTAPELD: 'EUROGESTAPELD',
  EUROPALLET: 'EUROPALLET',
  PALLETOVERIG: 'PALLETOVERIG',
  WEGWERPGES: 'WEGWERPGES',
  WEGWERPPALLET: 'WEGWERPPALLET',
})

export class PackingSlipPackageDto extends Dto implements IPackingSlipPackageDto {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  _path = '/packing_slip_package/{id}'
  _reference = '/PackingSlipPackage/{id}'

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
    
    dimensions: IDimensions = {
      height: new Dimension(),
      length: new Dimension(),
      width: new Dimension(),
    }

    
    
    _kind = "PackingSlipPackage"
    
    
    name = ""
    
    
    transsmartCode = TranssmartCodeValues.BOX
    
    
    type = ""
    
    
    weight = new Dimension()
    

  /// ///////////////////////////////////////
  /// ///////////// METHODS /////////////////
  /// ///////////////////////////////////////

  static FromJSON(data: unknown): PackingSlipPackageDto {
    // Create a new instance of the class
    const obj = new PackingSlipPackageDto()
    obj.set(data as JSONData)
    return obj
  }

  static FromJSONArray(data: unknown[]): PackingSlipPackageDto[] {
    return data.map((item) => PackingSlipPackageDto.FromJSON(item))
  }

  override set(data: JSONData): void {
    super.set(data)
 
    // Implement variables here
    if (data.dimensions !== undefined && data.dimensions !== null ) {
      if (data.dimensions) {
        this.dimensions = Dimensions.FromJSON(data.dimensions)
      }
    }
    if (data._kind !== undefined && data._kind !== null ) {
      this._kind = data._kind as string ?? this._kind
      
    }
    if (data.name !== undefined && data.name !== null ) {
      this.name = data.name as string ?? this.name
      
    }
    if (data.transsmartCode !== undefined && data.transsmartCode !== null ) {
        this.transsmartCode = TranssmartCodeValues[(data?.transsmartCode as string)?.toUpperCase() ?? this.transsmartCode.toUpperCase()]
    }
    if (data.type !== undefined && data.type !== null ) {
      this.type = data.type as string ?? this.type
      
    }
    if (data.weight !== undefined && data.weight !== null ) {
      if (!this.weight) {
        this.weight = new Dimension()
      }
      this.weight.set(data.weight as JSONData ?? this.weight.toJSON())
    }
  }

}

export class Dimensions implements IDimensions {
  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
    
      height = new Dimension()
    
    
      length = new Dimension()
    
    
      width = new Dimension()
    

  /// ///////////////////////////////////////
  /// ///////////// METHODS /////////////////
  /// ///////////////////////////////////////

  static FromJSON(data: unknown): Dimensions {
    // Create a new instance of the class
    const obj = new Dimensions()
    obj.set(data as JSONData)
    return obj
  }

  static FromJSONArray(data: unknown[]): Dimensions[] {
    return data.map((item) => Dimensions.FromJSON(item))
  }


  set(data: JSONData): void {
    // Implement variables here
    if (data.height !== undefined && data.height !== null ) {
      if (!this.height) {
        this.height = new Dimension()
      }
      this.height.set(data.height as JSONData ?? this.height.toJSON())
    }
    if (data.length !== undefined && data.length !== null ) {
      if (!this.length) {
        this.length = new Dimension()
      }
      this.length.set(data.length as JSONData ?? this.length.toJSON())
    }
    if (data.width !== undefined && data.width !== null ) {
      if (!this.width) {
        this.width = new Dimension()
      }
      this.width.set(data.width as JSONData ?? this.width.toJSON())
    }
  }
}
