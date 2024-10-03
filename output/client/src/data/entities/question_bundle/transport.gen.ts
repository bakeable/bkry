import request from "../../api/request";
import { type QuestionBundle } from './interface.gen'
import { type IPagination, Pagination } from "../../general/pagination";
import { Query } from "../../general/query";

type QuestionBundlePathVariables = {
  id?: string,
}

export class QuestionBundleTransporter {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  static config = Object.freeze({
    endpoint: "/question_bundle",
  })

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  private static Endpoint(params?: QuestionBundlePathVariables): string {
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
  }: QuestionBundlePathVariables): Promise<QuestionBundle> {
    // Retrieve
    const { item }: { item: QuestionBundle } = await request({
      url: this.Endpoint({
        id,
      }),
      method: "get",
    });

    // Parse
    return item
  }

  public static async Create({
  }: QuestionBundlePathVariables, 
    data: QuestionBundle
  ): Promise<QuestionBundle> {
    // Create
    const { item }: { item: QuestionBundle } = await request({
      url: this.Endpoint({
      }),
      method: "post",
      data,
    });

    // Parse
    return item
  }

  public static async Update(
    data: QuestionBundle
  ): Promise<QuestionBundle> {
    // Update
    const { item }: { item: QuestionBundle } = await request({
      url: this.Endpoint(data),
      method: "put",
      data,
    });

    // Parse
    return item
  }

  public static async Delete({
    id,
  }: QuestionBundlePathVariables): Promise<void> {
    // Delete
    await request({
      url: this.Endpoint({
        id,
      }),
      method: "delete",
    });
  }

  public static async GetAll(): Promise<QuestionBundle[]> {
    // Retrieve
    const { items }: { items: QuestionBundle[] } = await request({
      url: this.Endpoint(),
      method: "get",
    });

    return items
  }

  public static async Query(pagination: Pagination, queries: Query[]): Promise<{ items: QuestionBundle[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: QuestionBundle[], page_info: IPagination } = await request({
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

  public static async GetPaginated(pagination: Pagination): Promise<{ items: QuestionBundle[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: QuestionBundle[], page_info: IPagination } = await request({
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
  ): Promise<{ items: QuestionBundle[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: QuestionBundle[], page_info: IPagination }  = await request({
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

