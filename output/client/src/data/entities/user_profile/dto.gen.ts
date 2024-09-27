import { type JSONData } from '../../base_classes/dto.d'
import { Dto } from '../../base_classes/dto'
import type { IUserProfileDto, IAccessLevel, Permission, Language,  } from './dto.gen.d'

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


// Types
export const PermissionValues: Record<string,Permission> = Object.freeze({
  NONE: 'none',
  READ: 'read',
  READ_WRITE: 'read_write',
})
export const LanguageValues: Record<string,Language> = Object.freeze({
  EN: 'en',
  NL: 'nl',
})

export class UserProfileDto extends Dto implements IUserProfileDto {
  /// ///////////////////////////////////////
  /// //////// PRIVATE VARIABLES ////////////
  /// ///////////////////////////////////////
  _path = '/user_profile/{id}'
  _reference = '/UserProfile/{id}'

  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
    
    accessLevel: IAccessLevel = {
      admin: PermissionValues.none,
      airlines: PermissionValues.none,
      apiDocs: PermissionValues.none,
      packingSlips: PermissionValues.none,
      priceConfiguration: PermissionValues.none,
      printingOrders: PermissionValues.none,
      quotation: PermissionValues.none,
      sampleTest: PermissionValues.none,
      userManagement: PermissionValues.none,
    }

    
    
    address = {} satisfies Record<string, any>
    
    
    avatar = new media.Media()
    
    
    avatarUrl = ""
    
    
    email = ""
    
    
    firstName = ""
    
    
    infix = ""
    
    
    isAdmin = false
    
    
    _kind = "UserProfile"
    
    
    language = LanguageValues.nl
    
    
    lastName = ""
    
    
    prefix = ""
    

  /// ///////////////////////////////////////
  /// ///////////// METHODS /////////////////
  /// ///////////////////////////////////////

  static FromJSON(data: unknown): UserProfileDto {
    // Create a new instance of the class
    const obj = new UserProfileDto()
    obj.set(data as JSONData)
    return obj
  }

  static FromJSONArray(data: unknown[]): UserProfileDto[] {
    return data.map((item) => UserProfileDto.FromJSON(item))
  }

  override set(data: JSONData): void {
    super.set(data)
 
    // Implement variables here
    if (data.accessLevel !== undefined && data.accessLevel !== null ) {
      if (data.accessLevel) {
        this.accessLevel = AccessLevel.FromJSON(data.accessLevel)
      }
    }
    if (data.address !== undefined && data.address !== null ) {
      this.address = data.address as Record<string, any> ?? this.address
      
    }
    if (data.avatar !== undefined && data.avatar !== null ) {
      if (!this.avatar) {
        this.avatar = new media.Media()
      }
      this.avatar.set(data.avatar as JSONData ?? this.avatar.toJSON())
    }
    if (data.avatarUrl !== undefined && data.avatarUrl !== null ) {
      this.avatarUrl = data.avatarUrl as string ?? this.avatarUrl
      
    }
    if (data.email !== undefined && data.email !== null ) {
      this.email = data.email as string ?? this.email
      
    }
    if (data.firstName !== undefined && data.firstName !== null ) {
      this.firstName = data.firstName as string ?? this.firstName
      
    }
    if (data.infix !== undefined && data.infix !== null ) {
      this.infix = data.infix as string ?? this.infix
      
    }
    if (data.isAdmin !== undefined && data.isAdmin !== null ) {
      this.isAdmin = data.isAdmin as boolean ?? this.isAdmin
      
    }
    if (data._kind !== undefined && data._kind !== null ) {
      this._kind = data._kind as string ?? this._kind
      
    }
    if (data.language !== undefined && data.language !== null ) {
        this.language = LanguageValues[(data?.language as string)?.toUpperCase() ?? this.language.toUpperCase()]
    }
    if (data.lastName !== undefined && data.lastName !== null ) {
      this.lastName = data.lastName as string ?? this.lastName
      
    }
    if (data.prefix !== undefined && data.prefix !== null ) {
      this.prefix = data.prefix as string ?? this.prefix
      
    }
  }

}

export class AccessLevel implements IAccessLevel {
  /// ///////////////////////////////////////
  /// ///////// PUBLIC VARIABLES ////////////
  /// ///////////////////////////////////////
    
      admin = PermissionValues.none
    
    
      airlines = PermissionValues.none
    
    
      apiDocs = PermissionValues.none
    
    
      packingSlips = PermissionValues.none
    
    
      priceConfiguration = PermissionValues.none
    
    
      printingOrders = PermissionValues.none
    
    
      quotation = PermissionValues.none
    
    
      sampleTest = PermissionValues.none
    
    
      userManagement = PermissionValues.none
    

  /// ///////////////////////////////////////
  /// ///////////// METHODS /////////////////
  /// ///////////////////////////////////////

  static FromJSON(data: unknown): AccessLevel {
    // Create a new instance of the class
    const obj = new AccessLevel()
    obj.set(data as JSONData)
    return obj
  }

  static FromJSONArray(data: unknown[]): AccessLevel[] {
    return data.map((item) => AccessLevel.FromJSON(item))
  }


  set(data: JSONData): void {
    // Implement variables here
    if (data.admin !== undefined && data.admin !== null ) {
        this.admin = PermissionValues[(data?.admin as string)?.toUpperCase() ?? this.admin.toUpperCase()]
    }
    if (data.airlines !== undefined && data.airlines !== null ) {
        this.airlines = PermissionValues[(data?.airlines as string)?.toUpperCase() ?? this.airlines.toUpperCase()]
    }
    if (data.apiDocs !== undefined && data.apiDocs !== null ) {
        this.apiDocs = PermissionValues[(data?.apiDocs as string)?.toUpperCase() ?? this.apiDocs.toUpperCase()]
    }
    if (data.packingSlips !== undefined && data.packingSlips !== null ) {
        this.packingSlips = PermissionValues[(data?.packingSlips as string)?.toUpperCase() ?? this.packingSlips.toUpperCase()]
    }
    if (data.priceConfiguration !== undefined && data.priceConfiguration !== null ) {
        this.priceConfiguration = PermissionValues[(data?.priceConfiguration as string)?.toUpperCase() ?? this.priceConfiguration.toUpperCase()]
    }
    if (data.printingOrders !== undefined && data.printingOrders !== null ) {
        this.printingOrders = PermissionValues[(data?.printingOrders as string)?.toUpperCase() ?? this.printingOrders.toUpperCase()]
    }
    if (data.quotation !== undefined && data.quotation !== null ) {
        this.quotation = PermissionValues[(data?.quotation as string)?.toUpperCase() ?? this.quotation.toUpperCase()]
    }
    if (data.sampleTest !== undefined && data.sampleTest !== null ) {
        this.sampleTest = PermissionValues[(data?.sampleTest as string)?.toUpperCase() ?? this.sampleTest.toUpperCase()]
    }
    if (data.userManagement !== undefined && data.userManagement !== null ) {
        this.userManagement = PermissionValues[(data?.userManagement as string)?.toUpperCase() ?? this.userManagement.toUpperCase()]
    }
  }
}
