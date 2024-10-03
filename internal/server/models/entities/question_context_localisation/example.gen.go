package question_context_localisation

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
    language_id: string, // The default language code
    main_category: {
      id: "", // A unique identifier that is equal across localizations 
      label: "", // The localized label for the main category 
    }, // The main category of the question
    question: {
      explanation: "", // An optional further explanation of the question line 
      line: "", // The question line, basically the question itself 
      media: , // DocumentReference to the media associated with the question 
      text_media_type: TextMediaTypeValues.TextOnly, // The type of content supported by the question: TextOnly, TextAndMedia, or MediaOnly (an enum with possible values: [TextOnly TextAndMedia MediaOnly])
      type: QuestionTypeValues.MultipleChoice, // The type of the question: MultipleChoice, TrueFalse, or Estimation (an enum with possible values: [MultipleChoice TrueFalse Estimation])
    }, // The question object
    question_context_id: string, // The ID of the parent entity for which this is the child document
    status: Status, // The current status of the question context
    tags: {
      id: "", // A unique identifier that is equal across localizations 
      label: "", // The localized label for the tag 
    }[], // A list of tags used to identify the question
}`