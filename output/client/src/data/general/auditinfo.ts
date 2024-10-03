import { type JSONData } from "../base_classes/dto";

export class AuditInfo {
  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
  /**
   * The timestamp when the object was audited.
   */
  timestamp: number = 0;

  /**
   * The date object
   */
  get date(): Date {
    return new Date(this.timestamp * 1000);
  }

  /**
   * The ID of the user who audited the object.
   */
  user_id: string = "";

  /// //////////////////////////////
  /// ///////// METHODS ////////////
  /// //////////////////////////////

  /**
   * Format timestamp as sortable datestring
   */
  toSortableDate(): string {
    return `${this.date.getFullYear()}-${(this.date.getMonth() + 1).toString().padStart(2, "0")}-${this.date.getDate().toString().padStart(2, "0")}`;
  }

  /**
   * Format timestamp as locale datetime string
   */
  toSortableDateTime(): string {
    return `${this.date.getFullYear()}-${(this.date.getMonth() + 1).toString().padStart(2, "0")}-${this.date.getDate().toString().padStart(2, "0")} ${this.date.getHours().toString().padStart(2, "0")}:${this.date.getMinutes().toString().padStart(2, "0")}:${this.date.getSeconds().toString().padStart(2, "0")}`;
  }

  /**
   * Set the AuditInfo's properties
   */
  set(data: unknown): void {
    if (typeof data !== "object") {
      return;
    }
    const json = data as JSONData;
    this.timestamp =
      parseInt((json.timestamp?.toString() ?? "0") as string) ?? 0;
    this.user_id = json.user_id as string;
  }

  /**
   * Parse to JSON
   */
  toJSON(): JSONData {
    return {
      timestamp: this.timestamp,
      user_id: this.user_id,
    };
  }
}
