import { defineStore } from 'pinia'
import { Product } from './entity'
import { ProductList } from './list.gen'
import type { ProductStore, ProductStoreState } from './store.gen.d'
import { analytics, auth } from '@/firebase'
import { type JSONData } from '../../base_classes/dto.d'
import { Pagination } from "../../base_classes/pagination";
import { Query } from "../../base_classes/query";
import { paginationAndQueriesToCacheKey } from "../../base_classes/dto_list";


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// IF YOU WANT TO EDIT THIS FILE, REMOVE THE .gen. FROM THE FILENAME. THE FILE WILL THEN BE IGNORED.
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

const productList = new ProductList()
export const useProductStore = defineStore('productStore', {
  state: (): ProductStoreState => ({
    activeID: null,
    cache: {},
    list: [] as Product[],
    retrievalTimestamp: 0,
    updateKey: 0,
    storeUpdateTimeout: null,
  }),
  actions: {
    initializeStore(): void {
      this.convertList()
    },
    emitFormUpdateEvent(id: string): void {
      window.dispatchEvent(new CustomEvent("form-update-Product", { detail: { id } }));
    },
    emitStoreUpdateEvent(): void {
      // Clear existing timeout
      if (this.storeUpdateTimeout !== null) {
        clearTimeout(this.storeUpdateTimeout)
      }

      // Set new timeout
      this.storeUpdateTimeout = setTimeout(() => {
        window.dispatchEvent(new CustomEvent("store.update.Product"));
        this.storeUpdateTimeout = null
      }, 1500)
    },
    async activate(product: Product): Promise<void> {
      // Set active
      await product
        .activate(auth.firebaseUser?.uid)
        .then(() => {
          this.activeID = product.id
        })
        .catch((error) => {
          console.error(error)
          this.activeID = null
        })
    },
    add(product: Product): void {
      const id = product.id
      if (typeof id === 'string' && !this.contains(id)) {
        this.list.push(product)

        // Set active
        if (this.active === null) {
          this.activeID = product.id
        }
      } else {
        this.reinsert(product)
      }
      this.updateKey += 1;
    },
    convertList(): void {
      this.list = (this.list as Product[]).map((x: any) => {
        const product = new Product({})
        product.set(x as JSONData)
        return product
      })
    },
    async delete(product: Product): Promise<void> {
      await product.delete().then(() => {
        this.remove(product)
      }).finally(() => {
        this.emitStoreUpdateEvent()
      });
    },
    async deleteById(id: string): Promise<void> {
      const product = (this.list as Product[]).find((x: Product) => x.id === id)
      if (product !== undefined) {
        await this.delete(product)
      } else {
        throw new Error('Could not find Product with id: ' + id)
      }
    },
    contains(id: string): boolean {
      return (this.list as Product[]).some((x: Product) => x.id === id)
    },
    get(id: string): Product | null {
      return (this.list as Product[]).find((x: Product) => x.id === id) || null
    },
    reinsert(product: Product): void {
      const index = (this.list as Product[]).findIndex((x: Product) => x.id === product.id)
      if (index !== -1) {
        this.list = [
          ...this.list.slice(0, index),
          product,
          ...this.list.slice(index + 1),
        ]
      } else {
        this.add(product)
      }
      this.updateKey += 1;
    },
    remove({ id }: Product): void {
      const index = (this.list as Product[]).findIndex((x: Product) => x.id === id)
      if (index !== -1) {
        this.list = [
          ...this.list.slice(0, index),
          ...this.list.slice(index + 1),
        ]
      }
      this.updateKey += 1;
    },
    async retrieve(id: string, force = false): Promise<Product> {
      // Check if a refresh is required
      if (!force && !this.refreshRequired) {
        const product = this.get(id)
        if (product !== null) {
          return product
        }
      }

      // Retrieve product
      const product = new Product({})

      // Get product
      await product.get([id]).then(async (entity) => {
        this.reinsert(entity)

        // Activate product
        await this.activate(entity)
      })

      return product
    },
    async retrieveAll(force = false): Promise<{
      active: string | null
      items: Product[]
    }> {
      // Check if a refresh is required
      if (!force && !this.refreshRequired) {
        return {
          active: this.activeID,
          items: this.list as Product[],
        }
      }

      // Clear list
      while (this.list.length > 0) {
        this.list.pop()
      }

      // Retrieve list
      await productList.getAll().then(async (list) => {
        list.forEach((product) => {
          this.reinsert(product)
        })

        // Activate product
        if (list.length > 0) {
          if (productList.active !== null) {
            this.activeID = productList.active
          } else {
            await this.activate(list[0])
          }
        }
      })

      // Set retrieval timestamp
      this.retrievalTimestamp = Date.now()

      // Set user property
      if (this.activeID !== null) {
        analytics.setUserProperty('product_id', this.activeID)
      }

      return {
        active: this.activeID,
        items: this.list as Product[],
      }
    },
    async retrievePaginated(pagination: Pagination, force = false): Promise<{
      pagination: Pagination;
      items: Product[];
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
          items: cached.ids.map((id: string) => this.get(id) as Product),
        };
      }

      // Retrieve list
      const items = await productList.getPaginated(
        [],
        pagination
      ).then(async ({items, pagination }) => {
        items.forEach((product) => {
          this.reinsert(product);
        });

        // Set new pagination
        newPagination = pagination.deepCopy();

        return items;
      });

      // Set cache
      this.cache[cacheKey] = {
        ids: items.map((x: Product) => x.id),
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
      items: Product[];
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
          items: cached.ids.map((id: string) => this.get(id) as Product),
        };
      }

      // Retrieve list
      const items = await productList.query(
        queries,
        pagination,
      ).then(async ({ items, pagination }) => {
        items.forEach((product) => {
          this.reinsert(product);
        });

        // Set new pagination
        newPagination = pagination.deepCopy();;

        return items
      });

      // Set cache
      this.cache[cacheKey] = {
        ids: items.map((x: Product) => x.id),
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
      items: Product[];
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
            (id: string) => this.get(id) as Product,
          ),
          pagination: cached.pagination.deepCopy(),
        };
      }

      // Retrieve list
      const items = await productList
        .search(query, pagination)
        .then(async ({ items, pagination }) => {
          items.forEach((product) => {
            this.reinsert(product);
          });

          // Set new pagination
          newPagination = pagination.deepCopy();

          return items;
        });

      // Set cache
      this.cache[cacheKey] = {
        ids: items.map((x: Product) => x.id),
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
    async update(product: Product): Promise<string> {
      // Update product
      await product.update()

      // Reinsert product
      this.reinsert(product)

      // Activate product
      if (this.active === null) {
        await this.activate(product)
      }
      
      this.emitStoreUpdateEvent()

      return product.id
    },
  },
  getters: {
    active(): Product | null {
      if (this.activeID === null) {
        return null
      }

      const product = (this.list as Product[]).find((x: Product) => x.id === this.activeID)
      if (product === undefined) {
        return null
      }

      return product
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
export function createProductStore(): ProductStore {
  const store = useProductStore()
  // Run the initializeStore action when the store is created
  store.initializeStore()
  return store as unknown as ProductStore // TODO: Fix this
}
