import { type JSONData } from '../../base_classes/dto.d'
import { AttributeDto } from './dto.gen'
import { type IAttribute } from './entity.d'
import AttributeParsers from './parsers'
import AttributeValidators from './custom_validators'

export class Attribute extends AttributeDto implements IAttribute {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, AttributeParsers, AttributeValidators)

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
  ): Promise<Attribute> {
    return (await super.get(pathVars)) as Attribute
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): Attribute {
    return new Attribute(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): Attribute[] {
    return data.map((item) => Attribute.FromJSON(item))
  }

  static FromDto(dto: AttributeDto): Attribute {
    return new Attribute(undefined, dto.toJSON())
  }
}

