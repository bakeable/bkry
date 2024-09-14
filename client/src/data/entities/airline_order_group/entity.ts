import { type JSONData } from '../../base_classes/dto.d'
import { AirlineOrderGroupDto } from './dto.gen'
import { type IAirlineOrderGroup } from './entity.d'
import AirlineOrderGroupParsers from './parsers'
import AirlineOrderGroupValidators from './custom_validators'

export class AirlineOrderGroup extends AirlineOrderGroupDto implements IAirlineOrderGroup {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, AirlineOrderGroupParsers, AirlineOrderGroupValidators)

    if (data !== undefined) {
      // Set data
      this.set(data)
    }

    this.buildPathObject(pathVars);
  }
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  public async get(
    pathVars?: string | string[] | {
    id?: string,
  } | undefined
  ): Promise<AirlineOrderGroup> {
    return (await super.get(pathVars)) as AirlineOrderGroup
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): AirlineOrderGroup {
    return new AirlineOrderGroup(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): AirlineOrderGroup[] {
    return data.map((item) => AirlineOrderGroup.FromJSON(item))
  }

  static FromDto(dto: AirlineOrderGroupDto): AirlineOrderGroup {
    return new AirlineOrderGroup(undefined, dto.toJSON())
  }
}

