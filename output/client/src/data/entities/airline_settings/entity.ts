import { type JSONData } from '../../base_classes/dto.d'
import { AirlineSettingsDto } from './dto.gen'
import { type IAirlineSettings } from './entity.d'
import AirlineSettingsParsers from './parsers'
import AirlineSettingsValidators from './custom_validators'

export class AirlineSettings extends AirlineSettingsDto implements IAirlineSettings {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, AirlineSettingsParsers, AirlineSettingsValidators)

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
  ): Promise<AirlineSettings> {
    return (await super.get(pathVars)) as AirlineSettings
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): AirlineSettings {
    return new AirlineSettings(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): AirlineSettings[] {
    return data.map((item) => AirlineSettings.FromJSON(item))
  }

  static FromDto(dto: AirlineSettingsDto): AirlineSettings {
    return new AirlineSettings(undefined, dto.toJSON())
  }
}

