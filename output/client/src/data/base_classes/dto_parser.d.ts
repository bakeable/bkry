import type IUserDto from "./user_dto.d";

export type Parser = {
  customParser?: (value: unknown, data?: unknown) => unknown; // The custom parser
  field: string; // The field to parse
  parser: string; // 'custom' | 'string' | 'number' | 'integer' | 'float' | 'boolean' | 'date' | 'hide', // Choose a standard parser
};

export default interface IDtoParser {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /**
   * The validators to run
   */
  parsers: Parser[];

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  /**
   * Parse the object
   * @param {Record<string, unknown>} data The data to be parsed
   * @returns {Record<string, unknown> }  The parsed object
   */
  parse(data: Record<string, unknown>): Record<string, unknown>;

  /**
   * Parse a specific field
   * @param {unknown} value The value to parse
   * @param {Parser} parserOptions The options in parsing the value
   * @returns {unknown}  The parsed value
   */
  parseField(value: unknown, parserOptions: Parser): unknown;

  /**
   * A set of standard parsers
   * @param {string} parser The parser to use
   * @param {unknown} value The value to parse
   * @returns {unknown }  The parsed value
   */
  standardParser(parser: string, value: unknown): unknown;
}
