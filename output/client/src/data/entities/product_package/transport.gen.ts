import request from "../../api/request";
import { type ProductPackage } from './interface.gen'
import { type IPagination, Pagination } from "../../general/pagination";
import { Query } from "../../general/query";

type ProductPackagePathVariables = {
  id?: string,
}

export class ProductPackageTransporter {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  static config = Object.freeze({
    endpoint: "/product_package",
  })

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  private static Endpoint(params?: ProductPackagePathVariables): string {
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
  }: ProductPackagePathVariables): Promise<ProductPackage> {
    // Retrieve
    const { item }: { item: ProductPackage } = await request({
      url: this.Endpoint({
        id,
      }),
      method: "get",
    });

    // Parse
    return item
  }

  public static async Create({
  }: ProductPackagePathVariables, 
    data: ProductPackage
  ): Promise<ProductPackage> {
    // Create
    const { item }: { item: ProductPackage } = await request({
      url: this.Endpoint({
      }),
      method: "post",
      data,
    });

    // Parse
    return item
  }

  public static async Update(
    data: ProductPackage
  ): Promise<ProductPackage> {
    // Update
    const { item }: { item: ProductPackage } = await request({
      url: this.Endpoint(data),
      method: "put",
      data,
    });

    // Parse
    return item
  }

  public static async Delete({
    id,
  }: ProductPackagePathVariables): Promise<void> {
    // Delete
    await request({
      url: this.Endpoint({
        id,
      }),
      method: "delete",
    });
  }

  public static async GetAll(): Promise<ProductPackage[]> {
    // Retrieve
    const { items }: { items: ProductPackage[] } = await request({
      url: this.Endpoint(),
      method: "get",
    });

    return items
  }

  public static async Query(pagination: Pagination, queries: Query[]): Promise<{ items: ProductPackage[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: ProductPackage[], page_info: IPagination } = await request({
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

  public static async GetPaginated(pagination: Pagination): Promise<{ items: ProductPackage[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: ProductPackage[], page_info: IPagination } = await request({
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
  ): Promise<{ items: ProductPackage[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: ProductPackage[], page_info: IPagination }  = await request({
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

