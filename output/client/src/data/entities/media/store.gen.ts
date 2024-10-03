import { defineStore } from 'pinia'
import { Media } from './interface.gen'
import type { MediaStore, MediaStoreState } from './store.gen.d'
import { Pagination, paginationAndQueriesToCacheKey } from "../../general/pagination";
import { Query } from "../../general/query";
import { MediaTransporter } from './transport.gen'

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// IF YOU WANT TO EDIT THIS FILE, REMOVE THE .gen. FROM THE FILENAME. THE FILE WILL THEN BE IGNORED.
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

export const useMediaStore = defineStore('mediaStore', {
  state: (): MediaStoreState => ({
    activeID: null,
    cache: {},
    list: [] as Media[],
    retrievalTimestamp: 0,
    updateKey: 0,
    storeUpdateTimeout: null,
  }),
  actions: {
    emitFormUpdateEvent(id: string): void {
      window.dispatchEvent(new CustomEvent("form.update.media", { detail: { id } }));
    },
    emitStoreUpdateEvent(): void {
      // Clear existing timeout
      if (this.storeUpdateTimeout !== null) {
        clearTimeout(this.storeUpdateTimeout)
      }

      // Set new timeout
      this.storeUpdateTimeout = setTimeout(() => {
        window.dispatchEvent(new CustomEvent("store.update.media"));
        this.storeUpdateTimeout = null
      }, 1500)
    },
    async activate(media: Media): Promise<void> {
      // Set active
      this.activeID = media.id
    },
    add(media: Media): void {
      const id = media.id
      if (typeof id === 'string' && !this.contains(id)) {
        this.list.push(media)

        // Set active
        if (this.active === null) {
          this.activeID = media.id
        }
      } else {
        this.reinsert(media)
      }
      this.updateKey += 1;
    },
    async delete(media: Media): Promise<void> {
      await MediaTransporter.Delete(media).then(() => {
        this.remove(media)
      }).finally(() => {
        this.emitStoreUpdateEvent()
      });
    },
    async deleteById(id: string): Promise<void> {
      const media = (this.list as Media[]).find((x: Media) => x.id === id)
      if (media !== undefined) {
        await this.delete(media)
      } else {
        throw new Error('Could not find Media with id: ' + id)
      }
    },
    contains(id: string): boolean {
      return (this.list as Media[]).some((x: Media) => x.id === id)
    },
    async create(media: Media): Promise<Media> {
      // Create media
      const item = await MediaTransporter.Create({
      }, media)

      // Reinsert media
      this.reinsert(item)

      // Activate media
      if (this.active === null) {
        await this.activate(item)
      }
      
      this.emitStoreUpdateEvent()

      return item
    },
    get(id: string): Media | null {
      return (this.list as Media[]).find((x: Media) => x.id === id) || null
    },
    reinsert(media: Media): void {
      const index = (this.list as Media[]).findIndex((x: Media) => x.id === media.id)
      if (index !== -1) {
        this.list = [
          ...this.list.slice(0, index),
          media,
          ...this.list.slice(index + 1),
        ]
      } else {
        this.add(media)
      }
      this.updateKey += 1;
    },
    remove({ id }: Media): void {
      const index = (this.list as Media[]).findIndex((x: Media) => x.id === id)
      if (index !== -1) {
        this.list = [
          ...this.list.slice(0, index),
          ...this.list.slice(index + 1),
        ]
      }
      this.updateKey += 1;
    },
    async retrieve(id: string, force = false): Promise<Media> {
      // Check if a refresh is required
      if (!force && !this.refreshRequired) {
        const media = this.get(id)
        if (media !== null) {
          return media
        }
      }

      // Get media
      const media = await MediaTransporter.Get({
        id,
      }).then(async (entity) => {
        this.reinsert(entity)

        // Activate media
        await this.activate(entity)

        return entity
      })

      return media
    },
    async retrieveAll(force = false): Promise<{
      active: string | null
      items: Media[]
    }> {
      // Check if a refresh is required
      if (!force && !this.refreshRequired) {
        return {
          active: this.activeID,
          items: this.list as Media[],
        }
      }

      // Clear list
      while (this.list.length > 0) {
        this.list.pop()
      }

      // Retrieve list
      await MediaTransporter.GetAll({
      }).then(async (list) => {
        list.forEach((media) => {
          this.reinsert(media)
        })

        // Activate media
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
        items: this.list as Media[],
      }
    },
    async retrievePaginated(pagination: Pagination, force = false): Promise<{
      pagination: Pagination;
      items: Media[];
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
          items: cached.ids.map((id: string) => this.get(id) as Media),
        };
      }

      // Retrieve list
      const items = await MediaTransporter.GetPaginated(pagination).then(async ({items, pagination }) => {
        items.forEach((media) => {
          this.reinsert(media);
        });

        // Set new pagination
        newPagination = pagination.deepCopy();

        return items;
      });

      // Set cache
      this.cache[cacheKey] = {
        ids: items.map((x: Media) => x.id),
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
      items: Media[];
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
          items: cached.ids.map((id: string) => this.get(id) as Media),
        };
      }

      // Retrieve list
      const items = await MediaTransporter.Query(
        pagination,
        queries,
      ).then(async ({ items, pagination }) => {
        items.forEach((media) => {
          this.reinsert(media);
        });

        // Set new pagination
        newPagination = pagination.deepCopy();;

        return items
      });

      // Set cache
      this.cache[cacheKey] = {
        ids: items.map((x: Media) => x.id),
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
      items: Media[];
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
            (id: string) => this.get(id) as Media,
          ),
          pagination: cached.pagination.deepCopy(),
        };
      }

      // Retrieve list
      const items = await MediaTransporter.Search(
        pagination,
        search,
      ).then(async ({ items, pagination }) => {
        items.forEach((media) => {
          this.reinsert(media);
        });

        // Set new pagination
        newPagination = pagination.deepCopy();;

        return items
      });

      // Set cache
      this.cache[cacheKey] = {
        ids: items.map((x: Media) => x.id),
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
    async update(media: Media): Promise<Media> {
      // Update media
      const item = await MediaTransporter.Update(media)

      // Reinsert media
      this.reinsert(item)

      // Activate media
      if (this.active === null) {
        await this.activate(item)
      }
      
      this.emitStoreUpdateEvent()

      return item
    },
  },
  getters: {
    active(): Media | null {
      if (this.activeID === null) {
        return null
      }

      const media = (this.list as Media[]).find((x: Media) => x.id === this.activeID)
      if (media === undefined) {
        return null
      }

      return media
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
export function createMediaStore(): MediaStore {
  const store = useMediaStore()
  // Run the initializeStore action when the store is created
  store.initializeStore()
  return store as unknown as MediaStore // TODO: Fix this
}
