export class ValidationError extends Error {
  constructor(
    message: string,
    errors: Array<{
      field: string;
      message: string;
    }>,
  ) {
    super(message);
    this.name = "ValidationError";
    this.errors = errors;
  }

  /**
   * The validation errors that were returned
   */
  errors: Array<{
    field: string;
    message: string;
  }> = [];
}
