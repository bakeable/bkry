package question_context

var ExampleJSON = `{
    answer_list: {
      correct: false, // Indicates whether the answer is correct 
      disabled: false, // Indicates whether the answer is disabled 
      explanation: "", // An optional explanation on why the answer is correct/incorrect 
      key: "", // A unique key used to identify the answer, such that answer values and labels can be edited without losing analytical insights 
      label: "", // The label of the answer, to be displayed to the player 
      media: "", // The media file is optional on the Answer object 
      text_media_type: TextMediaTypeValues.TextOnly, // The type of content supported by the answer: TextOnly, TextAndMedia, or MediaOnly (an enum with possible values: [])
      value: "", // The value of the answer, parsed as text (might be equal to the label, it might not) 
      value_type: ValueTypeValues.String, // The true value type of the answer, can be used in Estimation or TrueFalse questions for example. Additionally, we could add types like Percentage or Range and parse the string Value to get the values needed to check if the answer is correct (an enum with possible values: [String Int Float Boolean Percentage Range])
    }[], // A list of possible answers for the question
    counters: {
      answers: {} satisfies Record<string, number>, // Keep a counter of which answer was given, if applicable to the question. The {Answers.Key} refers to the Key property of the Answer child object 
      total_answers: 0, // A counter that aggregates the number of times a question has been answered 
      total_correct: 0, // A counter that aggregates the number of times a question has been answered correctly 
    }, // Statistical counters for the question
    difficulty: {
      accuracy: 0, // The ratio of correct answers compared to all answers, ranging from 0 to 1 
      default_score: 0, // A default measure of how difficult the question is estimated, ranging from 0 to 100 
      label: DifficultyLabelValues.Medium, // The label of the difficulty level (an enum with possible values: [VeryEasy Easy Medium Hard VeryHard])
    }, // The difficulty level of the question
    language_id: string, // The default language code
    language_ids: string[], // A list of language id's in which the question is available
    main_category: {
      id: "", // A unique identifier that is equal across localizations 
      label: "", // The localized label for the main category 
    }, // The main category of the question
    quality: {
      default_rating: 0, // A quality rating from 0 to 5, given by the admin 
      rating: 0, // A quality rating from 0 to 5, averaged by user_rating_sum/user_rating_count 
      user_rating_count: 0, // The number of  user ratings 
      user_rating_sum: 0, // The sum of the user ratings 
    }, // Quality measures for the question
    question: {
      explanation: "", // An optional further explanation of the question line 
      line: "", // The question line, basically the question itself 
      media: , // DocumentReference to the media associated with the question 
      text_media_type: TextMediaTypeValues.TextOnly, // The type of content supported by the question: TextOnly, TextAndMedia, or MediaOnly (an enum with possible values: [TextOnly TextAndMedia MediaOnly])
      type: QuestionTypeValues.MultipleChoice, // The type of the question: MultipleChoice, TrueFalse, or Estimation (an enum with possible values: [MultipleChoice TrueFalse Estimation])
    }, // The question object
    question_bundle_ids: string[], // A list of QuestionBundleID strings
    status: Status, // The current status of the question context
    tag_ids: string[], // A list of tag Id strings
    tags: {
      id: "", // A unique identifier that is equal across localizations 
      label: "", // The localized label for the tag 
    }[], // A list of tags used to identify the question
}`