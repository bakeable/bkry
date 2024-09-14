import { type JSONData } from '../../base_classes/dto.d'
import { PortalEventDto } from './dto.gen'
import { type IPortalEvent } from './entity.d'
import PortalEventParsers from './parsers'
import PortalEventValidators from './custom_validators'

export class PortalEvent extends PortalEventDto implements IPortalEvent {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, PortalEventParsers, PortalEventValidators)

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
  ): Promise<PortalEvent> {
    return (await super.get(pathVars)) as PortalEvent
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): PortalEvent {
    return new PortalEvent(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): PortalEvent[] {
    return data.map((item) => PortalEvent.FromJSON(item))
  }

  static FromDto(dto: PortalEventDto): PortalEvent {
    return new PortalEvent(undefined, dto.toJSON())
  }
}

