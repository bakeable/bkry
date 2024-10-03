
import {type Gender, } from './enums.gen'
import { type DocRef } from '../../general/docref'
import { AuditInfo } from '../../general/auditinfo'

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


// User Interface
export interface User {
  // The ID of the User
  id: string

  // The audit information on the creation of the User
  created: AuditInfo

  // The audit information on the modification of the User
  modified: AuditInfo

  // The address of the user
  address: Address

  // Reference to the user's avatar media
  avatar_media: DocRef

  // The actual e-mail of the user
  email: string

  // The preferred locale settings of the user
  locale: Locale

  // Indicates whether the sound is muted for the user
  mute_sound: boolean

  // The profile information of the user
  profile: Profile

  // The user's settings
  settings: string

  // A nick-name for the user to use in games
  user_name: string
}



// Property Interfaces
export interface Address {
    // The city of the address
      city: string
    // The country code of the address
      country_code: string
    // The first line of the address
      line1: string
    // The second line of the address
      line2: string
    // The postal code of the address
      postal_code: string
    // The state of the address
      state: string
    // The street of the address
      street: string
}
export interface Locale {
    // The preferred locale extended with accent
      code: string
    // Fallback on this locale if the preferred accent is not available
      two_digit_code: string
}
export interface Profile {
    // The user's date of birth
      date_of_birth: string
    // The user's first name
      first_name: string
    // The user's gender
      gender: Gender
    // The user's infix
      infix: string
    // Indicates whether the user is an adult
      is_grown_up: boolean
    // The user's last name
      last_name: string
    // The user's prefix
      prefix: string
}
