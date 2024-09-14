import { type JSONData } from '../../base_classes/dto.d'
import { PriceLayerDto } from './dto.gen'
import { type IPriceLayer } from './entity.d'
import PriceLayerParsers from './parsers'
import PriceLayerValidators from './custom_validators'

export class PriceLayer extends PriceLayerDto implements IPriceLayer {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, PriceLayerParsers, PriceLayerValidators)

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
  ): Promise<PriceLayer> {
    return (await super.get(pathVars)) as PriceLayer
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): PriceLayer {
    return new PriceLayer(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): PriceLayer[] {
    return data.map((item) => PriceLayer.FromJSON(item))
  }

  static FromDto(dto: PriceLayerDto): PriceLayer {
    return new PriceLayer(undefined, dto.toJSON())
  }
}

