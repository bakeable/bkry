import type IDtoList from '../../base_classes/dto_list'
import { type JSONData } from '../../base_classes/dto.d';
import { type Pagination } from '../../base_classes/pagination';
import { type User } from '.';


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

export interface IUserList extends IDtoList {

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
     * Build an instance of type User from a JSON object
     * @param pathVars - The path variables
     * @returns A promise of User[]
     */
    _buildInstance: (data: JSONData) => User

    /**
     * Get all Users
     * @param pathVars - The path variables
     * @returns A promise of User[]
     */
    getAll: (pathVars?: string | string[] | Record<string, string>) => Promise<User[]>

    /**
     * Get paginated Users
     * @param pathVars - The path variables
     * @param pagination - The pagination object
     * @returns A promise of {
     *      items: User[],
     *      pagination: Pagination,
     *  }
     */
    getPaginated: (pathVars: string | string[], pagination: Pagination) => Promise<{
        items: User[],
        pagination: Pagination,
    }>

    /**
     * Query Users
     * @param queries - The filters to query on
     * @returns A promise of User[]
     */
    query: (queries: Query[], pagination?: Pagination) => Promise<{
        items: User[],
        pagination: Pagination,
    }>

    /**
     * Search Users
     * @param query - The search query
     * @returns A promise of User[]
     */
    search: (query: string, pagination: Pagination) => Promise<{
        items: User[],
        pagination: Pagination,
    }>
}