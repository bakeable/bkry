
export interface IQuery {
  field: string;
  operator: Operator;
  value: unknown;
}


/**
 * Represents a page information configuration.
 */
export class Query implements IQuery {
  constructor(field: string, operator: Operator, value: unknown) {
    this.field = field;
    this.operator = operator;
    this.value = value;
  }

  /// ////////////////////////////////////////////////
  /// ////////////// PUBLIC VARIABLES ////////////////
  /// ////////////////////////////////////////////////

  /**
   * The field to query.
   */
  field: string = "id";

  /**
   * The operator to use for the query.
   */
  operator: Operator = Operator.EQUALS;

  /**
   * The value to compare against.
   */
  value: unknown;

  /// //////////////////////////////////////////////
  /// /////////////// METHODS //////////////////////
  /// //////////////////////////////////////////////

  /*
   * Converts the query to a cache key.
   */
  toCacheKey(): string {
    return [this.field, this.operator, this.value].join("-");
  }

  /**
   * Converts the query configuration to query parameters.
   * @returns The query parameters as a string
   */
  toQueryParams(): string {
    return JSON.stringify({
      field: this.field,
      operator: this.operator,
      value: this.value,
    });
  }

  /// //////////////////////////////////////////////
  /// ////////////// STATIC METHODS ////////////////
  /// //////////////////////////////////////////////

  /**
   * Converts an array of type IQuery to a string cache key.
   * @param {IQuery[]} queries - The array to convert.
   * @returns The converted string
   */
  static ToCacheKey(queries: IQuery[]): string {
    return queries
      .map((q) => [q.field, q.operator, q.value].join("-"))
      .join("-");
  }

  /**
   * Converts an array of Query objects to a string of query parameters.
   */
  static ToQueryParams(queries: Query[]): string {
    return JSON.stringify(queries);
  }
}

/**
 * Represents the available operators for a query.
 */
export enum Operator {
  EQUALS = "=",
  LESS_THAN = "<",
  GREATER_THAN = ">",
  LESS_THAN_OR_EQUAL = "<=",
  GREATER_THAN_OR_EQUAL = ">=",
  HAS_ANCESTOR = "HAS_ANCESTOR",
  NOT_EQUALS = "!=",
  IN = "IN",
  NOT_IN = "NOT_IN",
}
