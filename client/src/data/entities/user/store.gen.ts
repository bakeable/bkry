import { defineStore } from 'pinia'
import { User } from './entity'
import { UserList } from './list.gen'
import type { UserStore, UserStoreState } from './store.gen.d'
import { analytics, auth } from '@/firebase'
import { type JSONData } from '../../base_classes/dto.d'
import { Pagination } from "../../base_classes/pagination";
import { Query } from "../../base_classes/query";
import { paginationAndQueriesToCacheKey } from "../../base_classes/dto_list";


/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// IF YOU WANT TO EDIT THIS FILE, REMOVE THE .gen. FROM THE FILENAME. THE FILE WILL THEN BE IGNORED.
/// ///////////////////////////////////////////////////////////////////////////////////////////////////

const userList = new UserList()
export const useUserStore = defineStore('userStore', {
  state: (): UserStoreState => ({
    activeID: null,
    cache: {},
    list: [] as User[],
    retrievalTimestamp: 0,
    updateKey: 0,
    storeUpdateTimeout: null,
  }),
  actions: {
    initializeStore(): void {
      this.convertList()
    },
    emitFormUpdateEvent(id: string): void {
      window.dispatchEvent(new CustomEvent("form-update-User", { detail: { id } }));
    },
    emitStoreUpdateEvent(): void {
      // Clear existing timeout
      if (this.storeUpdateTimeout !== null) {
        clearTimeout(this.storeUpdateTimeout)
      }

      // Set new timeout
      this.storeUpdateTimeout = setTimeout(() => {
        window.dispatchEvent(new CustomEvent("store.update.User"));
        this.storeUpdateTimeout = null
      }, 1500)
    },
    async activate(user: User): Promise<void> {
      // Set active
      await user
        .activate(auth.firebaseUser?.uid)
        .then(() => {
          this.activeID = user.id
        })
        .catch((error) => {
          console.error(error)
          this.activeID = null
        })
    },
    add(user: User): void {
      const id = user.id
      if (typeof id === 'string' && !this.contains(id)) {
        this.list.push(user)

        // Set active
        if (this.active === null) {
          this.activeID = user.id
        }
      } else {
        this.reinsert(user)
      }
      this.updateKey += 1;
    },
    convertList(): void {
      this.list = (this.list as User[]).map((x: any) => {
        const user = new User({})
        user.set(x as JSONData)
        return user
      })
    },
    async delete(user: User): Promise<void> {
      await user.delete().then(() => {
        this.remove(user)
      }).finally(() => {
        this.emitStoreUpdateEvent()
      });
    },
    async deleteById(id: string): Promise<void> {
      const user = (this.list as User[]).find((x: User) => x.id === id)
      if (user !== undefined) {
        await this.delete(user)
      } else {
        throw new Error('Could not find User with id: ' + id)
      }
    },
    contains(id: string): boolean {
      return (this.list as User[]).some((x: User) => x.id === id)
    },
    get(id: string): User | null {
      return (this.list as User[]).find((x: User) => x.id === id) || null
    },
    reinsert(user: User): void {
      const index = (this.list as User[]).findIndex((x: User) => x.id === user.id)
      if (index !== -1) {
        this.list = [
          ...this.list.slice(0, index),
          user,
          ...this.list.slice(index + 1),
        ]
      } else {
        this.add(user)
      }
      this.updateKey += 1;
    },
    remove({ id }: User): void {
      const index = (this.list as User[]).findIndex((x: User) => x.id === id)
      if (index !== -1) {
        this.list = [
          ...this.list.slice(0, index),
          ...this.list.slice(index + 1),
        ]
      }
      this.updateKey += 1;
    },
    async retrieve(id: string, force = false): Promise<User> {
      // Check if a refresh is required
      if (!force && !this.refreshRequired) {
        const user = this.get(id)
        if (user !== null) {
          return user
        }
      }

      // Retrieve user
      const user = new User({})

      // Get user
      await user.get([id]).then(async (entity) => {
        this.reinsert(entity)

        // Activate user
        await this.activate(entity)
      })

      return user
    },
    async retrieveAll(force = false): Promise<{
      active: string | null
      items: User[]
    }> {
      // Check if a refresh is required
      if (!force && !this.refreshRequired) {
        return {
          active: this.activeID,
          items: this.list as User[],
        }
      }

      // Clear list
      while (this.list.length > 0) {
        this.list.pop()
      }

      // Retrieve list
      await userList.getAll().then(async (list) => {
        list.forEach((user) => {
          this.reinsert(user)
        })

        // Activate user
        if (list.length > 0) {
          if (userList.active !== null) {
            this.activeID = userList.active
          } else {
            await this.activate(list[0])
          }
        }
      })

      // Set retrieval timestamp
      this.retrievalTimestamp = Date.now()

      // Set user property
      if (this.activeID !== null) {
        analytics.setUserProperty('user_id', this.activeID)
      }

      return {
        active: this.activeID,
        items: this.list as User[],
      }
    },
    async retrievePaginated(pagination: Pagination, force = false): Promise<{
      pagination: Pagination;
      items: User[];
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
          items: cached.ids.map((id: string) => this.get(id) as User),
        };
      }

      // Retrieve list
      const items = await userList.getPaginated(
        [],
        pagination
      ).then(async ({items, pagination }) => {
        items.forEach((user) => {
          this.reinsert(user);
        });

        // Set new pagination
        newPagination = pagination.deepCopy();

        return items;
      });

      // Set cache
      this.cache[cacheKey] = {
        ids: items.map((x: User) => x.id),
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
      items: User[];
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
          items: cached.ids.map((id: string) => this.get(id) as User),
        };
      }

      // Retrieve list
      const items = await userList.query(
        queries,
        pagination,
      ).then(async ({ items, pagination }) => {
        items.forEach((user) => {
          this.reinsert(user);
        });

        // Set new pagination
        newPagination = pagination.deepCopy();;

        return items
      });

      // Set cache
      this.cache[cacheKey] = {
        ids: items.map((x: User) => x.id),
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
      items: User[];
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
            (id: string) => this.get(id) as User,
          ),
          pagination: cached.pagination.deepCopy(),
        };
      }

      // Retrieve list
      const items = await userList
        .search(query, pagination)
        .then(async ({ items, pagination }) => {
          items.forEach((user) => {
            this.reinsert(user);
          });

          // Set new pagination
          newPagination = pagination.deepCopy();

          return items;
        });

      // Set cache
      this.cache[cacheKey] = {
        ids: items.map((x: User) => x.id),
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
    async update(user: User): Promise<string> {
      // Update user
      await user.update()

      // Reinsert user
      this.reinsert(user)

      // Activate user
      if (this.active === null) {
        await this.activate(user)
      }
      
      this.emitStoreUpdateEvent()

      return user.id
    },
  },
  getters: {
    active(): User | null {
      if (this.activeID === null) {
        return null
      }

      const user = (this.list as User[]).find((x: User) => x.id === this.activeID)
      if (user === undefined) {
        return null
      }

      return user
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
export function createUserStore(): UserStore {
  const store = useUserStore()
  // Run the initializeStore action when the store is created
  store.initializeStore()
  return store as unknown as UserStore // TODO: Fix this
}
