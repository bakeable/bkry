/**
 * Classes based on the Dto class are the base of the communication with the back-end.
 * The classes have the following notation requirements
 *
 * Private variables are denoted in camelCase
 * Public variables, which are usually externally retrieved information or calculated properties, are denoted in snake_case
 * Methods are denoted in camelCase
 *
 * Classes that retrieve back-end information should extend the Dto or List class to retrieve information
 */

import { Dto } from "./dto";

export type Query = {
  field: string;
  operator:
    | "=="
    | "!="
    | "<="
    | "<"
    | ">"
    | ">="
    | "in"
    | "array-contains-any"
    | "array-contains";
  value: unknown;
};

export type Pagination = {
  page_number: number;
  page_size: number;
  order_by: string;
  asc: boolean;
  count?: number;
  total_pages?: number;
};

export type DataListResponse = {
  active?: string;
  items: JSONData[];
  pagination?: Pagination;
};

export default interface IDtoList {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////

  /**
   * The path to the correspondng API endpoint
   */
  _path: string;

  /**
   * The path variables
   */
  _pathVars: {
    [key: string]: string;
  };

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
  /**
   * The ID of the data transfer object that is considered active
   */
  active: string | null;

  /**
   * The list of all  Data Transfer Objects
   */
  list: Dto[];

  /**
   * The pages of the queried list
   */
  pages: Record<string, Dto[]>;

  /**
   * The pagination queries
   */
  pagination: Pagination;

  /**
   * The queries applied to the queried list
   */
  queries: Query[];

  /**
   * The queried list of all  Data Transfer Objects
   */
  queriedList: Dto[];

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  /**
   * Applies the path variables
   * @param {string | string[] | {
   *  [key: string]: string;
   * }} pathVars A single, multiple or mapped path variables
   * @returns {string[]} The path variables
   */
  buildPathObject(
    pathVars?:
      | string
      | string[]
      | {
          [key: string]: string;
        },
  ): {
    [key: string]: string;
  };

  /**
   * Builds the full path to the corresponding API endpoint
   * @param {string | string[]} pathVars A single or multiple path variables
   * @returns {string} The path to the API endpoint
   */
  buildPath(pathVars?: string | string[] | { [key: string]: string }): string;

  /**
   * Delete all or multiple Data Transfer Objects
   * @param {unknown[]} entities A single or multiple entities to delete
   * @returns {void}
   */
  delete(entities?: unknown[]): Promise<void>;

  /**
   * Retrieve all Data Transfer Objects
   * @returns {Dto[]} The list of Data Transfer Objects
   */
  getAll(
    pathVars?: string | string[] | { [key: string]: string },
  ): Promise<Dto[]>;

  /**
   * Retrieve the Data Transfer Objects in a paginated manner
   * @returns {Record<number, Dto[]>} The object of pages
   */
  getPaginated(
    pathVars?: string | string[] | { [key: string]: string },
    pagination?: Pagination,
  ): Promise<{
    items: Dto[];
    pagination: Pagination;
  }>;

  /**
   * Query the Data Transfer Objects based on various attributes
   * @param {Query[]} entities A single or multiple queries
   * @param {{ field: string, desc: boolean }} sort How to handle the sorting of the results
   * @returns {void}
   */
  query(queries: Query[], pagination?: Pagination): Promise<Dto[]>;

  /**
   * Update all or multiple Data Transfer Objects
   * @param {unknown[]} entities A single or multiple entities to update
   * @returns {void}
   */
  update(entities?: Dto[]): Promise<void>;

  /**
   * Build an instance of the object in the list
   * @param {Record<string, unknown>} data The given data
   * @returns {Dto} The data object
   */
  _buildInstance(data?: Record<string, unknown>): Dto;
}
