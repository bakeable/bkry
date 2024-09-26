import { type JSONData } from '../../base_classes/dto.d'
import { UserDto } from './dto.gen'
import { type IUser } from './entity.d'
import UserParsers from './parsers'
import UserValidators from './custom_validators'

export class User extends UserDto implements IUser {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, UserParsers, UserValidators)

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
  ): Promise<User> {
    return (await super.get(pathVars)) as User
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): User {
    return new User(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): User[] {
    return data.map((item) => User.FromJSON(item))
  }

  static FromDto(dto: UserDto): User {
    return new User(undefined, dto.toJSON())
  }
}

