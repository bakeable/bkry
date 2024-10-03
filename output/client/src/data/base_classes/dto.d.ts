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

import DtoValidator from "./dto_validator";
export type JSONData = {
  [key: string]:
    | JSONData
    | JSONData[]
    | string
    | number
    | boolean
    | string[]
    | number[]
    | boolean[];
};
export default interface IDto {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////

  /**
   * Whether the object has been retrieved from the database
   */
  _hasBeenRetrieved: boolean;

  /**
   * Whether the current object is newly created
   */
  _isNew: boolean;

  /**
   * The parser to use to parse data transfer objects
   */
  _parser: DtoParser;

  /**
   * The path to the corresponding API endpoint with variables marked as {?}
   */
  _path: string;

  /**
   * The unformatted reference to the document in the database
   */
  _reference: string;

  /**
   * The path variables in order
   */
  _pathVars: {
    [key: string]: string;
  };

  /**
   * The timestamp of retrieval from the database
   */
  _retrievedTimestamp: number;

  /**
   * A random key that forces the update of the entity
   */
  _updateKey: string;

  /**
   * The validator to use before updating entities
   */
  _validator: DtoValidator;

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
  /**
   * The document ID by of the corresponding document was created in the database
   */
  id: string;

  /// ///////////////////////////////////////
  /// ////////////  METHODS /////////////////
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
   * Retrieves the corresponding object from the database
   * @param {string | string[]} pathVars A single or multiple path variables
   * @returns {Dto} The data transfer object
   */
  get(pathVars?: string | string[] | { [key: string]: string }): Promise<Dto>;

  /**
   * Sets the data of the corresponding object from the database
   * @param {Record<string, unknown>} data The data to be set on the data transfer object
   * @returns {void} The data transfer object
   */
  set(data: Record<string, unknown>): void;

  /**
   * Updates the corresponding object from the database
   * @param {string | string[]} pathVars A single or multiple path variables
   * @returns {void}
   */
  update(
    pathVars?: string | string[] | { [key: string]: string },
  ): Promise<void>;

  /**
   * Deletes the corresponding object from the database
   * @param {string | string[]} pathVars A single or multiple path variables
   * @returns {void}
   */
  delete(pathVars: string | string[]): Promise<void>;

  /**
   * Converts the current object to JSON
   * @returns {Record<string, unknown>} The formatted JSON
   */
  toJSON(): Record<string, unknown>;

  /**
   * Converts the current object to the given type
   * @param {ClassConstructor<never>} classType The type to convert the class to
   * @return {never} The class with data
   */
  toType(classType: ClassConstructor<never>): never;
}
