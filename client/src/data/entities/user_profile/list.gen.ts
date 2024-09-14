import { UserProfile } from './entity'
import DtoList from '../../base_classes/dto_list'
import { type JSONData } from '../../base_classes/dto.d'
import { type IUserProfileList } from './list.gen.d'
import { Pagination } from "../../base_classes/pagination";
import { Query } from "../../base_classes/query";


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


export class UserProfileList 
extends DtoList 
implements IUserProfileList 
{
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  _path = '/user_profile'

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
  list: UserProfile[] = []
  queriedList: UserProfile[] = []

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  _buildInstance(data: JSONData): UserProfile {
    return new UserProfile([], data)
  }

  override async getAll(
    pathVars?: string | string[] | Record<string, string>
  ): Promise<UserProfile[]> {
    return (await super.getAll(pathVars)) as UserProfile[]
  }

  override async getPaginated(
    pathVars: string | string[],
    pagination: Pagination
  ): Promise<{
    items: UserProfile[],
    pagination: Pagination,
  }> {
    // eslint-disable-next-line @typescript-eslint/no-unsafe-argument
    return (await super.getPaginated(pathVars, pagination)) as {
      items: UserProfile[],
      pagination: Pagination,
    }
  }

  override async query(queries: Query[], pagination?: Pagination): Promise<{
    items: UserProfile[],
    pagination: Pagination,
  }> {
    return (await super.query(queries, pagination)) as {
      items: UserProfile[],
      pagination: Pagination,
    }
  }

  override async search(query: string, pagination?: Pagination): Promise<{
    items: UserProfile[],
    pagination: Pagination,
  }> {
    return (await super.search(query, pagination)) as {
      items: UserProfile[],
      pagination: Pagination,
    }
  }
}
