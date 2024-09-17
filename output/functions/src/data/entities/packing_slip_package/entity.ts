import { type JSONData } from '../../base_classes/dto.d'
import { PackingSlipPackageDto } from './dto.gen'
import { type IPackingSlipPackage } from './entity.d'
import PackingSlipPackageParsers from './parsers'
import PackingSlipPackageValidators from './custom_validators'

export class PackingSlipPackage extends PackingSlipPackageDto implements IPackingSlipPackage {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, PackingSlipPackageParsers, PackingSlipPackageValidators)

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
  ): Promise<PackingSlipPackage> {
    return (await super.get(pathVars)) as PackingSlipPackage
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): PackingSlipPackage {
    return new PackingSlipPackage(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): PackingSlipPackage[] {
    return data.map((item) => PackingSlipPackage.FromJSON(item))
  }

  static FromDto(dto: PackingSlipPackageDto): PackingSlipPackage {
    return new PackingSlipPackage(undefined, dto.toJSON())
  }
}

