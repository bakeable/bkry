import { type JSONData } from '../../base_classes/dto.d'
import { AirlineOrderDto } from './dto.gen'
import { type IAirlineOrder } from './entity.d'
import AirlineOrderParsers from './parsers'
import AirlineOrderValidators from './custom_validators'

export class AirlineOrder extends AirlineOrderDto implements IAirlineOrder {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, AirlineOrderParsers, AirlineOrderValidators)

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
  ): Promise<AirlineOrder> {
    return (await super.get(pathVars)) as AirlineOrder
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): AirlineOrder {
    return new AirlineOrder(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): AirlineOrder[] {
    return data.map((item) => AirlineOrder.FromJSON(item))
  }

  static FromDto(dto: AirlineOrderDto): AirlineOrder {
    return new AirlineOrder(undefined, dto.toJSON())
  }
}

