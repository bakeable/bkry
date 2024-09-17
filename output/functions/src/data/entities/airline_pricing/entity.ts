import { type JSONData } from '../../base_classes/dto.d'
import { AirlinePricingDto } from './dto.gen'
import { type IAirlinePricing } from './entity.d'
import AirlinePricingParsers from './parsers'
import AirlinePricingValidators from './custom_validators'

export class AirlinePricing extends AirlinePricingDto implements IAirlinePricing {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, AirlinePricingParsers, AirlinePricingValidators)

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
  ): Promise<AirlinePricing> {
    return (await super.get(pathVars)) as AirlinePricing
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): AirlinePricing {
    return new AirlinePricing(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): AirlinePricing[] {
    return data.map((item) => AirlinePricing.FromJSON(item))
  }

  static FromDto(dto: AirlinePricingDto): AirlinePricing {
    return new AirlinePricing(undefined, dto.toJSON())
  }
}

