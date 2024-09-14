import { type JSONData } from '../../base_classes/dto.d'
import { DeliveryEntryDto } from './dto.gen'
import { type IDeliveryEntry } from './entity.d'
import DeliveryEntryParsers from './parsers'
import DeliveryEntryValidators from './custom_validators'

export class DeliveryEntry extends DeliveryEntryDto implements IDeliveryEntry {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, DeliveryEntryParsers, DeliveryEntryValidators)

    if (data !== undefined) {
      // Set data
      this.set(data)
    }

    this.buildPathObject(pathVars);
  }
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////

  /// ///////////////////////////////////////
  /// ////////////// METHODS ////////////////
  /// ///////////////////////////////////////

  public async get(
    pathVars?: string | string[] | {
    id?: string,
  } | undefined
  ): Promise<DeliveryEntry> {
    return (await super.get(pathVars)) as DeliveryEntry
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): DeliveryEntry {
    return new DeliveryEntry(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): DeliveryEntry[] {
    return data.map((item) => DeliveryEntry.FromJSON(item))
  }

  static FromDto(dto: DeliveryEntryDto): DeliveryEntry {
    return new DeliveryEntry(undefined, dto.toJSON())
  }
}

