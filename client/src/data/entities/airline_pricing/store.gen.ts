import { defineStore } from 'pinia'
import { AirlinePricing } from './entity'
import { AirlinePricingList } from './list.gen'
import type { AirlinePricingStore, AirlinePricingStoreState } from './store.gen.d'
import { analytics, auth } from '@/firebase'
import { type JSONData } from '../../base_classes/dto.d'
import { Pagination } from "../../base_classes/pagination";
import { Query } from "../../base_classes/query";
import { paginationAndQueriesToCacheKey } from "../../base_classes/dto_list";


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// IF YOU WANT TO EDIT THIS FILE, REMOVE THE .gen. FROM THE FILENAME. THE FILE WILL THEN BE IGNORED.
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

const airlinePricingList = new AirlinePricingList()
export const useAirlinePricingStore = defineStore('airlinePricingStore', {
  state: (): AirlinePricingStoreState => ({
    activeID: null,
    cache: {},
    list: [] as AirlinePricing[],
    retrievalTimestamp: 0,
    updateKey: 0,
    storeUpdateTimeout: null,
  }),
  actions: {
    initializeStore(): void {
      this.convertList()
    },
    emitFormUpdateEvent(id: string): void {
      window.dispatchEvent(new CustomEvent("form-update-AirlinePricing", { detail: { id } }));
    },
    emitStoreUpdateEvent(): void {
      // Clear existing timeout
      if (this.storeUpdateTimeout !== null) {
        clearTimeout(this.storeUpdateTimeout)
      }

      // Set new timeout
      this.storeUpdateTimeout = setTimeout(() => {
        window.dispatchEvent(new CustomEvent("store.update.AirlinePricing"));
        this.storeUpdateTimeout = null
      }, 1500)
    },
    async activate(airlinePricing: AirlinePricing): Promise<void> {
      // Set active
      await airlinePricing
        .activate(auth.firebaseUser?.uid)
        .then(() => {
          this.activeID = airlinePricing.id
        })
        .catch((error) => {
          console.error(error)
          this.activeID = null
        })
    },
    add(airlinePricing: AirlinePricing): void {
      const id = airlinePricing.id
      if (typeof id === 'string' && !this.contains(id)) {
        this.list.push(airlinePricing)

        // Set active
        if (this.active === null) {
          this.activeID = airlinePricing.id
        }
      } else {
        this.reinsert(airlinePricing)
      }
      this.updateKey += 1;
    },
    convertList(): void {
      this.list = (this.list as AirlinePricing[]).map((x: any) => {
        const airlinePricing = new AirlinePricing({})
        airlinePricing.set(x as JSONData)
        return airlinePricing
      })
    },
    async delete(airlinePricing: AirlinePricing): Promise<void> {
      await airlinePricing.delete().then(() => {
        this.remove(airlinePricing)
      }).finally(() => {
        this.emitStoreUpdateEvent()
      });
    },
    async deleteById(id: string): Promise<void> {
      const airlinePricing = (this.list as AirlinePricing[]).find((x: AirlinePricing) => x.id === id)
      if (airlinePricing !== undefined) {
        await this.delete(airlinePricing)
      } else {
        throw new Error('Could not find AirlinePricing with id: ' + id)
      }
    },
    contains(id: string): boolean {
      return (this.list as AirlinePricing[]).some((x: AirlinePricing) => x.id === id)
    },
    get(id: string): AirlinePricing | null {
      return (this.list as AirlinePricing[]).find((x: AirlinePricing) => x.id === id) || null
    },
    reinsert(airlinePricing: AirlinePricing): void {
      const index = (this.list as AirlinePricing[]).findIndex((x: AirlinePricing) => x.id === airlinePricing.id)
      if (index !== -1) {
        this.list = [
          ...this.list.slice(0, index),
          airlinePricing,
          ...this.list.slice(index + 1),
        ]
      } else {
        this.add(airlinePricing)
      }
      this.updateKey += 1;
    },
    remove({ id }: AirlinePricing): void {
      const index = (this.list as AirlinePricing[]).findIndex((x: AirlinePricing) => x.id === id)
      if (index !== -1) {
        this.list = [
          ...this.list.slice(0, index),
          ...this.list.slice(index + 1),
        ]
      }
      this.updateKey += 1;
    },
    async retrieve(id: string, force = false): Promise<AirlinePricing> {
      // Check if a refresh is required
      if (!force && !this.refreshRequired) {
        const airlinePricing = this.get(id)
        if (airlinePricing !== null) {
          return airlinePricing
        }
      }

      // Retrieve airlinePricing
      const airlinePricing = new AirlinePricing({})

      // Get airlinePricing
      await airlinePricing.get([id]).then(async (entity) => {
        this.reinsert(entity)

        // Activate airlinePricing
        await this.activate(entity)
      })

      return airlinePricing
    },
    async retrieveAll(force = false): Promise<{
      active: string | null
      items: AirlinePricing[]
    }> {
      // Check if a refresh is required
      if (!force && !this.refreshRequired) {
        return {
          active: this.activeID,
          items: this.list as AirlinePricing[],
        }
      }

      // Clear list
      while (this.list.length > 0) {
        this.list.pop()
      }

      // Retrieve list
      await airlinePricingList.getAll().then(async (list) => {
        list.forEach((airlinePricing) => {
          this.reinsert(airlinePricing)
        })

        // Activate airlinePricing
        if (list.length > 0) {
          if (airlinePricingList.active !== null) {
            this.activeID = airlinePricingList.active
          } else {
            await this.activate(list[0])
          }
        }
      })

      // Set retrieval timestamp
      this.retrievalTimestamp = Date.now()

      // Set user property
      if (this.activeID !== null) {
        analytics.setUserProperty('airline_pricing_id', this.activeID)
      }

      return {
        active: this.activeID,
        items: this.list as AirlinePricing[],
      }
    },
    async retrievePaginated(pagination: Pagination, force = false): Promise<{
      pagination: Pagination;
      items: AirlinePricing[];
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
          items: cached.ids.map((id: string) => this.get(id) as AirlinePricing),
        };
      }

      // Retrieve list
      const items = await airlinePricingList.getPaginated(
        [],
        pagination
      ).then(async ({items, pagination }) => {
        items.forEach((airlinePricing) => {
          this.reinsert(airlinePricing);
        });

        // Set new pagination
        newPagination = pagination.deepCopy();

        return items;
      });

      // Set cache
      this.cache[cacheKey] = {
        ids: items.map((x: AirlinePricing) => x.id),
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
      items: AirlinePricing[];
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
          items: cached.ids.map((id: string) => this.get(id) as AirlinePricing),
        };
      }

      // Retrieve list
      const items = await airlinePricingList.query(
        queries,
        pagination,
      ).then(async ({ items, pagination }) => {
        items.forEach((airlinePricing) => {
          this.reinsert(airlinePricing);
        });

        // Set new pagination
        newPagination = pagination.deepCopy();;

        return items
      });

      // Set cache
      this.cache[cacheKey] = {
        ids: items.map((x: AirlinePricing) => x.id),
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
      items: AirlinePricing[];
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
            (id: string) => this.get(id) as AirlinePricing,
          ),
          pagination: cached.pagination.deepCopy(),
        };
      }

      // Retrieve list
      const items = await airlinePricingList
        .search(query, pagination)
        .then(async ({ items, pagination }) => {
          items.forEach((airlinePricing) => {
            this.reinsert(airlinePricing);
          });

          // Set new pagination
          newPagination = pagination.deepCopy();

          return items;
        });

      // Set cache
      this.cache[cacheKey] = {
        ids: items.map((x: AirlinePricing) => x.id),
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
    async update(airlinePricing: AirlinePricing): Promise<string> {
      // Update airlinePricing
      await airlinePricing.update()

      // Reinsert airlinePricing
      this.reinsert(airlinePricing)

      // Activate airlinePricing
      if (this.active === null) {
        await this.activate(airlinePricing)
      }
      
      this.emitStoreUpdateEvent()

      return airlinePricing.id
    },
  },
  getters: {
    active(): AirlinePricing | null {
      if (this.activeID === null) {
        return null
      }

      const airlinePricing = (this.list as AirlinePricing[]).find((x: AirlinePricing) => x.id === this.activeID)
      if (airlinePricing === undefined) {
        return null
      }

      return airlinePricing
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
export function createAirlinePricingStore(): AirlinePricingStore {
  const store = useAirlinePricingStore()
  // Run the initializeStore action when the store is created
  store.initializeStore()
  return store as unknown as AirlinePricingStore // TODO: Fix this
}
