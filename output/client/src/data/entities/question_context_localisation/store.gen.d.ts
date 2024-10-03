import { type QuestionContextLocalisation } from './interface.gen'
import { type Store } from 'pinia'
import { Pagination } from "../../general/pagination"
import { Query } from "../../general/query"

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// IF YOU WANT TO EDIT THIS FILE, REMOVE THE .gen. FROM THE FILENAME. THE FILE WILL THEN BE IGNORED.
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Represents the state of the questionContextLocalisation store.
 */
export interface QuestionContextLocalisationStoreState {
  /**
   * The currently active questionContextLocalisation ID
   */
  activeID: string | null

  /**
   * A cache of questionContextLocalisations.
   */
  cache: Record<string, {
    ids: string[]
    pagination: Pagination
    retrievalTimestamp: number
  }>

  /**
   * The list of questionContextLocalisations.
   */
  list: QuestionContextLocalisation[]

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
 * Represents an QuestionContextLocalisationStore that extends the Store interface.
 */
export interface QuestionContextLocalisationStore extends QuestionContextLocalisationStoreState, Store {
  /**
   * The currently active questionContextLocalisation
   */
  active: QuestionContextLocalisation | null

  /**
   * Adds an questionContextLocalisation to the local QuestionContextLocalisationStore.
   * @param questionContextLocalisation The questionContextLocalisation to be added.
   */
  add: (questionContextLocalisation: QuestionContextLocalisation) => void

  /**
   * Checks whether the list contains an instance with the provided ID.
   * @param id The id to be searched.
   * @returns {boolean} A boolean indicating whether the list contains the ID.
   */
  contains: (id: string) => boolean

  /**
   * Creates a new questionContextLocalisation in the external database and updates the store.
   * @param questionContextLocalisation The questionContextLocalisation to be created.
   * @returns {Promise<QuestionContextLocalisation>} A promise that resolves with the created questionContextLocalisation.
   */
  create: (questionContextId: string, questionContextLocalisation: QuestionContextLocalisation) => Promise<QuestionContextLocalisation>

  /**
   * Deletes an questionContextLocalisation from the QuestionContextLocalisationStore from the external database.
   * @param questionContextLocalisation The questionContextLocalisation to be deleted.
   * @returns {Promise<void>} A promise that resolves when the questionContextLocalisation is deleted.
   */
  delete: (questionContextLocalisation: QuestionContextLocalisation) => Promise<void>

  /**
   * Deletes an questionContextLocalisation by ID from the QuestionContextLocalisationStore from the external database.
   * @param id The ID to be deleted.
   * @returns {Promise<void>} A promise that resolves when the questionContextLocalisation is deleted.
   */
  deleteById: (id: string) => Promise<void>

  /**
   * Fire a form update event for the QuestionContextLocalisation.
   */
  emitFormUpdateEvent: (id: string) => void

  /**
   * Fire a store update event for all QuestionContextLocalisations.
   */
  emitStoreUpdateEvent: () => void

  /**
   * Get a QuestionContextLocalisation from the local QuestionContextLocalisationStore.
   * @returns { QuestionContextLocalisation | null } The questionContextLocalisation.
   */
  get: (id: string) => QuestionContextLocalisation | null

  /**
   * Queries questionContextLocalisations from the external source
   * @param queries The queries to be used for filtering.
   * @param pagination The pagination to be used for the query.
   * @param force Whether to force the query to the external source.
   * @returns {Promise<{ pagination: Pagination; items: QuestionContextLocalisation[] }>} A promise that resolves with queried questionContextLocalisations.
   */
  query: (questionContextId: string, queries: Query[], pagination: Pagination, force?: boolean) => Promise<{
    pagination: Pagination
    items: QuestionContextLocalisation[]
  }>

  /**
   * Reinserts an questionContextLocalisation to the QuestionContextLocalisationStore list.
   * @param questionContextLocalisation The questionContextLocalisation to be reinserted.
   * @returns {void}
   */
  reinsert: (questionContextLocalisation: QuestionContextLocalisation) => void

  /**
   * Removes an questionContextLocalisation from the local QuestionContextLocalisationStore.
   * @param questionContextLocalisation The questionContextLocalisation to be removed.
   * @returns {void}
   */
  remove: ({ id }: QuestionContextLocalisation) => void

  /**
   * Determines whether a refresh of the store is required
   */
  refreshRequired: boolean

  /**
   * Retrieves a single QuestionContextLocalisation from either the local QuestionContextLocalisationStore or external source, if forced or required.
   * @param id The ID of the QuestionContextLocalisation to retrieve.
   * @param force Indicates whether a refresh is forced.
   * @returns {Promise<QuestionContextLocalisation>} A promise that resolves when the questionContextLocalisation is set as active.
   */
  retrieve: (questionContextId: string, id: string, force?: boolean) => Promise<QuestionContextLocalisation>

  /**
   * Retrieves all questionContextLocalisations from the QuestionContextLocalisationStore or external source, depending if a refresh is required.
   * @param force Indicates whether a refresh is forced.
   * @returns {Promise<{ active: string | null; items: QuestionContextLocalisation[] }>} A promise that resolves with an array of questionContextLocalisations.
   */
  retrieveAll: (questionContextId: string, force?: boolean) => Promise<{
    active: string | null
    items: QuestionContextLocalisation[]
  }>

  /**
   * Retrieves questionContextLocalisations paginated from the external source.
   * @param pagination The pagination to be used for the query.
   * @param force Whether to force the query to the external source.
   * @returns {Promise<{ pagination: Pagination; items: QuestionContextLocalisation[] }>} A promise that resolves with paginated questionContextLocalisations.
   */
  retrievePaginated: (questionContextId: string, pagination?: Pagination, force?: boolean) => Promise<{
    pagination: Pagination
    items: QuestionContextLocalisation[]
  }>

  /**
   * Searches for questionContextLocalisations in the external source.
   * @param search The search query string.
   * @param pagination The pagination to be used for the search.
   * @param force Whether to force the search to the external source.
   * @returns {Promise<{ pagination: Pagination; items: QuestionContextLocalisation[] }>} A promise that resolves with search results.
   */
  search: (questionContextId: string, search: string, pagination: Pagination, force?: boolean) => Promise<{
    pagination: Pagination
    items: QuestionContextLocalisation[]
  }>

  /**
   * Update an questionContextLocalisation from the QuestionContextLocalisationStore in the external database.
   * @param questionContextLocalisation The questionContextLocalisation to be updated.
   * @returns {Promise<QuestionContextLocalisation>} A promise that resolves with the updated questionContextLocalisation.
   */
  update: (questionContextLocalisation: QuestionContextLocalisation) => Promise<QuestionContextLocalisation>
}
