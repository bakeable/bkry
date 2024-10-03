import { defineStore } from 'pinia'
import { Category } from './interface.gen'
import type { CategoryStore, CategoryStoreState } from './store.gen.d'
import { Pagination, paginationAndQueriesToCacheKey } from "../../general/pagination";
import { Query } from "../../general/query";
import { CategoryTransporter } from './transport.gen'

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// IF YOU WANT TO EDIT THIS FILE, REMOVE THE .gen. FROM THE FILENAME. THE FILE WILL THEN BE IGNORED.
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

export const useCategoryStore = defineStore('categoryStore', {
  state: (): CategoryStoreState => ({
    activeID: null,
    cache: {},
    list: [] as Category[],
    retrievalTimestamp: 0,
    updateKey: 0,
    storeUpdateTimeout: null,
  }),
  actions: {
    emitFormUpdateEvent(id: string): void {
      window.dispatchEvent(new CustomEvent("form.update.category", { detail: { id } }));
    },
    emitStoreUpdateEvent(): void {
      // Clear existing timeout
      if (this.storeUpdateTimeout !== null) {
        clearTimeout(this.storeUpdateTimeout)
      }

      // Set new timeout
      this.storeUpdateTimeout = setTimeout(() => {
        window.dispatchEvent(new CustomEvent("store.update.category"));
        this.storeUpdateTimeout = null
      }, 1500)
    },
    async activate(category: Category): Promise<void> {
      // Set active
      this.activeID = category.id
    },
    add(category: Category): void {
      const id = category.id
      if (typeof id === 'string' && !this.contains(id)) {
        this.list.push(category)

        // Set active
        if (this.active === null) {
          this.activeID = category.id
        }
      } else {
        this.reinsert(category)
      }
      this.updateKey += 1;
    },
    async delete(category: Category): Promise<void> {
      await CategoryTransporter.Delete(category).then(() => {
        this.remove(category)
      }).finally(() => {
        this.emitStoreUpdateEvent()
      });
    },
    async deleteById(id: string): Promise<void> {
      const category = (this.list as Category[]).find((x: Category) => x.id === id)
      if (category !== undefined) {
        await this.delete(category)
      } else {
        throw new Error('Could not find Category with id: ' + id)
      }
    },
    contains(id: string): boolean {
      return (this.list as Category[]).some((x: Category) => x.id === id)
    },
    async create(category: Category): Promise<Category> {
      // Create category
      const item = await CategoryTransporter.Create({
      }, category)

      // Reinsert category
      this.reinsert(item)

      // Activate category
      if (this.active === null) {
        await this.activate(item)
      }
      
      this.emitStoreUpdateEvent()

      return item
    },
    get(id: string): Category | null {
      return (this.list as Category[]).find((x: Category) => x.id === id) || null
    },
    reinsert(category: Category): void {
      const index = (this.list as Category[]).findIndex((x: Category) => x.id === category.id)
      if (index !== -1) {
        this.list = [
          ...this.list.slice(0, index),
          category,
          ...this.list.slice(index + 1),
        ]
      } else {
        this.add(category)
      }
      this.updateKey += 1;
    },
    remove({ id }: Category): void {
      const index = (this.list as Category[]).findIndex((x: Category) => x.id === id)
      if (index !== -1) {
        this.list = [
          ...this.list.slice(0, index),
          ...this.list.slice(index + 1),
        ]
      }
      this.updateKey += 1;
    },
    async retrieve(id: string, force = false): Promise<Category> {
      // Check if a refresh is required
      if (!force && !this.refreshRequired) {
        const category = this.get(id)
        if (category !== null) {
          return category
        }
      }

      // Get category
      const category = await CategoryTransporter.Get({
        id,
      }).then(async (entity) => {
        this.reinsert(entity)

        // Activate category
        await this.activate(entity)

        return entity
      })

      return category
    },
    async retrieveAll(force = false): Promise<{
      active: string | null
      items: Category[]
    }> {
      // Check if a refresh is required
      if (!force && !this.refreshRequired) {
        return {
          active: this.activeID,
          items: this.list as Category[],
        }
      }

      // Clear list
      while (this.list.length > 0) {
        this.list.pop()
      }

      // Retrieve list
      await CategoryTransporter.GetAll({
      }).then(async (list) => {
        list.forEach((category) => {
          this.reinsert(category)
        })

        // Activate category
        if (list.length > 0) {
          if (this.activeID === null) {
            await this.activate(list[0])
          }
        }
      })

      // Set retrieval timestamp
      this.retrievalTimestamp = Date.now()

      return {
        active: this.activeID,
        items: this.list as Category[],
      }
    },
    async retrievePaginated(pagination: Pagination, force = false): Promise<{
      pagination: Pagination;
      items: Category[];
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
          items: cached.ids.map((id: string) => this.get(id) as Category),
        };
      }

      // Retrieve list
      const items = await CategoryTransporter.GetPaginated(pagination).then(async ({items, pagination }) => {
        items.forEach((category) => {
          this.reinsert(category);
        });

        // Set new pagination
        newPagination = pagination.deepCopy();

        return items;
      });

      // Set cache
      this.cache[cacheKey] = {
        ids: items.map((x: Category) => x.id),
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
    async query(pagination: Pagination, queries: Query[], force = false): Promise<{
      pagination: Pagination;
      items: Category[];
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
          items: cached.ids.map((id: string) => this.get(id) as Category),
        };
      }

      // Retrieve list
      const items = await CategoryTransporter.Query(
        pagination,
        queries,
      ).then(async ({ items, pagination }) => {
        items.forEach((category) => {
          this.reinsert(category);
        });

        // Set new pagination
        newPagination = pagination.deepCopy();;

        return items
      });

      // Set cache
      this.cache[cacheKey] = {
        ids: items.map((x: Category) => x.id),
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
      
      pagination: Pagination,
      search: string,
      force = false,
    ): Promise<{
      pagination: Pagination;
      items: Category[];
    }> {
      // Check pagination
      let newPagination = new Pagination(pagination).deepCopy();

      // Get cachekey
      const cacheKey = search.replaceAll(" ", "_") + '-' + pagination.toCacheKey();

      // Check if a refresh is required
      if (!force && this.cache[cacheKey]) {
        const cached = this.cache[cacheKey]
        return {
          items: cached.ids.map(
            (id: string) => this.get(id) as Category,
          ),
          pagination: cached.pagination.deepCopy(),
        };
      }

      // Retrieve list
      const items = await CategoryTransporter.Search(
        pagination,
        search,
      ).then(async ({ items, pagination }) => {
        items.forEach((category) => {
          this.reinsert(category);
        });

        // Set new pagination
        newPagination = pagination.deepCopy();;

        return items
      });

      // Set cache
      this.cache[cacheKey] = {
        ids: items.map((x: Category) => x.id),
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
    async update(category: Category): Promise<Category> {
      // Update category
      const item = await CategoryTransporter.Update(category)

      // Reinsert category
      this.reinsert(item)

      // Activate category
      if (this.active === null) {
        await this.activate(item)
      }
      
      this.emitStoreUpdateEvent()

      return item
    },
  },
  getters: {
    active(): Category | null {
      if (this.activeID === null) {
        return null
      }

      const category = (this.list as Category[]).find((x: Category) => x.id === this.activeID)
      if (category === undefined) {
        return null
      }

      return category
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
export function createCategoryStore(): CategoryStore {
  const store = useCategoryStore()
  // Run the initializeStore action when the store is created
  store.initializeStore()
  return store as unknown as CategoryStore // TODO: Fix this
}
