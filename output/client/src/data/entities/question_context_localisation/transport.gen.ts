import request from "../../api/request";
import { type QuestionContextLocalisation } from './interface.gen'
import { type IPagination, Pagination } from "../../general/pagination";
import { Query } from "../../general/query";

type QuestionContextLocalisationPathVariables = {
  question_context_id: string,
  id?: string,
}

export class QuestionContextLocalisationTransporter {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  static config = Object.freeze({
    endpoint: "/question_context/:questionContextID/question_context_localisation",
  })

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  private static Endpoint(params?: QuestionContextLocalisationPathVariables): string {
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
    question_context_id,
    id,
  }: QuestionContextLocalisationPathVariables): Promise<QuestionContextLocalisation> {
    // Retrieve
    const { item }: { item: QuestionContextLocalisation } = await request({
      url: this.Endpoint({
        question_context_id,
        id,
      }),
      method: "get",
    });

    // Parse
    return item
  }

  public static async Create({
    question_context_id,
  }: QuestionContextLocalisationPathVariables, 
    data: QuestionContextLocalisation
  ): Promise<QuestionContextLocalisation> {
    // Create
    const { item }: { item: QuestionContextLocalisation } = await request({
      url: this.Endpoint({
        question_context_id,
      }),
      method: "post",
      data,
    });

    // Parse
    return item
  }

  public static async Update(
    data: QuestionContextLocalisation
  ): Promise<QuestionContextLocalisation> {
    // Update
    const { item }: { item: QuestionContextLocalisation } = await request({
      url: this.Endpoint(data),
      method: "put",
      data,
    });

    // Parse
    return item
  }

  public static async Delete({
    question_context_id,
    id,
  }: QuestionContextLocalisationPathVariables): Promise<void> {
    // Delete
    await request({
      url: this.Endpoint({
        question_context_id,
        id,
      }),
      method: "delete",
    });
  }

  public static async GetAll({
    question_context_id,
  }: QuestionContextLocalisationPathVariables): Promise<QuestionContextLocalisation[]> {
    // Retrieve
    const { items }: { items: QuestionContextLocalisation[] } = await request({
      url: this.Endpoint({
        question_context_id,
      }),
      method: "get",
    });

    return items
  }

  public static async Query({
    question_context_id,
  }: QuestionContextLocalisationPathVariables,pagination: Pagination, queries: Query[]): Promise<{ items: QuestionContextLocalisation[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: QuestionContextLocalisation[], page_info: IPagination } = await request({
      url: this.Endpoint({
        question_context_id,
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
    question_context_id,
  }: QuestionContextLocalisationPathVariables,pagination: Pagination): Promise<{ items: QuestionContextLocalisation[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: QuestionContextLocalisation[], page_info: IPagination } = await request({
      url: this.Endpoint({
        question_context_id,
      }),
      method: "get",
      params: {
        pagination: pagination.toRouteQuery(),
      },
    });

    return { items, pagination: Pagination.FromObject(page_info) }
  }

  public static async Search({
    question_context_id,
  }: QuestionContextLocalisationPathVariables,
  pagination: Pagination, 
  search: string
  ): Promise<{ items: QuestionContextLocalisation[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: QuestionContextLocalisation[], page_info: IPagination }  = await request({
      url: this.Endpoint({
        question_context_id,
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

