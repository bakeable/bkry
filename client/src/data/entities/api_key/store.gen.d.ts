import { type ApiKey } from './entity'
import { type Store } from 'pinia'


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// IF YOU WANT TO EDIT THIS FILE, REMOVE THE .gen. FROM THE FILENAME. THE FILE WILL THEN BE IGNORED.
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Represents the state of the apiKey store.
 */
export interface ApiKeyStoreState {
  /**
   * The currently active apiKey ID
   */
  activeID: string | null

  /**
   * A cache of apiKeys.
   */
  cache: Record<string, {
    ids: string[]
    pagination: Pagination
    retrievalTimestamp: number
  }>

  /**
   * The list of apiKeys.
   */
  list: ApiKey[]

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
 * Represents an ApiKeyStore that extends the Store interface.
 */
export interface ApiKeyStore extends ApiKeyStoreState, Store {
  /**
   * Fire a form update event for the ApiKey.
   */
  emitFormUpdateEvent: (id: string) => void

  /**
   * Fire a store update event for all ApiKey.
   */
  emitStoreUpdateEvent: () => void

  /**
   * The currently active apiKey
   */
  active: ApiKey

  /**
   * Adds an apiKey to the local ApiKeyStore.
   * @param apiKey The apiKey to be added.
   */
  add: (apiKey: ApiKey) => void

  /**
   * Checks whether the list contains an instance with the provided ID.
   * @param id The id to be searched.
   * @returns {boolean} A boolean indicating whether the list contains the ID.
   */
  contains: (id: string) => boolean

  /**
   * Deletes an apiKey from the ApiKeyStore from the external database.
   * @param apiKey The apiKey to be deleted.
   * @returns {Promise<void>} A promise that resolves when the apiKey is deleted.
   */
  delete: (apiKey: ApiKey) => Promise<void>

  /**
   * Deletes an apiKey by ID from the ApiKeyStore from the external database.
   * @param id The ID to be deleted.
   * @returns {Promise<void>} A promise that resolves when the apiKey is deleted.
   */
  deleteById: (id: string) => Promise<void>

  /**
   * Get a ApiKey from the local ApiKeyStore.
   * @returns { ApiKey | null } The apiKey.
   */
  get: (id: string) => ApiKey | null

  /**
   * Reinserts an apiKey to the ApiKeyStore list.
   * @param apiKey The apiKey to be reinserted.
   * @returns {void}
   */
  reinsert: (apiKey: ApiKey) => void

  /**
   * Removes an apiKey from the local ApiKeyStore.
   * @param apiKey The apiKey to be removed.
   * @returns {void}
   */
  remove: (apiKey: ApiKey) => void

  /**
   * Determines whether a refresh of the store is required
   */
  refreshRequired: boolean

  /**
   * Retrieves a single ApiKey from either the local ApiKeyStore or external source, if forced or required.
    * @param id The ID of the ApiKey to retrieve.
    * @param force Indicates whether a refresh is forced.
   * @returns {Promise<void>} A promise that resolves when the apiKey is set as active.
   */
  retrieve: (id: string, force?: boolean) => Promise<ApiKey>

  /**
   * Retrieves all apiKeys from the ApiKeyStore or external source, depending if a refresh is required.
   * @param force Indicates whether a refresh is forced.
   * @returns {Promise<void>} A promise that resolves with an array of apiKeys.
   */
  retrieveAll: (force?: boolean) => Promise<{
    active: string | null
    items: ApiKey[]
  }>

  /**
    * Retrieves apiKeys paginated from the external source
    * @param pagination The pagination to be used for the query.
    * @param force Whether to force the query to the external source.
    * @returns {Promise<void>} A promise that resolves with an array of apiKeys.
    */
  retrievePaginated: (pagination?: Pagination, force?: boolean) => Promise<{
    pagination: Pagination
    items: ApiKey[]
  }>

  /**
    * Queries apiKeys from the external source
    * @param queries The queries to be used for filtering.
    * @param pagination The pagination to be used for the query.
    * @param force Whether to force the query to the external source.
    * @returns {Promise<void>} A promise that resolves with an array of apiKeys.
    */
  query: (queries: Query[], pagination: Pagination, force?: boolean) => Promise<{
    pagination: Pagination
    items: ApiKey[]
  }>

  /**
   * Update an apiKey from the ApiKeyStore in the external database.
   * @param apiKey The apiKey to be updated.
   * @param connect Indicates whether the apiKey should be connected to the user on update.
   * @returns {Promise<string>} A promise that resolves to the ID as a string when the apiKey is updated.
   */
  update: (apiKey: ApiKey) => Promise<string>

  /**
    * Searches for apiKeys in the external source.
    * @param query The query to be used for searching.
    * @returns {Promise<void>} A promise that resolves with an array of apiKeys.
    */
  search: (query: string, pagination: Pagination, force?: boolean) => Promise<ApiKey[]>
}
