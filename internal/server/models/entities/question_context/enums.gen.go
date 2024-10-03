package question_context
type ValueType string
type DifficultyLabel string
type TextMediaType string
type QuestionType string
type Status string
const (
    ValueType_String ValueType = "String"
    ValueType_Int ValueType = "Int"
    ValueType_Float ValueType = "Float"
    ValueType_Boolean ValueType = "Boolean"
    ValueType_Percentage ValueType = "Percentage"
    ValueType_Range ValueType = "Range"
    DifficultyLabel_VeryEasy DifficultyLabel = "VeryEasy"
    DifficultyLabel_Easy DifficultyLabel = "Easy"
    DifficultyLabel_Medium DifficultyLabel = "Medium"
    DifficultyLabel_Hard DifficultyLabel = "Hard"
    DifficultyLabel_VeryHard DifficultyLabel = "VeryHard"
    TextMediaType_TextOnly TextMediaType = "TextOnly"
    TextMediaType_TextAndMedia TextMediaType = "TextAndMedia"
    TextMediaType_MediaOnly TextMediaType = "MediaOnly"
    QuestionType_MultipleChoice QuestionType = "MultipleChoice"
    QuestionType_TrueFalse QuestionType = "TrueFalse"
    QuestionType_Estimation QuestionType = "Estimation"
    Status_Concept Status = "Concept"
    Status_Approved Status = "Approved"
    Status_Published Status = "Published"
)