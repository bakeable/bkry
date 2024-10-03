package question_context_localisation
import (
	"github.com/bakeable/bkry/internal/server/models/general/docref"
   "github.com/bakeable/bkry/internal/server/models/general/auditinfo"
)

type QuestionContextLocalisation struct {
	// The unique identifier for the QuestionContextLocalisation
	ID string `json:"id"`
	// The audit information for the modification of the QuestionContextLocalisation
	Modified auditinfo.AuditInfo `json:"modified"`
	// The audit information for the creation of the QuestionContextLocalisation
	Created auditinfo.AuditInfo `json:"created"`
	// A list of possible answers for the question
	AnswerList []AnswerList `json:"answer_list"`
	// The default language code
	LanguageID string `json:"language_id"`
	// The main category of the question
	MainCategory Category `json:"main_category"`
	// The question object
	Question Question `json:"question"`
	// The ID of the parent entity for which this is the child document
	QuestionContextID string `json:"question_context_id"`
	// The current status of the question context
	Status Status `json:"status"`
	// A list of tags used to identify the question
	Tags []Tag `json:"tags"`
}




type AnswerList struct {
	// Indicates whether the answer is correct
	Correct bool `json:"correct"`
	// Indicates whether the answer is disabled
	Disabled bool `json:"disabled"`
	// An optional explanation on why the answer is correct/incorrect
	Explanation string `json:"explanation"`
	// A unique key used to identify the answer, such that answer values and labels can be edited without losing analytical insights
	Key string `json:"key"`
	// The label of the answer, to be displayed to the player
	Label string `json:"label"`
	// The media file is optional on the Answer object
	Media string `json:"media"`
	// The type of content supported by the answer: TextOnly, TextAndMedia, or MediaOnly
	TextMediaType TextMediaType `json:"text_media_type"`
	// The value of the answer, parsed as text (might be equal to the label, it might not)
	Value string `json:"value"`
	// The true value type of the answer, can be used in Estimation or TrueFalse questions for example. Additionally, we could add types like Percentage or Range and parse the string Value to get the values needed to check if the answer is correct
	ValueType ValueType `json:"value_type"`
}

type Category struct {
	// A unique identifier that is equal across localizations
	ID string `json:"id"`
	// The localized label for the main category
	Label string `json:"label"`
}

type Question struct {
	// An optional further explanation of the question line
	Explanation string `json:"explanation"`
	// The question line, basically the question itself
	Line string `json:"line"`
	// DocumentReference to the media associated with the question
	Media docref.DocRef `json:"media"`
	// The type of content supported by the question: TextOnly, TextAndMedia, or MediaOnly
	TextMediaType TextMediaType `json:"text_media_type"`
	// The type of the question: MultipleChoice, TrueFalse, or Estimation
	Type QuestionType `json:"type"`
}

type Tag struct {
	// A unique identifier that is equal across localizations
	ID string `json:"id"`
	// The localized label for the tag
	Label string `json:"label"`
}



