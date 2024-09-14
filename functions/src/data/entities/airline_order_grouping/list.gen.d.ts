import type IDtoList from '../../base_classes/dto_list'
import { type JSONData } from '../../base_classes/dto.d';
import { type Pagination } from '../../base_classes/dto_list.d';
import { type AirlineOrderGrouping } from '.';


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

export interface IAirlineOrderGroupingList extends IDtoList {

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
     * Build an instance of type AirlineOrderGrouping from a JSON object
     * @param pathVars - The path variables
     * @returns A promise of AirlineOrderGrouping[]
     */
    _buildInstance: (data: JSONData) => AirlineOrderGrouping

    /**
     * Get all AirlineOrderGroupings
     * @param pathVars - The path variables
     * @returns A promise of AirlineOrderGrouping[]
     */
    getAll: (pathVars?: string | string[] | Record<string, string>) => Promise<AirlineOrderGrouping[]>

    /**
     * Get paginated AirlineOrderGroupings
     * @param pathVars - The path variables
     * @param pagination - The pagination object
     * @returns A promise of {
     *      items: AirlineOrderGrouping[],
     *      pagination: Pagination,
     *  }
     */
    getPaginated: (pathVars: string | string[], pagination: Pagination) => Promise<{
        items: AirlineOrderGrouping[],
        pagination: Pagination,
    }>

    /**
     * Search AirlineOrderGroupings
     * @param query - The search query
     * @returns A promise of AirlineOrderGrouping[]
     */
    search: (query: string, pagination?: Pagination) => Promise<AirlineOrderGrouping[]>
}