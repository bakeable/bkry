import { Query } from "./query";

export interface IPagination {
  page_number: number;
  page_size: number;
  order_by: string;
  asc: boolean;
  count?: number;
  total_pages?: number;
  has_next_page: boolean;
  has_previous_page: boolean;
}

/**
 * Represents a pagination configuration.
 */
export class Pagination implements IPagination {
  constructor(data?: IPagination) {
    if (data) {
      this.fromObject(data);
    }
  }

  /**
   * The current page number.
   */
  page_number: number = 1;

  /**
   * The number of items per page.
   */
  page_size: number = 10;

  /**
   * The field to order the items by.
   */
  order_by: string = "id";

  /**
   * Specifies whether the items should be ordered in ascending order.
   */
  asc: boolean = false;

  /**
   * Specifies whether there is a next page available.
   */
  has_next_page: boolean = false;

  /**
   * Specifies whether there is a previous page available.
   */
  has_previous_page: boolean = false;

  /// //////////////////////////////////////////////
  /// ////////////// METHODS ///////////////////////
  /// //////////////////////////////////////////////

  /**
   * Converts an object of type IPagination to a Pagination instance.
   *
   * @param obj - The object to convert.
   * @returns The converted Pagination instance.
   */
  fromObject(obj: IPagination): Pagination {
    this.page_number = obj.page_number;
    this.page_size = obj.page_size;
    this.order_by = obj.order_by;
    this.asc = !!obj.asc;
    this.has_next_page = obj.has_next_page ?? false;
    this.has_previous_page = obj.has_previous_page ?? false;

    return this;
  }

  /**
   * Converts a incomplete pagination object to defaults
   */
  withDefaults({
    page_number = 1,
    page_size = 10,
    order_by = "id",
    asc = false,
  }): Pagination {
    this.page_number = page_number;
    this.page_size = page_size;
    this.order_by = order_by;
    this.asc = asc;
    this.has_next_page = false;
    this.has_previous_page = false;

    return this;
  }

  /**
   * Converts the pagination configuration to a cache key.
   * @returns The cache key representing the pagination configuration.
   */
  toCacheKey(): string {
    return `${this.page_number}-${this.page_size}-${this.order_by}-${this.asc}`;
  }

  /**
   * Converts the pagination configuration to a route query.
   * @returns The route query representing the pagination configuration.
   */
  toRouteQuery(): string {
    return JSON.stringify({
      page_number: this.page_number,
      page_size: this.page_size,
      order_by: this.order_by,
      asc: this.asc,
    });
  }

  fromRouteQuery(query: string): Pagination {
    const parsed = JSON.parse(query);
    return this.fromObject(parsed);
  }

  equals(pag: IPagination): boolean {
    return (
      this.page_number === pag.page_number &&
      this.page_size === pag.page_size &&
      this.order_by === pag.order_by &&
      this.asc === pag.asc
    );
  }

  deepCopy(): Pagination {
    return new Pagination().fromObject(this);
  }

  /// //////////////////////////////////////////////
  /// ////////////// STATIC METHODS ////////////////
  /// //////////////////////////////////////////////

  /**
   * Creates a Pagination instance from an IPagination object.
   * @param obj - The IPagination object containing the pagination configuration.
   * @returns The created Pagination instance.
   */
  static FromObject(obj: IPagination): Pagination {
    return new Pagination().fromObject(obj);
  }

  /**
   * Converts a pagination configuration to a cache key.
   * @param pag - The pagination configuration to convert.
   * @returns The cache key representing the pagination configuration.
   */
  static ToCacheKey(pag: IPagination): string {
    return `${pag?.page_number ?? 1}-${pag?.page_size ?? 10}-${pag?.order_by ?? "id"}-${pag?.asc ?? false}`;
  }

  /**
   * Creates a Pagination instance from a route query.
   * @param query - The route query containing the pagination configuration.
   * @returns The created Pagination instance.
   */
  static FromRouteQuery(query: string): Pagination {
    const parsed = JSON.parse(query);
    return new Pagination().fromObject(parsed);
  }

  /**
   * Creates a Pagination instance with defaults
   */
  static WithDefaults({
    page_number = 1,
    page_size = 10,
    order_by = "id",
    asc = false,
  }): Pagination {
    return new Pagination().withDefaults({
      page_number,
      page_size,
      order_by,
      asc,
    });
  }
}



export function paginationToCacheKey(pag: Pagination): string {
  return `${pag?.page_number ?? 1}-${pag?.page_size ?? 10}-${pag?.order_by ?? "id"}-${pag?.asc ?? false}`;
}

export function queriesToCachKey(queries: Query[]): string {
  return queries.map((q) => [q.field, q.operator, q.value].join("-")).join("-");
}

export function paginationAndQueriesToCacheKey(
  pag: Pagination,
  queries: Query[],
): string {
  return `${paginationToCacheKey(pag)}-${queriesToCachKey(queries)}`;
}