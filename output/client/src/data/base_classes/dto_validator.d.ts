import type IUserDto from "./user_dto.d";

export type Validator = {
  customValidator?: (value: unknown, data?: unknown) => boolean | string; // The custom validator
  errorMessage?: string; // The standard error message to show
  field: string; // The field to validate
  parse?: boolean; // Whether to parse the field if possible
  validator: string; // 'custom' | 'string' | 'email' | 'number' | 'integer' | 'float' | 'not_falsy' | 'date', // Choose a standard validator
};

export default interface IDtoValidator {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /**
   * The errors found in validation
   */
  errors: {
    field: string;
    message: string;
  }[];

  /**
   * Whether the supplied data is valid
   */
  valid: boolean;

  /**
   * The validators to run
   */
  validators: Validator[];

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  /**
   * Validate the object
   * @param {Record<string, unknown>} data The data to be validated
   * @returns {boolean} Whether the object is valid
   */
  validate(data: Record<string, unknown>): boolean;

  /**
   * A set of standard validators
   * @param {string} validator The validator to use
   * @param {unknown} value The value to validate
   * @returns {boolean | string} Returns valid or error message
   */
  standardValidator(parser: string, value: unknown): boolean | string;
}
