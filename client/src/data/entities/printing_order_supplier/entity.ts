import { type JSONData } from '../../base_classes/dto.d'
import { PrintingOrderSupplierDto } from './dto.gen'
import { type IPrintingOrderSupplier } from './entity.d'
import PrintingOrderSupplierParsers from './parsers'
import PrintingOrderSupplierValidators from './custom_validators'

export class PrintingOrderSupplier extends PrintingOrderSupplierDto implements IPrintingOrderSupplier {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, PrintingOrderSupplierParsers, PrintingOrderSupplierValidators)

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
  ): Promise<PrintingOrderSupplier> {
    return (await super.get(pathVars)) as PrintingOrderSupplier
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): PrintingOrderSupplier {
    return new PrintingOrderSupplier(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): PrintingOrderSupplier[] {
    return data.map((item) => PrintingOrderSupplier.FromJSON(item))
  }

  static FromDto(dto: PrintingOrderSupplierDto): PrintingOrderSupplier {
    return new PrintingOrderSupplier(undefined, dto.toJSON())
  }
}

