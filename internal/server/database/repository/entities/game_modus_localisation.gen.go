package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/game_modus_localisation"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/internal/server/database"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var GameModusLocalisationRepo = repository.NewFirestoreRepository[*game_modus_localisation.GameModusLocalisation]()

// GetGameModusLocalisation retrieves a single GameModusLocalisation entity by its ID and gameModusLocalisationID.
func GetGameModusLocalisation(GameModusID string, gameModusLocalisationID string) (game_modus_localisation.GameModusLocalisation, error) {
	entityMap, err := GameModusLocalisationRepo.Read(game_modus_localisation.GetDocumentPath(GameModusID,  gameModusLocalisationID))
	return game_modus_localisation.Decode(entityMap), err
}

// GetGameModusLocalisationOrNew retrieves a single GameModusLocalisation entity by its ID and gameModusLocalisationID.
func GetGameModusLocalisationOrNew(GameModusID string, gameModusLocalisationID string) (game_modus_localisation.GameModusLocalisation, error) {
	entityMap, err := GameModusLocalisationRepo.Read(game_modus_localisation.GetDocumentPath(GameModusID,  gameModusLocalisationID))
	if err != nil || entityMap == nil {
		return game_modus_localisation.GameModusLocalisation{}, err
	}
	return game_modus_localisation.Decode(entityMap), err
}

// GetGameModusLocalisation retrieves a single GameModusLocalisation entity by its document path.
func GetGameModusLocalisationByPath(path string) (game_modus_localisation.GameModusLocalisation, error) {
	entityMap, err := GameModusLocalisationRepo.Read(path)
	return game_modus_localisation.Decode(entityMap), err
}

// FindGameModusLocalisation retrieves a GameModusLocalisation entity according to given queries.
func FindGameModusLocalisation(gameModusID string, queries []database.Query) (game_modus_localisation.GameModusLocalisation, error) {
	entityMap, err := GameModusLocalisationRepo.Find(game_modus_localisation.GetCollectionPath(gameModusID), queries)
	if err != nil || entityMap == nil {
		return game_modus_localisation.GameModusLocalisation{}, err
	}
	return game_modus_localisation.Decode(entityMap), err
}

// GetAllGameModusLocalisations retrieves all GameModusLocalisation entities.
func GetAllGameModusLocalisations(gameModusID string) ([]game_modus_localisation.GameModusLocalisation, error) {
	entityMaps, err := GameModusLocalisationRepo.ReadAll(game_modus_localisation.GetCollectionPath(gameModusID))
	if err != nil {
		return []game_modus_localisation.GameModusLocalisation{}, err
	}
	return game_modus_localisation.DecodeAll(entityMaps), nil
}


// GetAllGameModusLocalisationsPaginated retrieves all GameModusLocalisation entities in a paginated manner.
func GetAllGameModusLocalisationsPaginated(gameModusID string, pagination database.Pagination) ([]game_modus_localisation.GameModusLocalisation, database.Pagination, error) {
	entityMaps, pagination, err := GameModusLocalisationRepo.ReadPaginated(game_modus_localisation.GetCollectionPath(gameModusID), pagination)
	if err != nil {
		return []game_modus_localisation.GameModusLocalisation{}, pagination, err
	}
	return game_modus_localisation.DecodeAll(entityMaps), pagination, nil
}

// QueryGameModusLocalisations retrieves all GameModusLocalisation entities according to given queries.
func QueryGameModusLocalisations(gameModusID string, queries []database.Query, pagination database.Pagination) ([]game_modus_localisation.GameModusLocalisation, error) {
	entityMaps, err := GameModusLocalisationRepo.Query(game_modus_localisation.GetCollectionPath(gameModusID), queries, pagination)
	if err != nil {
		return []game_modus_localisation.GameModusLocalisation{}, err
	}
	return game_modus_localisation.DecodeAll(entityMaps), nil
}

// QueryGameModusLocalisationsGroup retrieves all GameModusLocalisation entities according to given queries.
func QueryGameModusLocalisationsGroup(queries []database.Query) ([]game_modus_localisation.GameModusLocalisation, error) {
	entityMaps, err := GameModusLocalisationRepo.QueryGroup("game_modus_localisations", queries)
	if err != nil {
		return []game_modus_localisation.GameModusLocalisation{}, err
	}
	return game_modus_localisation.DecodeAll(entityMaps), nil
}

// CreateGameModusLocalisation creates a new GameModusLocalisation entity.
func CreateGameModusLocalisation(GameModusID string, entity game_modus_localisation.GameModusLocalisation, editorID *string) (string, error) {
	entity.GameModusID = GameModusID
	return GameModusLocalisationRepo.Create(&entity, editorID)
}

// UpdateGameModusLocalisation updates an existing GameModusLocalisation entity.
func UpdateGameModusLocalisation(GameModusID string, entity game_modus_localisation.GameModusLocalisation, editorID *string) (string, error) {
	entity.GameModusID = GameModusID
	return GameModusLocalisationRepo.Update(&entity, editorID)
}

// SaveGameModusLocalisation saves a GameModusLocalisation entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveGameModusLocalisation(GameModusID string, entity game_modus_localisation.GameModusLocalisation, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateGameModusLocalisation(GameModusID, entity, editorID)
	} else {
		return UpdateGameModusLocalisation(GameModusID, entity, editorID)
	}
}

// DeleteGameModusLocalisation deletes a GameModusLocalisation entity by its parents' path and gameModusLocalisationID.
func DeleteGameModusLocalisation(GameModusID string, gameModusLocalisationID string) error {
	return GameModusLocalisationRepo.Delete(game_modus_localisation.GetDocumentPath(GameModusID, gameModusLocalisationID))
}
