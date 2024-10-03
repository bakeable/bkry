package question_context
import (
	"github.com/bakeable/bkry/internal/server/models/general/docref"
   "github.com/bakeable/bkry/internal/server/models/general/auditinfo"
)

type QuestionContext struct {
	// The unique identifier for the QuestionContext
	ID string `json:"id"`
	// The audit information for the modification of the QuestionContext
	Modified auditinfo.AuditInfo `json:"modified"`
	// The audit information for the creation of the QuestionContext
	Created auditinfo.AuditInfo `json:"created"`
	// A list of possible answers for the question
	AnswerList []AnswerList `json:"answer_list"`
	// DocumentReference to the Author
	Author docref.DocRef `json:"author"`
	// Statistical counters for the question
	Counters Counters `json:"counters"`
	// The difficulty level of the question
	Difficulty Difficulty `json:"difficulty"`
	// The default language code
	LanguageID string `json:"language_id"`
	// A list of language id's in which the question is available
	LanguageIds []string `json:"language_ids"`
	// The main category of the question
	MainCategory Category `json:"main_category"`
	// Quality measures for the question
	Quality Quality `json:"quality"`
	// The question object
	Question Question `json:"question"`
	// A list of QuestionBundleID strings
	QuestionBundleIds []string `json:"question_bundle_ids"`
	// The current status of the question context
	Status Status `json:"status"`
	// A list of tag Id strings
	TagIds []string `json:"tag_ids"`
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

type Counters struct {
	// Keep a counter of which answer was given, if applicable to the question. The {Answers.Key} refers to the Key property of the Answer child object
	Answers map[string]int `json:"answers"`
	// A counter that aggregates the number of times a question has been answered
	TotalAnswers int `json:"total_answers"`
	// A counter that aggregates the number of times a question has been answered correctly
	TotalCorrect int `json:"total_correct"`
}

type Difficulty struct {
	// The ratio of correct answers compared to all answers, ranging from 0 to 1
	Accuracy float64 `json:"accuracy"`
	// A default measure of how difficult the question is estimated, ranging from 0 to 100
	DefaultScore float64 `json:"default_score"`
	// The label of the difficulty level
	Label DifficultyLabel `json:"label"`
}

type Category struct {
	// A unique identifier that is equal across localizations
	ID string `json:"id"`
	// The localized label for the main category
	Label string `json:"label"`
}

type Quality struct {
	// A quality rating from 0 to 5, given by the admin
	DefaultRating float64 `json:"default_rating"`
	// A quality rating from 0 to 5, averaged by user_rating_sum/user_rating_count
	Rating float64 `json:"rating"`
	// The number of  user ratings
	UserRatingCount int `json:"user_rating_count"`
	// The sum of the user ratings
	UserRatingSum int `json:"user_rating_sum"`
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



