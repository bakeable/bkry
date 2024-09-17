import { type JSONData } from '../../base_classes/dto.d'
import { UserProfileDto } from './dto.gen'
import { type IUserProfile } from './entity.d'
import UserProfileParsers from './parsers'
import UserProfileValidators from './custom_validators'

export class UserProfile extends UserProfileDto implements IUserProfile {
  constructor(pathVars?: string | string[] | {
    id?: string,
  }, data?: JSONData) {
    super(data, UserProfileParsers, UserProfileValidators)

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
  ): Promise<UserProfile> {
    return (await super.get(pathVars)) as UserProfile
  }

  /// ///////////////////////////////////////
  /// /////////// STATIC METHODS ////////////
  /// ///////////////////////////////////////

  static override FromJSON(data?: JSONData): UserProfile {
    return new UserProfile(undefined, data)
  }

  static override FromJSONArray(data: JSONData[]): UserProfile[] {
    return data.map((item) => UserProfile.FromJSON(item))
  }

  static FromDto(dto: UserProfileDto): UserProfile {
    return new UserProfile(undefined, dto.toJSON())
  }
}

