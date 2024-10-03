/* eslint-disable @typescript-eslint/no-unsafe-argument */
import { formatDate, formatString } from "@/utils";
import type IDtoParser from "./dto_parser.d";
import { type Parser } from "./dto_parser.d";
import { type JSONData } from "./dto.d";

export default class DtoParser implements IDtoParser {
  constructor(parsers: Parser[]) {
    this.parsers = parsers;
  }

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
  parsers: Parser[] = [
    {
      field: "created",
      parser: "date",
    },
  ];

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  getValue(field: string, data: Record<string, unknown>): unknown {
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

  setValue(field: string, data: Record<string, unknown>, value: unknown): void {
    const path = field.split(".");
    let parent = data;

    for (let i = 0; i < path.length; i++) {
      if (i === path.length - 1) {
        parent[path[i]] = value;
      } else {
        if (parent[path[i]] === undefined) {
          parent[path[i]] = {};
        }
        parent = parent[path[i]] as Record<string, unknown>;
      }
    }
  }

  parse(data: JSONData): JSONData {
    const newData: JSONData = {};

    // Map the keys to the new data object
    Object.keys(data).forEach((key) => {
      if (!Object.keys(newData).includes(key)) {
        newData[key] = data[key];
      }
    });

    // Loop through parsers
    this.parsers.forEach((p) => {
      const value = this.getValue(p.field, data);

      // Retrieve value
      const parsedValue = this.parseField(value, p, data);

      // Remove undefined values
      if (parsedValue !== undefined) {
        this.setValue(p.field, newData, parsedValue);
      }
    });

    return newData;
  }

  parseField(value: unknown, parserOptions: Parser, data?: unknown): unknown {
    if (
      parserOptions.parser === "custom" &&
      parserOptions.customParser !== undefined
    ) {
      return parserOptions.customParser(value, data);
    } else {
      return this.standardParser(parserOptions.parser, value);
    }
  }

  standardParser(parser: string, value: unknown): unknown {
    const strValue = formatString(value) ?? "";
    switch (parser) {
      case "date":
        return formatDate(value);
      case "float":
        return parseFloat(strValue);
      case "integer":
        return parseInt(strValue);
      case "number":
        return parseInt(strValue) !== parseFloat(strValue)
          ? parseFloat(strValue)
          : parseInt(strValue);
      case "string":
        return strValue;
      case "hide":
        return undefined;

      default:
        return value;
    }
  }
}
