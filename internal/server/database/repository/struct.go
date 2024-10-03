package repository

// Define your data model interface
type Entity interface {
	GetDocumentPath() string
	GetCollectionPath() string
	ToMap() map[string]interface{}
	GetID() string
}

type EntityForeignKeyMap interface {
	ToMap() map[string]interface{}
}

type Operator string

const (
	Equal       Operator = "=="
	NotEqual    Operator = "!="
	GreaterThan Operator = ">"
	LessThan    Operator = "<"
	In          Operator = "in"
)

type Query struct {
	Field    string
	Operator Operator
	Value    interface{}
}

type Pagination struct {
	PageNumber      int  `json:"page_number"`
	PageSize        int  `json:"page_size"`
	OrderBy         string  `json:"order_by"`
	Asc             bool  `json:"asc"`
	Count           *int  `json:"count,omitempty"`
	TotalPages      int  `json:"total_pages,omitempty"`
	HasNextPage     bool  `json:"has_next_page"`
	HasPreviousPage bool  `json:"has_previous_page"`
}

type Repository[T Entity] interface {
	Read(path string) (map[string]interface{}, error)
	Find(path string, queries []Query) (map[string]interface{}, error)
	ReadAll(path string) ([]map[string]interface{}, error)
	ReadPaginated(path string, pagination Pagination) ([]map[string]interface{}, Pagination, error)
	Query(path string, queries []Query, pagination Pagination) ([]map[string]interface{}, error)
	QueryGroup(collectionGroup string, queries []Query) ([]map[string]interface{}, error)
	Create(entity T, editorId *string) (string, error)
	Update(entity T, editorId *string) (string, error)
	Delete(path string) error
}

// GenericRepository provides CRUD operations on a generic entity type.
type DatastoreRepository[T Entity] struct{}
type FirestoreRepository[T Entity] struct{}
