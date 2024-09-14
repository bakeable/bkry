import { PriceConfiguration } from './entity'
import DtoList from '../../base_classes/dto_list'
import { type JSONData } from '../../base_classes/dto.d'
import { type IPriceConfigurationList } from './list.gen.d'
import { Pagination } from "../../base_classes/pagination";
import { Query } from "../../base_classes/query";


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


export class PriceConfigurationList 
extends DtoList 
implements IPriceConfigurationList 
{
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  _path = '/price_configuration'

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
  list: PriceConfiguration[] = []
  queriedList: PriceConfiguration[] = []

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  _buildInstance(data: JSONData): PriceConfiguration {
    return new PriceConfiguration([], data)
  }

  override async getAll(
    pathVars?: string | string[] | Record<string, string>
  ): Promise<PriceConfiguration[]> {
    return (await super.getAll(pathVars)) as PriceConfiguration[]
  }

  override async getPaginated(
    pathVars: string | string[],
    pagination: Pagination
  ): Promise<{
    items: PriceConfiguration[],
    pagination: Pagination,
  }> {
    // eslint-disable-next-line @typescript-eslint/no-unsafe-argument
    return (await super.getPaginated(pathVars, pagination)) as {
      items: PriceConfiguration[],
      pagination: Pagination,
    }
  }

  override async query(queries: Query[], pagination?: Pagination): Promise<{
    items: PriceConfiguration[],
    pagination: Pagination,
  }> {
    return (await super.query(queries, pagination)) as {
      items: PriceConfiguration[],
      pagination: Pagination,
    }
  }

  override async search(query: string, pagination?: Pagination): Promise<{
    items: PriceConfiguration[],
    pagination: Pagination,
  }> {
    return (await super.search(query, pagination)) as {
      items: PriceConfiguration[],
      pagination: Pagination,
    }
  }
}
