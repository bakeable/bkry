import { type JSONData } from '../../base_classes/dto.d'
import { PackingSlipDto } from './dto.gen'
import { type IPackingSlip } from './entity.d'
import PackingSlipParsers from './parsers'
import PackingSlipValidators from './custom_validators'

export class PackingSlip extends PackingSlipDto implements IPackingSlip {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, PackingSlipParsers, PackingSlipValidators)

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
  ): Promise<PackingSlip> {
    return (await super.get(pathVars)) as PackingSlip
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): PackingSlip {
    return new PackingSlip(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): PackingSlip[] {
    return data.map((item) => PackingSlip.FromJSON(item))
  }

  static FromDto(dto: PackingSlipDto): PackingSlip {
    return new PackingSlip(undefined, dto.toJSON())
  }
}

