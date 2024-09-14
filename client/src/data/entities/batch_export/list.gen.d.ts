import type IDtoList from '../../base_classes/dto_list'
import { type JSONData } from '../../base_classes/dto.d';
import { type Pagination } from '../../base_classes/pagination';
import { type BatchExport } from '.';


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

export interface IBatchExportList extends IDtoList {

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
     * Build an instance of type BatchExport from a JSON object
     * @param pathVars - The path variables
     * @returns A promise of BatchExport[]
     */
    _buildInstance: (data: JSONData) => BatchExport

    /**
     * Get all BatchExports
     * @param pathVars - The path variables
     * @returns A promise of BatchExport[]
     */
    getAll: (pathVars?: string | string[] | Record<string, string>) => Promise<BatchExport[]>

    /**
     * Get paginated BatchExports
     * @param pathVars - The path variables
     * @param pagination - The pagination object
     * @returns A promise of {
     *      items: BatchExport[],
     *      pagination: Pagination,
     *  }
     */
    getPaginated: (pathVars: string | string[], pagination: Pagination) => Promise<{
        items: BatchExport[],
        pagination: Pagination,
    }>

    /**
     * Query BatchExports
     * @param queries - The filters to query on
     * @returns A promise of BatchExport[]
     */
    query: (queries: Query[], pagination?: Pagination) => Promise<{
        items: BatchExport[],
        pagination: Pagination,
    }>

    /**
     * Search BatchExports
     * @param query - The search query
     * @returns A promise of BatchExport[]
     */
    search: (query: string, pagination: Pagination) => Promise<{
        items: BatchExport[],
        pagination: Pagination,
    }>
}