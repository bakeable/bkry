package repo

import (
	"github.com/bakeable/bkry/internal/server/database"
	"github.com/bakeable/bkry/internal/server/models/entities/author"
	"github.com/bakeable/bkry/internal/server/models/entities/category"
	"github.com/bakeable/bkry/internal/server/models/entities/game_modus"
	"github.com/bakeable/bkry/internal/server/models/entities/media"
	"github.com/bakeable/bkry/internal/server/models/entities/product_package"
	"github.com/bakeable/bkry/internal/server/models/entities/purchase"
	"github.com/bakeable/bkry/internal/server/models/entities/question_bundle"
	"github.com/bakeable/bkry/internal/server/models/entities/question_context"
	"github.com/bakeable/bkry/internal/server/models/entities/tag"
	"github.com/bakeable/bkry/internal/server/models/entities/user"
	"github.com/bakeable/bkry/internal/server/models/entities/category_localisation"
	"github.com/bakeable/bkry/internal/server/models/entities/game_modus_localisation"
	"github.com/bakeable/bkry/internal/server/models/entities/product_package_localisation"
	"github.com/bakeable/bkry/internal/server/models/entities/question_bundle_localisation"
	"github.com/bakeable/bkry/internal/server/models/entities/question_context_localisation"
	"github.com/bakeable/bkry/internal/server/models/entities/tag_localisation"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

type IRepository interface {
	// GetAuthor retrieves a single Author entity by its ID and authorID.
	GetAuthor(authorID string) (author.Author, error)
	// GetAuthorOrNew retrieves a single Author entity by its ID and authorID.
	GetAuthorOrNew(authorID string) (author.Author, error)
	// GetAuthor retrieves a single Author entity by its document path.
	GetAuthorByPath(path string) (author.Author, error)
	// FindAuthor retrieves a Author entity according to given queries.
	FindAuthor(queries []database.Query) (author.Author, error)
	// GetAllAuthors retrieves all Author entities.
	GetAllAuthors() ([]author.Author, error)
	// GetAllAuthorsPaginated retrieves all Author entities in a paginated manner.
	GetAllAuthorsPaginated(pagination database.Pagination) ([]author.Author, database.Pagination, error)
	// QueryAuthors retrieves all Author entities according to given queries.
	QueryAuthors(queries []database.Query, pagination database.Pagination) ([]author.Author, error)
	// QueryAuthorsGroup retrieves all Author entities according to given queries.
	QueryAuthorsGroup(queries []database.Query) ([]author.Author, error)
	// CreateAuthor creates a new Author entity.
	CreateAuthor(entity author.Author, editorID *string) (string, error)
	// UpdateAuthor updates an existing Author entity.
	UpdateAuthor(entity author.Author, editorID *string) (string, error)
	// SaveAuthor saves a Author entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveAuthor(entity author.Author, editorID *string) (string, error)
	// DeleteAuthor deletes a Author entity by its parents' path and authorID.
	DeleteAuthor(authorID string) error
	// GetCategory retrieves a single Category entity by its ID and categoryID.
	GetCategory(categoryID string) (category.Category, error)
	// GetCategoryOrNew retrieves a single Category entity by its ID and categoryID.
	GetCategoryOrNew(categoryID string) (category.Category, error)
	// GetCategory retrieves a single Category entity by its document path.
	GetCategoryByPath(path string) (category.Category, error)
	// FindCategory retrieves a Category entity according to given queries.
	FindCategory(queries []database.Query) (category.Category, error)
	// GetAllCategories retrieves all Category entities.
	GetAllCategories() ([]category.Category, error)
	// GetAllCategoriesPaginated retrieves all Category entities in a paginated manner.
	GetAllCategoriesPaginated(pagination database.Pagination) ([]category.Category, database.Pagination, error)
	// QueryCategories retrieves all Category entities according to given queries.
	QueryCategories(queries []database.Query, pagination database.Pagination) ([]category.Category, error)
	// QueryCategoriesGroup retrieves all Category entities according to given queries.
	QueryCategoriesGroup(queries []database.Query) ([]category.Category, error)
	// CreateCategory creates a new Category entity.
	CreateCategory(entity category.Category, editorID *string) (string, error)
	// UpdateCategory updates an existing Category entity.
	UpdateCategory(entity category.Category, editorID *string) (string, error)
	// SaveCategory saves a Category entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveCategory(entity category.Category, editorID *string) (string, error)
	// DeleteCategory deletes a Category entity by its parents' path and categoryID.
	DeleteCategory(categoryID string) error
	// GetGameModus retrieves a single GameModus entity by its ID and gameModusID.
	GetGameModus(gameModusID string) (game_modus.GameModus, error)
	// GetGameModusOrNew retrieves a single GameModus entity by its ID and gameModusID.
	GetGameModusOrNew(gameModusID string) (game_modus.GameModus, error)
	// GetGameModus retrieves a single GameModus entity by its document path.
	GetGameModusByPath(path string) (game_modus.GameModus, error)
	// FindGameModus retrieves a GameModus entity according to given queries.
	FindGameModus(queries []database.Query) (game_modus.GameModus, error)
	// GetAllGameModus retrieves all GameModus entities.
	GetAllGameModus() ([]game_modus.GameModus, error)
	// GetAllGameModusPaginated retrieves all GameModus entities in a paginated manner.
	GetAllGameModusPaginated(pagination database.Pagination) ([]game_modus.GameModus, database.Pagination, error)
	// QueryGameModus retrieves all GameModus entities according to given queries.
	QueryGameModus(queries []database.Query, pagination database.Pagination) ([]game_modus.GameModus, error)
	// QueryGameModusGroup retrieves all GameModus entities according to given queries.
	QueryGameModusGroup(queries []database.Query) ([]game_modus.GameModus, error)
	// CreateGameModus creates a new GameModus entity.
	CreateGameModus(entity game_modus.GameModus, editorID *string) (string, error)
	// UpdateGameModus updates an existing GameModus entity.
	UpdateGameModus(entity game_modus.GameModus, editorID *string) (string, error)
	// SaveGameModus saves a GameModus entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveGameModus(entity game_modus.GameModus, editorID *string) (string, error)
	// DeleteGameModus deletes a GameModus entity by its parents' path and gameModusID.
	DeleteGameModus(gameModusID string) error
	// GetMedia retrieves a single Media entity by its ID and mediaID.
	GetMedia(mediaID string) (media.Media, error)
	// GetMediaOrNew retrieves a single Media entity by its ID and mediaID.
	GetMediaOrNew(mediaID string) (media.Media, error)
	// GetMedia retrieves a single Media entity by its document path.
	GetMediaByPath(path string) (media.Media, error)
	// FindMedia retrieves a Media entity according to given queries.
	FindMedia(queries []database.Query) (media.Media, error)
	// GetAllMedia retrieves all Media entities.
	GetAllMedia() ([]media.Media, error)
	// GetAllMediaPaginated retrieves all Media entities in a paginated manner.
	GetAllMediaPaginated(pagination database.Pagination) ([]media.Media, database.Pagination, error)
	// QueryMedia retrieves all Media entities according to given queries.
	QueryMedia(queries []database.Query, pagination database.Pagination) ([]media.Media, error)
	// QueryMediaGroup retrieves all Media entities according to given queries.
	QueryMediaGroup(queries []database.Query) ([]media.Media, error)
	// CreateMedia creates a new Media entity.
	CreateMedia(entity media.Media, editorID *string) (string, error)
	// UpdateMedia updates an existing Media entity.
	UpdateMedia(entity media.Media, editorID *string) (string, error)
	// SaveMedia saves a Media entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveMedia(entity media.Media, editorID *string) (string, error)
	// DeleteMedia deletes a Media entity by its parents' path and mediaID.
	DeleteMedia(mediaID string) error
	// GetProductPackage retrieves a single ProductPackage entity by its ID and productPackageID.
	GetProductPackage(productPackageID string) (product_package.ProductPackage, error)
	// GetProductPackageOrNew retrieves a single ProductPackage entity by its ID and productPackageID.
	GetProductPackageOrNew(productPackageID string) (product_package.ProductPackage, error)
	// GetProductPackage retrieves a single ProductPackage entity by its document path.
	GetProductPackageByPath(path string) (product_package.ProductPackage, error)
	// FindProductPackage retrieves a ProductPackage entity according to given queries.
	FindProductPackage(queries []database.Query) (product_package.ProductPackage, error)
	// GetAllProductPackages retrieves all ProductPackage entities.
	GetAllProductPackages() ([]product_package.ProductPackage, error)
	// GetAllProductPackagesPaginated retrieves all ProductPackage entities in a paginated manner.
	GetAllProductPackagesPaginated(pagination database.Pagination) ([]product_package.ProductPackage, database.Pagination, error)
	// QueryProductPackages retrieves all ProductPackage entities according to given queries.
	QueryProductPackages(queries []database.Query, pagination database.Pagination) ([]product_package.ProductPackage, error)
	// QueryProductPackagesGroup retrieves all ProductPackage entities according to given queries.
	QueryProductPackagesGroup(queries []database.Query) ([]product_package.ProductPackage, error)
	// CreateProductPackage creates a new ProductPackage entity.
	CreateProductPackage(entity product_package.ProductPackage, editorID *string) (string, error)
	// UpdateProductPackage updates an existing ProductPackage entity.
	UpdateProductPackage(entity product_package.ProductPackage, editorID *string) (string, error)
	// SaveProductPackage saves a ProductPackage entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveProductPackage(entity product_package.ProductPackage, editorID *string) (string, error)
	// DeleteProductPackage deletes a ProductPackage entity by its parents' path and productPackageID.
	DeleteProductPackage(productPackageID string) error
	// GetPurchase retrieves a single Purchase entity by its ID and purchaseID.
	GetPurchase(userID string, purchaseID string) (purchase.Purchase, error)
	// GetPurchaseOrNew retrieves a single Purchase entity by its ID and purchaseID.
	GetPurchaseOrNew(userID string, purchaseID string) (purchase.Purchase, error)
	// GetPurchase retrieves a single Purchase entity by its document path.
	GetPurchaseByPath(path string) (purchase.Purchase, error)
	// FindPurchase retrieves a Purchase entity according to given queries.
	FindPurchase(userID string, queries []database.Query) (purchase.Purchase, error)
	// GetAllPurchases retrieves all Purchase entities.
	GetAllPurchases(userID string) ([]purchase.Purchase, error)
	// GetAllPurchasesPaginated retrieves all Purchase entities in a paginated manner.
	GetAllPurchasesPaginated(userID string, pagination database.Pagination) ([]purchase.Purchase, database.Pagination, error)
	// QueryPurchases retrieves all Purchase entities according to given queries.
	QueryPurchases(userID string, queries []database.Query, pagination database.Pagination) ([]purchase.Purchase, error)
	// QueryPurchasesGroup retrieves all Purchase entities according to given queries.
	QueryPurchasesGroup(queries []database.Query) ([]purchase.Purchase, error)
	// CreatePurchase creates a new Purchase entity.
	CreatePurchase(userID string, entity purchase.Purchase, editorID *string) (string, error)
	// UpdatePurchase updates an existing Purchase entity.
	UpdatePurchase(userID string, entity purchase.Purchase, editorID *string) (string, error)
	// SavePurchase saves a Purchase entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SavePurchase(userID string, entity purchase.Purchase, editorID *string) (string, error)
	// DeletePurchase deletes a Purchase entity by its parents' path and purchaseID.
	DeletePurchase(userID string, purchaseID string) error
	// GetQuestionBundle retrieves a single QuestionBundle entity by its ID and questionBundleID.
	GetQuestionBundle(questionBundleID string) (question_bundle.QuestionBundle, error)
	// GetQuestionBundleOrNew retrieves a single QuestionBundle entity by its ID and questionBundleID.
	GetQuestionBundleOrNew(questionBundleID string) (question_bundle.QuestionBundle, error)
	// GetQuestionBundle retrieves a single QuestionBundle entity by its document path.
	GetQuestionBundleByPath(path string) (question_bundle.QuestionBundle, error)
	// FindQuestionBundle retrieves a QuestionBundle entity according to given queries.
	FindQuestionBundle(queries []database.Query) (question_bundle.QuestionBundle, error)
	// GetAllQuestionBundles retrieves all QuestionBundle entities.
	GetAllQuestionBundles() ([]question_bundle.QuestionBundle, error)
	// GetAllQuestionBundlesPaginated retrieves all QuestionBundle entities in a paginated manner.
	GetAllQuestionBundlesPaginated(pagination database.Pagination) ([]question_bundle.QuestionBundle, database.Pagination, error)
	// QueryQuestionBundles retrieves all QuestionBundle entities according to given queries.
	QueryQuestionBundles(queries []database.Query, pagination database.Pagination) ([]question_bundle.QuestionBundle, error)
	// QueryQuestionBundlesGroup retrieves all QuestionBundle entities according to given queries.
	QueryQuestionBundlesGroup(queries []database.Query) ([]question_bundle.QuestionBundle, error)
	// CreateQuestionBundle creates a new QuestionBundle entity.
	CreateQuestionBundle(entity question_bundle.QuestionBundle, editorID *string) (string, error)
	// UpdateQuestionBundle updates an existing QuestionBundle entity.
	UpdateQuestionBundle(entity question_bundle.QuestionBundle, editorID *string) (string, error)
	// SaveQuestionBundle saves a QuestionBundle entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveQuestionBundle(entity question_bundle.QuestionBundle, editorID *string) (string, error)
	// DeleteQuestionBundle deletes a QuestionBundle entity by its parents' path and questionBundleID.
	DeleteQuestionBundle(questionBundleID string) error
	// GetQuestionContext retrieves a single QuestionContext entity by its ID and questionContextID.
	GetQuestionContext(questionContextID string) (question_context.QuestionContext, error)
	// GetQuestionContextOrNew retrieves a single QuestionContext entity by its ID and questionContextID.
	GetQuestionContextOrNew(questionContextID string) (question_context.QuestionContext, error)
	// GetQuestionContext retrieves a single QuestionContext entity by its document path.
	GetQuestionContextByPath(path string) (question_context.QuestionContext, error)
	// FindQuestionContext retrieves a QuestionContext entity according to given queries.
	FindQuestionContext(queries []database.Query) (question_context.QuestionContext, error)
	// GetAllQuestionContexts retrieves all QuestionContext entities.
	GetAllQuestionContexts() ([]question_context.QuestionContext, error)
	// GetAllQuestionContextsPaginated retrieves all QuestionContext entities in a paginated manner.
	GetAllQuestionContextsPaginated(pagination database.Pagination) ([]question_context.QuestionContext, database.Pagination, error)
	// QueryQuestionContexts retrieves all QuestionContext entities according to given queries.
	QueryQuestionContexts(queries []database.Query, pagination database.Pagination) ([]question_context.QuestionContext, error)
	// QueryQuestionContextsGroup retrieves all QuestionContext entities according to given queries.
	QueryQuestionContextsGroup(queries []database.Query) ([]question_context.QuestionContext, error)
	// CreateQuestionContext creates a new QuestionContext entity.
	CreateQuestionContext(entity question_context.QuestionContext, editorID *string) (string, error)
	// UpdateQuestionContext updates an existing QuestionContext entity.
	UpdateQuestionContext(entity question_context.QuestionContext, editorID *string) (string, error)
	// SaveQuestionContext saves a QuestionContext entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveQuestionContext(entity question_context.QuestionContext, editorID *string) (string, error)
	// DeleteQuestionContext deletes a QuestionContext entity by its parents' path and questionContextID.
	DeleteQuestionContext(questionContextID string) error
	// GetTag retrieves a single Tag entity by its ID and tagID.
	GetTag(tagID string) (tag.Tag, error)
	// GetTagOrNew retrieves a single Tag entity by its ID and tagID.
	GetTagOrNew(tagID string) (tag.Tag, error)
	// GetTag retrieves a single Tag entity by its document path.
	GetTagByPath(path string) (tag.Tag, error)
	// FindTag retrieves a Tag entity according to given queries.
	FindTag(queries []database.Query) (tag.Tag, error)
	// GetAllTags retrieves all Tag entities.
	GetAllTags() ([]tag.Tag, error)
	// GetAllTagsPaginated retrieves all Tag entities in a paginated manner.
	GetAllTagsPaginated(pagination database.Pagination) ([]tag.Tag, database.Pagination, error)
	// QueryTags retrieves all Tag entities according to given queries.
	QueryTags(queries []database.Query, pagination database.Pagination) ([]tag.Tag, error)
	// QueryTagsGroup retrieves all Tag entities according to given queries.
	QueryTagsGroup(queries []database.Query) ([]tag.Tag, error)
	// CreateTag creates a new Tag entity.
	CreateTag(entity tag.Tag, editorID *string) (string, error)
	// UpdateTag updates an existing Tag entity.
	UpdateTag(entity tag.Tag, editorID *string) (string, error)
	// SaveTag saves a Tag entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveTag(entity tag.Tag, editorID *string) (string, error)
	// DeleteTag deletes a Tag entity by its parents' path and tagID.
	DeleteTag(tagID string) error
	// GetUser retrieves a single User entity by its ID and userID.
	GetUser(userID string) (user.User, error)
	// GetUserOrNew retrieves a single User entity by its ID and userID.
	GetUserOrNew(userID string) (user.User, error)
	// GetUser retrieves a single User entity by its document path.
	GetUserByPath(path string) (user.User, error)
	// FindUser retrieves a User entity according to given queries.
	FindUser(queries []database.Query) (user.User, error)
	// GetAllUsers retrieves all User entities.
	GetAllUsers() ([]user.User, error)
	// GetAllUsersPaginated retrieves all User entities in a paginated manner.
	GetAllUsersPaginated(pagination database.Pagination) ([]user.User, database.Pagination, error)
	// QueryUsers retrieves all User entities according to given queries.
	QueryUsers(queries []database.Query, pagination database.Pagination) ([]user.User, error)
	// QueryUsersGroup retrieves all User entities according to given queries.
	QueryUsersGroup(queries []database.Query) ([]user.User, error)
	// CreateUser creates a new User entity.
	CreateUser(entity user.User, editorID *string) (string, error)
	// UpdateUser updates an existing User entity.
	UpdateUser(entity user.User, editorID *string) (string, error)
	// SaveUser saves a User entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveUser(entity user.User, editorID *string) (string, error)
	// DeleteUser deletes a User entity by its parents' path and userID.
	DeleteUser(userID string) error
	// GetCategoryLocalisation retrieves a single CategoryLocalisation entity by its ID and categoryLocalisationID.
	GetCategoryLocalisation(categoryID string, categoryLocalisationID string) (category_localisation.CategoryLocalisation, error)
	// GetCategoryLocalisationOrNew retrieves a single CategoryLocalisation entity by its ID and categoryLocalisationID.
	GetCategoryLocalisationOrNew(categoryID string, categoryLocalisationID string) (category_localisation.CategoryLocalisation, error)
	// GetCategoryLocalisation retrieves a single CategoryLocalisation entity by its document path.
	GetCategoryLocalisationByPath(path string) (category_localisation.CategoryLocalisation, error)
	// FindCategoryLocalisation retrieves a CategoryLocalisation entity according to given queries.
	FindCategoryLocalisation(categoryID string, queries []database.Query) (category_localisation.CategoryLocalisation, error)
	// GetAllCategoryLocalisations retrieves all CategoryLocalisation entities.
	GetAllCategoryLocalisations(categoryID string) ([]category_localisation.CategoryLocalisation, error)
	// GetAllCategoryLocalisationsPaginated retrieves all CategoryLocalisation entities in a paginated manner.
	GetAllCategoryLocalisationsPaginated(categoryID string, pagination database.Pagination) ([]category_localisation.CategoryLocalisation, database.Pagination, error)
	// QueryCategoryLocalisations retrieves all CategoryLocalisation entities according to given queries.
	QueryCategoryLocalisations(categoryID string, queries []database.Query, pagination database.Pagination) ([]category_localisation.CategoryLocalisation, error)
	// QueryCategoryLocalisationsGroup retrieves all CategoryLocalisation entities according to given queries.
	QueryCategoryLocalisationsGroup(queries []database.Query) ([]category_localisation.CategoryLocalisation, error)
	// CreateCategoryLocalisation creates a new CategoryLocalisation entity.
	CreateCategoryLocalisation(categoryID string, entity category_localisation.CategoryLocalisation, editorID *string) (string, error)
	// UpdateCategoryLocalisation updates an existing CategoryLocalisation entity.
	UpdateCategoryLocalisation(categoryID string, entity category_localisation.CategoryLocalisation, editorID *string) (string, error)
	// SaveCategoryLocalisation saves a CategoryLocalisation entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveCategoryLocalisation(categoryID string, entity category_localisation.CategoryLocalisation, editorID *string) (string, error)
	// DeleteCategoryLocalisation deletes a CategoryLocalisation entity by its parents' path and categoryLocalisationID.
	DeleteCategoryLocalisation(categoryID string, categoryLocalisationID string) error
	// GetGameModusLocalisation retrieves a single GameModusLocalisation entity by its ID and gameModusLocalisationID.
	GetGameModusLocalisation(gameModusID string, gameModusLocalisationID string) (game_modus_localisation.GameModusLocalisation, error)
	// GetGameModusLocalisationOrNew retrieves a single GameModusLocalisation entity by its ID and gameModusLocalisationID.
	GetGameModusLocalisationOrNew(gameModusID string, gameModusLocalisationID string) (game_modus_localisation.GameModusLocalisation, error)
	// GetGameModusLocalisation retrieves a single GameModusLocalisation entity by its document path.
	GetGameModusLocalisationByPath(path string) (game_modus_localisation.GameModusLocalisation, error)
	// FindGameModusLocalisation retrieves a GameModusLocalisation entity according to given queries.
	FindGameModusLocalisation(gameModusID string, queries []database.Query) (game_modus_localisation.GameModusLocalisation, error)
	// GetAllGameModusLocalisations retrieves all GameModusLocalisation entities.
	GetAllGameModusLocalisations(gameModusID string) ([]game_modus_localisation.GameModusLocalisation, error)
	// GetAllGameModusLocalisationsPaginated retrieves all GameModusLocalisation entities in a paginated manner.
	GetAllGameModusLocalisationsPaginated(gameModusID string, pagination database.Pagination) ([]game_modus_localisation.GameModusLocalisation, database.Pagination, error)
	// QueryGameModusLocalisations retrieves all GameModusLocalisation entities according to given queries.
	QueryGameModusLocalisations(gameModusID string, queries []database.Query, pagination database.Pagination) ([]game_modus_localisation.GameModusLocalisation, error)
	// QueryGameModusLocalisationsGroup retrieves all GameModusLocalisation entities according to given queries.
	QueryGameModusLocalisationsGroup(queries []database.Query) ([]game_modus_localisation.GameModusLocalisation, error)
	// CreateGameModusLocalisation creates a new GameModusLocalisation entity.
	CreateGameModusLocalisation(gameModusID string, entity game_modus_localisation.GameModusLocalisation, editorID *string) (string, error)
	// UpdateGameModusLocalisation updates an existing GameModusLocalisation entity.
	UpdateGameModusLocalisation(gameModusID string, entity game_modus_localisation.GameModusLocalisation, editorID *string) (string, error)
	// SaveGameModusLocalisation saves a GameModusLocalisation entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveGameModusLocalisation(gameModusID string, entity game_modus_localisation.GameModusLocalisation, editorID *string) (string, error)
	// DeleteGameModusLocalisation deletes a GameModusLocalisation entity by its parents' path and gameModusLocalisationID.
	DeleteGameModusLocalisation(gameModusID string, gameModusLocalisationID string) error
	// GetProductPackageLocalisation retrieves a single ProductPackageLocalisation entity by its ID and productPackageLocalisationID.
	GetProductPackageLocalisation(productPackageID string, productPackageLocalisationID string) (product_package_localisation.ProductPackageLocalisation, error)
	// GetProductPackageLocalisationOrNew retrieves a single ProductPackageLocalisation entity by its ID and productPackageLocalisationID.
	GetProductPackageLocalisationOrNew(productPackageID string, productPackageLocalisationID string) (product_package_localisation.ProductPackageLocalisation, error)
	// GetProductPackageLocalisation retrieves a single ProductPackageLocalisation entity by its document path.
	GetProductPackageLocalisationByPath(path string) (product_package_localisation.ProductPackageLocalisation, error)
	// FindProductPackageLocalisation retrieves a ProductPackageLocalisation entity according to given queries.
	FindProductPackageLocalisation(productPackageID string, queries []database.Query) (product_package_localisation.ProductPackageLocalisation, error)
	// GetAllProductPackageLocalisations retrieves all ProductPackageLocalisation entities.
	GetAllProductPackageLocalisations(productPackageID string) ([]product_package_localisation.ProductPackageLocalisation, error)
	// GetAllProductPackageLocalisationsPaginated retrieves all ProductPackageLocalisation entities in a paginated manner.
	GetAllProductPackageLocalisationsPaginated(productPackageID string, pagination database.Pagination) ([]product_package_localisation.ProductPackageLocalisation, database.Pagination, error)
	// QueryProductPackageLocalisations retrieves all ProductPackageLocalisation entities according to given queries.
	QueryProductPackageLocalisations(productPackageID string, queries []database.Query, pagination database.Pagination) ([]product_package_localisation.ProductPackageLocalisation, error)
	// QueryProductPackageLocalisationsGroup retrieves all ProductPackageLocalisation entities according to given queries.
	QueryProductPackageLocalisationsGroup(queries []database.Query) ([]product_package_localisation.ProductPackageLocalisation, error)
	// CreateProductPackageLocalisation creates a new ProductPackageLocalisation entity.
	CreateProductPackageLocalisation(productPackageID string, entity product_package_localisation.ProductPackageLocalisation, editorID *string) (string, error)
	// UpdateProductPackageLocalisation updates an existing ProductPackageLocalisation entity.
	UpdateProductPackageLocalisation(productPackageID string, entity product_package_localisation.ProductPackageLocalisation, editorID *string) (string, error)
	// SaveProductPackageLocalisation saves a ProductPackageLocalisation entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveProductPackageLocalisation(productPackageID string, entity product_package_localisation.ProductPackageLocalisation, editorID *string) (string, error)
	// DeleteProductPackageLocalisation deletes a ProductPackageLocalisation entity by its parents' path and productPackageLocalisationID.
	DeleteProductPackageLocalisation(productPackageID string, productPackageLocalisationID string) error
	// GetQuestionBundleLocalisation retrieves a single QuestionBundleLocalisation entity by its ID and questionBundleLocalisationID.
	GetQuestionBundleLocalisation(questionBundleID string, questionBundleLocalisationID string) (question_bundle_localisation.QuestionBundleLocalisation, error)
	// GetQuestionBundleLocalisationOrNew retrieves a single QuestionBundleLocalisation entity by its ID and questionBundleLocalisationID.
	GetQuestionBundleLocalisationOrNew(questionBundleID string, questionBundleLocalisationID string) (question_bundle_localisation.QuestionBundleLocalisation, error)
	// GetQuestionBundleLocalisation retrieves a single QuestionBundleLocalisation entity by its document path.
	GetQuestionBundleLocalisationByPath(path string) (question_bundle_localisation.QuestionBundleLocalisation, error)
	// FindQuestionBundleLocalisation retrieves a QuestionBundleLocalisation entity according to given queries.
	FindQuestionBundleLocalisation(questionBundleID string, queries []database.Query) (question_bundle_localisation.QuestionBundleLocalisation, error)
	// GetAllQuestionBundleLocalisations retrieves all QuestionBundleLocalisation entities.
	GetAllQuestionBundleLocalisations(questionBundleID string) ([]question_bundle_localisation.QuestionBundleLocalisation, error)
	// GetAllQuestionBundleLocalisationsPaginated retrieves all QuestionBundleLocalisation entities in a paginated manner.
	GetAllQuestionBundleLocalisationsPaginated(questionBundleID string, pagination database.Pagination) ([]question_bundle_localisation.QuestionBundleLocalisation, database.Pagination, error)
	// QueryQuestionBundleLocalisations retrieves all QuestionBundleLocalisation entities according to given queries.
	QueryQuestionBundleLocalisations(questionBundleID string, queries []database.Query, pagination database.Pagination) ([]question_bundle_localisation.QuestionBundleLocalisation, error)
	// QueryQuestionBundleLocalisationsGroup retrieves all QuestionBundleLocalisation entities according to given queries.
	QueryQuestionBundleLocalisationsGroup(queries []database.Query) ([]question_bundle_localisation.QuestionBundleLocalisation, error)
	// CreateQuestionBundleLocalisation creates a new QuestionBundleLocalisation entity.
	CreateQuestionBundleLocalisation(questionBundleID string, entity question_bundle_localisation.QuestionBundleLocalisation, editorID *string) (string, error)
	// UpdateQuestionBundleLocalisation updates an existing QuestionBundleLocalisation entity.
	UpdateQuestionBundleLocalisation(questionBundleID string, entity question_bundle_localisation.QuestionBundleLocalisation, editorID *string) (string, error)
	// SaveQuestionBundleLocalisation saves a QuestionBundleLocalisation entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveQuestionBundleLocalisation(questionBundleID string, entity question_bundle_localisation.QuestionBundleLocalisation, editorID *string) (string, error)
	// DeleteQuestionBundleLocalisation deletes a QuestionBundleLocalisation entity by its parents' path and questionBundleLocalisationID.
	DeleteQuestionBundleLocalisation(questionBundleID string, questionBundleLocalisationID string) error
	// GetQuestionContextLocalisation retrieves a single QuestionContextLocalisation entity by its ID and questionContextLocalisationID.
	GetQuestionContextLocalisation(questionContextID string, questionContextLocalisationID string) (question_context_localisation.QuestionContextLocalisation, error)
	// GetQuestionContextLocalisationOrNew retrieves a single QuestionContextLocalisation entity by its ID and questionContextLocalisationID.
	GetQuestionContextLocalisationOrNew(questionContextID string, questionContextLocalisationID string) (question_context_localisation.QuestionContextLocalisation, error)
	// GetQuestionContextLocalisation retrieves a single QuestionContextLocalisation entity by its document path.
	GetQuestionContextLocalisationByPath(path string) (question_context_localisation.QuestionContextLocalisation, error)
	// FindQuestionContextLocalisation retrieves a QuestionContextLocalisation entity according to given queries.
	FindQuestionContextLocalisation(questionContextID string, queries []database.Query) (question_context_localisation.QuestionContextLocalisation, error)
	// GetAllQuestionContextLocalisations retrieves all QuestionContextLocalisation entities.
	GetAllQuestionContextLocalisations(questionContextID string) ([]question_context_localisation.QuestionContextLocalisation, error)
	// GetAllQuestionContextLocalisationsPaginated retrieves all QuestionContextLocalisation entities in a paginated manner.
	GetAllQuestionContextLocalisationsPaginated(questionContextID string, pagination database.Pagination) ([]question_context_localisation.QuestionContextLocalisation, database.Pagination, error)
	// QueryQuestionContextLocalisations retrieves all QuestionContextLocalisation entities according to given queries.
	QueryQuestionContextLocalisations(questionContextID string, queries []database.Query, pagination database.Pagination) ([]question_context_localisation.QuestionContextLocalisation, error)
	// QueryQuestionContextLocalisationsGroup retrieves all QuestionContextLocalisation entities according to given queries.
	QueryQuestionContextLocalisationsGroup(queries []database.Query) ([]question_context_localisation.QuestionContextLocalisation, error)
	// CreateQuestionContextLocalisation creates a new QuestionContextLocalisation entity.
	CreateQuestionContextLocalisation(questionContextID string, entity question_context_localisation.QuestionContextLocalisation, editorID *string) (string, error)
	// UpdateQuestionContextLocalisation updates an existing QuestionContextLocalisation entity.
	UpdateQuestionContextLocalisation(questionContextID string, entity question_context_localisation.QuestionContextLocalisation, editorID *string) (string, error)
	// SaveQuestionContextLocalisation saves a QuestionContextLocalisation entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveQuestionContextLocalisation(questionContextID string, entity question_context_localisation.QuestionContextLocalisation, editorID *string) (string, error)
	// DeleteQuestionContextLocalisation deletes a QuestionContextLocalisation entity by its parents' path and questionContextLocalisationID.
	DeleteQuestionContextLocalisation(questionContextID string, questionContextLocalisationID string) error
	// GetTagLocalisation retrieves a single TagLocalisation entity by its ID and tagLocalisationID.
	GetTagLocalisation(tagID string, tagLocalisationID string) (tag_localisation.TagLocalisation, error)
	// GetTagLocalisationOrNew retrieves a single TagLocalisation entity by its ID and tagLocalisationID.
	GetTagLocalisationOrNew(tagID string, tagLocalisationID string) (tag_localisation.TagLocalisation, error)
	// GetTagLocalisation retrieves a single TagLocalisation entity by its document path.
	GetTagLocalisationByPath(path string) (tag_localisation.TagLocalisation, error)
	// FindTagLocalisation retrieves a TagLocalisation entity according to given queries.
	FindTagLocalisation(tagID string, queries []database.Query) (tag_localisation.TagLocalisation, error)
	// GetAllTagLocalisations retrieves all TagLocalisation entities.
	GetAllTagLocalisations(tagID string) ([]tag_localisation.TagLocalisation, error)
	// GetAllTagLocalisationsPaginated retrieves all TagLocalisation entities in a paginated manner.
	GetAllTagLocalisationsPaginated(tagID string, pagination database.Pagination) ([]tag_localisation.TagLocalisation, database.Pagination, error)
	// QueryTagLocalisations retrieves all TagLocalisation entities according to given queries.
	QueryTagLocalisations(tagID string, queries []database.Query, pagination database.Pagination) ([]tag_localisation.TagLocalisation, error)
	// QueryTagLocalisationsGroup retrieves all TagLocalisation entities according to given queries.
	QueryTagLocalisationsGroup(queries []database.Query) ([]tag_localisation.TagLocalisation, error)
	// CreateTagLocalisation creates a new TagLocalisation entity.
	CreateTagLocalisation(tagID string, entity tag_localisation.TagLocalisation, editorID *string) (string, error)
	// UpdateTagLocalisation updates an existing TagLocalisation entity.
	UpdateTagLocalisation(tagID string, entity tag_localisation.TagLocalisation, editorID *string) (string, error)
	// SaveTagLocalisation saves a TagLocalisation entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveTagLocalisation(tagID string, entity tag_localisation.TagLocalisation, editorID *string) (string, error)
	// DeleteTagLocalisation deletes a TagLocalisation entity by its parents' path and tagLocalisationID.
	DeleteTagLocalisation(tagID string, tagLocalisationID string) error
}


type Repository struct {}
func NewRepository() IRepository {
	return Repository{}
}
func (Repository) GetAuthor(authorID string) (author.Author, error) {
	return GetAuthor(authorID)
}
func (Repository) GetAuthorOrNew(authorID string) (author.Author, error) {
	return GetAuthorOrNew(authorID)
}
func (Repository) GetAuthorByPath(path string) (author.Author, error) {
	return GetAuthorByPath(path)
}
func (Repository) FindAuthor(queries []database.Query) (author.Author, error) {
	return FindAuthor(queries)
}
func (Repository) GetAllAuthors() ([]author.Author, error) {
	return GetAllAuthors()
}
func (Repository) GetAllAuthorsPaginated(pagination database.Pagination) ([]author.Author, database.Pagination, error) {
	return GetAllAuthorsPaginated(pagination)
}
func (Repository) QueryAuthors(queries []database.Query, pagination database.Pagination) ([]author.Author, error) {
	return QueryAuthors(queries, pagination)
}
func (Repository) QueryAuthorsGroup(queries []database.Query) ([]author.Author, error) {
	return QueryAuthorsGroup(queries)
}
func (Repository) CreateAuthor(entity author.Author, editorID *string) (string, error) {
	return CreateAuthor(entity, editorID)
}
func (Repository) UpdateAuthor(entity author.Author, editorID *string) (string, error) {
	return UpdateAuthor(entity, editorID)
}
func (Repository) SaveAuthor(entity author.Author, editorID *string) (string, error) {
	return SaveAuthor(entity, editorID)
}
func (Repository) DeleteAuthor(authorID string) error {
	return DeleteAuthor(authorID)
}
func (Repository) GetCategory(categoryID string) (category.Category, error) {
	return GetCategory(categoryID)
}
func (Repository) GetCategoryOrNew(categoryID string) (category.Category, error) {
	return GetCategoryOrNew(categoryID)
}
func (Repository) GetCategoryByPath(path string) (category.Category, error) {
	return GetCategoryByPath(path)
}
func (Repository) FindCategory(queries []database.Query) (category.Category, error) {
	return FindCategory(queries)
}
func (Repository) GetAllCategories() ([]category.Category, error) {
	return GetAllCategories()
}
func (Repository) GetAllCategoriesPaginated(pagination database.Pagination) ([]category.Category, database.Pagination, error) {
	return GetAllCategoriesPaginated(pagination)
}
func (Repository) QueryCategories(queries []database.Query, pagination database.Pagination) ([]category.Category, error) {
	return QueryCategories(queries, pagination)
}
func (Repository) QueryCategoriesGroup(queries []database.Query) ([]category.Category, error) {
	return QueryCategoriesGroup(queries)
}
func (Repository) CreateCategory(entity category.Category, editorID *string) (string, error) {
	return CreateCategory(entity, editorID)
}
func (Repository) UpdateCategory(entity category.Category, editorID *string) (string, error) {
	return UpdateCategory(entity, editorID)
}
func (Repository) SaveCategory(entity category.Category, editorID *string) (string, error) {
	return SaveCategory(entity, editorID)
}
func (Repository) DeleteCategory(categoryID string) error {
	return DeleteCategory(categoryID)
}
func (Repository) GetGameModus(gameModusID string) (game_modus.GameModus, error) {
	return GetGameModus(gameModusID)
}
func (Repository) GetGameModusOrNew(gameModusID string) (game_modus.GameModus, error) {
	return GetGameModusOrNew(gameModusID)
}
func (Repository) GetGameModusByPath(path string) (game_modus.GameModus, error) {
	return GetGameModusByPath(path)
}
func (Repository) FindGameModus(queries []database.Query) (game_modus.GameModus, error) {
	return FindGameModus(queries)
}
func (Repository) GetAllGameModus() ([]game_modus.GameModus, error) {
	return GetAllGameModus()
}
func (Repository) GetAllGameModusPaginated(pagination database.Pagination) ([]game_modus.GameModus, database.Pagination, error) {
	return GetAllGameModusPaginated(pagination)
}
func (Repository) QueryGameModus(queries []database.Query, pagination database.Pagination) ([]game_modus.GameModus, error) {
	return QueryGameModus(queries, pagination)
}
func (Repository) QueryGameModusGroup(queries []database.Query) ([]game_modus.GameModus, error) {
	return QueryGameModusGroup(queries)
}
func (Repository) CreateGameModus(entity game_modus.GameModus, editorID *string) (string, error) {
	return CreateGameModus(entity, editorID)
}
func (Repository) UpdateGameModus(entity game_modus.GameModus, editorID *string) (string, error) {
	return UpdateGameModus(entity, editorID)
}
func (Repository) SaveGameModus(entity game_modus.GameModus, editorID *string) (string, error) {
	return SaveGameModus(entity, editorID)
}
func (Repository) DeleteGameModus(gameModusID string) error {
	return DeleteGameModus(gameModusID)
}
func (Repository) GetMedia(mediaID string) (media.Media, error) {
	return GetMedia(mediaID)
}
func (Repository) GetMediaOrNew(mediaID string) (media.Media, error) {
	return GetMediaOrNew(mediaID)
}
func (Repository) GetMediaByPath(path string) (media.Media, error) {
	return GetMediaByPath(path)
}
func (Repository) FindMedia(queries []database.Query) (media.Media, error) {
	return FindMedia(queries)
}
func (Repository) GetAllMedia() ([]media.Media, error) {
	return GetAllMedia()
}
func (Repository) GetAllMediaPaginated(pagination database.Pagination) ([]media.Media, database.Pagination, error) {
	return GetAllMediaPaginated(pagination)
}
func (Repository) QueryMedia(queries []database.Query, pagination database.Pagination) ([]media.Media, error) {
	return QueryMedia(queries, pagination)
}
func (Repository) QueryMediaGroup(queries []database.Query) ([]media.Media, error) {
	return QueryMediaGroup(queries)
}
func (Repository) CreateMedia(entity media.Media, editorID *string) (string, error) {
	return CreateMedia(entity, editorID)
}
func (Repository) UpdateMedia(entity media.Media, editorID *string) (string, error) {
	return UpdateMedia(entity, editorID)
}
func (Repository) SaveMedia(entity media.Media, editorID *string) (string, error) {
	return SaveMedia(entity, editorID)
}
func (Repository) DeleteMedia(mediaID string) error {
	return DeleteMedia(mediaID)
}
func (Repository) GetProductPackage(productPackageID string) (product_package.ProductPackage, error) {
	return GetProductPackage(productPackageID)
}
func (Repository) GetProductPackageOrNew(productPackageID string) (product_package.ProductPackage, error) {
	return GetProductPackageOrNew(productPackageID)
}
func (Repository) GetProductPackageByPath(path string) (product_package.ProductPackage, error) {
	return GetProductPackageByPath(path)
}
func (Repository) FindProductPackage(queries []database.Query) (product_package.ProductPackage, error) {
	return FindProductPackage(queries)
}
func (Repository) GetAllProductPackages() ([]product_package.ProductPackage, error) {
	return GetAllProductPackages()
}
func (Repository) GetAllProductPackagesPaginated(pagination database.Pagination) ([]product_package.ProductPackage, database.Pagination, error) {
	return GetAllProductPackagesPaginated(pagination)
}
func (Repository) QueryProductPackages(queries []database.Query, pagination database.Pagination) ([]product_package.ProductPackage, error) {
	return QueryProductPackages(queries, pagination)
}
func (Repository) QueryProductPackagesGroup(queries []database.Query) ([]product_package.ProductPackage, error) {
	return QueryProductPackagesGroup(queries)
}
func (Repository) CreateProductPackage(entity product_package.ProductPackage, editorID *string) (string, error) {
	return CreateProductPackage(entity, editorID)
}
func (Repository) UpdateProductPackage(entity product_package.ProductPackage, editorID *string) (string, error) {
	return UpdateProductPackage(entity, editorID)
}
func (Repository) SaveProductPackage(entity product_package.ProductPackage, editorID *string) (string, error) {
	return SaveProductPackage(entity, editorID)
}
func (Repository) DeleteProductPackage(productPackageID string) error {
	return DeleteProductPackage(productPackageID)
}
func (Repository) GetPurchase(userID string, purchaseID string) (purchase.Purchase, error) {
	return GetPurchase(userID, purchaseID)
}
func (Repository) GetPurchaseOrNew(userID string, purchaseID string) (purchase.Purchase, error) {
	return GetPurchaseOrNew(userID, purchaseID)
}
func (Repository) GetPurchaseByPath(path string) (purchase.Purchase, error) {
	return GetPurchaseByPath(path)
}
func (Repository) FindPurchase(userID string, queries []database.Query) (purchase.Purchase, error) {
	return FindPurchase(userID, queries)
}
func (Repository) GetAllPurchases(userID string) ([]purchase.Purchase, error) {
	return GetAllPurchases(userID)
}
func (Repository) GetAllPurchasesPaginated(userID string, pagination database.Pagination) ([]purchase.Purchase, database.Pagination, error) {
	return GetAllPurchasesPaginated(userID, pagination)
}
func (Repository) QueryPurchases(userID string, queries []database.Query, pagination database.Pagination) ([]purchase.Purchase, error) {
	return QueryPurchases(userID, queries, pagination)
}
func (Repository) QueryPurchasesGroup(queries []database.Query) ([]purchase.Purchase, error) {
	return QueryPurchasesGroup(queries)
}
func (Repository) CreatePurchase(userID string, entity purchase.Purchase, editorID *string) (string, error) {
	return CreatePurchase(userID, entity, editorID)
}
func (Repository) UpdatePurchase(userID string, entity purchase.Purchase, editorID *string) (string, error) {
	return UpdatePurchase(userID, entity, editorID)
}
func (Repository) SavePurchase(userID string, entity purchase.Purchase, editorID *string) (string, error) {
	return SavePurchase(userID, entity, editorID)
}
func (Repository) DeletePurchase(userID string, purchaseID string) error {
	return DeletePurchase(userID, purchaseID)
}
func (Repository) GetQuestionBundle(questionBundleID string) (question_bundle.QuestionBundle, error) {
	return GetQuestionBundle(questionBundleID)
}
func (Repository) GetQuestionBundleOrNew(questionBundleID string) (question_bundle.QuestionBundle, error) {
	return GetQuestionBundleOrNew(questionBundleID)
}
func (Repository) GetQuestionBundleByPath(path string) (question_bundle.QuestionBundle, error) {
	return GetQuestionBundleByPath(path)
}
func (Repository) FindQuestionBundle(queries []database.Query) (question_bundle.QuestionBundle, error) {
	return FindQuestionBundle(queries)
}
func (Repository) GetAllQuestionBundles() ([]question_bundle.QuestionBundle, error) {
	return GetAllQuestionBundles()
}
func (Repository) GetAllQuestionBundlesPaginated(pagination database.Pagination) ([]question_bundle.QuestionBundle, database.Pagination, error) {
	return GetAllQuestionBundlesPaginated(pagination)
}
func (Repository) QueryQuestionBundles(queries []database.Query, pagination database.Pagination) ([]question_bundle.QuestionBundle, error) {
	return QueryQuestionBundles(queries, pagination)
}
func (Repository) QueryQuestionBundlesGroup(queries []database.Query) ([]question_bundle.QuestionBundle, error) {
	return QueryQuestionBundlesGroup(queries)
}
func (Repository) CreateQuestionBundle(entity question_bundle.QuestionBundle, editorID *string) (string, error) {
	return CreateQuestionBundle(entity, editorID)
}
func (Repository) UpdateQuestionBundle(entity question_bundle.QuestionBundle, editorID *string) (string, error) {
	return UpdateQuestionBundle(entity, editorID)
}
func (Repository) SaveQuestionBundle(entity question_bundle.QuestionBundle, editorID *string) (string, error) {
	return SaveQuestionBundle(entity, editorID)
}
func (Repository) DeleteQuestionBundle(questionBundleID string) error {
	return DeleteQuestionBundle(questionBundleID)
}
func (Repository) GetQuestionContext(questionContextID string) (question_context.QuestionContext, error) {
	return GetQuestionContext(questionContextID)
}
func (Repository) GetQuestionContextOrNew(questionContextID string) (question_context.QuestionContext, error) {
	return GetQuestionContextOrNew(questionContextID)
}
func (Repository) GetQuestionContextByPath(path string) (question_context.QuestionContext, error) {
	return GetQuestionContextByPath(path)
}
func (Repository) FindQuestionContext(queries []database.Query) (question_context.QuestionContext, error) {
	return FindQuestionContext(queries)
}
func (Repository) GetAllQuestionContexts() ([]question_context.QuestionContext, error) {
	return GetAllQuestionContexts()
}
func (Repository) GetAllQuestionContextsPaginated(pagination database.Pagination) ([]question_context.QuestionContext, database.Pagination, error) {
	return GetAllQuestionContextsPaginated(pagination)
}
func (Repository) QueryQuestionContexts(queries []database.Query, pagination database.Pagination) ([]question_context.QuestionContext, error) {
	return QueryQuestionContexts(queries, pagination)
}
func (Repository) QueryQuestionContextsGroup(queries []database.Query) ([]question_context.QuestionContext, error) {
	return QueryQuestionContextsGroup(queries)
}
func (Repository) CreateQuestionContext(entity question_context.QuestionContext, editorID *string) (string, error) {
	return CreateQuestionContext(entity, editorID)
}
func (Repository) UpdateQuestionContext(entity question_context.QuestionContext, editorID *string) (string, error) {
	return UpdateQuestionContext(entity, editorID)
}
func (Repository) SaveQuestionContext(entity question_context.QuestionContext, editorID *string) (string, error) {
	return SaveQuestionContext(entity, editorID)
}
func (Repository) DeleteQuestionContext(questionContextID string) error {
	return DeleteQuestionContext(questionContextID)
}
func (Repository) GetTag(tagID string) (tag.Tag, error) {
	return GetTag(tagID)
}
func (Repository) GetTagOrNew(tagID string) (tag.Tag, error) {
	return GetTagOrNew(tagID)
}
func (Repository) GetTagByPath(path string) (tag.Tag, error) {
	return GetTagByPath(path)
}
func (Repository) FindTag(queries []database.Query) (tag.Tag, error) {
	return FindTag(queries)
}
func (Repository) GetAllTags() ([]tag.Tag, error) {
	return GetAllTags()
}
func (Repository) GetAllTagsPaginated(pagination database.Pagination) ([]tag.Tag, database.Pagination, error) {
	return GetAllTagsPaginated(pagination)
}
func (Repository) QueryTags(queries []database.Query, pagination database.Pagination) ([]tag.Tag, error) {
	return QueryTags(queries, pagination)
}
func (Repository) QueryTagsGroup(queries []database.Query) ([]tag.Tag, error) {
	return QueryTagsGroup(queries)
}
func (Repository) CreateTag(entity tag.Tag, editorID *string) (string, error) {
	return CreateTag(entity, editorID)
}
func (Repository) UpdateTag(entity tag.Tag, editorID *string) (string, error) {
	return UpdateTag(entity, editorID)
}
func (Repository) SaveTag(entity tag.Tag, editorID *string) (string, error) {
	return SaveTag(entity, editorID)
}
func (Repository) DeleteTag(tagID string) error {
	return DeleteTag(tagID)
}
func (Repository) GetUser(userID string) (user.User, error) {
	return GetUser(userID)
}
func (Repository) GetUserOrNew(userID string) (user.User, error) {
	return GetUserOrNew(userID)
}
func (Repository) GetUserByPath(path string) (user.User, error) {
	return GetUserByPath(path)
}
func (Repository) FindUser(queries []database.Query) (user.User, error) {
	return FindUser(queries)
}
func (Repository) GetAllUsers() ([]user.User, error) {
	return GetAllUsers()
}
func (Repository) GetAllUsersPaginated(pagination database.Pagination) ([]user.User, database.Pagination, error) {
	return GetAllUsersPaginated(pagination)
}
func (Repository) QueryUsers(queries []database.Query, pagination database.Pagination) ([]user.User, error) {
	return QueryUsers(queries, pagination)
}
func (Repository) QueryUsersGroup(queries []database.Query) ([]user.User, error) {
	return QueryUsersGroup(queries)
}
func (Repository) CreateUser(entity user.User, editorID *string) (string, error) {
	return CreateUser(entity, editorID)
}
func (Repository) UpdateUser(entity user.User, editorID *string) (string, error) {
	return UpdateUser(entity, editorID)
}
func (Repository) SaveUser(entity user.User, editorID *string) (string, error) {
	return SaveUser(entity, editorID)
}
func (Repository) DeleteUser(userID string) error {
	return DeleteUser(userID)
}
func (Repository) GetCategoryLocalisation(categoryID string, categoryLocalisationID string) (category_localisation.CategoryLocalisation, error) {
	return GetCategoryLocalisation(categoryID, categoryLocalisationID)
}
func (Repository) GetCategoryLocalisationOrNew(categoryID string, categoryLocalisationID string) (category_localisation.CategoryLocalisation, error) {
	return GetCategoryLocalisationOrNew(categoryID, categoryLocalisationID)
}
func (Repository) GetCategoryLocalisationByPath(path string) (category_localisation.CategoryLocalisation, error) {
	return GetCategoryLocalisationByPath(path)
}
func (Repository) FindCategoryLocalisation(categoryID string, queries []database.Query) (category_localisation.CategoryLocalisation, error) {
	return FindCategoryLocalisation(categoryID, queries)
}
func (Repository) GetAllCategoryLocalisations(categoryID string) ([]category_localisation.CategoryLocalisation, error) {
	return GetAllCategoryLocalisations(categoryID)
}
func (Repository) GetAllCategoryLocalisationsPaginated(categoryID string, pagination database.Pagination) ([]category_localisation.CategoryLocalisation, database.Pagination, error) {
	return GetAllCategoryLocalisationsPaginated(categoryID, pagination)
}
func (Repository) QueryCategoryLocalisations(categoryID string, queries []database.Query, pagination database.Pagination) ([]category_localisation.CategoryLocalisation, error) {
	return QueryCategoryLocalisations(categoryID, queries, pagination)
}
func (Repository) QueryCategoryLocalisationsGroup(queries []database.Query) ([]category_localisation.CategoryLocalisation, error) {
	return QueryCategoryLocalisationsGroup(queries)
}
func (Repository) CreateCategoryLocalisation(categoryID string, entity category_localisation.CategoryLocalisation, editorID *string) (string, error) {
	return CreateCategoryLocalisation(categoryID, entity, editorID)
}
func (Repository) UpdateCategoryLocalisation(categoryID string, entity category_localisation.CategoryLocalisation, editorID *string) (string, error) {
	return UpdateCategoryLocalisation(categoryID, entity, editorID)
}
func (Repository) SaveCategoryLocalisation(categoryID string, entity category_localisation.CategoryLocalisation, editorID *string) (string, error) {
	return SaveCategoryLocalisation(categoryID, entity, editorID)
}
func (Repository) DeleteCategoryLocalisation(categoryID string, categoryLocalisationID string) error {
	return DeleteCategoryLocalisation(categoryID, categoryLocalisationID)
}
func (Repository) GetGameModusLocalisation(gameModusID string, gameModusLocalisationID string) (game_modus_localisation.GameModusLocalisation, error) {
	return GetGameModusLocalisation(gameModusID, gameModusLocalisationID)
}
func (Repository) GetGameModusLocalisationOrNew(gameModusID string, gameModusLocalisationID string) (game_modus_localisation.GameModusLocalisation, error) {
	return GetGameModusLocalisationOrNew(gameModusID, gameModusLocalisationID)
}
func (Repository) GetGameModusLocalisationByPath(path string) (game_modus_localisation.GameModusLocalisation, error) {
	return GetGameModusLocalisationByPath(path)
}
func (Repository) FindGameModusLocalisation(gameModusID string, queries []database.Query) (game_modus_localisation.GameModusLocalisation, error) {
	return FindGameModusLocalisation(gameModusID, queries)
}
func (Repository) GetAllGameModusLocalisations(gameModusID string) ([]game_modus_localisation.GameModusLocalisation, error) {
	return GetAllGameModusLocalisations(gameModusID)
}
func (Repository) GetAllGameModusLocalisationsPaginated(gameModusID string, pagination database.Pagination) ([]game_modus_localisation.GameModusLocalisation, database.Pagination, error) {
	return GetAllGameModusLocalisationsPaginated(gameModusID, pagination)
}
func (Repository) QueryGameModusLocalisations(gameModusID string, queries []database.Query, pagination database.Pagination) ([]game_modus_localisation.GameModusLocalisation, error) {
	return QueryGameModusLocalisations(gameModusID, queries, pagination)
}
func (Repository) QueryGameModusLocalisationsGroup(queries []database.Query) ([]game_modus_localisation.GameModusLocalisation, error) {
	return QueryGameModusLocalisationsGroup(queries)
}
func (Repository) CreateGameModusLocalisation(gameModusID string, entity game_modus_localisation.GameModusLocalisation, editorID *string) (string, error) {
	return CreateGameModusLocalisation(gameModusID, entity, editorID)
}
func (Repository) UpdateGameModusLocalisation(gameModusID string, entity game_modus_localisation.GameModusLocalisation, editorID *string) (string, error) {
	return UpdateGameModusLocalisation(gameModusID, entity, editorID)
}
func (Repository) SaveGameModusLocalisation(gameModusID string, entity game_modus_localisation.GameModusLocalisation, editorID *string) (string, error) {
	return SaveGameModusLocalisation(gameModusID, entity, editorID)
}
func (Repository) DeleteGameModusLocalisation(gameModusID string, gameModusLocalisationID string) error {
	return DeleteGameModusLocalisation(gameModusID, gameModusLocalisationID)
}
func (Repository) GetProductPackageLocalisation(productPackageID string, productPackageLocalisationID string) (product_package_localisation.ProductPackageLocalisation, error) {
	return GetProductPackageLocalisation(productPackageID, productPackageLocalisationID)
}
func (Repository) GetProductPackageLocalisationOrNew(productPackageID string, productPackageLocalisationID string) (product_package_localisation.ProductPackageLocalisation, error) {
	return GetProductPackageLocalisationOrNew(productPackageID, productPackageLocalisationID)
}
func (Repository) GetProductPackageLocalisationByPath(path string) (product_package_localisation.ProductPackageLocalisation, error) {
	return GetProductPackageLocalisationByPath(path)
}
func (Repository) FindProductPackageLocalisation(productPackageID string, queries []database.Query) (product_package_localisation.ProductPackageLocalisation, error) {
	return FindProductPackageLocalisation(productPackageID, queries)
}
func (Repository) GetAllProductPackageLocalisations(productPackageID string) ([]product_package_localisation.ProductPackageLocalisation, error) {
	return GetAllProductPackageLocalisations(productPackageID)
}
func (Repository) GetAllProductPackageLocalisationsPaginated(productPackageID string, pagination database.Pagination) ([]product_package_localisation.ProductPackageLocalisation, database.Pagination, error) {
	return GetAllProductPackageLocalisationsPaginated(productPackageID, pagination)
}
func (Repository) QueryProductPackageLocalisations(productPackageID string, queries []database.Query, pagination database.Pagination) ([]product_package_localisation.ProductPackageLocalisation, error) {
	return QueryProductPackageLocalisations(productPackageID, queries, pagination)
}
func (Repository) QueryProductPackageLocalisationsGroup(queries []database.Query) ([]product_package_localisation.ProductPackageLocalisation, error) {
	return QueryProductPackageLocalisationsGroup(queries)
}
func (Repository) CreateProductPackageLocalisation(productPackageID string, entity product_package_localisation.ProductPackageLocalisation, editorID *string) (string, error) {
	return CreateProductPackageLocalisation(productPackageID, entity, editorID)
}
func (Repository) UpdateProductPackageLocalisation(productPackageID string, entity product_package_localisation.ProductPackageLocalisation, editorID *string) (string, error) {
	return UpdateProductPackageLocalisation(productPackageID, entity, editorID)
}
func (Repository) SaveProductPackageLocalisation(productPackageID string, entity product_package_localisation.ProductPackageLocalisation, editorID *string) (string, error) {
	return SaveProductPackageLocalisation(productPackageID, entity, editorID)
}
func (Repository) DeleteProductPackageLocalisation(productPackageID string, productPackageLocalisationID string) error {
	return DeleteProductPackageLocalisation(productPackageID, productPackageLocalisationID)
}
func (Repository) GetQuestionBundleLocalisation(questionBundleID string, questionBundleLocalisationID string) (question_bundle_localisation.QuestionBundleLocalisation, error) {
	return GetQuestionBundleLocalisation(questionBundleID, questionBundleLocalisationID)
}
func (Repository) GetQuestionBundleLocalisationOrNew(questionBundleID string, questionBundleLocalisationID string) (question_bundle_localisation.QuestionBundleLocalisation, error) {
	return GetQuestionBundleLocalisationOrNew(questionBundleID, questionBundleLocalisationID)
}
func (Repository) GetQuestionBundleLocalisationByPath(path string) (question_bundle_localisation.QuestionBundleLocalisation, error) {
	return GetQuestionBundleLocalisationByPath(path)
}
func (Repository) FindQuestionBundleLocalisation(questionBundleID string, queries []database.Query) (question_bundle_localisation.QuestionBundleLocalisation, error) {
	return FindQuestionBundleLocalisation(questionBundleID, queries)
}
func (Repository) GetAllQuestionBundleLocalisations(questionBundleID string) ([]question_bundle_localisation.QuestionBundleLocalisation, error) {
	return GetAllQuestionBundleLocalisations(questionBundleID)
}
func (Repository) GetAllQuestionBundleLocalisationsPaginated(questionBundleID string, pagination database.Pagination) ([]question_bundle_localisation.QuestionBundleLocalisation, database.Pagination, error) {
	return GetAllQuestionBundleLocalisationsPaginated(questionBundleID, pagination)
}
func (Repository) QueryQuestionBundleLocalisations(questionBundleID string, queries []database.Query, pagination database.Pagination) ([]question_bundle_localisation.QuestionBundleLocalisation, error) {
	return QueryQuestionBundleLocalisations(questionBundleID, queries, pagination)
}
func (Repository) QueryQuestionBundleLocalisationsGroup(queries []database.Query) ([]question_bundle_localisation.QuestionBundleLocalisation, error) {
	return QueryQuestionBundleLocalisationsGroup(queries)
}
func (Repository) CreateQuestionBundleLocalisation(questionBundleID string, entity question_bundle_localisation.QuestionBundleLocalisation, editorID *string) (string, error) {
	return CreateQuestionBundleLocalisation(questionBundleID, entity, editorID)
}
func (Repository) UpdateQuestionBundleLocalisation(questionBundleID string, entity question_bundle_localisation.QuestionBundleLocalisation, editorID *string) (string, error) {
	return UpdateQuestionBundleLocalisation(questionBundleID, entity, editorID)
}
func (Repository) SaveQuestionBundleLocalisation(questionBundleID string, entity question_bundle_localisation.QuestionBundleLocalisation, editorID *string) (string, error) {
	return SaveQuestionBundleLocalisation(questionBundleID, entity, editorID)
}
func (Repository) DeleteQuestionBundleLocalisation(questionBundleID string, questionBundleLocalisationID string) error {
	return DeleteQuestionBundleLocalisation(questionBundleID, questionBundleLocalisationID)
}
func (Repository) GetQuestionContextLocalisation(questionContextID string, questionContextLocalisationID string) (question_context_localisation.QuestionContextLocalisation, error) {
	return GetQuestionContextLocalisation(questionContextID, questionContextLocalisationID)
}
func (Repository) GetQuestionContextLocalisationOrNew(questionContextID string, questionContextLocalisationID string) (question_context_localisation.QuestionContextLocalisation, error) {
	return GetQuestionContextLocalisationOrNew(questionContextID, questionContextLocalisationID)
}
func (Repository) GetQuestionContextLocalisationByPath(path string) (question_context_localisation.QuestionContextLocalisation, error) {
	return GetQuestionContextLocalisationByPath(path)
}
func (Repository) FindQuestionContextLocalisation(questionContextID string, queries []database.Query) (question_context_localisation.QuestionContextLocalisation, error) {
	return FindQuestionContextLocalisation(questionContextID, queries)
}
func (Repository) GetAllQuestionContextLocalisations(questionContextID string) ([]question_context_localisation.QuestionContextLocalisation, error) {
	return GetAllQuestionContextLocalisations(questionContextID)
}
func (Repository) GetAllQuestionContextLocalisationsPaginated(questionContextID string, pagination database.Pagination) ([]question_context_localisation.QuestionContextLocalisation, database.Pagination, error) {
	return GetAllQuestionContextLocalisationsPaginated(questionContextID, pagination)
}
func (Repository) QueryQuestionContextLocalisations(questionContextID string, queries []database.Query, pagination database.Pagination) ([]question_context_localisation.QuestionContextLocalisation, error) {
	return QueryQuestionContextLocalisations(questionContextID, queries, pagination)
}
func (Repository) QueryQuestionContextLocalisationsGroup(queries []database.Query) ([]question_context_localisation.QuestionContextLocalisation, error) {
	return QueryQuestionContextLocalisationsGroup(queries)
}
func (Repository) CreateQuestionContextLocalisation(questionContextID string, entity question_context_localisation.QuestionContextLocalisation, editorID *string) (string, error) {
	return CreateQuestionContextLocalisation(questionContextID, entity, editorID)
}
func (Repository) UpdateQuestionContextLocalisation(questionContextID string, entity question_context_localisation.QuestionContextLocalisation, editorID *string) (string, error) {
	return UpdateQuestionContextLocalisation(questionContextID, entity, editorID)
}
func (Repository) SaveQuestionContextLocalisation(questionContextID string, entity question_context_localisation.QuestionContextLocalisation, editorID *string) (string, error) {
	return SaveQuestionContextLocalisation(questionContextID, entity, editorID)
}
func (Repository) DeleteQuestionContextLocalisation(questionContextID string, questionContextLocalisationID string) error {
	return DeleteQuestionContextLocalisation(questionContextID, questionContextLocalisationID)
}
func (Repository) GetTagLocalisation(tagID string, tagLocalisationID string) (tag_localisation.TagLocalisation, error) {
	return GetTagLocalisation(tagID, tagLocalisationID)
}
func (Repository) GetTagLocalisationOrNew(tagID string, tagLocalisationID string) (tag_localisation.TagLocalisation, error) {
	return GetTagLocalisationOrNew(tagID, tagLocalisationID)
}
func (Repository) GetTagLocalisationByPath(path string) (tag_localisation.TagLocalisation, error) {
	return GetTagLocalisationByPath(path)
}
func (Repository) FindTagLocalisation(tagID string, queries []database.Query) (tag_localisation.TagLocalisation, error) {
	return FindTagLocalisation(tagID, queries)
}
func (Repository) GetAllTagLocalisations(tagID string) ([]tag_localisation.TagLocalisation, error) {
	return GetAllTagLocalisations(tagID)
}
func (Repository) GetAllTagLocalisationsPaginated(tagID string, pagination database.Pagination) ([]tag_localisation.TagLocalisation, database.Pagination, error) {
	return GetAllTagLocalisationsPaginated(tagID, pagination)
}
func (Repository) QueryTagLocalisations(tagID string, queries []database.Query, pagination database.Pagination) ([]tag_localisation.TagLocalisation, error) {
	return QueryTagLocalisations(tagID, queries, pagination)
}
func (Repository) QueryTagLocalisationsGroup(queries []database.Query) ([]tag_localisation.TagLocalisation, error) {
	return QueryTagLocalisationsGroup(queries)
}
func (Repository) CreateTagLocalisation(tagID string, entity tag_localisation.TagLocalisation, editorID *string) (string, error) {
	return CreateTagLocalisation(tagID, entity, editorID)
}
func (Repository) UpdateTagLocalisation(tagID string, entity tag_localisation.TagLocalisation, editorID *string) (string, error) {
	return UpdateTagLocalisation(tagID, entity, editorID)
}
func (Repository) SaveTagLocalisation(tagID string, entity tag_localisation.TagLocalisation, editorID *string) (string, error) {
	return SaveTagLocalisation(tagID, entity, editorID)
}
func (Repository) DeleteTagLocalisation(tagID string, tagLocalisationID string) error {
	return DeleteTagLocalisation(tagID, tagLocalisationID)
}