import { type JSONData } from '../../base_classes/dto.d'
import { CloudFunctionDto } from './dto.gen'
import { type ICloudFunction } from './entity.d'
import CloudFunctionParsers from './parsers'
import CloudFunctionValidators from './custom_validators'

export class CloudFunction extends CloudFunctionDto implements ICloudFunction {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, CloudFunctionParsers, CloudFunctionValidators)

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
  ): Promise<CloudFunction> {
    return (await super.get(pathVars)) as CloudFunction
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): CloudFunction {
    return new CloudFunction(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): CloudFunction[] {
    return data.map((item) => CloudFunction.FromJSON(item))
  }

  static FromDto(dto: CloudFunctionDto): CloudFunction {
    return new CloudFunction(undefined, dto.toJSON())
  }
}

