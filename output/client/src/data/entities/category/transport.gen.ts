import request from "../../api/request";
import { type Category } from './interface.gen'
import { type IPagination, Pagination } from "../../general/pagination";
import { Query } from "../../general/query";

type CategoryPathVariables = {
  id?: string,
}

export class CategoryTransporter {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  static config = Object.freeze({
    endpoint: "/category",
  })

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  private static Endpoint(params?: CategoryPathVariables): string {
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
  }: CategoryPathVariables): Promise<Category> {
    // Retrieve
    const { item }: { item: Category } = await request({
      url: this.Endpoint({
        id,
      }),
      method: "get",
    });

    // Parse
    return item
  }

  public static async Create({
  }: CategoryPathVariables, 
    data: Category
  ): Promise<Category> {
    // Create
    const { item }: { item: Category } = await request({
      url: this.Endpoint({
      }),
      method: "post",
      data,
    });

    // Parse
    return item
  }

  public static async Update(
    data: Category
  ): Promise<Category> {
    // Update
    const { item }: { item: Category } = await request({
      url: this.Endpoint(data),
      method: "put",
      data,
    });

    // Parse
    return item
  }

  public static async Delete({
    id,
  }: CategoryPathVariables): Promise<void> {
    // Delete
    await request({
      url: this.Endpoint({
        id,
      }),
      method: "delete",
    });
  }

  public static async GetAll(): Promise<Category[]> {
    // Retrieve
    const { items }: { items: Category[] } = await request({
      url: this.Endpoint(),
      method: "get",
    });

    return items
  }

  public static async Query(pagination: Pagination, queries: Query[]): Promise<{ items: Category[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: Category[], page_info: IPagination } = await request({
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

  public static async GetPaginated(pagination: Pagination): Promise<{ items: Category[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: Category[], page_info: IPagination } = await request({
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
  ): Promise<{ items: Category[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: Category[], page_info: IPagination }  = await request({
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

