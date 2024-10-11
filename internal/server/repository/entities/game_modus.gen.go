package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/game_modus"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/internal/server/database"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var GameModusRepo = repository.NewFirestoreRepository[*game_modus.GameModus]()

// GetGameModus retrieves a single GameModus entity by its ID and gameModusID.
func GetGameModus(gameModusID string) (game_modus.GameModus, error) {
	entityMap, err := GameModusRepo.Read(game_modus.GetDocumentPath( gameModusID))
	return game_modus.Decode(entityMap), err
}

// GetGameModusOrNew retrieves a single GameModus entity by its ID and gameModusID.
func GetGameModusOrNew(gameModusID string) (game_modus.GameModus, error) {
	entityMap, err := GameModusRepo.Read(game_modus.GetDocumentPath( gameModusID))
	if err != nil || entityMap == nil {
		return game_modus.GameModus{}, err
	}
	return game_modus.Decode(entityMap), err
}

// GetGameModus retrieves a single GameModus entity by its document path.
func GetGameModusByPath(path string) (game_modus.GameModus, error) {
	entityMap, err := GameModusRepo.Read(path)
	return game_modus.Decode(entityMap), err
}

// FindGameModus retrieves a GameModus entity according to given queries.
func FindGameModus(queries []database.Query) (game_modus.GameModus, error) {
	entityMap, err := GameModusRepo.Find(game_modus.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return game_modus.GameModus{}, err
	}
	return game_modus.Decode(entityMap), err
}

// GetAllGameModus retrieves all GameModus entities.
func GetAllGameModus() ([]game_modus.GameModus, error) {
	entityMaps, err := GameModusRepo.ReadAll(game_modus.GetCollectionPath())
	if err != nil {
		return []game_modus.GameModus{}, err
	}
	return game_modus.DecodeAll(entityMaps), nil
}


// GetAllGameModusPaginated retrieves all GameModus entities in a paginated manner.
func GetAllGameModusPaginated(pagination database.Pagination) ([]game_modus.GameModus, database.Pagination, error) {
	entityMaps, pagination, err := GameModusRepo.ReadPaginated(game_modus.GetCollectionPath(), pagination)
	if err != nil {
		return []game_modus.GameModus{}, pagination, err
	}
	return game_modus.DecodeAll(entityMaps), pagination, nil
}

// QueryGameModus retrieves all GameModus entities according to given queries.
func QueryGameModus(queries []database.Query, pagination database.Pagination) ([]game_modus.GameModus, error) {
	entityMaps, err := GameModusRepo.Query(game_modus.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []game_modus.GameModus{}, err
	}
	return game_modus.DecodeAll(entityMaps), nil
}

// QueryGameModusGroup retrieves all GameModus entities according to given queries.
func QueryGameModusGroup(queries []database.Query) ([]game_modus.GameModus, error) {
	entityMaps, err := GameModusRepo.QueryGroup("game_modus", queries)
	if err != nil {
		return []game_modus.GameModus{}, err
	}
	return game_modus.DecodeAll(entityMaps), nil
}

// CreateGameModus creates a new GameModus entity.
func CreateGameModus(entity game_modus.GameModus, editorID *string) (string, error) {
	return GameModusRepo.Create(&entity, editorID)
}

// UpdateGameModus updates an existing GameModus entity.
func UpdateGameModus(entity game_modus.GameModus, editorID *string) (string, error) {
	return GameModusRepo.Update(&entity, editorID)
}

// SaveGameModus saves a GameModus entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveGameModus(entity game_modus.GameModus, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateGameModus(entity, editorID)
	} else {
		return UpdateGameModus(entity, editorID)
	}
}

// DeleteGameModus deletes a GameModus entity by its parents' path and gameModusID.
func DeleteGameModus(gameModusID string) error {
	return GameModusRepo.Delete(game_modus.GetDocumentPath(gameModusID))
}
