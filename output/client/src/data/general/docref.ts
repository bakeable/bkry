import {
  entityBuilder,
  entityCollectionNameToEntityName,
  entityNameToTypeName,
} from "../entities";

// Type alias for the IDocRef using the conditional type
type IDocRef<T extends EntityTypeName> = T extends keyof IEntityClassMap
  ? DocRef<T>
  : never;

export class DocRef<T extends EntityTypeName> {
  constructor(data?: unknown) {
    this.set(data);
  }
  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
  /**
   * The path to the document in Firestore
   */
  document_path: string = "";

  /**
   * The type of the document
   * */
  type?: T;

  /**
   * The class of the document
   * */
  class?: IEntityClassMap[T];

  /// //////////////////////////////
  /// ///////// METHODS ////////////
  /// //////////////////////////////

  /**
   * Determine the type and class based on the document_path
   */
  private determineTypeAndClass(): void {
    // Example logic to determine type based on document_path
    const [_, entityName] =
      Object.entries(entityCollectionNameToEntityName).find(([key, _]) =>
        this.document_path.includes("/" + key + "/"),
      ) || [];

    // Gather the path variables by taking the even indexes of the splitted string
    const pathVars = this.document_path
      .split("/")
      .filter((x, i) => i % 2 === 0 && !!x);

    // Based on the entityName, set the type and class
    if (entityName && pathVars.length > 0) {
      this.type = entityNameToTypeName[entityName] as T;
      this.class = entityBuilder[entityName](pathVars, {
        id: pathVars[pathVars.length - 1],
      }) as IEntityClassMap[T];
    }
  }

  /**
   * Set the AuditInfo's properties
   */
  set(data: unknown): void {
    if (typeof data === "string") {
      this.document_path = data;
    } else if (
      typeof data === "object" &&
      data?.hasOwnProperty("document_path")
    ) {
      this.document_path = (data as any)["document_path"];
    }
    this.determineTypeAndClass();
  }
  /**
   * Parse to JSON
   */
  toJSON(): string {
    return this.document_path;
  }
}
