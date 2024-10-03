
import {type Status, } from './enums.gen'
import { AuditInfo } from '../../general/auditinfo'

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


// QuestionBundleLocalisation Interface
export interface QuestionBundleLocalisation {
  // The ID of the QuestionBundleLocalisation
  id: string

  // The audit information on the creation of the QuestionBundleLocalisation
  created: AuditInfo

  // The audit information on the modification of the QuestionBundleLocalisation
  modified: AuditInfo

  // Optional further description of the contents of the bundle
  description: string

  // The default language code for the question bundle
  language_id: string

  // The ID of the parent entity for which this is the child document
  question_bundle_id: string

  // The current status of the question bundle
  status: Status

  // The subtitle of the question bundle
  subtitle: string

  // The title of the question bundle
  title: string
}



// Property Interfaces
