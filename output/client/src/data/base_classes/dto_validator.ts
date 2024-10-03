import { formatString, isFalsy } from "@/utils";
import type IDtoValidator from "./dto_validator.d";
import { type Validator } from "./dto_validator.d";
import { type JSONData } from "./dto.d";

export default class DtoValidator implements IDtoValidator {
  constructor(validators: Validator[]) {
    this.validators = validators;
  }

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
  errors: Array<{
    field: string;
    message: string;
  }> = [];

  valid = false;
  validators: Validator[] = [
    {
      field: "created",
      validator: "date",
    },
  ];

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////
  validate(data: JSONData): boolean {
    // Empty error
    this.errors = [];

    // Loop through validators
    this.validators.forEach((v) => {
      // Retrieve value
      const value = this.getValue(v.field, data);
      let valid: boolean | string = true;
      if (v.validator === "custom" && v.customValidator !== undefined) {
        valid = v.customValidator(value, data);
      } else {
        valid = this.standardValidator(v.validator, value);
      }

      if (valid !== true) {
        this.errors.push({
          field: v.field,
          message:
            "error:" + (typeof valid === "string" ? valid : v.errorMessage),
        });
      }
    });

    // Determine whether data is valid
    this.valid = this.errors.length === 0;
    return this.valid;
  }

  getValue(field: string, data: JSONData): unknown {
    const path = field.split(".");
    let value: unknown = data;

    for (const key of path) {
      if (value !== undefined && typeof value === "object") {
        value = (value as Record<string, unknown>)[key];
      } else {
        value = undefined;
        break;
      }
    }

    return value;
  }

  standardValidator(validator: string, value: unknown): boolean | string {
    const strValue = formatString(value) ?? "";
    switch (validator) {
      case "date":
        return value instanceof Date ? true : "invalid_date";
      case "float":
        return value === parseFloat(strValue) ? true : "invalid_float";
      case "integer":
        return value === parseInt(strValue) ? true : "invalid_integer";
      case "not_falsy":
        return !isFalsy(value) ? true : "invalid_falsy";
      case "number":
        return typeof value === "number" ? true : "invalid_number";
      case "string":
        return typeof value === "string" ? true : "invalid_string";
      case "email":
        return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(strValue)
          ? true
          : "invalid_email";

      default:
        return true;
    }
  }
}
