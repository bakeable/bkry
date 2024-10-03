package game_modus
import "github.com/bakeable/bkry/internal/server/models/general/auditinfo"

type GameModus struct {
	// The unique identifier for the GameModus
	ID string `json:"id"`
	// The audit information for the modification of the GameModus
	Modified auditinfo.AuditInfo `json:"modified"`
	// The audit information for the creation of the GameModus
	Created auditinfo.AuditInfo `json:"created"`
	// Detailed explanation of the game mode
	Explanation string `json:"explanation"`
	// The default language code for the game mode
	LanguageID string `json:"language_id"`
	// A list of language id's in which the question bundle is available
	LanguageIds []string `json:"language_ids"`
	// Sections used to elaborate the game mode further
	Sections []Section `json:"sections"`
	// A short summarized explanation of the game mode
	Short string `json:"short"`
	// The subtitle of the game mode
	Subtitle string `json:"subtitle"`
	// The title of the game mode
	Title string `json:"title"`
}




type Section struct {
	// A fixed key across languages
	Key string `json:"key"`
	// The ordering of the explanation sections
	Order string `json:"order"`
	// The subtitle of the section
	Subtitle string `json:"subtitle"`
	// The text content of the section
	Text string `json:"text"`
	// The title of the section
	Title string `json:"title"`
}



