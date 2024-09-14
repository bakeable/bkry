import { type JSONData } from '../../base_classes/dto.d'
import { EmailDto } from './dto.gen'
import { type IEmail } from './entity.d'
import EmailParsers from './parsers'
import EmailValidators from './custom_validators'

export class Email extends EmailDto implements IEmail {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, EmailParsers, EmailValidators)

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
  ): Promise<Email> {
    return (await super.get(pathVars)) as Email
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): Email {
    return new Email(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): Email[] {
    return data.map((item) => Email.FromJSON(item))
  }

  static FromDto(dto: EmailDto): Email {
    return new Email(undefined, dto.toJSON())
  }
}

