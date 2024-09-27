import { type JSONData } from '../../base_classes/dto.d'
import { Dto } from '../../base_classes/dto'
import type { IUserDto, IAddress, ILocale, IProfile, Gender,  } from './dto.gen.d'

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


// Types
export const GenderValues: Record<string,Gender> = Object.freeze({
  UNKNOWN: 'Unknown',
  MALE: 'Male',
  FEMALE: 'Female',
})

export class UserDto extends Dto implements IUserDto {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  _path = '/user/{id}'
  _reference = '/User/{id}'

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
    
    address: IAddress = {
      city: "",
      country_code: "",
      line1: "",
      line2: "",
      postal_code: "",
      state: "",
      street: "",
    }

    
    
    avatar = new media.Media()
    
    
    email = ""
    
    
    locale: ILocale = {
      code: "nl-NL",
      two_digit_code: "nl",
    }

    
    
    mute_sound = false
    
    
    profile: IProfile = {
      date_of_birth: "",
      first_name: "",
      gender: GenderValues.Unknown,
      infix: "",
      is_grown_up: false,
      last_name: "",
      prefix: "",
    }

    
    
    settings = ""
    
    
    user_name = ""
    

  /// ///////////////////////////////////////
  /// ///////////// METHODS /////////////////
  /// ///////////////////////////////////////

  static FromJSON(data: unknown): UserDto {
    // Create a new instance of the class
    const obj = new UserDto()
    obj.set(data as JSONData)
    return obj
  }

  static FromJSONArray(data: unknown[]): UserDto[] {
    return data.map((item) => UserDto.FromJSON(item))
  }

  override set(data: JSONData): void {
    super.set(data)
 
    // Implement variables here
    if (data.address !== undefined && data.address !== null ) {
      if (data.address) {
        this.address = Address.FromJSON(data.address)
      }
    }
    if (data.avatar !== undefined && data.avatar !== null ) {
      if (!this.avatar) {
        this.avatar = new media.Media()
      }
      this.avatar.set(data.avatar as JSONData ?? this.avatar.toJSON())
    }
    if (data.email !== undefined && data.email !== null ) {
      this.email = data.email as string ?? this.email
      
    }
    if (data.locale !== undefined && data.locale !== null ) {
      if (data.locale) {
        this.locale = Locale.FromJSON(data.locale)
      }
    }
    if (data.mute_sound !== undefined && data.mute_sound !== null ) {
      this.mute_sound = data.mute_sound as boolean ?? this.mute_sound
      
    }
    if (data.profile !== undefined && data.profile !== null ) {
      if (data.profile) {
        this.profile = Profile.FromJSON(data.profile)
      }
    }
    if (data.settings !== undefined && data.settings !== null ) {
      this.settings = data.settings as string ?? this.settings
      
    }
    if (data.user_name !== undefined && data.user_name !== null ) {
      this.user_name = data.user_name as string ?? this.user_name
      
    }
  }

}

export class Address implements IAddress {
  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
    
      city = ""
    
    
      country_code = ""
    
    
      line1 = ""
    
    
      line2 = ""
    
    
      postal_code = ""
    
    
      state = ""
    
    
      street = ""
    

  /// ///////////////////////////////////////
  /// ///////////// METHODS /////////////////
  /// ///////////////////////////////////////

  static FromJSON(data: unknown): Address {
    // Create a new instance of the class
    const obj = new Address()
    obj.set(data as JSONData)
    return obj
  }

  static FromJSONArray(data: unknown[]): Address[] {
    return data.map((item) => Address.FromJSON(item))
  }


  set(data: JSONData): void {
    // Implement variables here
    if (data.city !== undefined && data.city !== null ) {
      this.city = data.city as string ?? this.city
      
    }
    if (data.country_code !== undefined && data.country_code !== null ) {
      this.country_code = data.country_code as string ?? this.country_code
      
    }
    if (data.line1 !== undefined && data.line1 !== null ) {
      this.line1 = data.line1 as string ?? this.line1
      
    }
    if (data.line2 !== undefined && data.line2 !== null ) {
      this.line2 = data.line2 as string ?? this.line2
      
    }
    if (data.postal_code !== undefined && data.postal_code !== null ) {
      this.postal_code = data.postal_code as string ?? this.postal_code
      
    }
    if (data.state !== undefined && data.state !== null ) {
      this.state = data.state as string ?? this.state
      
    }
    if (data.street !== undefined && data.street !== null ) {
      this.street = data.street as string ?? this.street
      
    }
  }
}

export class Locale implements ILocale {
  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
    
      code = "nl-NL"
    
    
      two_digit_code = "nl"
    

  /// ///////////////////////////////////////
  /// ///////////// METHODS /////////////////
  /// ///////////////////////////////////////

  static FromJSON(data: unknown): Locale {
    // Create a new instance of the class
    const obj = new Locale()
    obj.set(data as JSONData)
    return obj
  }

  static FromJSONArray(data: unknown[]): Locale[] {
    return data.map((item) => Locale.FromJSON(item))
  }


  set(data: JSONData): void {
    // Implement variables here
    if (data.code !== undefined && data.code !== null ) {
      this.code = data.code as string ?? this.code
      
    }
    if (data.two_digit_code !== undefined && data.two_digit_code !== null ) {
      this.two_digit_code = data.two_digit_code as string ?? this.two_digit_code
      
    }
  }
}

export class Profile implements IProfile {
  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
    
      date_of_birth = ""
    
    
      first_name = ""
    
    
      gender = GenderValues.Unknown
    
    
      infix = ""
    
    
      is_grown_up = false
    
    
      last_name = ""
    
    
      prefix = ""
    

  /// ///////////////////////////////////////
  /// ///////////// METHODS /////////////////
  /// ///////////////////////////////////////

  static FromJSON(data: unknown): Profile {
    // Create a new instance of the class
    const obj = new Profile()
    obj.set(data as JSONData)
    return obj
  }

  static FromJSONArray(data: unknown[]): Profile[] {
    return data.map((item) => Profile.FromJSON(item))
  }


  set(data: JSONData): void {
    // Implement variables here
    if (data.date_of_birth !== undefined && data.date_of_birth !== null ) {
      this.date_of_birth = data.date_of_birth as string ?? this.date_of_birth
      
    }
    if (data.first_name !== undefined && data.first_name !== null ) {
      this.first_name = data.first_name as string ?? this.first_name
      
    }
    if (data.gender !== undefined && data.gender !== null ) {
        this.gender = GenderValues[(data?.gender as string)?.toUpperCase() ?? this.gender.toUpperCase()]
    }
    if (data.infix !== undefined && data.infix !== null ) {
      this.infix = data.infix as string ?? this.infix
      
    }
    if (data.is_grown_up !== undefined && data.is_grown_up !== null ) {
      this.is_grown_up = data.is_grown_up as boolean ?? this.is_grown_up
      
    }
    if (data.last_name !== undefined && data.last_name !== null ) {
      this.last_name = data.last_name as string ?? this.last_name
      
    }
    if (data.prefix !== undefined && data.prefix !== null ) {
      this.prefix = data.prefix as string ?? this.prefix
      
    }
  }
}
