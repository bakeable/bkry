import type IDtoList from '../../base_classes/dto_list'
import { type JSONData } from '../../base_classes/dto.d';
import { type Pagination } from '../../base_classes/pagination';
import { type UserProfile } from '.';


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

export interface IUserProfileList extends IDtoList {

    /// ///////////////////////////////////////
    /// //////// PRIVATE VARIABLES ////////////
    /// ///////////////////////////////////////

    /// ///////////////////////////////////////
    /// ///////// PUBLIC VARIABLES ////////////
    /// ///////////////////////////////////////

    /// ///////////////////////////////////////
    /// ////////////// METHODS ////////////////
    /// ///////////////////////////////////////

    /**
     * Build an instance of type UserProfile from a JSON object
     * @param pathVars - The path variables
     * @returns A promise of UserProfile[]
     */
    _buildInstance: (data: JSONData) => UserProfile

    /**
     * Get all UserProfiles
     * @param pathVars - The path variables
     * @returns A promise of UserProfile[]
     */
    getAll: (pathVars?: string | string[] | Record<string, string>) => Promise<UserProfile[]>

    /**
     * Get paginated UserProfiles
     * @param pathVars - The path variables
     * @param pagination - The pagination object
     * @returns A promise of {
     *      items: UserProfile[],
     *      pagination: Pagination,
     *  }
     */
    getPaginated: (pathVars: string | string[], pagination: Pagination) => Promise<{
        items: UserProfile[],
        pagination: Pagination,
    }>

    /**
     * Query UserProfiles
     * @param queries - The filters to query on
     * @returns A promise of UserProfile[]
     */
    query: (queries: Query[], pagination?: Pagination) => Promise<{
        items: UserProfile[],
        pagination: Pagination,
    }>

    /**
     * Search UserProfiles
     * @param query - The search query
     * @returns A promise of UserProfile[]
     */
    search: (query: string, pagination: Pagination) => Promise<{
        items: UserProfile[],
        pagination: Pagination,
    }>
}