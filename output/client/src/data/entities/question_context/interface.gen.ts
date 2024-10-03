
import {type ValueType,type DifficultyLabel,type TextMediaType,type QuestionType,type Status, } from './enums.gen'
import { type DocRef } from '../../general/docref'
import { AuditInfo } from '../../general/auditinfo'

/// ///////////////////////////////////////////////////////////////////////////////////////////////////
/// WARNING! THIS FILE IS AUTOMATICALLY GENERATED, ANY CHANGES WILL BE LOST ON NEXT GENERATION
/// ///////////////////////////////////////////////////////////////////////////////////////////////////


// QuestionContext Interface
export interface QuestionContext {
  // The ID of the QuestionContext
  id: string

  // The audit information on the creation of the QuestionContext
  created: AuditInfo

  // The audit information on the modification of the QuestionContext
  modified: AuditInfo

  // A list of possible answers for the question
  answer_list: AnswerList[]

  // DocumentReference to the Author
  author: DocRef

  // Statistical counters for the question
  counters: Counters

  // The difficulty level of the question
  difficulty: Difficulty

  // The default language code
  language_id: string

  // A list of language id's in which the question is available
  language_ids: string[]

  // The main category of the question
  main_category: MainCategory

  // Quality measures for the question
  quality: Quality

  // The question object
  question: Question

  // A list of QuestionBundleID strings
  question_bundle_ids: string[]

  // The current status of the question context
  status: Status

  // A list of tag Id strings
  tag_ids: string[]

  // A list of tags used to identify the question
  tags: Tags[]
}



// Property Interfaces
export interface AnswerList {
    // Indicates whether the answer is correct
      correct: boolean
    // Indicates whether the answer is disabled
      disabled: boolean
    // An optional explanation on why the answer is correct/incorrect
      explanation: string
    // A unique key used to identify the answer, such that answer values and labels can be edited without losing analytical insights
      key: string
    // The label of the answer, to be displayed to the player
      label: string
    // The media file is optional on the Answer object
      media: string
    // The type of content supported by the answer: TextOnly, TextAndMedia, or MediaOnly
      text_media_type: TextMediaType
    // The value of the answer, parsed as text (might be equal to the label, it might not)
      value: string
    // The true value type of the answer, can be used in Estimation or TrueFalse questions for example. Additionally, we could add types like Percentage or Range and parse the string Value to get the values needed to check if the answer is correct
      value_type: ValueType
}
export interface Counters {
    // Keep a counter of which answer was given, if applicable to the question. The {Answers.Key} refers to the Key property of the Answer child object
      answers: Record<string, number>
    // A counter that aggregates the number of times a question has been answered
      total_answers: number
    // A counter that aggregates the number of times a question has been answered correctly
      total_correct: number
}
export interface Difficulty {
    // The ratio of correct answers compared to all answers, ranging from 0 to 1
      accuracy: number
    // A default measure of how difficult the question is estimated, ranging from 0 to 100
      default_score: number
    // The label of the difficulty level
      label: DifficultyLabel
}
export interface MainCategory {
    // A unique identifier that is equal across localizations
      id: string
    // The localized label for the main category
      label: string
}
export interface Quality {
    // A quality rating from 0 to 5, given by the admin
      default_rating: number
    // A quality rating from 0 to 5, averaged by user_rating_sum/user_rating_count
      rating: number
    // The number of  user ratings
      user_rating_count: number
    // The sum of the user ratings
      user_rating_sum: number
}
export interface Question {
    // An optional further explanation of the question line
      explanation: string
    // The question line, basically the question itself
      line: string
    // DocumentReference to the media associated with the question
      media: DocRef
    // The type of content supported by the question: TextOnly, TextAndMedia, or MediaOnly
      text_media_type: TextMediaType
    // The type of the question: MultipleChoice, TrueFalse, or Estimation
      type: QuestionType
}
export interface Tags {
    // A unique identifier that is equal across localizations
      id: string
    // The localized label for the tag
      label: string
}
