import { type JSONData } from '../../base_classes/dto.d'
import { ExaminationPropertyDto } from './dto.gen'
import { type IExaminationProperty } from './entity.d'
import ExaminationPropertyParsers from './parsers'
import ExaminationPropertyValidators from './custom_validators'

export class ExaminationProperty extends ExaminationPropertyDto implements IExaminationProperty {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, ExaminationPropertyParsers, ExaminationPropertyValidators)

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
  ): Promise<ExaminationProperty> {
    return (await super.get(pathVars)) as ExaminationProperty
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): ExaminationProperty {
    return new ExaminationProperty(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): ExaminationProperty[] {
    return data.map((item) => ExaminationProperty.FromJSON(item))
  }

  static FromDto(dto: ExaminationPropertyDto): ExaminationProperty {
    return new ExaminationProperty(undefined, dto.toJSON())
  }
}

