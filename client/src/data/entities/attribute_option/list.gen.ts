import { AttributeOption } from './entity'
import DtoList from '../../base_classes/dto_list'
import { type JSONData } from '../../base_classes/dto.d'
import { type IAttributeOptionList } from './list.gen.d'
import { Pagination } from "../../base_classes/pagination";
import { Query } from "../../base_classes/query";


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


export class AttributeOptionList 
extends DtoList 
implements IAttributeOptionList 
{
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  _path = '/attribute_option'

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
  list: AttributeOption[] = []
  queriedList: AttributeOption[] = []

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  _buildInstance(data: JSONData): AttributeOption {
    return new AttributeOption([], data)
  }

  override async getAll(
    pathVars?: string | string[] | Record<string, string>
  ): Promise<AttributeOption[]> {
    return (await super.getAll(pathVars)) as AttributeOption[]
  }

  override async getPaginated(
    pathVars: string | string[],
    pagination: Pagination
  ): Promise<{
    items: AttributeOption[],
    pagination: Pagination,
  }> {
    // eslint-disable-next-line @typescript-eslint/no-unsafe-argument
    return (await super.getPaginated(pathVars, pagination)) as {
      items: AttributeOption[],
      pagination: Pagination,
    }
  }

  override async query(queries: Query[], pagination?: Pagination): Promise<{
    items: AttributeOption[],
    pagination: Pagination,
  }> {
    return (await super.query(queries, pagination)) as {
      items: AttributeOption[],
      pagination: Pagination,
    }
  }

  override async search(query: string, pagination?: Pagination): Promise<{
    items: AttributeOption[],
    pagination: Pagination,
  }> {
    return (await super.search(query, pagination)) as {
      items: AttributeOption[],
      pagination: Pagination,
    }
  }
}
