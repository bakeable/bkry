import type IDtoList from '../../base_classes/dto_list'
import { type JSONData } from '../../base_classes/dto.d';
import { type Pagination } from '../../base_classes/pagination';
import { type AirlineOrderGroup } from '.';


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

export interface IAirlineOrderGroupList extends IDtoList {

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
     * Build an instance of type AirlineOrderGroup from a JSON object
     * @param pathVars - The path variables
     * @returns A promise of AirlineOrderGroup[]
     */
    _buildInstance: (data: JSONData) => AirlineOrderGroup

    /**
     * Get all AirlineOrderGroups
     * @param pathVars - The path variables
     * @returns A promise of AirlineOrderGroup[]
     */
    getAll: (pathVars?: string | string[] | Record<string, string>) => Promise<AirlineOrderGroup[]>

    /**
     * Get paginated AirlineOrderGroups
     * @param pathVars - The path variables
     * @param pagination - The pagination object
     * @returns A promise of {
     *      items: AirlineOrderGroup[],
     *      pagination: Pagination,
     *  }
     */
    getPaginated: (pathVars: string | string[], pagination: Pagination) => Promise<{
        items: AirlineOrderGroup[],
        pagination: Pagination,
    }>

    /**
     * Query AirlineOrderGroups
     * @param queries - The filters to query on
     * @returns A promise of AirlineOrderGroup[]
     */
    query: (queries: Query[], pagination?: Pagination) => Promise<{
        items: AirlineOrderGroup[],
        pagination: Pagination,
    }>

    /**
     * Search AirlineOrderGroups
     * @param query - The search query
     * @returns A promise of AirlineOrderGroup[]
     */
    search: (query: string, pagination: Pagination) => Promise<{
        items: AirlineOrderGroup[],
        pagination: Pagination,
    }>
}