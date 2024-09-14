package repository

import (
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
)

// Define your data model interface
type Entity interface {
	GetDocumentPath() string
	GetCollectionPath() string
	Encode() map[string]interface{}
	GetID() string
}

type EntityForeignKeyMap interface {
	Encode() map[string]interface{}
}

type Query datastore.Query

// GenericRepository provides CRUD operations on a generic entity type.
type Repository[T Entity] struct{}
