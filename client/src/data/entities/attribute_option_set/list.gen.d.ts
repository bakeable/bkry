import type IDtoList from '../../base_classes/dto_list'
import { type JSONData } from '../../base_classes/dto.d';
import { type Pagination } from '../../base_classes/pagination';
import { type AttributeOptionSet } from '.';


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

export interface IAttributeOptionSetList extends IDtoList {

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
     * Build an instance of type AttributeOptionSet from a JSON object
     * @param pathVars - The path variables
     * @returns A promise of AttributeOptionSet[]
     */
    _buildInstance: (data: JSONData) => AttributeOptionSet

    /**
     * Get all AttributeOptionSets
     * @param pathVars - The path variables
     * @returns A promise of AttributeOptionSet[]
     */
    getAll: (pathVars?: string | string[] | Record<string, string>) => Promise<AttributeOptionSet[]>

    /**
     * Get paginated AttributeOptionSets
     * @param pathVars - The path variables
     * @param pagination - The pagination object
     * @returns A promise of {
     *      items: AttributeOptionSet[],
     *      pagination: Pagination,
     *  }
     */
    getPaginated: (pathVars: string | string[], pagination: Pagination) => Promise<{
        items: AttributeOptionSet[],
        pagination: Pagination,
    }>

    /**
     * Query AttributeOptionSets
     * @param queries - The filters to query on
     * @returns A promise of AttributeOptionSet[]
     */
    query: (queries: Query[], pagination?: Pagination) => Promise<{
        items: AttributeOptionSet[],
        pagination: Pagination,
    }>

    /**
     * Search AttributeOptionSets
     * @param query - The search query
     * @returns A promise of AttributeOptionSet[]
     */
    search: (query: string, pagination: Pagination) => Promise<{
        items: AttributeOptionSet[],
        pagination: Pagination,
    }>
}