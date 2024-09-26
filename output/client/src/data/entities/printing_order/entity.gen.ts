import { type JSONData } from '../../base_classes/dto.d'
import { PrintingOrderDto } from './dto.gen'
import { type IPrintingOrder } from './entity.d'
import PrintingOrderParsers from './parsers'
import PrintingOrderValidators from './custom_validators'

export class PrintingOrder extends PrintingOrderDto implements IPrintingOrder {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, PrintingOrderParsers, PrintingOrderValidators)

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
  ): Promise<PrintingOrder> {
    return (await super.get(pathVars)) as PrintingOrder
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): PrintingOrder {
    return new PrintingOrder(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): PrintingOrder[] {
    return data.map((item) => PrintingOrder.FromJSON(item))
  }

  static FromDto(dto: PrintingOrderDto): PrintingOrder {
    return new PrintingOrder(undefined, dto.toJSON())
  }
}

