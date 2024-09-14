import { PrintingOrderItem } from './entity'
import DtoList from '../../base_classes/dto_list'
import { type JSONData } from '../../base_classes/dto.d'
import { type IPrintingOrderItemList } from './list.gen.d'
import { type Pagination, type Query } from '../../base_classes/dto_list.d'


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


export class PrintingOrderItemList 
extends DtoList 
implements IPrintingOrderItemList 
{
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  _path = '/printing_order_item'
  _kind = 'PrintingOrderItem'
  _name = 'printing_order_item'

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
  list: PrintingOrderItem[] = []
  queriedList: PrintingOrderItem[] = []

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  _buildInstance(data: JSONData): PrintingOrderItem {
    return new PrintingOrderItem([], data)
  }

  override async getAll(
    pathVars?: string | string[] | Record<string, string>
  ): Promise<PrintingOrderItem[]> {
    return (await super.getAll(pathVars)) as PrintingOrderItem[]
  }

  override async getPaginated(
    pathVars: string | string[],
    pagination: Pagination
  ): Promise<{
    items: PrintingOrderItem[],
    pagination: Pagination,
  }> {
    // eslint-disable-next-line @typescript-eslint/no-unsafe-argument
    return (await super.getPaginated(pathVars, pagination)) as {
      items: PrintingOrderItem[],
      pagination: Pagination,
    }
  }

  override async query(queries: Query[], pagination?: Pagination): Promise<PrintingOrderItem[]> {
    return (await super.query(queries, pagination)) as PrintingOrderItem[];
  }

  override async search(query: string, pagination?: Pagination): Promise<PrintingOrderItem[]> {
    return (await super.search(query, pagination)) as PrintingOrderItem[];
  }
}
