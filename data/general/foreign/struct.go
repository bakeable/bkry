package foreign

// Interface to grab any entity with an ID string
type Entity interface {
	GetId() string
}

// Interface to grab any entity with a DocumentPath string
type IKey interface {
	GetId() string
	GetDocumentPath() string
}

// Key is used to store relationships
type Key struct {
	ID string `json:"id"`
	DocumentPath string `json:"documentPath"`
}

type Keys []Key

// ActivationKey is used to store relationships which can be activated and deactivated
type ActivationKey struct {
	ID string `json:"id"`
	DocumentPath string `json:"documentPath"`
	Active bool `json:"active"`
}
type ActivationKeys []ActivationKey