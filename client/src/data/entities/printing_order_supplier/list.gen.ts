import { PrintingOrderSupplier } from './entity'
import DtoList from '../../base_classes/dto_list'
import { type JSONData } from '../../base_classes/dto.d'
import { type IPrintingOrderSupplierList } from './list.gen.d'
import { Pagination } from "../../base_classes/pagination";
import { Query } from "../../base_classes/query";


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


export class PrintingOrderSupplierList 
extends DtoList 
implements IPrintingOrderSupplierList 
{
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  _path = '/printing_order_supplier'

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
  list: PrintingOrderSupplier[] = []
  queriedList: PrintingOrderSupplier[] = []

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  _buildInstance(data: JSONData): PrintingOrderSupplier {
    return new PrintingOrderSupplier([], data)
  }

  override async getAll(
    pathVars?: string | string[] | Record<string, string>
  ): Promise<PrintingOrderSupplier[]> {
    return (await super.getAll(pathVars)) as PrintingOrderSupplier[]
  }

  override async getPaginated(
    pathVars: string | string[],
    pagination: Pagination
  ): Promise<{
    items: PrintingOrderSupplier[],
    pagination: Pagination,
  }> {
    // eslint-disable-next-line @typescript-eslint/no-unsafe-argument
    return (await super.getPaginated(pathVars, pagination)) as {
      items: PrintingOrderSupplier[],
      pagination: Pagination,
    }
  }

  override async query(queries: Query[], pagination?: Pagination): Promise<{
    items: PrintingOrderSupplier[],
    pagination: Pagination,
  }> {
    return (await super.query(queries, pagination)) as {
      items: PrintingOrderSupplier[],
      pagination: Pagination,
    }
  }

  override async search(query: string, pagination?: Pagination): Promise<{
    items: PrintingOrderSupplier[],
    pagination: Pagination,
  }> {
    return (await super.search(query, pagination)) as {
      items: PrintingOrderSupplier[],
      pagination: Pagination,
    }
  }
}
