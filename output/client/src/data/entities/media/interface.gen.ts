
import {type ContentType, } from './enums.gen'
import { type DocRef } from '../../general/docref'
import { AuditInfo } from '../../general/auditinfo'

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


// Media Interface
export interface Media {
  // The ID of the Media
  id: string

  // The audit information on the creation of the Media
  created: AuditInfo

  // The audit information on the modification of the Media
  modified: AuditInfo

  // DocumentReference to the Author
  author: DocRef

  // The type of content: Image, Video, or Audio
  content_type: ContentType

  // The description of the media file
  description: Description

  // The file extension
  extension: string

  // The array of language codes to which this media applies
  language_ids: string[]

  // The official file mimetype
  mime_type: string

  // Indicates if the media file is multilingual
  multilingual: boolean

  // The size of the file in Kb
  size: string

  // The Firebase Storage location
  storage_path: string

  // The public URL to get media directly from the internet
  url: string
}



// Property Interfaces
export interface Description {
    // The default locale of the description
      default_locale: string
    // The default description of the Media file
      default_value: string
    // A mapping containing the LanguageCode => Description values
      locales: Record<string, string>
}
