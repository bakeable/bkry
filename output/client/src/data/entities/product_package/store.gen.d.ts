import { type ProductPackage } from './interface.gen'
import { type Store } from 'pinia'
import { Pagination } from "../../general/pagination"
import { Query } from "../../general/query"

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// IF YOU WANT TO EDIT THIS FILE, REMOVE THE .gen. FROM THE FILENAME. THE FILE WILL THEN BE IGNORED.
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Represents the state of the productPackage store.
 */
export interface ProductPackageStoreState {
  /**
   * The currently active productPackage ID
   */
  activeID: string | null

  /**
   * A cache of productPackages.
   */
  cache: Record<string, {
    ids: string[]
    pagination: Pagination
    retrievalTimestamp: number
  }>

  /**
   * The list of productPackages.
   */
  list: ProductPackage[]

  /**
   * The retrieval timestamp
   */
  retrievalTimestamp: number

  /**
   * A key that is updated to force the update of the entity.
   */
  updateKey: number

  /**
   * A timeout before emitting a store update event.
   */
  storeUpdateTimeout: window.Timeout | null
}

/**
 * Represents an ProductPackageStore that extends the Store interface.
 */
export interface ProductPackageStore extends ProductPackageStoreState, Store {
  /**
   * The currently active productPackage
   */
  active: ProductPackage | null

  /**
   * Adds an productPackage to the local ProductPackageStore.
   * @param productPackage The productPackage to be added.
   */
  add: (productPackage: ProductPackage) => void

  /**
   * Checks whether the list contains an instance with the provided ID.
   * @param id The id to be searched.
   * @returns {boolean} A boolean indicating whether the list contains the ID.
   */
  contains: (id: string) => boolean

  /**
   * Creates a new productPackage in the external database and updates the store.
   * @param productPackage The productPackage to be created.
   * @returns {Promise<ProductPackage>} A promise that resolves with the created productPackage.
   */
  create: (productPackage: ProductPackage) => Promise<ProductPackage>

  /**
   * Deletes an productPackage from the ProductPackageStore from the external database.
   * @param productPackage The productPackage to be deleted.
   * @returns {Promise<void>} A promise that resolves when the productPackage is deleted.
   */
  delete: (productPackage: ProductPackage) => Promise<void>

  /**
   * Deletes an productPackage by ID from the ProductPackageStore from the external database.
   * @param id The ID to be deleted.
   * @returns {Promise<void>} A promise that resolves when the productPackage is deleted.
   */
  deleteById: (id: string) => Promise<void>

  /**
   * Fire a form update event for the ProductPackage.
   */
  emitFormUpdateEvent: (id: string) => void

  /**
   * Fire a store update event for all ProductPackages.
   */
  emitStoreUpdateEvent: () => void

  /**
   * Get a ProductPackage from the local ProductPackageStore.
   * @returns { ProductPackage | null } The productPackage.
   */
  get: (id: string) => ProductPackage | null

  /**
   * Queries productPackages from the external source
   * @param queries The queries to be used for filtering.
   * @param pagination The pagination to be used for the query.
   * @param force Whether to force the query to the external source.
   * @returns {Promise<{ pagination: Pagination; items: ProductPackage[] }>} A promise that resolves with queried productPackages.
   */
  query: (queries: Query[], pagination: Pagination, force?: boolean) => Promise<{
    pagination: Pagination
    items: ProductPackage[]
  }>

  /**
   * Reinserts an productPackage to the ProductPackageStore list.
   * @param productPackage The productPackage to be reinserted.
   * @returns {void}
   */
  reinsert: (productPackage: ProductPackage) => void

  /**
   * Removes an productPackage from the local ProductPackageStore.
   * @param productPackage The productPackage to be removed.
   * @returns {void}
   */
  remove: ({ id }: ProductPackage) => void

  /**
   * Determines whether a refresh of the store is required
   */
  refreshRequired: boolean

  /**
   * Retrieves a single ProductPackage from either the local ProductPackageStore or external source, if forced or required.
   * @param id The ID of the ProductPackage to retrieve.
   * @param force Indicates whether a refresh is forced.
   * @returns {Promise<ProductPackage>} A promise that resolves when the productPackage is set as active.
   */
  retrieve: (id: string, force?: boolean) => Promise<ProductPackage>

  /**
   * Retrieves all productPackages from the ProductPackageStore or external source, depending if a refresh is required.
   * @param force Indicates whether a refresh is forced.
   * @returns {Promise<{ active: string | null; items: ProductPackage[] }>} A promise that resolves with an array of productPackages.
   */
  retrieveAll: (force?: boolean) => Promise<{
    active: string | null
    items: ProductPackage[]
  }>

  /**
   * Retrieves productPackages paginated from the external source.
   * @param pagination The pagination to be used for the query.
   * @param force Whether to force the query to the external source.
   * @returns {Promise<{ pagination: Pagination; items: ProductPackage[] }>} A promise that resolves with paginated productPackages.
   */
  retrievePaginated: (pagination?: Pagination, force?: boolean) => Promise<{
    pagination: Pagination
    items: ProductPackage[]
  }>

  /**
   * Searches for productPackages in the external source.
   * @param search The search query string.
   * @param pagination The pagination to be used for the search.
   * @param force Whether to force the search to the external source.
   * @returns {Promise<{ pagination: Pagination; items: ProductPackage[] }>} A promise that resolves with search results.
   */
  search: (search: string, pagination: Pagination, force?: boolean) => Promise<{
    pagination: Pagination
    items: ProductPackage[]
  }>

  /**
   * Update an productPackage from the ProductPackageStore in the external database.
   * @param productPackage The productPackage to be updated.
   * @returns {Promise<ProductPackage>} A promise that resolves with the updated productPackage.
   */
  update: (productPackage: ProductPackage) => Promise<ProductPackage>
}
