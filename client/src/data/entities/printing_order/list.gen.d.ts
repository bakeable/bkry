import type IDtoList from '../../base_classes/dto_list'
import { type JSONData } from '../../base_classes/dto.d';
import { type Pagination } from '../../base_classes/pagination';
import { type PrintingOrder } from '.';


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

export interface IPrintingOrderList extends IDtoList {

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
     * Build an instance of type PrintingOrder from a JSON object
     * @param pathVars - The path variables
     * @returns A promise of PrintingOrder[]
     */
    _buildInstance: (data: JSONData) => PrintingOrder

    /**
     * Get all PrintingOrders
     * @param pathVars - The path variables
     * @returns A promise of PrintingOrder[]
     */
    getAll: (pathVars?: string | string[] | Record<string, string>) => Promise<PrintingOrder[]>

    /**
     * Get paginated PrintingOrders
     * @param pathVars - The path variables
     * @param pagination - The pagination object
     * @returns A promise of {
     *      items: PrintingOrder[],
     *      pagination: Pagination,
     *  }
     */
    getPaginated: (pathVars: string | string[], pagination: Pagination) => Promise<{
        items: PrintingOrder[],
        pagination: Pagination,
    }>

    /**
     * Query PrintingOrders
     * @param queries - The filters to query on
     * @returns A promise of PrintingOrder[]
     */
    query: (queries: Query[], pagination?: Pagination) => Promise<{
        items: PrintingOrder[],
        pagination: Pagination,
    }>

    /**
     * Search PrintingOrders
     * @param query - The search query
     * @returns A promise of PrintingOrder[]
     */
    search: (query: string, pagination: Pagination) => Promise<{
        items: PrintingOrder[],
        pagination: Pagination,
    }>
}