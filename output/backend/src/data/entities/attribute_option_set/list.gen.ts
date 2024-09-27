import { AttributeOptionSet } from './entity'
import DtoList from '../../base_classes/dto_list'
import { type JSONData } from '../../base_classes/dto.d'
import { type IAttributeOptionSetList } from './list.gen.d'
import { type Pagination, type Query } from '../../base_classes/dto_list.d'


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


export class AttributeOptionSetList 
extends DtoList 
implements IAttributeOptionSetList 
{
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  _path = '/attribute_option_set'
  _kind = 'AttributeOptionSet'
  _name = 'attribute_option_set'

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
  list: AttributeOptionSet[] = []
  queriedList: AttributeOptionSet[] = []

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  _buildInstance(data: JSONData): AttributeOptionSet {
    return new AttributeOptionSet([], data)
  }

  override async getAll(
    pathVars?: string | string[] | Record<string, string>
  ): Promise<AttributeOptionSet[]> {
    return (await super.getAll(pathVars)) as AttributeOptionSet[]
  }

  override async getPaginated(
    pathVars: string | string[],
    pagination: Pagination
  ): Promise<{
    items: AttributeOptionSet[],
    pagination: Pagination,
  }> {
    // eslint-disable-next-line @typescript-eslint/no-unsafe-argument
    return (await super.getPaginated(pathVars, pagination)) as {
      items: AttributeOptionSet[],
      pagination: Pagination,
    }
  }

  override async query(queries: Query[], pagination?: Pagination): Promise<AttributeOptionSet[]> {
    return (await super.query(queries, pagination)) as AttributeOptionSet[];
  }

  override async search(query: string, pagination?: Pagination): Promise<AttributeOptionSet[]> {
    return (await super.search(query, pagination)) as AttributeOptionSet[];
  }
}
