package repo

import (
	"github.com/bakeable/bkry/data/entities/user"
	"github.com/bakeable/bkry/data/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var UserRepo = repository.NewRepository[*user.User]()

// GetUser retrieves a single User entity by its ID and userID.
func GetUser(userID string) (user.User, error) {
	entityMap, err := UserRepo.Read(user.GetDocumentPath( userID))
	return user.Decode(entityMap), err
}

// GetUserOrNew retrieves a single User entity by its ID and userID.
func GetUserOrNew(userID string) (user.User, error) {
	entityMap, err := UserRepo.Read(user.GetDocumentPath( userID))
	if err != nil || entityMap == nil {
		return user.User{}, err
	}
	return user.Decode(entityMap), err
}

// GetUser retrieves a single User entity by its document path.
func GetUserByPath(path string) (user.User, error) {
	entityMap, err := UserRepo.Read(path)
	return user.Decode(entityMap), err
}

// FindUser retrieves a User entity according to given queries.
func FindUser(queries []datastore.Query) (user.User, error) {
	entityMap, err := UserRepo.Find(user.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return user.User{}, err
	}
	return user.Decode(entityMap), err
}

// GetAllUsers retrieves all User entities.
func GetAllUsers() ([]user.User, error) {
	entityMaps, err := UserRepo.ReadAll(user.GetCollectionPath())
	if err != nil {
		return []user.User{}, err
	}
	return user.DecodeAll(entityMaps), nil
}


// GetAllUsersPaginated retrieves all User entities in a paginated manner.
func GetAllUsersPaginated(pagination datastore.Pagination) ([]user.User, datastore.Pagination, error) {
	entityMaps, pagination, err := UserRepo.ReadPaginated(user.GetCollectionPath(), pagination)
	if err != nil {
		return []user.User{}, pagination, err
	}
	return user.DecodeAll(entityMaps), pagination, nil
}

// QueryUsers retrieves all User entities according to given queries.
func QueryUsers(queries []datastore.Query, pagination datastore.Pagination) ([]user.User, error) {
	entityMaps, err := UserRepo.Query(user.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []user.User{}, err
	}
	return user.DecodeAll(entityMaps), nil
}

// QueryUsersGroup retrieves all User entities according to given queries.
func QueryUsersGroup(queries []datastore.Query) ([]user.User, error) {
	entityMaps, err := UserRepo.QueryGroup("users", queries)
	if err != nil {
		return []user.User{}, err
	}
	return user.DecodeAll(entityMaps), nil
}

// CreateUser creates a new User entity.
func CreateUser(entity user.User, editorID *string) (string, error) {
	return UserRepo.Create(&entity, editorID)
}

// UpdateUser updates an existing User entity.
func UpdateUser(entity user.User, editorID *string) (string, error) {
	return UserRepo.Update(&entity, editorID)
}

// SaveUser saves a User entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveUser(entity user.User, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateUser(entity, editorID)
	} else {
		return UpdateUser(entity, editorID)
	}
}

// DeleteUser deletes a User entity by its parents' path and userID.
func DeleteUser(userID string) error {
	return UserRepo.Delete(user.GetDocumentPath(userID))
}
