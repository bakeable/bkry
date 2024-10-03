import request from "../../api/request";
import { type TagLocalisation } from './interface.gen'
import { type IPagination, Pagination } from "../../general/pagination";
import { Query } from "../../general/query";

type TagLocalisationPathVariables = {
  tag_id: string,
  id?: string,
}

export class TagLocalisationTransporter {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  static config = Object.freeze({
    endpoint: "/tag/:tagID/tag_localisation",
  })

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  private static Endpoint(params?: TagLocalisationPathVariables): string {
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
    tag_id,
    id,
  }: TagLocalisationPathVariables): Promise<TagLocalisation> {
    // Retrieve
    const { item }: { item: TagLocalisation } = await request({
      url: this.Endpoint({
        tag_id,
        id,
      }),
      method: "get",
    });

    // Parse
    return item
  }

  public static async Create({
    tag_id,
  }: TagLocalisationPathVariables, 
    data: TagLocalisation
  ): Promise<TagLocalisation> {
    // Create
    const { item }: { item: TagLocalisation } = await request({
      url: this.Endpoint({
        tag_id,
      }),
      method: "post",
      data,
    });

    // Parse
    return item
  }

  public static async Update(
    data: TagLocalisation
  ): Promise<TagLocalisation> {
    // Update
    const { item }: { item: TagLocalisation } = await request({
      url: this.Endpoint(data),
      method: "put",
      data,
    });

    // Parse
    return item
  }

  public static async Delete({
    tag_id,
    id,
  }: TagLocalisationPathVariables): Promise<void> {
    // Delete
    await request({
      url: this.Endpoint({
        tag_id,
        id,
      }),
      method: "delete",
    });
  }

  public static async GetAll({
    tag_id,
  }: TagLocalisationPathVariables): Promise<TagLocalisation[]> {
    // Retrieve
    const { items }: { items: TagLocalisation[] } = await request({
      url: this.Endpoint({
        tag_id,
      }),
      method: "get",
    });

    return items
  }

  public static async Query({
    tag_id,
  }: TagLocalisationPathVariables,pagination: Pagination, queries: Query[]): Promise<{ items: TagLocalisation[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: TagLocalisation[], page_info: IPagination } = await request({
      url: this.Endpoint({
        tag_id,
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
    tag_id,
  }: TagLocalisationPathVariables,pagination: Pagination): Promise<{ items: TagLocalisation[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: TagLocalisation[], page_info: IPagination } = await request({
      url: this.Endpoint({
        tag_id,
      }),
      method: "get",
      params: {
        pagination: pagination.toRouteQuery(),
      },
    });

    return { items, pagination: Pagination.FromObject(page_info) }
  }

  public static async Search({
    tag_id,
  }: TagLocalisationPathVariables,
  pagination: Pagination, 
  search: string
  ): Promise<{ items: TagLocalisation[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: TagLocalisation[], page_info: IPagination }  = await request({
      url: this.Endpoint({
        tag_id,
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

