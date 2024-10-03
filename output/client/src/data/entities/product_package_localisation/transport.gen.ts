import request from "../../api/request";
import { type ProductPackageLocalisation } from './interface.gen'
import { type IPagination, Pagination } from "../../general/pagination";
import { Query } from "../../general/query";

type ProductPackageLocalisationPathVariables = {
  product_package_id: string,
  id?: string,
}

export class ProductPackageLocalisationTransporter {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  static config = Object.freeze({
    endpoint: "/product_package/:productPackageID/product_package_localisation",
  })

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  private static Endpoint(params?: ProductPackageLocalisationPathVariables): string {
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
    product_package_id,
    id,
  }: ProductPackageLocalisationPathVariables): Promise<ProductPackageLocalisation> {
    // Retrieve
    const { item }: { item: ProductPackageLocalisation } = await request({
      url: this.Endpoint({
        product_package_id,
        id,
      }),
      method: "get",
    });

    // Parse
    return item
  }

  public static async Create({
    product_package_id,
  }: ProductPackageLocalisationPathVariables, 
    data: ProductPackageLocalisation
  ): Promise<ProductPackageLocalisation> {
    // Create
    const { item }: { item: ProductPackageLocalisation } = await request({
      url: this.Endpoint({
        product_package_id,
      }),
      method: "post",
      data,
    });

    // Parse
    return item
  }

  public static async Update(
    data: ProductPackageLocalisation
  ): Promise<ProductPackageLocalisation> {
    // Update
    const { item }: { item: ProductPackageLocalisation } = await request({
      url: this.Endpoint(data),
      method: "put",
      data,
    });

    // Parse
    return item
  }

  public static async Delete({
    product_package_id,
    id,
  }: ProductPackageLocalisationPathVariables): Promise<void> {
    // Delete
    await request({
      url: this.Endpoint({
        product_package_id,
        id,
      }),
      method: "delete",
    });
  }

  public static async GetAll({
    product_package_id,
  }: ProductPackageLocalisationPathVariables): Promise<ProductPackageLocalisation[]> {
    // Retrieve
    const { items }: { items: ProductPackageLocalisation[] } = await request({
      url: this.Endpoint({
        product_package_id,
      }),
      method: "get",
    });

    return items
  }

  public static async Query({
    product_package_id,
  }: ProductPackageLocalisationPathVariables,pagination: Pagination, queries: Query[]): Promise<{ items: ProductPackageLocalisation[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: ProductPackageLocalisation[], page_info: IPagination } = await request({
      url: this.Endpoint({
        product_package_id,
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
    product_package_id,
  }: ProductPackageLocalisationPathVariables,pagination: Pagination): Promise<{ items: ProductPackageLocalisation[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: ProductPackageLocalisation[], page_info: IPagination } = await request({
      url: this.Endpoint({
        product_package_id,
      }),
      method: "get",
      params: {
        pagination: pagination.toRouteQuery(),
      },
    });

    return { items, pagination: Pagination.FromObject(page_info) }
  }

  public static async Search({
    product_package_id,
  }: ProductPackageLocalisationPathVariables,
  pagination: Pagination, 
  search: string
  ): Promise<{ items: ProductPackageLocalisation[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: ProductPackageLocalisation[], page_info: IPagination }  = await request({
      url: this.Endpoint({
        product_package_id,
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

