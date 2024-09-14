import { Media } from './entity'
import DtoList from '../../base_classes/dto_list'
import { type JSONData } from '../../base_classes/dto.d'
import { type IMediaList } from './list.gen.d'
import { type Pagination, type Query } from '../../base_classes/dto_list.d'


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


export class MediaList 
extends DtoList 
implements IMediaList 
{
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  _path = '/media'
  _kind = 'Media'
  _name = 'media'

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
  list: Media[] = []
  queriedList: Media[] = []

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  _buildInstance(data: JSONData): Media {
    return new Media([], data)
  }

  override async getAll(
    pathVars?: string | string[] | Record<string, string>
  ): Promise<Media[]> {
    return (await super.getAll(pathVars)) as Media[]
  }

  override async getPaginated(
    pathVars: string | string[],
    pagination: Pagination
  ): Promise<{
    items: Media[],
    pagination: Pagination,
  }> {
    // eslint-disable-next-line @typescript-eslint/no-unsafe-argument
    return (await super.getPaginated(pathVars, pagination)) as {
      items: Media[],
      pagination: Pagination,
    }
  }

  override async query(queries: Query[], pagination?: Pagination): Promise<Media[]> {
    return (await super.query(queries, pagination)) as Media[];
  }

  override async search(query: string, pagination?: Pagination): Promise<Media[]> {
    return (await super.search(query, pagination)) as Media[];
  }
}
