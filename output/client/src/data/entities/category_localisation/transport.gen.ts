import request from "../../api/request";
import { type CategoryLocalisation } from './interface.gen'
import { type IPagination, Pagination } from "../../general/pagination";
import { Query } from "../../general/query";

type CategoryLocalisationPathVariables = {
  category_id: string,
  id?: string,
}

export class CategoryLocalisationTransporter {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  static config = Object.freeze({
    endpoint: "/category/:categoryID/category_localisation",
  })

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  private static Endpoint(params?: CategoryLocalisationPathVariables): string {
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
    category_id,
    id,
  }: CategoryLocalisationPathVariables): Promise<CategoryLocalisation> {
    // Retrieve
    const { item }: { item: CategoryLocalisation } = await request({
      url: this.Endpoint({
        category_id,
        id,
      }),
      method: "get",
    });

    // Parse
    return item
  }

  public static async Create({
    category_id,
  }: CategoryLocalisationPathVariables, 
    data: CategoryLocalisation
  ): Promise<CategoryLocalisation> {
    // Create
    const { item }: { item: CategoryLocalisation } = await request({
      url: this.Endpoint({
        category_id,
      }),
      method: "post",
      data,
    });

    // Parse
    return item
  }

  public static async Update(
    data: CategoryLocalisation
  ): Promise<CategoryLocalisation> {
    // Update
    const { item }: { item: CategoryLocalisation } = await request({
      url: this.Endpoint(data),
      method: "put",
      data,
    });

    // Parse
    return item
  }

  public static async Delete({
    category_id,
    id,
  }: CategoryLocalisationPathVariables): Promise<void> {
    // Delete
    await request({
      url: this.Endpoint({
        category_id,
        id,
      }),
      method: "delete",
    });
  }

  public static async GetAll({
    category_id,
  }: CategoryLocalisationPathVariables): Promise<CategoryLocalisation[]> {
    // Retrieve
    const { items }: { items: CategoryLocalisation[] } = await request({
      url: this.Endpoint({
        category_id,
      }),
      method: "get",
    });

    return items
  }

  public static async Query({
    category_id,
  }: CategoryLocalisationPathVariables,pagination: Pagination, queries: Query[]): Promise<{ items: CategoryLocalisation[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: CategoryLocalisation[], page_info: IPagination } = await request({
      url: this.Endpoint({
        category_id,
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
    category_id,
  }: CategoryLocalisationPathVariables,pagination: Pagination): Promise<{ items: CategoryLocalisation[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: CategoryLocalisation[], page_info: IPagination } = await request({
      url: this.Endpoint({
        category_id,
      }),
      method: "get",
      params: {
        pagination: pagination.toRouteQuery(),
      },
    });

    return { items, pagination: Pagination.FromObject(page_info) }
  }

  public static async Search({
    category_id,
  }: CategoryLocalisationPathVariables,
  pagination: Pagination, 
  search: string
  ): Promise<{ items: CategoryLocalisation[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: CategoryLocalisation[], page_info: IPagination }  = await request({
      url: this.Endpoint({
        category_id,
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

