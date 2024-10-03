import request from "../../api/request";
import { type GameModusLocalisation } from './interface.gen'
import { type IPagination, Pagination } from "../../general/pagination";
import { Query } from "../../general/query";

type GameModusLocalisationPathVariables = {
  game_modus_id: string,
  id?: string,
}

export class GameModusLocalisationTransporter {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  static config = Object.freeze({
    endpoint: "/game_modus/:gameModusID/game_modus_localisation",
  })

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  private static Endpoint(params?: GameModusLocalisationPathVariables): string {
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
    game_modus_id,
    id,
  }: GameModusLocalisationPathVariables): Promise<GameModusLocalisation> {
    // Retrieve
    const { item }: { item: GameModusLocalisation } = await request({
      url: this.Endpoint({
        game_modus_id,
        id,
      }),
      method: "get",
    });

    // Parse
    return item
  }

  public static async Create({
    game_modus_id,
  }: GameModusLocalisationPathVariables, 
    data: GameModusLocalisation
  ): Promise<GameModusLocalisation> {
    // Create
    const { item }: { item: GameModusLocalisation } = await request({
      url: this.Endpoint({
        game_modus_id,
      }),
      method: "post",
      data,
    });

    // Parse
    return item
  }

  public static async Update(
    data: GameModusLocalisation
  ): Promise<GameModusLocalisation> {
    // Update
    const { item }: { item: GameModusLocalisation } = await request({
      url: this.Endpoint(data),
      method: "put",
      data,
    });

    // Parse
    return item
  }

  public static async Delete({
    game_modus_id,
    id,
  }: GameModusLocalisationPathVariables): Promise<void> {
    // Delete
    await request({
      url: this.Endpoint({
        game_modus_id,
        id,
      }),
      method: "delete",
    });
  }

  public static async GetAll({
    game_modus_id,
  }: GameModusLocalisationPathVariables): Promise<GameModusLocalisation[]> {
    // Retrieve
    const { items }: { items: GameModusLocalisation[] } = await request({
      url: this.Endpoint({
        game_modus_id,
      }),
      method: "get",
    });

    return items
  }

  public static async Query({
    game_modus_id,
  }: GameModusLocalisationPathVariables,pagination: Pagination, queries: Query[]): Promise<{ items: GameModusLocalisation[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: GameModusLocalisation[], page_info: IPagination } = await request({
      url: this.Endpoint({
        game_modus_id,
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
    game_modus_id,
  }: GameModusLocalisationPathVariables,pagination: Pagination): Promise<{ items: GameModusLocalisation[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: GameModusLocalisation[], page_info: IPagination } = await request({
      url: this.Endpoint({
        game_modus_id,
      }),
      method: "get",
      params: {
        pagination: pagination.toRouteQuery(),
      },
    });

    return { items, pagination: Pagination.FromObject(page_info) }
  }

  public static async Search({
    game_modus_id,
  }: GameModusLocalisationPathVariables,
  pagination: Pagination, 
  search: string
  ): Promise<{ items: GameModusLocalisation[], pagination: Pagination }> {
    // Retrieve
    const { items, page_info }: { items: GameModusLocalisation[], page_info: IPagination }  = await request({
      url: this.Endpoint({
        game_modus_id,
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

