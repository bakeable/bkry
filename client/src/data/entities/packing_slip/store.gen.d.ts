import { type PackingSlip } from './entity'
import { type Store } from 'pinia'


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// IF YOU WANT TO EDIT THIS FILE, REMOVE THE .gen. FROM THE FILENAME. THE FILE WILL THEN BE IGNORED.
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Represents the state of the packingSlip store.
 */
export interface PackingSlipStoreState {
  /**
   * The currently active packingSlip ID
   */
  activeID: string | null

  /**
   * A cache of packingSlips.
   */
  cache: Record<string, {
    ids: string[]
    pagination: Pagination
    retrievalTimestamp: number
  }>

  /**
   * The list of packingSlips.
   */
  list: PackingSlip[]

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
 * Represents an PackingSlipStore that extends the Store interface.
 */
export interface PackingSlipStore extends PackingSlipStoreState, Store {
  /**
   * Fire a form update event for the PackingSlip.
   */
  emitFormUpdateEvent: (id: string) => void

  /**
   * Fire a store update event for all PackingSlip.
   */
  emitStoreUpdateEvent: () => void

  /**
   * The currently active packingSlip
   */
  active: PackingSlip

  /**
   * Adds an packingSlip to the local PackingSlipStore.
   * @param packingSlip The packingSlip to be added.
   */
  add: (packingSlip: PackingSlip) => void

  /**
   * Checks whether the list contains an instance with the provided ID.
   * @param id The id to be searched.
   * @returns {boolean} A boolean indicating whether the list contains the ID.
   */
  contains: (id: string) => boolean

  /**
   * Deletes an packingSlip from the PackingSlipStore from the external database.
   * @param packingSlip The packingSlip to be deleted.
   * @returns {Promise<void>} A promise that resolves when the packingSlip is deleted.
   */
  delete: (packingSlip: PackingSlip) => Promise<void>

  /**
   * Deletes an packingSlip by ID from the PackingSlipStore from the external database.
   * @param id The ID to be deleted.
   * @returns {Promise<void>} A promise that resolves when the packingSlip is deleted.
   */
  deleteById: (id: string) => Promise<void>

  /**
   * Get a PackingSlip from the local PackingSlipStore.
   * @returns { PackingSlip | null } The packingSlip.
   */
  get: (id: string) => PackingSlip | null

  /**
   * Reinserts an packingSlip to the PackingSlipStore list.
   * @param packingSlip The packingSlip to be reinserted.
   * @returns {void}
   */
  reinsert: (packingSlip: PackingSlip) => void

  /**
   * Removes an packingSlip from the local PackingSlipStore.
   * @param packingSlip The packingSlip to be removed.
   * @returns {void}
   */
  remove: (packingSlip: PackingSlip) => void

  /**
   * Determines whether a refresh of the store is required
   */
  refreshRequired: boolean

  /**
   * Retrieves a single PackingSlip from either the local PackingSlipStore or external source, if forced or required.
    * @param id The ID of the PackingSlip to retrieve.
    * @param force Indicates whether a refresh is forced.
   * @returns {Promise<void>} A promise that resolves when the packingSlip is set as active.
   */
  retrieve: (id: string, force?: boolean) => Promise<PackingSlip>

  /**
   * Retrieves all packingSlips from the PackingSlipStore or external source, depending if a refresh is required.
   * @param force Indicates whether a refresh is forced.
   * @returns {Promise<void>} A promise that resolves with an array of packingSlips.
   */
  retrieveAll: (force?: boolean) => Promise<{
    active: string | null
    items: PackingSlip[]
  }>

  /**
    * Retrieves packingSlips paginated from the external source
    * @param pagination The pagination to be used for the query.
    * @param force Whether to force the query to the external source.
    * @returns {Promise<void>} A promise that resolves with an array of packingSlips.
    */
  retrievePaginated: (pagination?: Pagination, force?: boolean) => Promise<{
    pagination: Pagination
    items: PackingSlip[]
  }>

  /**
    * Queries packingSlips from the external source
    * @param queries The queries to be used for filtering.
    * @param pagination The pagination to be used for the query.
    * @param force Whether to force the query to the external source.
    * @returns {Promise<void>} A promise that resolves with an array of packingSlips.
    */
  query: (queries: Query[], pagination: Pagination, force?: boolean) => Promise<{
    pagination: Pagination
    items: PackingSlip[]
  }>

  /**
   * Update an packingSlip from the PackingSlipStore in the external database.
   * @param packingSlip The packingSlip to be updated.
   * @param connect Indicates whether the packingSlip should be connected to the user on update.
   * @returns {Promise<string>} A promise that resolves to the ID as a string when the packingSlip is updated.
   */
  update: (packingSlip: PackingSlip) => Promise<string>

  /**
    * Searches for packingSlips in the external source.
    * @param query The query to be used for searching.
    * @returns {Promise<void>} A promise that resolves with an array of packingSlips.
    */
  search: (query: string, pagination: Pagination, force?: boolean) => Promise<PackingSlip[]>
}
