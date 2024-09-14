import type IDtoList from '../../base_classes/dto_list'
import { type JSONData } from '../../base_classes/dto.d';
import { type Pagination } from '../../base_classes/pagination';
import { type AirlinePricing } from '.';


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

export interface IAirlinePricingList extends IDtoList {

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
     * Build an instance of type AirlinePricing from a JSON object
     * @param pathVars - The path variables
     * @returns A promise of AirlinePricing[]
     */
    _buildInstance: (data: JSONData) => AirlinePricing

    /**
     * Get all AirlinePricings
     * @param pathVars - The path variables
     * @returns A promise of AirlinePricing[]
     */
    getAll: (pathVars?: string | string[] | Record<string, string>) => Promise<AirlinePricing[]>

    /**
     * Get paginated AirlinePricings
     * @param pathVars - The path variables
     * @param pagination - The pagination object
     * @returns A promise of {
     *      items: AirlinePricing[],
     *      pagination: Pagination,
     *  }
     */
    getPaginated: (pathVars: string | string[], pagination: Pagination) => Promise<{
        items: AirlinePricing[],
        pagination: Pagination,
    }>

    /**
     * Query AirlinePricings
     * @param queries - The filters to query on
     * @returns A promise of AirlinePricing[]
     */
    query: (queries: Query[], pagination?: Pagination) => Promise<{
        items: AirlinePricing[],
        pagination: Pagination,
    }>

    /**
     * Search AirlinePricings
     * @param query - The search query
     * @returns A promise of AirlinePricing[]
     */
    search: (query: string, pagination: Pagination) => Promise<{
        items: AirlinePricing[],
        pagination: Pagination,
    }>
}