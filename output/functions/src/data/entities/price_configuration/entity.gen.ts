import { type JSONData } from '../../base_classes/dto.d'
import { PriceConfigurationDto } from './dto.gen'
import { type IPriceConfiguration } from './entity.d'
import PriceConfigurationParsers from './parsers'
import PriceConfigurationValidators from './custom_validators'

export class PriceConfiguration extends PriceConfigurationDto implements IPriceConfiguration {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, PriceConfigurationParsers, PriceConfigurationValidators)

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
  ): Promise<PriceConfiguration> {
    return (await super.get(pathVars)) as PriceConfiguration
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): PriceConfiguration {
    return new PriceConfiguration(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): PriceConfiguration[] {
    return data.map((item) => PriceConfiguration.FromJSON(item))
  }

  static FromDto(dto: PriceConfigurationDto): PriceConfiguration {
    return new PriceConfiguration(undefined, dto.toJSON())
  }
}

