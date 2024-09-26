package repo

import (
	"github.com/bakeable/bkry/data/entities/user_profile"
	"github.com/bakeable/bkry/data/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var UserProfileRepo = repository.NewRepository[*user_profile.UserProfile]()

// GetUserProfile retrieves a single UserProfile entity by its ID and userProfileID.
func GetUserProfile(userProfileID string) (user_profile.UserProfile, error) {
	entityMap, err := UserProfileRepo.Read(user_profile.GetDocumentPath( userProfileID))
	return user_profile.Decode(entityMap), err
}

// GetUserProfileOrNew retrieves a single UserProfile entity by its ID and userProfileID.
func GetUserProfileOrNew(userProfileID string) (user_profile.UserProfile, error) {
	entityMap, err := UserProfileRepo.Read(user_profile.GetDocumentPath( userProfileID))
	if err != nil || entityMap == nil {
		return user_profile.UserProfile{}, err
	}
	return user_profile.Decode(entityMap), err
}

// GetUserProfile retrieves a single UserProfile entity by its document path.
func GetUserProfileByPath(path string) (user_profile.UserProfile, error) {
	entityMap, err := UserProfileRepo.Read(path)
	return user_profile.Decode(entityMap), err
}

// FindUserProfile retrieves a UserProfile entity according to given queries.
func FindUserProfile(queries []datastore.Query) (user_profile.UserProfile, error) {
	entityMap, err := UserProfileRepo.Find(user_profile.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return user_profile.UserProfile{}, err
	}
	return user_profile.Decode(entityMap), err
}

// GetAllUserProfiles retrieves all UserProfile entities.
func GetAllUserProfiles() ([]user_profile.UserProfile, error) {
	entityMaps, err := UserProfileRepo.ReadAll(user_profile.GetCollectionPath())
	if err != nil {
		return []user_profile.UserProfile{}, err
	}
	return user_profile.DecodeAll(entityMaps), nil
}


// GetAllUserProfilesPaginated retrieves all UserProfile entities in a paginated manner.
func GetAllUserProfilesPaginated(pagination datastore.Pagination) ([]user_profile.UserProfile, datastore.Pagination, error) {
	entityMaps, pagination, err := UserProfileRepo.ReadPaginated(user_profile.GetCollectionPath(), pagination)
	if err != nil {
		return []user_profile.UserProfile{}, pagination, err
	}
	return user_profile.DecodeAll(entityMaps), pagination, nil
}

// QueryUserProfiles retrieves all UserProfile entities according to given queries.
func QueryUserProfiles(queries []datastore.Query, pagination datastore.Pagination) ([]user_profile.UserProfile, error) {
	entityMaps, err := UserProfileRepo.Query(user_profile.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []user_profile.UserProfile{}, err
	}
	return user_profile.DecodeAll(entityMaps), nil
}

// QueryUserProfilesGroup retrieves all UserProfile entities according to given queries.
func QueryUserProfilesGroup(queries []datastore.Query) ([]user_profile.UserProfile, error) {
	entityMaps, err := UserProfileRepo.QueryGroup("user_profiles", queries)
	if err != nil {
		return []user_profile.UserProfile{}, err
	}
	return user_profile.DecodeAll(entityMaps), nil
}

// CreateUserProfile creates a new UserProfile entity.
func CreateUserProfile(entity user_profile.UserProfile, editorID *string) (string, error) {
	return UserProfileRepo.Create(&entity, editorID)
}

// UpdateUserProfile updates an existing UserProfile entity.
func UpdateUserProfile(entity user_profile.UserProfile, editorID *string) (string, error) {
	return UserProfileRepo.Update(&entity, editorID)
}

// SaveUserProfile saves a UserProfile entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveUserProfile(entity user_profile.UserProfile, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateUserProfile(entity, editorID)
	} else {
		return UpdateUserProfile(entity, editorID)
	}
}

// DeleteUserProfile deletes a UserProfile entity by its parents' path and userProfileID.
func DeleteUserProfile(userProfileID string) error {
	return UserProfileRepo.Delete(user_profile.GetDocumentPath(userProfileID))
}
