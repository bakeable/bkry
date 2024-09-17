import { type JSONData } from '../../base_classes/dto.d'
import { AttributeOptionSetDto } from './dto.gen'
import { type IAttributeOptionSet } from './entity.d'
import AttributeOptionSetParsers from './parsers'
import AttributeOptionSetValidators from './custom_validators'

export class AttributeOptionSet extends AttributeOptionSetDto implements IAttributeOptionSet {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, AttributeOptionSetParsers, AttributeOptionSetValidators)

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
  ): Promise<AttributeOptionSet> {
    return (await super.get(pathVars)) as AttributeOptionSet
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): AttributeOptionSet {
    return new AttributeOptionSet(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): AttributeOptionSet[] {
    return data.map((item) => AttributeOptionSet.FromJSON(item))
  }

  static FromDto(dto: AttributeOptionSetDto): AttributeOptionSet {
    return new AttributeOptionSet(undefined, dto.toJSON())
  }
}

