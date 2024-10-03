import request from "../../api/request";
import { type Purchase } from './interface.gen'
import { type IPagination, Pagination } from "../../general/pagination";
import { Query } from "../../general/query";

type PurchasePathVariables = {
  user_id: string,
  id?: string,
}

export class PurchaseTransporter {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  static config = Object.freeze({
    endpoint: "/user/:userID/purchase",
  })

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  private static Endpoint(params?: PurchasePathVariables): string {
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
    user_id,
    id,
  }: PurchasePathVariables): Promise<Purchase> {
    // Retrieve
    const { item }: { item: Purchase } = await request({
      url: this.Endpoint({
        user_id,
        id,
      }),
      method: "get",
    });

    // Parse
    return item
  }

  public static async Create({
    user_id,
  }: PurchasePathVariables, 
    data: Purchase
  ): Promise<Purchase> {
    // Create
    const { item }: { item: Purchase } = await request({
      url: this.Endpoint({
        user_id,
      }),
      method: "post",
      data,
    });

    // Parse
    return item
  }

  public static async Update(
    data: Purchase
  ): Promise<Purchase> {
    // Update
    const { item }: { item: Purchase } = await request({
      url: this.Endpoint(data),
      method: "put",
      data,
    });

    // Parse
    return item
  }

  public static async Delete({
    user_id,
    id,
  }: PurchasePathVariables): Promise<void> {
    // Delete
    await request({
      url: this.Endpoint({
        user_id,
        id,
      }),
      method: "delete",
    });
  }

  public static async GetAll({
    user_id,
  }: PurchasePathVariables): Promise<Purchase[]> {
    // Retrieve
    const { items }: { items: Purchase[] } = await request({
      url: this.Endpoint({
        user_id,
      }),
      method: "get",
    });

    return items
  }

  public static async Query({
    user_id,
  }: PurchasePathVariables,pagination: Pagination, queries: Query[]): Promise<{ items: Purchase[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: Purchase[], page_info: IPagination } = await request({
      url: this.Endpoint({
        user_id,
      }),
      method: "get",
      params: {
        pagination: pagination.toRouteQuery(),
        queries: Query.ToQueryParams(queries),
      },
    });

    return { items, pagination: Pagination.FromObject(page_info) }
  }

  public static async GetPaginated({
    user_id,
  }: PurchasePathVariables,pagination: Pagination): Promise<{ items: Purchase[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: Purchase[], page_info: IPagination } = await request({
      url: this.Endpoint({
        user_id,
      }),
      method: "get",
      params: {
        pagination: pagination.toRouteQuery(),
      },
    });

    return { items, pagination: Pagination.FromObject(page_info) }
  }

  public static async Search({
    user_id,
  }: PurchasePathVariables,
  pagination: Pagination, 
  search: string
  ): Promise<{ items: Purchase[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: Purchase[], page_info: IPagination }  = await request({
      url: this.Endpoint({
        user_id,
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

