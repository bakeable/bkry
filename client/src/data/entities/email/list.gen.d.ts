import type IDtoList from '../../base_classes/dto_list'
import { type JSONData } from '../../base_classes/dto.d';
import { type Pagination } from '../../base_classes/pagination';
import { type Email } from '.';


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

export interface IEmailList extends IDtoList {

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
     * Build an instance of type Email from a JSON object
     * @param pathVars - The path variables
     * @returns A promise of Email[]
     */
    _buildInstance: (data: JSONData) => Email

    /**
     * Get all Emails
     * @param pathVars - The path variables
     * @returns A promise of Email[]
     */
    getAll: (pathVars?: string | string[] | Record<string, string>) => Promise<Email[]>

    /**
     * Get paginated Emails
     * @param pathVars - The path variables
     * @param pagination - The pagination object
     * @returns A promise of {
     *      items: Email[],
     *      pagination: Pagination,
     *  }
     */
    getPaginated: (pathVars: string | string[], pagination: Pagination) => Promise<{
        items: Email[],
        pagination: Pagination,
    }>

    /**
     * Query Emails
     * @param queries - The filters to query on
     * @returns A promise of Email[]
     */
    query: (queries: Query[], pagination?: Pagination) => Promise<{
        items: Email[],
        pagination: Pagination,
    }>

    /**
     * Search Emails
     * @param query - The search query
     * @returns A promise of Email[]
     */
    search: (query: string, pagination: Pagination) => Promise<{
        items: Email[],
        pagination: Pagination,
    }>
}