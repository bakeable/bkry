import { type JSONData } from '../../base_classes/dto.d'
import { QuotationDto } from './dto.gen'
import { type IQuotation } from './entity.d'
import QuotationParsers from './parsers'
import QuotationValidators from './custom_validators'

export class Quotation extends QuotationDto implements IQuotation {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, QuotationParsers, QuotationValidators)

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
  ): Promise<Quotation> {
    return (await super.get(pathVars)) as Quotation
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): Quotation {
    return new Quotation(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): Quotation[] {
    return data.map((item) => Quotation.FromJSON(item))
  }

  static FromDto(dto: QuotationDto): Quotation {
    return new Quotation(undefined, dto.toJSON())
  }
}

