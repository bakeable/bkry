import { type JSONData } from '../../base_classes/dto.d'
import { ProductDto } from './dto.gen'
import { type IProduct } from './entity.d'
import ProductParsers from './parsers'
import ProductValidators from './custom_validators'

export class Product extends ProductDto implements IProduct {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, ProductParsers, ProductValidators)

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
  ): Promise<Product> {
    return (await super.get(pathVars)) as Product
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): Product {
    return new Product(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): Product[] {
    return data.map((item) => Product.FromJSON(item))
  }

  static FromDto(dto: ProductDto): Product {
    return new Product(undefined, dto.toJSON())
  }
}

