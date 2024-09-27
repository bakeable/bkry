import { type JSONData } from '../../base_classes/dto.d'
import { PrintingOrderItemDto } from './dto.gen'
import { type IPrintingOrderItem } from './entity.d'
import PrintingOrderItemParsers from './parsers'
import PrintingOrderItemValidators from './custom_validators'

export class PrintingOrderItem extends PrintingOrderItemDto implements IPrintingOrderItem {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, PrintingOrderItemParsers, PrintingOrderItemValidators)

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
  ): Promise<PrintingOrderItem> {
    return (await super.get(pathVars)) as PrintingOrderItem
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): PrintingOrderItem {
    return new PrintingOrderItem(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): PrintingOrderItem[] {
    return data.map((item) => PrintingOrderItem.FromJSON(item))
  }

  static FromDto(dto: PrintingOrderItemDto): PrintingOrderItem {
    return new PrintingOrderItem(undefined, dto.toJSON())
  }
}

