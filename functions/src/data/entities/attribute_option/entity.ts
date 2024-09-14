import { type JSONData } from '../../base_classes/dto.d'
import { AttributeOptionDto } from './dto.gen'
import { type IAttributeOption } from './entity.d'
import AttributeOptionParsers from './parsers'
import AttributeOptionValidators from './custom_validators'

export class AttributeOption extends AttributeOptionDto implements IAttributeOption {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, AttributeOptionParsers, AttributeOptionValidators)

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
  ): Promise<AttributeOption> {
    return (await super.get(pathVars)) as AttributeOption
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): AttributeOption {
    return new AttributeOption(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): AttributeOption[] {
    return data.map((item) => AttributeOption.FromJSON(item))
  }

  static FromDto(dto: AttributeOptionDto): AttributeOption {
    return new AttributeOption(undefined, dto.toJSON())
  }
}

