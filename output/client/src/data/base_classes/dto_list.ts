import request from "@/data/api/request";
import { Dto } from "./dto";
import type IDtoList from "./dto_list.d";
import {
  type DataListResponse,
  type Pagination,
  type Query,
} from "./dto_list.d";
import { type JSONData } from "./dto.d";

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

export default class DtoList implements IDtoList {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  _path = "example/{example_id}/path";
  _pathVars: {
    [key: string]: string;
  } = {};

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
  active: string | null = "";
  list: Dto[] = [];
  pages: Record<string, Dto[]> = {};
  pagination: Pagination = {
    page_number: 1,
    order_by: "id",
    page_size: 10,
    asc: false,
  };

  queries: Query[] = [];
  queriedList: Dto[] = [];

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  get getPathKeys(): string[] {
    // Get path keys
    return (this._path.match(/{[^}]*}/g) ?? []).map((key) =>
      key.replace("{", "").replace("}", ""),
    );
  }

  buildPathObject(
    pathVars?:
      | string
      | string[]
      | {
          [key: string]: string;
        },
  ): {
    [key: string]: string;
  } {
    // Failsafe
    if (typeof pathVars === "string") {
      pathVars = [pathVars];
    } else if (pathVars === undefined || pathVars === null) {
      return this._pathVars;
    }

    // Build the path object
    this._pathVars = this.getPathKeys.reduce(
      (acc, key, i) => {
        if (typeof pathVars === "object" && !Array.isArray(pathVars)) {
          if (pathVars[key] === undefined) {
            throw new Error(`Missing path variable: ${key}`);
          } else {
            acc[key] = pathVars[key];
          }
        } else if (pathVars.length < i + 1) {
          throw new Error(`Missing path variable: ${key}`);
        } else {
          acc[key] = pathVars[i];
        }

        return acc;
      },
      {} as {
        [key: string]: string;
      },
    );

    return this._pathVars;
  }

  buildPath(
    pathVars?:
      | string
      | string[]
      | {
          [key: string]: string;
        },
  ): string {
    let path = this._path;

    // Format path variables
    this._pathVars = this.buildPathObject(pathVars);

    // Replace path variables
    Object.entries(this._pathVars).forEach(([key, value]) => {
      path = path.replace(`{${key}}`, value);
    });

    return this._path;
  }

  async delete(entities?: Dto[]): Promise<void> {
    // Get entities
    if (entities === undefined) {
      entities = this.list;
    }

    // Loop through entities
    await Promise.all(
      entities.map(async (entity: Dto) => {
        await entity.delete(this._pathVars);
      }),
    );
  }

  async getAll(
    pathVars?: string | string[] | { [key: string]: string },
  ): Promise<Dto[]> {
    // Build path
    const path = this.buildPath(pathVars);

    // Get data
    const data: DataListResponse =
      (await request({
        url: path,
        method: "get",
      })) ?? [];

    // Unpack
    const { items } = data;
    if (data.active !== undefined) {
      this.active = data.active;
    }

    // Exists
    if (typeof items !== "object") {
      throw new Error("Incorrect data");
    } else if (items === null) {
      console.info("No items found");
      return [];
    }

    // Build
    items.forEach((item: JSONData) => {
      // Create dto
      const dto = this._buildInstance(item);

      // Reinsert into list
      this.reinsert(dto);
    });

    return this.list;
  }

  async getPaginated(
    pathVars?: string | string[] | { [key: string]: string },
    pag?: Pagination,
  ): Promise<{
    items: Dto[];
    pagination: Pagination;
  }> {
    // Build path
    const path = this.buildPath(pathVars);

    // Set pagination
    if (pag !== undefined) {
      this.pagination = pag;
    }

    // Get data
    const data: DataListResponse =
      (await request({
        url: path,
        method: "get",
        params: this.pagination,
      })) ?? [];

    // Unpack
    const { items, pagination } = data;
    if (data.active !== undefined) {
      this.active = data.active;
    }

    // Exists
    if (typeof items !== "object" || items === null) {
      throw new Error("Incorrect data");
    }

    // Set pagination equal to server
    if (pagination !== undefined) {
      this.pagination = pagination;
    }

    // Build
    const dtoList = [] as Dto[];
    items.forEach((item: JSONData) => {
      // Create dto
      const dto = this._buildInstance(item);

      // Set page number
      dto._page = this.pagination.page_number;

      // Reinsert into list, cache and dtoList
      this.reinsert(dto);
      dtoList.push(dto);
    });

    return {
      items: dtoList,
      pagination: this.pagination,
    };
  }

  async query(queries: Query[], pagination?: Pagination): Promise<Dto[]> {
    // Build path
    const path = this.buildPath();

    // Get data
    const { items }: DataListResponse = await request({
      url: path + "/query",
      method: "post",
      data: {
        queries,
        pagination,
      },
    });

    // Exists
    if (typeof items !== "object" || items === null) {
      throw new Error("Incorrect data");
    }

    // Build
    const dtoList = [] as Dto[];
    items.forEach((entity: JSONData) => {
      // Create dto
      const dto = this._buildInstance(entity);

      // Reinsert into list
      this.reinsert(dto);
      dtoList.push(dto);
    });

    return dtoList;
  }

  async update(entities?: Dto[] | undefined): Promise<void> {
    // Get entities
    if (entities === undefined) {
      entities = this.list;
    }

    // Loop through entities
    await Promise.all(
      entities.map(async (entity: Dto) => {
        await entity.update(this._pathVars);
      }),
    );
  }

  private reinsert(entity: Dto): void {
    // Find entity
    const index = this.list.findIndex((e) => e.id === entity.id);

    // Reinsert
    if (index === -1) {
      this.list.push(entity);
    } else {
      this.list[index] = entity;
    }
  }

  _buildInstance(data: JSONData): Dto {
    return new Dto(data);
  }
}
