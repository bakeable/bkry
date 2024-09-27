import type IDtoList from '../../base_classes/dto_list'
import { type JSONData } from '../../base_classes/dto.d';
import { type Pagination } from '../../base_classes/dto_list.d';
import { type PackingSlip } from '.';


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

export interface IPackingSlipList extends IDtoList {

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
     * Build an instance of type PackingSlip from a JSON object
     * @param pathVars - The path variables
     * @returns A promise of PackingSlip[]
     */
    _buildInstance: (data: JSONData) => PackingSlip

    /**
     * Get all PackingSlips
     * @param pathVars - The path variables
     * @returns A promise of PackingSlip[]
     */
    getAll: (pathVars?: string | string[] | Record<string, string>) => Promise<PackingSlip[]>

    /**
     * Get paginated PackingSlips
     * @param pathVars - The path variables
     * @param pagination - The pagination object
     * @returns A promise of {
     *      items: PackingSlip[],
     *      pagination: Pagination,
     *  }
     */
    getPaginated: (pathVars: string | string[], pagination: Pagination) => Promise<{
        items: PackingSlip[],
        pagination: Pagination,
    }>

    /**
     * Search PackingSlips
     * @param query - The search query
     * @returns A promise of PackingSlip[]
     */
    search: (query: string, pagination?: Pagination) => Promise<PackingSlip[]>
}