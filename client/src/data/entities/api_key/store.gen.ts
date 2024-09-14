import { defineStore } from 'pinia'
import { ApiKey } from './entity'
import { ApiKeyList } from './list.gen'
import type { ApiKeyStore, ApiKeyStoreState } from './store.gen.d'
import { analytics, auth } from '@/firebase'
import { type JSONData } from '../../base_classes/dto.d'
import { Pagination } from "../../base_classes/pagination";
import { Query } from "../../base_classes/query";
import { paginationAndQueriesToCacheKey } from "../../base_classes/dto_list";


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// IF YOU WANT TO EDIT THIS FILE, REMOVE THE .gen. FROM THE FILENAME. THE FILE WILL THEN BE IGNORED.
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

const apiKeyList = new ApiKeyList()
export const useApiKeyStore = defineStore('apiKeyStore', {
  state: (): ApiKeyStoreState => ({
    activeID: null,
    cache: {},
    list: [] as ApiKey[],
    retrievalTimestamp: 0,
    updateKey: 0,
    storeUpdateTimeout: null,
  }),
  actions: {
    initializeStore(): void {
      this.convertList()
    },
    emitFormUpdateEvent(id: string): void {
      window.dispatchEvent(new CustomEvent("form-update-ApiKey", { detail: { id } }));
    },
    emitStoreUpdateEvent(): void {
      // Clear existing timeout
      if (this.storeUpdateTimeout !== null) {
        clearTimeout(this.storeUpdateTimeout)
      }

      // Set new timeout
      this.storeUpdateTimeout = setTimeout(() => {
        window.dispatchEvent(new CustomEvent("store.update.ApiKey"));
        this.storeUpdateTimeout = null
      }, 1500)
    },
    async activate(apiKey: ApiKey): Promise<void> {
      // Set active
      await apiKey
        .activate(auth.firebaseUser?.uid)
        .then(() => {
          this.activeID = apiKey.id
        })
        .catch((error) => {
          console.error(error)
          this.activeID = null
        })
    },
    add(apiKey: ApiKey): void {
      const id = apiKey.id
      if (typeof id === 'string' && !this.contains(id)) {
        this.list.push(apiKey)

        // Set active
        if (this.active === null) {
          this.activeID = apiKey.id
        }
      } else {
        this.reinsert(apiKey)
      }
      this.updateKey += 1;
    },
    convertList(): void {
      this.list = (this.list as ApiKey[]).map((x: any) => {
        const apiKey = new ApiKey({})
        apiKey.set(x as JSONData)
        return apiKey
      })
    },
    async delete(apiKey: ApiKey): Promise<void> {
      await apiKey.delete().then(() => {
        this.remove(apiKey)
      }).finally(() => {
        this.emitStoreUpdateEvent()
      });
    },
    async deleteById(id: string): Promise<void> {
      const apiKey = (this.list as ApiKey[]).find((x: ApiKey) => x.id === id)
      if (apiKey !== undefined) {
        await this.delete(apiKey)
      } else {
        throw new Error('Could not find ApiKey with id: ' + id)
      }
    },
    contains(id: string): boolean {
      return (this.list as ApiKey[]).some((x: ApiKey) => x.id === id)
    },
    get(id: string): ApiKey | null {
      return (this.list as ApiKey[]).find((x: ApiKey) => x.id === id) || null
    },
    reinsert(apiKey: ApiKey): void {
      const index = (this.list as ApiKey[]).findIndex((x: ApiKey) => x.id === apiKey.id)
      if (index !== -1) {
        this.list = [
          ...this.list.slice(0, index),
          apiKey,
          ...this.list.slice(index + 1),
        ]
      } else {
        this.add(apiKey)
      }
      this.updateKey += 1;
    },
    remove({ id }: ApiKey): void {
      const index = (this.list as ApiKey[]).findIndex((x: ApiKey) => x.id === id)
      if (index !== -1) {
        this.list = [
          ...this.list.slice(0, index),
          ...this.list.slice(index + 1),
        ]
      }
      this.updateKey += 1;
    },
    async retrieve(id: string, force = false): Promise<ApiKey> {
      // Check if a refresh is required
      if (!force && !this.refreshRequired) {
        const apiKey = this.get(id)
        if (apiKey !== null) {
          return apiKey
        }
      }

      // Retrieve apiKey
      const apiKey = new ApiKey({})

      // Get apiKey
      await apiKey.get([id]).then(async (entity) => {
        this.reinsert(entity)

        // Activate apiKey
        await this.activate(entity)
      })

      return apiKey
    },
    async retrieveAll(force = false): Promise<{
      active: string | null
      items: ApiKey[]
    }> {
      // Check if a refresh is required
      if (!force && !this.refreshRequired) {
        return {
          active: this.activeID,
          items: this.list as ApiKey[],
        }
      }

      // Clear list
      while (this.list.length > 0) {
        this.list.pop()
      }

      // Retrieve list
      await apiKeyList.getAll().then(async (list) => {
        list.forEach((apiKey) => {
          this.reinsert(apiKey)
        })

        // Activate apiKey
        if (list.length > 0) {
          if (apiKeyList.active !== null) {
            this.activeID = apiKeyList.active
          } else {
            await this.activate(list[0])
          }
        }
      })

      // Set retrieval timestamp
      this.retrievalTimestamp = Date.now()

      // Set user property
      if (this.activeID !== null) {
        analytics.setUserProperty('api_key_id', this.activeID)
      }

      return {
        active: this.activeID,
        items: this.list as ApiKey[],
      }
    },
    async retrievePaginated(pagination: Pagination, force = false): Promise<{
      pagination: Pagination;
      items: ApiKey[];
    }> {
      // Check pagination
      let newPagination = new Pagination(pagination).deepCopy();

      // Get cachekey
      const cacheKey = pagination.toCacheKey();

      // Check if a refresh is required
      if (!force && this.cache[cacheKey]) {
        const cached = this.cache[cacheKey]
        return {
          pagination: cached.pagination.deepCopy(),
          items: cached.ids.map((id: string) => this.get(id) as ApiKey),
        };
      }

      // Retrieve list
      const items = await apiKeyList.getPaginated(
        [],
        pagination
      ).then(async ({items, pagination }) => {
        items.forEach((apiKey) => {
          this.reinsert(apiKey);
        });

        // Set new pagination
        newPagination = pagination.deepCopy();

        return items;
      });

      // Set cache
      this.cache[cacheKey] = {
        ids: items.map((x: ApiKey) => x.id),
        pagination: newPagination,
        retrievalTimestamp: Date.now(),
      };

      // Set retrieval timestamp
      this.retrievalTimestamp = Date.now();

      return {
        pagination: newPagination,
        items,
      };
    },
    async query(queries: Query[], pagination: Pagination, force = false): Promise<{
      pagination: Pagination;
      items: ApiKey[];
    }> {
      // Check pagination
      let newPagination = new Pagination(pagination).deepCopy();

      // Get cachekey
      const cacheKey = paginationAndQueriesToCacheKey(pagination, queries);

      // Check if a refresh is required
      if (!force && this.cache[cacheKey]) {
        const cached = this.cache[cacheKey]
        return {
          pagination: cached.pagination.deepCopy(),
          items: cached.ids.map((id: string) => this.get(id) as ApiKey),
        };
      }

      // Retrieve list
      const items = await apiKeyList.query(
        queries,
        pagination,
      ).then(async ({ items, pagination }) => {
        items.forEach((apiKey) => {
          this.reinsert(apiKey);
        });

        // Set new pagination
        newPagination = pagination.deepCopy();;

        return items
      });

      // Set cache
      this.cache[cacheKey] = {
        ids: items.map((x: ApiKey) => x.id),
        pagination: newPagination.deepCopy(),
        retrievalTimestamp: Date.now(),
      };

      // Set retrieval timestamp
      this.retrievalTimestamp = Date.now();

      return {
        pagination: newPagination,
        items: items,
      };
    },
    async search(
      
      query: string,
      pagination: Pagination,
      force = false,
    ): Promise<{
      pagination: Pagination;
      items: ApiKey[];
    }> {
      // Check pagination
      let newPagination = new Pagination(pagination).deepCopy();

      // Get cachekey
      const cacheKey = query.replaceAll(" ", "_") + '-' + pagination.toCacheKey();

      // Check if a refresh is required
      if (!force && this.cache[cacheKey]) {
        const cached = this.cache[cacheKey]
        return {
          items: cached.ids.map(
            (id: string) => this.get(id) as ApiKey,
          ),
          pagination: cached.pagination.deepCopy(),
        };
      }

      // Retrieve list
      const items = await apiKeyList
        .search(query, pagination)
        .then(async ({ items, pagination }) => {
          items.forEach((apiKey) => {
            this.reinsert(apiKey);
          });

          // Set new pagination
          newPagination = pagination.deepCopy();

          return items;
        });

      // Set cache
      this.cache[cacheKey] = {
        ids: items.map((x: ApiKey) => x.id),
        pagination: newPagination.deepCopy(),
        retrievalTimestamp: Date.now(),
      };

      // Set retrieval timestamp
      this.retrievalTimestamp = Date.now();

      return {
        items: items,
        pagination: newPagination,
      };
    },
    async update(apiKey: ApiKey): Promise<string> {
      // Update apiKey
      await apiKey.update()

      // Reinsert apiKey
      this.reinsert(apiKey)

      // Activate apiKey
      if (this.active === null) {
        await this.activate(apiKey)
      }
      
      this.emitStoreUpdateEvent()

      return apiKey.id
    },
  },
  getters: {
    active(): ApiKey | null {
      if (this.activeID === null) {
        return null
      }

      const apiKey = (this.list as ApiKey[]).find((x: ApiKey) => x.id === this.activeID)
      if (apiKey === undefined) {
        return null
      }

      return apiKey
    },
    refreshRequired(): boolean {
      return (
        this.retrievalTimestamp === 0 ||
        this.retrievalTimestamp < Date.now() - 30 * 60000 ||
        this.list.length === 0
      )
    },
  },
})

// Export a function to create a new instance of the store
export function createApiKeyStore(): ApiKeyStore {
  const store = useApiKeyStore()
  // Run the initializeStore action when the store is created
  store.initializeStore()
  return store as unknown as ApiKeyStore // TODO: Fix this
}
