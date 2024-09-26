import { type JSONData } from '../../base_classes/dto.d'
import { ExaminationSettingsDto } from './dto.gen'
import { type IExaminationSettings } from './entity.d'
import ExaminationSettingsParsers from './parsers'
import ExaminationSettingsValidators from './custom_validators'

export class ExaminationSettings extends ExaminationSettingsDto implements IExaminationSettings {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, ExaminationSettingsParsers, ExaminationSettingsValidators)

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
  ): Promise<ExaminationSettings> {
    return (await super.get(pathVars)) as ExaminationSettings
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): ExaminationSettings {
    return new ExaminationSettings(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): ExaminationSettings[] {
    return data.map((item) => ExaminationSettings.FromJSON(item))
  }

  static FromDto(dto: ExaminationSettingsDto): ExaminationSettings {
    return new ExaminationSettings(undefined, dto.toJSON())
  }
}

