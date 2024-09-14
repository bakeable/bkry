import { type JSONData } from '../../base_classes/dto.d'
import { ApiKeyDto } from './dto.gen'
import { type IApiKey } from './entity.d'
import ApiKeyParsers from './parsers'
import ApiKeyValidators from './custom_validators'

export class ApiKey extends ApiKeyDto implements IApiKey {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, ApiKeyParsers, ApiKeyValidators)

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
  ): Promise<ApiKey> {
    return (await super.get(pathVars)) as ApiKey
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): ApiKey {
    return new ApiKey(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): ApiKey[] {
    return data.map((item) => ApiKey.FromJSON(item))
  }

  static FromDto(dto: ApiKeyDto): ApiKey {
    return new ApiKey(undefined, dto.toJSON())
  }
}

