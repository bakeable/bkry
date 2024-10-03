import request from "../../api/request";
import { type QuestionContext } from './interface.gen'
import { type IPagination, Pagination } from "../../general/pagination";
import { Query } from "../../general/query";

type QuestionContextPathVariables = {
  id?: string,
}

export class QuestionContextTransporter {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  static config = Object.freeze({
    endpoint: "/question_context",
  })

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  private static Endpoint(params?: QuestionContextPathVariables): string {
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
  }: QuestionContextPathVariables): Promise<QuestionContext> {
    // Retrieve
    const { item }: { item: QuestionContext } = await request({
      url: this.Endpoint({
        id,
      }),
      method: "get",
    });

    // Parse
    return item
  }

  public static async Create({
  }: QuestionContextPathVariables, 
    data: QuestionContext
  ): Promise<QuestionContext> {
    // Create
    const { item }: { item: QuestionContext } = await request({
      url: this.Endpoint({
      }),
      method: "post",
      data,
    });

    // Parse
    return item
  }

  public static async Update(
    data: QuestionContext
  ): Promise<QuestionContext> {
    // Update
    const { item }: { item: QuestionContext } = await request({
      url: this.Endpoint(data),
      method: "put",
      data,
    });

    // Parse
    return item
  }

  public static async Delete({
    id,
  }: QuestionContextPathVariables): Promise<void> {
    // Delete
    await request({
      url: this.Endpoint({
        id,
      }),
      method: "delete",
    });
  }

  public static async GetAll(): Promise<QuestionContext[]> {
    // Retrieve
    const { items }: { items: QuestionContext[] } = await request({
      url: this.Endpoint(),
      method: "get",
    });

    return items
  }

  public static async Query(pagination: Pagination, queries: Query[]): Promise<{ items: QuestionContext[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: QuestionContext[], page_info: IPagination } = await request({
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

  public static async GetPaginated(pagination: Pagination): Promise<{ items: QuestionContext[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: QuestionContext[], page_info: IPagination } = await request({
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
  ): Promise<{ items: QuestionContext[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: QuestionContext[], page_info: IPagination }  = await request({
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

