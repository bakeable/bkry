import {
  type ClassConstructor,
  Expose,
  instanceToPlain,
  plainToInstance,
} from "class-transformer";
import { formatString, isFalsy } from "@/utils";
import request from "@/data/api/request";
import { type AxiosError } from "axios";
import type IDto from "./dto.d";
import { type Validator } from "./dto_validator.d";
import { type Parser } from "./dto_parser.d";
import DtoValidator from "./dto_validator";
import DtoParser from "./dto_parser";
import { ValidationError } from "./validation_error";
import { cleanJSON } from "@/utils";
import { type JSONData } from "./dto.d";

@Expose()
export class Dto implements IDto {
  constructor(
    data?: JSONData,
    parsers: Parser[] = [],
    validators: Validator[] = [],
  ) {
    if (data !== undefined) {
      // Set data
      this.set(data);
    }

    this._parser = new DtoParser(parsers);
    this._validator = new DtoValidator(validators);
  }

  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  _hasBeenRetrieved = false;
  _isNew = true;
  _parser: DtoParser;
  _path = "example/{exampl_id}/path";
  _reference = "examples/{example_id}/path";
  _pathVars: {
    [key: string]: string;
  } = {};
  _page = 1;
  _retrievedTimestamp = 0;
  _updateKey = "";
  _validator: DtoValidator;

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
  id = "";

  /// ///////////////////////////////////////
  /// ////////////  METHODS /////////////////
  /// ///////////////////////////////////////

