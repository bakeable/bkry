import request from "../../api/request";
import { type QuestionBundleLocalisation } from './interface.gen'
import { type IPagination, Pagination } from "../../general/pagination";
import { Query } from "../../general/query";

type QuestionBundleLocalisationPathVariables = {
  question_bundle_id: string,
  id?: string,
}

export class QuestionBundleLocalisationTransporter {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  static config = Object.freeze({
    endpoint: "/question_bundle/:questionBundleID/question_bundle_localisation",
  })

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  private static Endpoint(params?: QuestionBundleLocalisationPathVariables): string {
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
    question_bundle_id,
    id,
  }: QuestionBundleLocalisationPathVariables): Promise<QuestionBundleLocalisation> {
    // Retrieve
    const { item }: { item: QuestionBundleLocalisation } = await request({
      url: this.Endpoint({
        question_bundle_id,
        id,
      }),
      method: "get",
    });

    // Parse
    return item
  }

  public static async Create({
    question_bundle_id,
  }: QuestionBundleLocalisationPathVariables, 
    data: QuestionBundleLocalisation
  ): Promise<QuestionBundleLocalisation> {
    // Create
    const { item }: { item: QuestionBundleLocalisation } = await request({
      url: this.Endpoint({
        question_bundle_id,
      }),
      method: "post",
      data,
    });

    // Parse
    return item
  }

  public static async Update(
    data: QuestionBundleLocalisation
  ): Promise<QuestionBundleLocalisation> {
    // Update
    const { item }: { item: QuestionBundleLocalisation } = await request({
      url: this.Endpoint(data),
      method: "put",
      data,
    });

    // Parse
    return item
  }

  public static async Delete({
    question_bundle_id,
    id,
  }: QuestionBundleLocalisationPathVariables): Promise<void> {
    // Delete
    await request({
      url: this.Endpoint({
        question_bundle_id,
        id,
      }),
      method: "delete",
    });
  }

  public static async GetAll({
    question_bundle_id,
  }: QuestionBundleLocalisationPathVariables): Promise<QuestionBundleLocalisation[]> {
    // Retrieve
    const { items }: { items: QuestionBundleLocalisation[] } = await request({
      url: this.Endpoint({
        question_bundle_id,
      }),
      method: "get",
    });

    return items
  }

  public static async Query({
    question_bundle_id,
  }: QuestionBundleLocalisationPathVariables,pagination: Pagination, queries: Query[]): Promise<{ items: QuestionBundleLocalisation[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: QuestionBundleLocalisation[], page_info: IPagination } = await request({
      url: this.Endpoint({
        question_bundle_id,
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
    question_bundle_id,
  }: QuestionBundleLocalisationPathVariables,pagination: Pagination): Promise<{ items: QuestionBundleLocalisation[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: QuestionBundleLocalisation[], page_info: IPagination } = await request({
      url: this.Endpoint({
        question_bundle_id,
      }),
      method: "get",
      params: {
        pagination: pagination.toRouteQuery(),
      },
    });

    return { items, pagination: Pagination.FromObject(page_info) }
  }

  public static async Search({
    question_bundle_id,
  }: QuestionBundleLocalisationPathVariables,
  pagination: Pagination, 
  search: string
  ): Promise<{ items: QuestionBundleLocalisation[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: QuestionBundleLocalisation[], page_info: IPagination }  = await request({
      url: this.Endpoint({
        question_bundle_id,
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

