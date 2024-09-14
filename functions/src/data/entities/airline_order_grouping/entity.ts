import { type JSONData } from '../../base_classes/dto.d'
import { AirlineOrderGroupingDto } from './dto.gen'
import { type IAirlineOrderGrouping } from './entity.d'
import AirlineOrderGroupingParsers from './parsers'
import AirlineOrderGroupingValidators from './custom_validators'

export class AirlineOrderGrouping extends AirlineOrderGroupingDto implements IAirlineOrderGrouping {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, AirlineOrderGroupingParsers, AirlineOrderGroupingValidators)

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
  ): Promise<AirlineOrderGrouping> {
    return (await super.get(pathVars)) as AirlineOrderGrouping
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): AirlineOrderGrouping {
    return new AirlineOrderGrouping(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): AirlineOrderGrouping[] {
    return data.map((item) => AirlineOrderGrouping.FromJSON(item))
  }

  static FromDto(dto: AirlineOrderGroupingDto): AirlineOrderGrouping {
    return new AirlineOrderGrouping(undefined, dto.toJSON())
  }
}