  get getPathKeys(): string[] {
    // Get path keys
    return (this._path.match(/{[^}]*}/g) ?? []).map((key) =>
      key.replace("{", "").replace("}", ""),
    );
  }

  setPathVars(pathVars: string[]): void {
    this._pathVars = this.getPathKeys.reduce(
      (acc, key, i) => {
        acc[key] = pathVars[i];
        return acc;
      },
      {} as {
        [key: string]: string;
      },
    );
  }

  buildPathObject(
    pathVars?:
      | string
      | string[]
      | {
          [key: string]: string;
        },
  ): {
    [key: string]: string;
  } {
    // Failsafe
    if (typeof pathVars === "string") {
      pathVars = [pathVars];
    } else if (pathVars === undefined || pathVars === null) {
      pathVars = {};
    } else if (typeof pathVars === "object" && !Array.isArray(pathVars)) {
      this.set(pathVars);
    }

    // Build the path object
    this._pathVars = this.getPathKeys.reduce(
      (acc, key, i) => {
        if (typeof pathVars === "object" && !Array.isArray(pathVars)) {
          if (key !== "id" && !pathVars[key] && !this.getProperty(key)) {
            throw new Error(`Missing path variable in object: ${key}`);
          } else {
            acc[key] = pathVars[key] ?? this.getProperty(key);
          }
        } else if (key !== "id" && pathVars.length < i + 1) {
          throw new Error(`Missing path variable in array: ${key}`);
        } else {
          acc[key] = pathVars[i];
        }

        return acc;
      },
      {} as {
        [key: string]: string;
      },
    );

    // Apply _pathVars to object too
    for (const [key, value] of Object.entries(this._pathVars)) {
      if (key !== undefined && key in this && value !== undefined) {
        (this as any)[key] = value;
      }
    }

    return this._pathVars;
  }

  buildPath(
    pathVars?:
      | string
      | string[]
      | {
          [key: string]: string;
        },
  ): string {
    let path = this._path.slice(0, this._path.length);

    // Format path variables
    const pathObject = this.buildPathObject(pathVars);

    // Replace path variables
    for (const [key, value] of Object.entries(pathObject)) {
      if (key === "id" && !value) {
        path = path.replace("/{id}", "");
      } else {
        path = path.replace(`{${key}}`, value);
      }
    }

    return path;
  }

  buildReference(): string {
    let reference = this._reference.slice(0, this._reference.length);

    // Replace path variables
    const pathObject = this.buildPathObject();

    // Replace path variables
    for (const [key, value] of Object.entries(pathObject)) {
      if (key === "id" && !value) {
        reference = reference.replace("/{id}", "");
      } else {
        reference = reference.replace(`{${key}}`, value);
      }
    }

    return reference;
  }

  public async get(
    pathVars?: string | string[] | { [key: string]: string },
  ): Promise<Dto> {
    // Build path
    const path = this.buildPath(pathVars);

    // Get snapshot
    const { item }: { item: JSONData } = await request({
      url: path,
      method: "get",
    });

    // Exists
    if (typeof item !== "object") {
      throw new Error("Incorrect data");
    } else if (item === null) {
      console.info("No item found");
      return this;
    }

    // Parse data
    const parsedData = this._parser.parse(item);

    // Get data
    this.set(parsedData);

    // Set metadata
    this._hasBeenRetrieved = true;
    this._isNew = false;
    this._retrievedTimestamp = Date.now();

    // Return result
    return this;
  }

  public getParentIds(): string[] {
    return Object.entries(this._pathVars)
      .filter(([_key, value]) => {
        return value !== this.id;
      })
      .map(([_key, value]) => value);
  }

  public set(data: JSONData): void {
    // Apply data
    this.id = formatString(data.id) ?? this.id;

    // Update key
    this._updateKey = this.id + "_" + Math.floor(Date.now() * Math.random());
  }

  public async update(
    pathVars?: string | string[] | { [key: string]: string },
  ): Promise<void> {
    // Build path
    const path = this.buildPath(pathVars);

    // Set creation information
    if (
      this._isNew &&
      (this.id === "" || this.id === undefined || this.id === null)
    ) {
    } else {
      this._isNew = false;
    }

    // Validate data
    const data = this.toJSON();
    if (!this._validator.validate(data)) {
      throw new ValidationError(
        "Data could not be validated",
        this._validator.errors,
      );
    }

    // Parse data
    const parsedData = this._parser.parse(data);
    this.set(parsedData);

    // Update document
    const entity: JSONData = await request({
      url: path,
      method: this._isNew ? "post" : "put",
      data: parsedData,
    });

    // Update object locally
    if (!isFalsy(entity)) {
      this.set(entity);

      // Object is saved
      this._isNew = false;
    }
  }

  public async delete(
    pathVars?: string | string[] | { [key: string]: string },
  ): Promise<void> {
    // Build path
    const path = this.buildPath(pathVars);

    // Delete document
    await request({
      url: path,
      method: "delete",
      data: this.toJSON(),
    }).catch((error) => {
      console.error(error);
      if ((error as AxiosError).response !== null) {
        throw (error as AxiosError).response?.data;
      }
    });
  }

  public async activate(userId?: string): Promise<void> {
    console.log("Not implemented", userId);
  }

  public fresh(): Dto {
    return new Dto({}, this._parser.parsers, this._validator.validators);
  }

  public toJSON(): JSONData {
    return instanceToPlain(this, {
      excludePrefixes: ["_"],
      excludeExtraneousValues: false,
    });
  }

  public toCleanJSON(): JSONData {
    const json = instanceToPlain(this, {
      excludePrefixes: ["_"],
      excludeExtraneousValues: false,
    });

    // Remove null, undefined or empty strings
    return cleanJSON(json) as JSONData;
  }

  public toType(classType: ClassConstructor<never>): never {
    return plainToInstance(classType, this, { excludeExtraneousValues: true });
  }

  public getProperty(property: string): unknown {
    return (this as any)[property];
  }

  public async translate(locale: string): Promise<Dto> {
    // Delete document
    const { item }: { item: JSONData } = await request({
      url: "/openai/translate",
      method: "post",
      data: {
        locale,
        data: this.toJSON(),
      },
    });

    // Update object locally
    if (!isFalsy(item)) {
      this.set(item);
    }

    return this;
  }
}
