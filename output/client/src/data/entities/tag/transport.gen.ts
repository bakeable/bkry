import request from "../../api/request";
import { type Tag } from './interface.gen'
import { type IPagination, Pagination } from "../../general/pagination";
import { Query } from "../../general/query";

type TagPathVariables = {
  id?: string,
}

export class TagTransporter {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  static config = Object.freeze({
    endpoint: "/tag",
  })

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  private static Endpoint(params?: TagPathVariables): string {
    let path = this.config.endpoint.toString()
    if (!params) return path
    for (const key in Object.keys(params)) {
      path = path.replace(`{${key}}`, params[key])
    }
    return path
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  public static async Get({
    id,
  }: TagPathVariables): Promise<Tag> {
    // Retrieve
    const { item }: { item: Tag } = await request({
      url: this.Endpoint({
        id,
      }),
      method: "get",
    });

    // Parse
    return item
  }

  public static async Create({
  }: TagPathVariables, 
    data: Tag
  ): Promise<Tag> {
    // Create
    const { item }: { item: Tag } = await request({
      url: this.Endpoint({
      }),
      method: "post",
      data,
    });

    // Parse
    return item
  }

  public static async Update(
    data: Tag
  ): Promise<Tag> {
    // Update
    const { item }: { item: Tag } = await request({
      url: this.Endpoint(data),
      method: "put",
      data,
    });

    // Parse
    return item
  }

  public static async Delete({
    id,
  }: TagPathVariables): Promise<void> {
    // Delete
    await request({
      url: this.Endpoint({
        id,
      }),
      method: "delete",
    });
  }

  public static async GetAll(): Promise<Tag[]> {
    // Retrieve
    const { items }: { items: Tag[] } = await request({
      url: this.Endpoint(),
      method: "get",
    });

    return items
  }

  public static async Query(pagination: Pagination, queries: Query[]): Promise<{ items: Tag[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: Tag[], page_info: IPagination } = await request({
      url: this.Endpoint({
      }),
      method: "get",
      params: {
        pagination: pagination.toRouteQuery(),
        queries: Query.ToQueryParams(queries),
      },
    });

    return { items, pagination: Pagination.FromObject(page_info) }
  }

  public static async GetPaginated(pagination: Pagination): Promise<{ items: Tag[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: Tag[], page_info: IPagination } = await request({
      url: this.Endpoint({
      }),
      method: "get",
      params: {
        pagination: pagination.toRouteQuery(),
      },
    });

    return { items, pagination: Pagination.FromObject(page_info) }
  }

  public static async Search(
  pagination: Pagination, 
  search: string
  ): Promise<{ items: Tag[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: Tag[], page_info: IPagination }  = await request({
      url: this.Endpoint({
      }),
      method: "get",
      params: {
        pagination: pagination.toRouteQuery(),
        search,
      },
    });

    return { items, pagination: Pagination.FromObject(page_info) }
  }
}

