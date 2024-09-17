package repo

import (
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	"github.com/bakeable/bkry/data/entities/user_profile"
	"github.com/bakeable/bkry/data/entities/user"
	"github.com/bakeable/bkry/data/entities/portal_event"
	"github.com/bakeable/bkry/data/entities/media"
	"github.com/bakeable/bkry/data/entities/airline_settings"
	"github.com/bakeable/bkry/data/entities/delivery_entry"
	"github.com/bakeable/bkry/data/entities/printing_order_supplier"
	"github.com/bakeable/bkry/data/entities/product"
	"github.com/bakeable/bkry/data/entities/examination_settings"
	"github.com/bakeable/bkry/data/entities/price_configuration"
	"github.com/bakeable/bkry/data/entities/printing_order_item"
	"github.com/bakeable/bkry/data/entities/packing_slip_package"
	"github.com/bakeable/bkry/data/entities/cloud_function"
	"github.com/bakeable/bkry/data/entities/api_key"
	"github.com/bakeable/bkry/data/entities/quotation"
	"github.com/bakeable/bkry/data/entities/examination_property"
	"github.com/bakeable/bkry/data/entities/airline_order_grouping"
	"github.com/bakeable/bkry/data/entities/price_layer"
	"github.com/bakeable/bkry/data/entities/airline_order_group"
	"github.com/bakeable/bkry/data/entities/airline_order"
	"github.com/bakeable/bkry/data/entities/airline_pricing"
	"github.com/bakeable/bkry/data/entities/printing_order"
	"github.com/bakeable/bkry/data/entities/packing_slip"
	"github.com/bakeable/bkry/data/entities/email"
	"github.com/bakeable/bkry/data/entities/batch_export"
	"github.com/bakeable/bkry/data/entities/examination_task"
	"github.com/bakeable/bkry/data/entities/attribute"
	"github.com/bakeable/bkry/data/entities/attribute_option_set"
	"github.com/bakeable/bkry/data/entities/attribute_option"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

type IRepository interface {
	// GetUserProfile retrieves a single UserProfile entity by its ID and userProfileID.
	GetUserProfile(userProfileID string) (user_profile.UserProfile, error)
	// GetUserProfileOrNew retrieves a single UserProfile entity by its ID and userProfileID.
	GetUserProfileOrNew(userProfileID string) (user_profile.UserProfile, error)
	// GetUserProfile retrieves a single UserProfile entity by its document path.
	GetUserProfileByPath(path string) (user_profile.UserProfile, error)
	// FindUserProfile retrieves a UserProfile entity according to given queries.
	FindUserProfile(queries []datastore.Query) (user_profile.UserProfile, error)
	// GetAllUserProfiles retrieves all UserProfile entities.
	GetAllUserProfiles() ([]user_profile.UserProfile, error)
	// GetAllUserProfilesPaginated retrieves all UserProfile entities in a paginated manner.
	GetAllUserProfilesPaginated(pagination datastore.Pagination) ([]user_profile.UserProfile, datastore.Pagination, error)
	// QueryUserProfiles retrieves all UserProfile entities according to given queries.
	QueryUserProfiles(queries []datastore.Query, pagination datastore.Pagination) ([]user_profile.UserProfile, error)
	// QueryUserProfilesGroup retrieves all UserProfile entities according to given queries.
	QueryUserProfilesGroup(queries []datastore.Query) ([]user_profile.UserProfile, error)
	// CreateUserProfile creates a new UserProfile entity.
	CreateUserProfile(entity user_profile.UserProfile, editorID *string) (string, error)
	// UpdateUserProfile updates an existing UserProfile entity.
	UpdateUserProfile(entity user_profile.UserProfile, editorID *string) (string, error)
	// SaveUserProfile saves a UserProfile entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveUserProfile(entity user_profile.UserProfile, editorID *string) (string, error)
	// DeleteUserProfile deletes a UserProfile entity by its parents' path and userProfileID.
	DeleteUserProfile(userProfileID string) error
	// GetUser retrieves a single User entity by its ID and userID.
	GetUser(userID string) (user.User, error)
	// GetUserOrNew retrieves a single User entity by its ID and userID.
	GetUserOrNew(userID string) (user.User, error)
	// GetUser retrieves a single User entity by its document path.
	GetUserByPath(path string) (user.User, error)
	// FindUser retrieves a User entity according to given queries.
	FindUser(queries []datastore.Query) (user.User, error)
	// GetAllUsers retrieves all User entities.
	GetAllUsers() ([]user.User, error)
	// GetAllUsersPaginated retrieves all User entities in a paginated manner.
	GetAllUsersPaginated(pagination datastore.Pagination) ([]user.User, datastore.Pagination, error)
	// QueryUsers retrieves all User entities according to given queries.
	QueryUsers(queries []datastore.Query, pagination datastore.Pagination) ([]user.User, error)
	// QueryUsersGroup retrieves all User entities according to given queries.
	QueryUsersGroup(queries []datastore.Query) ([]user.User, error)
	// CreateUser creates a new User entity.
	CreateUser(entity user.User, editorID *string) (string, error)
	// UpdateUser updates an existing User entity.
	UpdateUser(entity user.User, editorID *string) (string, error)
	// SaveUser saves a User entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveUser(entity user.User, editorID *string) (string, error)
	// DeleteUser deletes a User entity by its parents' path and userID.
	DeleteUser(userID string) error
	// GetPortalEvent retrieves a single PortalEvent entity by its ID and portalEventID.
	GetPortalEvent(portalEventID string) (portal_event.PortalEvent, error)
	// GetPortalEventOrNew retrieves a single PortalEvent entity by its ID and portalEventID.
	GetPortalEventOrNew(portalEventID string) (portal_event.PortalEvent, error)
	// GetPortalEvent retrieves a single PortalEvent entity by its document path.
	GetPortalEventByPath(path string) (portal_event.PortalEvent, error)
	// FindPortalEvent retrieves a PortalEvent entity according to given queries.
	FindPortalEvent(queries []datastore.Query) (portal_event.PortalEvent, error)
	// GetAllPortalEvents retrieves all PortalEvent entities.
	GetAllPortalEvents() ([]portal_event.PortalEvent, error)
	// GetAllPortalEventsPaginated retrieves all PortalEvent entities in a paginated manner.
	GetAllPortalEventsPaginated(pagination datastore.Pagination) ([]portal_event.PortalEvent, datastore.Pagination, error)
	// QueryPortalEvents retrieves all PortalEvent entities according to given queries.
	QueryPortalEvents(queries []datastore.Query, pagination datastore.Pagination) ([]portal_event.PortalEvent, error)
	// QueryPortalEventsGroup retrieves all PortalEvent entities according to given queries.
	QueryPortalEventsGroup(queries []datastore.Query) ([]portal_event.PortalEvent, error)
	// CreatePortalEvent creates a new PortalEvent entity.
	CreatePortalEvent(entity portal_event.PortalEvent, editorID *string) (string, error)
	// UpdatePortalEvent updates an existing PortalEvent entity.
	UpdatePortalEvent(entity portal_event.PortalEvent, editorID *string) (string, error)
	// SavePortalEvent saves a PortalEvent entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SavePortalEvent(entity portal_event.PortalEvent, editorID *string) (string, error)
	// DeletePortalEvent deletes a PortalEvent entity by its parents' path and portalEventID.
	DeletePortalEvent(portalEventID string) error
	// GetMedia retrieves a single Media entity by its ID and mediaID.
	GetMedia(mediaID string) (media.Media, error)
	// GetMediaOrNew retrieves a single Media entity by its ID and mediaID.
	GetMediaOrNew(mediaID string) (media.Media, error)
	// GetMedia retrieves a single Media entity by its document path.
	GetMediaByPath(path string) (media.Media, error)
	// FindMedia retrieves a Media entity according to given queries.
	FindMedia(queries []datastore.Query) (media.Media, error)
	// GetAllMedia retrieves all Media entities.
	GetAllMedia() ([]media.Media, error)
	// GetAllMediaPaginated retrieves all Media entities in a paginated manner.
	GetAllMediaPaginated(pagination datastore.Pagination) ([]media.Media, datastore.Pagination, error)
	// QueryMedia retrieves all Media entities according to given queries.
	QueryMedia(queries []datastore.Query, pagination datastore.Pagination) ([]media.Media, error)
	// QueryMediaGroup retrieves all Media entities according to given queries.
	QueryMediaGroup(queries []datastore.Query) ([]media.Media, error)
	// CreateMedia creates a new Media entity.
	CreateMedia(entity media.Media, editorID *string) (string, error)
	// UpdateMedia updates an existing Media entity.
	UpdateMedia(entity media.Media, editorID *string) (string, error)
	// SaveMedia saves a Media entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveMedia(entity media.Media, editorID *string) (string, error)
	// DeleteMedia deletes a Media entity by its parents' path and mediaID.
	DeleteMedia(mediaID string) error
	// GetAirlineSettings retrieves a single AirlineSettings entity by its ID and airlineSettingsID.
	GetAirlineSettings(airlineSettingsID string) (airline_settings.AirlineSettings, error)
	// GetAirlineSettingsOrNew retrieves a single AirlineSettings entity by its ID and airlineSettingsID.
	GetAirlineSettingsOrNew(airlineSettingsID string) (airline_settings.AirlineSettings, error)
	// GetAirlineSettings retrieves a single AirlineSettings entity by its document path.
	GetAirlineSettingsByPath(path string) (airline_settings.AirlineSettings, error)
	// FindAirlineSettings retrieves a AirlineSettings entity according to given queries.
	FindAirlineSettings(queries []datastore.Query) (airline_settings.AirlineSettings, error)
	// GetAllAirlineSettings retrieves all AirlineSettings entities.
	GetAllAirlineSettings() ([]airline_settings.AirlineSettings, error)
	// GetAllAirlineSettingsPaginated retrieves all AirlineSettings entities in a paginated manner.
	GetAllAirlineSettingsPaginated(pagination datastore.Pagination) ([]airline_settings.AirlineSettings, datastore.Pagination, error)
	// QueryAirlineSettings retrieves all AirlineSettings entities according to given queries.
	QueryAirlineSettings(queries []datastore.Query, pagination datastore.Pagination) ([]airline_settings.AirlineSettings, error)
	// QueryAirlineSettingsGroup retrieves all AirlineSettings entities according to given queries.
	QueryAirlineSettingsGroup(queries []datastore.Query) ([]airline_settings.AirlineSettings, error)
	// CreateAirlineSettings creates a new AirlineSettings entity.
	CreateAirlineSettings(entity airline_settings.AirlineSettings, editorID *string) (string, error)
	// UpdateAirlineSettings updates an existing AirlineSettings entity.
	UpdateAirlineSettings(entity airline_settings.AirlineSettings, editorID *string) (string, error)
	// SaveAirlineSettings saves a AirlineSettings entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveAirlineSettings(entity airline_settings.AirlineSettings, editorID *string) (string, error)
	// DeleteAirlineSettings deletes a AirlineSettings entity by its parents' path and airlineSettingsID.
	DeleteAirlineSettings(airlineSettingsID string) error
	// GetDeliveryEntry retrieves a single DeliveryEntry entity by its ID and deliveryEntryID.
	GetDeliveryEntry(deliveryEntryID string) (delivery_entry.DeliveryEntry, error)
	// GetDeliveryEntryOrNew retrieves a single DeliveryEntry entity by its ID and deliveryEntryID.
	GetDeliveryEntryOrNew(deliveryEntryID string) (delivery_entry.DeliveryEntry, error)
	// GetDeliveryEntry retrieves a single DeliveryEntry entity by its document path.
	GetDeliveryEntryByPath(path string) (delivery_entry.DeliveryEntry, error)
	// FindDeliveryEntry retrieves a DeliveryEntry entity according to given queries.
	FindDeliveryEntry(queries []datastore.Query) (delivery_entry.DeliveryEntry, error)
	// GetAllDeliveryEntries retrieves all DeliveryEntry entities.
	GetAllDeliveryEntries() ([]delivery_entry.DeliveryEntry, error)
	// GetAllDeliveryEntriesPaginated retrieves all DeliveryEntry entities in a paginated manner.
	GetAllDeliveryEntriesPaginated(pagination datastore.Pagination) ([]delivery_entry.DeliveryEntry, datastore.Pagination, error)
	// QueryDeliveryEntries retrieves all DeliveryEntry entities according to given queries.
	QueryDeliveryEntries(queries []datastore.Query, pagination datastore.Pagination) ([]delivery_entry.DeliveryEntry, error)
	// QueryDeliveryEntriesGroup retrieves all DeliveryEntry entities according to given queries.
	QueryDeliveryEntriesGroup(queries []datastore.Query) ([]delivery_entry.DeliveryEntry, error)
	// CreateDeliveryEntry creates a new DeliveryEntry entity.
	CreateDeliveryEntry(entity delivery_entry.DeliveryEntry, editorID *string) (string, error)
	// UpdateDeliveryEntry updates an existing DeliveryEntry entity.
	UpdateDeliveryEntry(entity delivery_entry.DeliveryEntry, editorID *string) (string, error)
	// SaveDeliveryEntry saves a DeliveryEntry entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveDeliveryEntry(entity delivery_entry.DeliveryEntry, editorID *string) (string, error)
	// DeleteDeliveryEntry deletes a DeliveryEntry entity by its parents' path and deliveryEntryID.
	DeleteDeliveryEntry(deliveryEntryID string) error
	// GetPrintingOrderSupplier retrieves a single PrintingOrderSupplier entity by its ID and printingOrderSupplierID.
	GetPrintingOrderSupplier(printingOrderSupplierID string) (printing_order_supplier.PrintingOrderSupplier, error)
	// GetPrintingOrderSupplierOrNew retrieves a single PrintingOrderSupplier entity by its ID and printingOrderSupplierID.
	GetPrintingOrderSupplierOrNew(printingOrderSupplierID string) (printing_order_supplier.PrintingOrderSupplier, error)
	// GetPrintingOrderSupplier retrieves a single PrintingOrderSupplier entity by its document path.
	GetPrintingOrderSupplierByPath(path string) (printing_order_supplier.PrintingOrderSupplier, error)
	// FindPrintingOrderSupplier retrieves a PrintingOrderSupplier entity according to given queries.
	FindPrintingOrderSupplier(queries []datastore.Query) (printing_order_supplier.PrintingOrderSupplier, error)
	// GetAllPrintingOrderSuppliers retrieves all PrintingOrderSupplier entities.
	GetAllPrintingOrderSuppliers() ([]printing_order_supplier.PrintingOrderSupplier, error)
	// GetAllPrintingOrderSuppliersPaginated retrieves all PrintingOrderSupplier entities in a paginated manner.
	GetAllPrintingOrderSuppliersPaginated(pagination datastore.Pagination) ([]printing_order_supplier.PrintingOrderSupplier, datastore.Pagination, error)
	// QueryPrintingOrderSuppliers retrieves all PrintingOrderSupplier entities according to given queries.
	QueryPrintingOrderSuppliers(queries []datastore.Query, pagination datastore.Pagination) ([]printing_order_supplier.PrintingOrderSupplier, error)
	// QueryPrintingOrderSuppliersGroup retrieves all PrintingOrderSupplier entities according to given queries.
	QueryPrintingOrderSuppliersGroup(queries []datastore.Query) ([]printing_order_supplier.PrintingOrderSupplier, error)
	// CreatePrintingOrderSupplier creates a new PrintingOrderSupplier entity.
	CreatePrintingOrderSupplier(entity printing_order_supplier.PrintingOrderSupplier, editorID *string) (string, error)
	// UpdatePrintingOrderSupplier updates an existing PrintingOrderSupplier entity.
	UpdatePrintingOrderSupplier(entity printing_order_supplier.PrintingOrderSupplier, editorID *string) (string, error)
	// SavePrintingOrderSupplier saves a PrintingOrderSupplier entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SavePrintingOrderSupplier(entity printing_order_supplier.PrintingOrderSupplier, editorID *string) (string, error)
	// DeletePrintingOrderSupplier deletes a PrintingOrderSupplier entity by its parents' path and printingOrderSupplierID.
	DeletePrintingOrderSupplier(printingOrderSupplierID string) error
	// GetProduct retrieves a single Product entity by its ID and productID.
	GetProduct(productID string) (product.Product, error)
	// GetProductOrNew retrieves a single Product entity by its ID and productID.
	GetProductOrNew(productID string) (product.Product, error)
	// GetProduct retrieves a single Product entity by its document path.
	GetProductByPath(path string) (product.Product, error)
	// FindProduct retrieves a Product entity according to given queries.
	FindProduct(queries []datastore.Query) (product.Product, error)
	// GetAllProducts retrieves all Product entities.
	GetAllProducts() ([]product.Product, error)
	// GetAllProductsPaginated retrieves all Product entities in a paginated manner.
	GetAllProductsPaginated(pagination datastore.Pagination) ([]product.Product, datastore.Pagination, error)
	// QueryProducts retrieves all Product entities according to given queries.
	QueryProducts(queries []datastore.Query, pagination datastore.Pagination) ([]product.Product, error)
	// QueryProductsGroup retrieves all Product entities according to given queries.
	QueryProductsGroup(queries []datastore.Query) ([]product.Product, error)
	// CreateProduct creates a new Product entity.
	CreateProduct(entity product.Product, editorID *string) (string, error)
	// UpdateProduct updates an existing Product entity.
	UpdateProduct(entity product.Product, editorID *string) (string, error)
	// SaveProduct saves a Product entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveProduct(entity product.Product, editorID *string) (string, error)
	// DeleteProduct deletes a Product entity by its parents' path and productID.
	DeleteProduct(productID string) error
	// GetExaminationSettings retrieves a single ExaminationSettings entity by its ID and examinationSettingsID.
	GetExaminationSettings(examinationSettingsID string) (examination_settings.ExaminationSettings, error)
	// GetExaminationSettingsOrNew retrieves a single ExaminationSettings entity by its ID and examinationSettingsID.
	GetExaminationSettingsOrNew(examinationSettingsID string) (examination_settings.ExaminationSettings, error)
	// GetExaminationSettings retrieves a single ExaminationSettings entity by its document path.
	GetExaminationSettingsByPath(path string) (examination_settings.ExaminationSettings, error)
	// FindExaminationSettings retrieves a ExaminationSettings entity according to given queries.
	FindExaminationSettings(queries []datastore.Query) (examination_settings.ExaminationSettings, error)
	// GetAllExaminationSettings retrieves all ExaminationSettings entities.
	GetAllExaminationSettings() ([]examination_settings.ExaminationSettings, error)
	// GetAllExaminationSettingsPaginated retrieves all ExaminationSettings entities in a paginated manner.
	GetAllExaminationSettingsPaginated(pagination datastore.Pagination) ([]examination_settings.ExaminationSettings, datastore.Pagination, error)
	// QueryExaminationSettings retrieves all ExaminationSettings entities according to given queries.
	QueryExaminationSettings(queries []datastore.Query, pagination datastore.Pagination) ([]examination_settings.ExaminationSettings, error)
	// QueryExaminationSettingsGroup retrieves all ExaminationSettings entities according to given queries.
	QueryExaminationSettingsGroup(queries []datastore.Query) ([]examination_settings.ExaminationSettings, error)
	// CreateExaminationSettings creates a new ExaminationSettings entity.
	CreateExaminationSettings(entity examination_settings.ExaminationSettings, editorID *string) (string, error)
	// UpdateExaminationSettings updates an existing ExaminationSettings entity.
	UpdateExaminationSettings(entity examination_settings.ExaminationSettings, editorID *string) (string, error)
	// SaveExaminationSettings saves a ExaminationSettings entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveExaminationSettings(entity examination_settings.ExaminationSettings, editorID *string) (string, error)
	// DeleteExaminationSettings deletes a ExaminationSettings entity by its parents' path and examinationSettingsID.
	DeleteExaminationSettings(examinationSettingsID string) error
	// GetPriceConfiguration retrieves a single PriceConfiguration entity by its ID and priceConfigurationID.
	GetPriceConfiguration(priceConfigurationID string) (price_configuration.PriceConfiguration, error)
	// GetPriceConfigurationOrNew retrieves a single PriceConfiguration entity by its ID and priceConfigurationID.
	GetPriceConfigurationOrNew(priceConfigurationID string) (price_configuration.PriceConfiguration, error)
	// GetPriceConfiguration retrieves a single PriceConfiguration entity by its document path.
	GetPriceConfigurationByPath(path string) (price_configuration.PriceConfiguration, error)
	// FindPriceConfiguration retrieves a PriceConfiguration entity according to given queries.
	FindPriceConfiguration(queries []datastore.Query) (price_configuration.PriceConfiguration, error)
	// GetAllPriceConfigurations retrieves all PriceConfiguration entities.
	GetAllPriceConfigurations() ([]price_configuration.PriceConfiguration, error)
	// GetAllPriceConfigurationsPaginated retrieves all PriceConfiguration entities in a paginated manner.
	GetAllPriceConfigurationsPaginated(pagination datastore.Pagination) ([]price_configuration.PriceConfiguration, datastore.Pagination, error)
	// QueryPriceConfigurations retrieves all PriceConfiguration entities according to given queries.
	QueryPriceConfigurations(queries []datastore.Query, pagination datastore.Pagination) ([]price_configuration.PriceConfiguration, error)
	// QueryPriceConfigurationsGroup retrieves all PriceConfiguration entities according to given queries.
	QueryPriceConfigurationsGroup(queries []datastore.Query) ([]price_configuration.PriceConfiguration, error)
	// CreatePriceConfiguration creates a new PriceConfiguration entity.
	CreatePriceConfiguration(entity price_configuration.PriceConfiguration, editorID *string) (string, error)
	// UpdatePriceConfiguration updates an existing PriceConfiguration entity.
	UpdatePriceConfiguration(entity price_configuration.PriceConfiguration, editorID *string) (string, error)
	// SavePriceConfiguration saves a PriceConfiguration entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SavePriceConfiguration(entity price_configuration.PriceConfiguration, editorID *string) (string, error)
	// DeletePriceConfiguration deletes a PriceConfiguration entity by its parents' path and priceConfigurationID.
	DeletePriceConfiguration(priceConfigurationID string) error
	// GetPrintingOrderItem retrieves a single PrintingOrderItem entity by its ID and printingOrderItemID.
	GetPrintingOrderItem(printingOrderItemID string) (printing_order_item.PrintingOrderItem, error)
	// GetPrintingOrderItemOrNew retrieves a single PrintingOrderItem entity by its ID and printingOrderItemID.
	GetPrintingOrderItemOrNew(printingOrderItemID string) (printing_order_item.PrintingOrderItem, error)
	// GetPrintingOrderItem retrieves a single PrintingOrderItem entity by its document path.
	GetPrintingOrderItemByPath(path string) (printing_order_item.PrintingOrderItem, error)
	// FindPrintingOrderItem retrieves a PrintingOrderItem entity according to given queries.
	FindPrintingOrderItem(queries []datastore.Query) (printing_order_item.PrintingOrderItem, error)
	// GetAllPrintingOrderItems retrieves all PrintingOrderItem entities.
	GetAllPrintingOrderItems() ([]printing_order_item.PrintingOrderItem, error)
	// GetAllPrintingOrderItemsPaginated retrieves all PrintingOrderItem entities in a paginated manner.
	GetAllPrintingOrderItemsPaginated(pagination datastore.Pagination) ([]printing_order_item.PrintingOrderItem, datastore.Pagination, error)
	// QueryPrintingOrderItems retrieves all PrintingOrderItem entities according to given queries.
	QueryPrintingOrderItems(queries []datastore.Query, pagination datastore.Pagination) ([]printing_order_item.PrintingOrderItem, error)
	// QueryPrintingOrderItemsGroup retrieves all PrintingOrderItem entities according to given queries.
	QueryPrintingOrderItemsGroup(queries []datastore.Query) ([]printing_order_item.PrintingOrderItem, error)
	// CreatePrintingOrderItem creates a new PrintingOrderItem entity.
	CreatePrintingOrderItem(entity printing_order_item.PrintingOrderItem, editorID *string) (string, error)
	// UpdatePrintingOrderItem updates an existing PrintingOrderItem entity.
	UpdatePrintingOrderItem(entity printing_order_item.PrintingOrderItem, editorID *string) (string, error)
	// SavePrintingOrderItem saves a PrintingOrderItem entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SavePrintingOrderItem(entity printing_order_item.PrintingOrderItem, editorID *string) (string, error)
	// DeletePrintingOrderItem deletes a PrintingOrderItem entity by its parents' path and printingOrderItemID.
	DeletePrintingOrderItem(printingOrderItemID string) error
	// GetPackingSlipPackage retrieves a single PackingSlipPackage entity by its ID and packingSlipPackageID.
	GetPackingSlipPackage(packingSlipPackageID string) (packing_slip_package.PackingSlipPackage, error)
	// GetPackingSlipPackageOrNew retrieves a single PackingSlipPackage entity by its ID and packingSlipPackageID.
	GetPackingSlipPackageOrNew(packingSlipPackageID string) (packing_slip_package.PackingSlipPackage, error)
	// GetPackingSlipPackage retrieves a single PackingSlipPackage entity by its document path.
	GetPackingSlipPackageByPath(path string) (packing_slip_package.PackingSlipPackage, error)
	// FindPackingSlipPackage retrieves a PackingSlipPackage entity according to given queries.
	FindPackingSlipPackage(queries []datastore.Query) (packing_slip_package.PackingSlipPackage, error)
	// GetAllPackingSlipPackages retrieves all PackingSlipPackage entities.
	GetAllPackingSlipPackages() ([]packing_slip_package.PackingSlipPackage, error)
	// GetAllPackingSlipPackagesPaginated retrieves all PackingSlipPackage entities in a paginated manner.
	GetAllPackingSlipPackagesPaginated(pagination datastore.Pagination) ([]packing_slip_package.PackingSlipPackage, datastore.Pagination, error)
	// QueryPackingSlipPackages retrieves all PackingSlipPackage entities according to given queries.
	QueryPackingSlipPackages(queries []datastore.Query, pagination datastore.Pagination) ([]packing_slip_package.PackingSlipPackage, error)
	// QueryPackingSlipPackagesGroup retrieves all PackingSlipPackage entities according to given queries.
	QueryPackingSlipPackagesGroup(queries []datastore.Query) ([]packing_slip_package.PackingSlipPackage, error)
	// CreatePackingSlipPackage creates a new PackingSlipPackage entity.
	CreatePackingSlipPackage(entity packing_slip_package.PackingSlipPackage, editorID *string) (string, error)
	// UpdatePackingSlipPackage updates an existing PackingSlipPackage entity.
	UpdatePackingSlipPackage(entity packing_slip_package.PackingSlipPackage, editorID *string) (string, error)
	// SavePackingSlipPackage saves a PackingSlipPackage entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SavePackingSlipPackage(entity packing_slip_package.PackingSlipPackage, editorID *string) (string, error)
	// DeletePackingSlipPackage deletes a PackingSlipPackage entity by its parents' path and packingSlipPackageID.
	DeletePackingSlipPackage(packingSlipPackageID string) error
	// GetCloudFunction retrieves a single CloudFunction entity by its ID and cloudFunctionID.
	GetCloudFunction(cloudFunctionID string) (cloud_function.CloudFunction, error)
	// GetCloudFunctionOrNew retrieves a single CloudFunction entity by its ID and cloudFunctionID.
	GetCloudFunctionOrNew(cloudFunctionID string) (cloud_function.CloudFunction, error)
	// GetCloudFunction retrieves a single CloudFunction entity by its document path.
	GetCloudFunctionByPath(path string) (cloud_function.CloudFunction, error)
	// FindCloudFunction retrieves a CloudFunction entity according to given queries.
	FindCloudFunction(queries []datastore.Query) (cloud_function.CloudFunction, error)
	// GetAllCloudFunctions retrieves all CloudFunction entities.
	GetAllCloudFunctions() ([]cloud_function.CloudFunction, error)
	// GetAllCloudFunctionsPaginated retrieves all CloudFunction entities in a paginated manner.
	GetAllCloudFunctionsPaginated(pagination datastore.Pagination) ([]cloud_function.CloudFunction, datastore.Pagination, error)
	// QueryCloudFunctions retrieves all CloudFunction entities according to given queries.
	QueryCloudFunctions(queries []datastore.Query, pagination datastore.Pagination) ([]cloud_function.CloudFunction, error)
	// QueryCloudFunctionsGroup retrieves all CloudFunction entities according to given queries.
	QueryCloudFunctionsGroup(queries []datastore.Query) ([]cloud_function.CloudFunction, error)
	// CreateCloudFunction creates a new CloudFunction entity.
	CreateCloudFunction(entity cloud_function.CloudFunction, editorID *string) (string, error)
	// UpdateCloudFunction updates an existing CloudFunction entity.
	UpdateCloudFunction(entity cloud_function.CloudFunction, editorID *string) (string, error)
	// SaveCloudFunction saves a CloudFunction entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveCloudFunction(entity cloud_function.CloudFunction, editorID *string) (string, error)
	// DeleteCloudFunction deletes a CloudFunction entity by its parents' path and cloudFunctionID.
	DeleteCloudFunction(cloudFunctionID string) error
	// GetApiKey retrieves a single ApiKey entity by its ID and apiKeyID.
	GetApiKey(apiKeyID string) (api_key.ApiKey, error)
	// GetApiKeyOrNew retrieves a single ApiKey entity by its ID and apiKeyID.
	GetApiKeyOrNew(apiKeyID string) (api_key.ApiKey, error)
	// GetApiKey retrieves a single ApiKey entity by its document path.
	GetApiKeyByPath(path string) (api_key.ApiKey, error)
	// FindApiKey retrieves a ApiKey entity according to given queries.
	FindApiKey(queries []datastore.Query) (api_key.ApiKey, error)
	// GetAllApiKeys retrieves all ApiKey entities.
	GetAllApiKeys() ([]api_key.ApiKey, error)
	// GetAllApiKeysPaginated retrieves all ApiKey entities in a paginated manner.
	GetAllApiKeysPaginated(pagination datastore.Pagination) ([]api_key.ApiKey, datastore.Pagination, error)
	// QueryApiKeys retrieves all ApiKey entities according to given queries.
	QueryApiKeys(queries []datastore.Query, pagination datastore.Pagination) ([]api_key.ApiKey, error)
	// QueryApiKeysGroup retrieves all ApiKey entities according to given queries.
	QueryApiKeysGroup(queries []datastore.Query) ([]api_key.ApiKey, error)
	// CreateApiKey creates a new ApiKey entity.
	CreateApiKey(entity api_key.ApiKey, editorID *string) (string, error)
	// UpdateApiKey updates an existing ApiKey entity.
	UpdateApiKey(entity api_key.ApiKey, editorID *string) (string, error)
	// SaveApiKey saves a ApiKey entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveApiKey(entity api_key.ApiKey, editorID *string) (string, error)
	// DeleteApiKey deletes a ApiKey entity by its parents' path and apiKeyID.
	DeleteApiKey(apiKeyID string) error
	// GetQuotation retrieves a single Quotation entity by its ID and quotationID.
	GetQuotation(quotationID string) (quotation.Quotation, error)
	// GetQuotationOrNew retrieves a single Quotation entity by its ID and quotationID.
	GetQuotationOrNew(quotationID string) (quotation.Quotation, error)
	// GetQuotation retrieves a single Quotation entity by its document path.
	GetQuotationByPath(path string) (quotation.Quotation, error)
	// FindQuotation retrieves a Quotation entity according to given queries.
	FindQuotation(queries []datastore.Query) (quotation.Quotation, error)
	// GetAllQuotations retrieves all Quotation entities.
	GetAllQuotations() ([]quotation.Quotation, error)
	// GetAllQuotationsPaginated retrieves all Quotation entities in a paginated manner.
	GetAllQuotationsPaginated(pagination datastore.Pagination) ([]quotation.Quotation, datastore.Pagination, error)
	// QueryQuotations retrieves all Quotation entities according to given queries.
	QueryQuotations(queries []datastore.Query, pagination datastore.Pagination) ([]quotation.Quotation, error)
	// QueryQuotationsGroup retrieves all Quotation entities according to given queries.
	QueryQuotationsGroup(queries []datastore.Query) ([]quotation.Quotation, error)
	// CreateQuotation creates a new Quotation entity.
	CreateQuotation(entity quotation.Quotation, editorID *string) (string, error)
	// UpdateQuotation updates an existing Quotation entity.
	UpdateQuotation(entity quotation.Quotation, editorID *string) (string, error)
	// SaveQuotation saves a Quotation entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveQuotation(entity quotation.Quotation, editorID *string) (string, error)
	// DeleteQuotation deletes a Quotation entity by its parents' path and quotationID.
	DeleteQuotation(quotationID string) error
	// GetExaminationProperty retrieves a single ExaminationProperty entity by its ID and examinationPropertyID.
	GetExaminationProperty(examinationPropertyID string) (examination_property.ExaminationProperty, error)
	// GetExaminationPropertyOrNew retrieves a single ExaminationProperty entity by its ID and examinationPropertyID.
	GetExaminationPropertyOrNew(examinationPropertyID string) (examination_property.ExaminationProperty, error)
	// GetExaminationProperty retrieves a single ExaminationProperty entity by its document path.
	GetExaminationPropertyByPath(path string) (examination_property.ExaminationProperty, error)
	// FindExaminationProperty retrieves a ExaminationProperty entity according to given queries.
	FindExaminationProperty(queries []datastore.Query) (examination_property.ExaminationProperty, error)
	// GetAllExaminationProperties retrieves all ExaminationProperty entities.
	GetAllExaminationProperties() ([]examination_property.ExaminationProperty, error)
	// GetAllExaminationPropertiesPaginated retrieves all ExaminationProperty entities in a paginated manner.
	GetAllExaminationPropertiesPaginated(pagination datastore.Pagination) ([]examination_property.ExaminationProperty, datastore.Pagination, error)
	// QueryExaminationProperties retrieves all ExaminationProperty entities according to given queries.
	QueryExaminationProperties(queries []datastore.Query, pagination datastore.Pagination) ([]examination_property.ExaminationProperty, error)
	// QueryExaminationPropertiesGroup retrieves all ExaminationProperty entities according to given queries.
	QueryExaminationPropertiesGroup(queries []datastore.Query) ([]examination_property.ExaminationProperty, error)
	// CreateExaminationProperty creates a new ExaminationProperty entity.
	CreateExaminationProperty(entity examination_property.ExaminationProperty, editorID *string) (string, error)
	// UpdateExaminationProperty updates an existing ExaminationProperty entity.
	UpdateExaminationProperty(entity examination_property.ExaminationProperty, editorID *string) (string, error)
	// SaveExaminationProperty saves a ExaminationProperty entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveExaminationProperty(entity examination_property.ExaminationProperty, editorID *string) (string, error)
	// DeleteExaminationProperty deletes a ExaminationProperty entity by its parents' path and examinationPropertyID.
	DeleteExaminationProperty(examinationPropertyID string) error
	// GetAirlineOrderGrouping retrieves a single AirlineOrderGrouping entity by its ID and airlineOrderGroupingID.
	GetAirlineOrderGrouping(airlineOrderGroupingID string) (airline_order_grouping.AirlineOrderGrouping, error)
	// GetAirlineOrderGroupingOrNew retrieves a single AirlineOrderGrouping entity by its ID and airlineOrderGroupingID.
	GetAirlineOrderGroupingOrNew(airlineOrderGroupingID string) (airline_order_grouping.AirlineOrderGrouping, error)
	// GetAirlineOrderGrouping retrieves a single AirlineOrderGrouping entity by its document path.
	GetAirlineOrderGroupingByPath(path string) (airline_order_grouping.AirlineOrderGrouping, error)
	// FindAirlineOrderGrouping retrieves a AirlineOrderGrouping entity according to given queries.
	FindAirlineOrderGrouping(queries []datastore.Query) (airline_order_grouping.AirlineOrderGrouping, error)
	// GetAllAirlineOrderGroupings retrieves all AirlineOrderGrouping entities.
	GetAllAirlineOrderGroupings() ([]airline_order_grouping.AirlineOrderGrouping, error)
	// GetAllAirlineOrderGroupingsPaginated retrieves all AirlineOrderGrouping entities in a paginated manner.
	GetAllAirlineOrderGroupingsPaginated(pagination datastore.Pagination) ([]airline_order_grouping.AirlineOrderGrouping, datastore.Pagination, error)
	// QueryAirlineOrderGroupings retrieves all AirlineOrderGrouping entities according to given queries.
	QueryAirlineOrderGroupings(queries []datastore.Query, pagination datastore.Pagination) ([]airline_order_grouping.AirlineOrderGrouping, error)
	// QueryAirlineOrderGroupingsGroup retrieves all AirlineOrderGrouping entities according to given queries.
	QueryAirlineOrderGroupingsGroup(queries []datastore.Query) ([]airline_order_grouping.AirlineOrderGrouping, error)
	// CreateAirlineOrderGrouping creates a new AirlineOrderGrouping entity.
	CreateAirlineOrderGrouping(entity airline_order_grouping.AirlineOrderGrouping, editorID *string) (string, error)
	// UpdateAirlineOrderGrouping updates an existing AirlineOrderGrouping entity.
	UpdateAirlineOrderGrouping(entity airline_order_grouping.AirlineOrderGrouping, editorID *string) (string, error)
	// SaveAirlineOrderGrouping saves a AirlineOrderGrouping entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveAirlineOrderGrouping(entity airline_order_grouping.AirlineOrderGrouping, editorID *string) (string, error)
	// DeleteAirlineOrderGrouping deletes a AirlineOrderGrouping entity by its parents' path and airlineOrderGroupingID.
	DeleteAirlineOrderGrouping(airlineOrderGroupingID string) error
	// GetPriceLayer retrieves a single PriceLayer entity by its ID and priceLayerID.
	GetPriceLayer(priceLayerID string) (price_layer.PriceLayer, error)
	// GetPriceLayerOrNew retrieves a single PriceLayer entity by its ID and priceLayerID.
	GetPriceLayerOrNew(priceLayerID string) (price_layer.PriceLayer, error)
	// GetPriceLayer retrieves a single PriceLayer entity by its document path.
	GetPriceLayerByPath(path string) (price_layer.PriceLayer, error)
	// FindPriceLayer retrieves a PriceLayer entity according to given queries.
	FindPriceLayer(queries []datastore.Query) (price_layer.PriceLayer, error)
	// GetAllPriceLayers retrieves all PriceLayer entities.
	GetAllPriceLayers() ([]price_layer.PriceLayer, error)
	// GetAllPriceLayersPaginated retrieves all PriceLayer entities in a paginated manner.
	GetAllPriceLayersPaginated(pagination datastore.Pagination) ([]price_layer.PriceLayer, datastore.Pagination, error)
	// QueryPriceLayers retrieves all PriceLayer entities according to given queries.
	QueryPriceLayers(queries []datastore.Query, pagination datastore.Pagination) ([]price_layer.PriceLayer, error)
	// QueryPriceLayersGroup retrieves all PriceLayer entities according to given queries.
	QueryPriceLayersGroup(queries []datastore.Query) ([]price_layer.PriceLayer, error)
	// CreatePriceLayer creates a new PriceLayer entity.
	CreatePriceLayer(entity price_layer.PriceLayer, editorID *string) (string, error)
	// UpdatePriceLayer updates an existing PriceLayer entity.
	UpdatePriceLayer(entity price_layer.PriceLayer, editorID *string) (string, error)
	// SavePriceLayer saves a PriceLayer entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SavePriceLayer(entity price_layer.PriceLayer, editorID *string) (string, error)
	// DeletePriceLayer deletes a PriceLayer entity by its parents' path and priceLayerID.
	DeletePriceLayer(priceLayerID string) error
	// GetAirlineOrderGroup retrieves a single AirlineOrderGroup entity by its ID and airlineOrderGroupID.
	GetAirlineOrderGroup(airlineOrderGroupID string) (airline_order_group.AirlineOrderGroup, error)
	// GetAirlineOrderGroupOrNew retrieves a single AirlineOrderGroup entity by its ID and airlineOrderGroupID.
	GetAirlineOrderGroupOrNew(airlineOrderGroupID string) (airline_order_group.AirlineOrderGroup, error)
	// GetAirlineOrderGroup retrieves a single AirlineOrderGroup entity by its document path.
	GetAirlineOrderGroupByPath(path string) (airline_order_group.AirlineOrderGroup, error)
	// FindAirlineOrderGroup retrieves a AirlineOrderGroup entity according to given queries.
	FindAirlineOrderGroup(queries []datastore.Query) (airline_order_group.AirlineOrderGroup, error)
	// GetAllAirlineOrderGroups retrieves all AirlineOrderGroup entities.
	GetAllAirlineOrderGroups() ([]airline_order_group.AirlineOrderGroup, error)
	// GetAllAirlineOrderGroupsPaginated retrieves all AirlineOrderGroup entities in a paginated manner.
	GetAllAirlineOrderGroupsPaginated(pagination datastore.Pagination) ([]airline_order_group.AirlineOrderGroup, datastore.Pagination, error)
	// QueryAirlineOrderGroups retrieves all AirlineOrderGroup entities according to given queries.
	QueryAirlineOrderGroups(queries []datastore.Query, pagination datastore.Pagination) ([]airline_order_group.AirlineOrderGroup, error)
	// QueryAirlineOrderGroupsGroup retrieves all AirlineOrderGroup entities according to given queries.
	QueryAirlineOrderGroupsGroup(queries []datastore.Query) ([]airline_order_group.AirlineOrderGroup, error)
	// CreateAirlineOrderGroup creates a new AirlineOrderGroup entity.
	CreateAirlineOrderGroup(entity airline_order_group.AirlineOrderGroup, editorID *string) (string, error)
	// UpdateAirlineOrderGroup updates an existing AirlineOrderGroup entity.
	UpdateAirlineOrderGroup(entity airline_order_group.AirlineOrderGroup, editorID *string) (string, error)
	// SaveAirlineOrderGroup saves a AirlineOrderGroup entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveAirlineOrderGroup(entity airline_order_group.AirlineOrderGroup, editorID *string) (string, error)
	// DeleteAirlineOrderGroup deletes a AirlineOrderGroup entity by its parents' path and airlineOrderGroupID.
	DeleteAirlineOrderGroup(airlineOrderGroupID string) error
	// GetAirlineOrder retrieves a single AirlineOrder entity by its ID and airlineOrderID.
	GetAirlineOrder(airlineOrderID string) (airline_order.AirlineOrder, error)
	// GetAirlineOrderOrNew retrieves a single AirlineOrder entity by its ID and airlineOrderID.
	GetAirlineOrderOrNew(airlineOrderID string) (airline_order.AirlineOrder, error)
	// GetAirlineOrder retrieves a single AirlineOrder entity by its document path.
	GetAirlineOrderByPath(path string) (airline_order.AirlineOrder, error)
	// FindAirlineOrder retrieves a AirlineOrder entity according to given queries.
	FindAirlineOrder(queries []datastore.Query) (airline_order.AirlineOrder, error)
	// GetAllAirlineOrders retrieves all AirlineOrder entities.
	GetAllAirlineOrders() ([]airline_order.AirlineOrder, error)
	// GetAllAirlineOrdersPaginated retrieves all AirlineOrder entities in a paginated manner.
	GetAllAirlineOrdersPaginated(pagination datastore.Pagination) ([]airline_order.AirlineOrder, datastore.Pagination, error)
	// QueryAirlineOrders retrieves all AirlineOrder entities according to given queries.
	QueryAirlineOrders(queries []datastore.Query, pagination datastore.Pagination) ([]airline_order.AirlineOrder, error)
	// QueryAirlineOrdersGroup retrieves all AirlineOrder entities according to given queries.
	QueryAirlineOrdersGroup(queries []datastore.Query) ([]airline_order.AirlineOrder, error)
	// CreateAirlineOrder creates a new AirlineOrder entity.
	CreateAirlineOrder(entity airline_order.AirlineOrder, editorID *string) (string, error)
	// UpdateAirlineOrder updates an existing AirlineOrder entity.
	UpdateAirlineOrder(entity airline_order.AirlineOrder, editorID *string) (string, error)
	// SaveAirlineOrder saves a AirlineOrder entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveAirlineOrder(entity airline_order.AirlineOrder, editorID *string) (string, error)
	// DeleteAirlineOrder deletes a AirlineOrder entity by its parents' path and airlineOrderID.
	DeleteAirlineOrder(airlineOrderID string) error
	// GetAirlinePricing retrieves a single AirlinePricing entity by its ID and airlinePricingID.
	GetAirlinePricing(airlinePricingID string) (airline_pricing.AirlinePricing, error)
	// GetAirlinePricingOrNew retrieves a single AirlinePricing entity by its ID and airlinePricingID.
	GetAirlinePricingOrNew(airlinePricingID string) (airline_pricing.AirlinePricing, error)
	// GetAirlinePricing retrieves a single AirlinePricing entity by its document path.
	GetAirlinePricingByPath(path string) (airline_pricing.AirlinePricing, error)
	// FindAirlinePricing retrieves a AirlinePricing entity according to given queries.
	FindAirlinePricing(queries []datastore.Query) (airline_pricing.AirlinePricing, error)
	// GetAllAirlinePricings retrieves all AirlinePricing entities.
	GetAllAirlinePricings() ([]airline_pricing.AirlinePricing, error)
	// GetAllAirlinePricingsPaginated retrieves all AirlinePricing entities in a paginated manner.
	GetAllAirlinePricingsPaginated(pagination datastore.Pagination) ([]airline_pricing.AirlinePricing, datastore.Pagination, error)
	// QueryAirlinePricings retrieves all AirlinePricing entities according to given queries.
	QueryAirlinePricings(queries []datastore.Query, pagination datastore.Pagination) ([]airline_pricing.AirlinePricing, error)
	// QueryAirlinePricingsGroup retrieves all AirlinePricing entities according to given queries.
	QueryAirlinePricingsGroup(queries []datastore.Query) ([]airline_pricing.AirlinePricing, error)
	// CreateAirlinePricing creates a new AirlinePricing entity.
	CreateAirlinePricing(entity airline_pricing.AirlinePricing, editorID *string) (string, error)
	// UpdateAirlinePricing updates an existing AirlinePricing entity.
	UpdateAirlinePricing(entity airline_pricing.AirlinePricing, editorID *string) (string, error)
	// SaveAirlinePricing saves a AirlinePricing entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveAirlinePricing(entity airline_pricing.AirlinePricing, editorID *string) (string, error)
	// DeleteAirlinePricing deletes a AirlinePricing entity by its parents' path and airlinePricingID.
	DeleteAirlinePricing(airlinePricingID string) error
	// GetPrintingOrder retrieves a single PrintingOrder entity by its ID and printingOrderID.
	GetPrintingOrder(printingOrderID string) (printing_order.PrintingOrder, error)
	// GetPrintingOrderOrNew retrieves a single PrintingOrder entity by its ID and printingOrderID.
	GetPrintingOrderOrNew(printingOrderID string) (printing_order.PrintingOrder, error)
	// GetPrintingOrder retrieves a single PrintingOrder entity by its document path.
	GetPrintingOrderByPath(path string) (printing_order.PrintingOrder, error)
	// FindPrintingOrder retrieves a PrintingOrder entity according to given queries.
	FindPrintingOrder(queries []datastore.Query) (printing_order.PrintingOrder, error)
	// GetAllPrintingOrders retrieves all PrintingOrder entities.
	GetAllPrintingOrders() ([]printing_order.PrintingOrder, error)
	// GetAllPrintingOrdersPaginated retrieves all PrintingOrder entities in a paginated manner.
	GetAllPrintingOrdersPaginated(pagination datastore.Pagination) ([]printing_order.PrintingOrder, datastore.Pagination, error)
	// QueryPrintingOrders retrieves all PrintingOrder entities according to given queries.
	QueryPrintingOrders(queries []datastore.Query, pagination datastore.Pagination) ([]printing_order.PrintingOrder, error)
	// QueryPrintingOrdersGroup retrieves all PrintingOrder entities according to given queries.
	QueryPrintingOrdersGroup(queries []datastore.Query) ([]printing_order.PrintingOrder, error)
	// CreatePrintingOrder creates a new PrintingOrder entity.
	CreatePrintingOrder(entity printing_order.PrintingOrder, editorID *string) (string, error)
	// UpdatePrintingOrder updates an existing PrintingOrder entity.
	UpdatePrintingOrder(entity printing_order.PrintingOrder, editorID *string) (string, error)
	// SavePrintingOrder saves a PrintingOrder entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SavePrintingOrder(entity printing_order.PrintingOrder, editorID *string) (string, error)
	// DeletePrintingOrder deletes a PrintingOrder entity by its parents' path and printingOrderID.
	DeletePrintingOrder(printingOrderID string) error
	// GetPackingSlip retrieves a single PackingSlip entity by its ID and packingSlipID.
	GetPackingSlip(packingSlipID string) (packing_slip.PackingSlip, error)
	// GetPackingSlipOrNew retrieves a single PackingSlip entity by its ID and packingSlipID.
	GetPackingSlipOrNew(packingSlipID string) (packing_slip.PackingSlip, error)
	// GetPackingSlip retrieves a single PackingSlip entity by its document path.
	GetPackingSlipByPath(path string) (packing_slip.PackingSlip, error)
	// FindPackingSlip retrieves a PackingSlip entity according to given queries.
	FindPackingSlip(queries []datastore.Query) (packing_slip.PackingSlip, error)
	// GetAllPackingSlips retrieves all PackingSlip entities.
	GetAllPackingSlips() ([]packing_slip.PackingSlip, error)
	// GetAllPackingSlipsPaginated retrieves all PackingSlip entities in a paginated manner.
	GetAllPackingSlipsPaginated(pagination datastore.Pagination) ([]packing_slip.PackingSlip, datastore.Pagination, error)
	// QueryPackingSlips retrieves all PackingSlip entities according to given queries.
	QueryPackingSlips(queries []datastore.Query, pagination datastore.Pagination) ([]packing_slip.PackingSlip, error)
	// QueryPackingSlipsGroup retrieves all PackingSlip entities according to given queries.
	QueryPackingSlipsGroup(queries []datastore.Query) ([]packing_slip.PackingSlip, error)
	// CreatePackingSlip creates a new PackingSlip entity.
	CreatePackingSlip(entity packing_slip.PackingSlip, editorID *string) (string, error)
	// UpdatePackingSlip updates an existing PackingSlip entity.
	UpdatePackingSlip(entity packing_slip.PackingSlip, editorID *string) (string, error)
	// SavePackingSlip saves a PackingSlip entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SavePackingSlip(entity packing_slip.PackingSlip, editorID *string) (string, error)
	// DeletePackingSlip deletes a PackingSlip entity by its parents' path and packingSlipID.
	DeletePackingSlip(packingSlipID string) error
	// GetEmail retrieves a single Email entity by its ID and emailID.
	GetEmail(emailID string) (email.Email, error)
	// GetEmailOrNew retrieves a single Email entity by its ID and emailID.
	GetEmailOrNew(emailID string) (email.Email, error)
	// GetEmail retrieves a single Email entity by its document path.
	GetEmailByPath(path string) (email.Email, error)
	// FindEmail retrieves a Email entity according to given queries.
	FindEmail(queries []datastore.Query) (email.Email, error)
	// GetAllEmails retrieves all Email entities.
	GetAllEmails() ([]email.Email, error)
	// GetAllEmailsPaginated retrieves all Email entities in a paginated manner.
	GetAllEmailsPaginated(pagination datastore.Pagination) ([]email.Email, datastore.Pagination, error)
	// QueryEmails retrieves all Email entities according to given queries.
	QueryEmails(queries []datastore.Query, pagination datastore.Pagination) ([]email.Email, error)
	// QueryEmailsGroup retrieves all Email entities according to given queries.
	QueryEmailsGroup(queries []datastore.Query) ([]email.Email, error)
	// CreateEmail creates a new Email entity.
	CreateEmail(entity email.Email, editorID *string) (string, error)
	// UpdateEmail updates an existing Email entity.
	UpdateEmail(entity email.Email, editorID *string) (string, error)
	// SaveEmail saves a Email entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveEmail(entity email.Email, editorID *string) (string, error)
	// DeleteEmail deletes a Email entity by its parents' path and emailID.
	DeleteEmail(emailID string) error
	// GetBatchExport retrieves a single BatchExport entity by its ID and batchExportID.
	GetBatchExport(batchExportID string) (batch_export.BatchExport, error)
	// GetBatchExportOrNew retrieves a single BatchExport entity by its ID and batchExportID.
	GetBatchExportOrNew(batchExportID string) (batch_export.BatchExport, error)
	// GetBatchExport retrieves a single BatchExport entity by its document path.
	GetBatchExportByPath(path string) (batch_export.BatchExport, error)
	// FindBatchExport retrieves a BatchExport entity according to given queries.
	FindBatchExport(queries []datastore.Query) (batch_export.BatchExport, error)
	// GetAllBatchExports retrieves all BatchExport entities.
	GetAllBatchExports() ([]batch_export.BatchExport, error)
	// GetAllBatchExportsPaginated retrieves all BatchExport entities in a paginated manner.
	GetAllBatchExportsPaginated(pagination datastore.Pagination) ([]batch_export.BatchExport, datastore.Pagination, error)
	// QueryBatchExports retrieves all BatchExport entities according to given queries.
	QueryBatchExports(queries []datastore.Query, pagination datastore.Pagination) ([]batch_export.BatchExport, error)
	// QueryBatchExportsGroup retrieves all BatchExport entities according to given queries.
	QueryBatchExportsGroup(queries []datastore.Query) ([]batch_export.BatchExport, error)
	// CreateBatchExport creates a new BatchExport entity.
	CreateBatchExport(entity batch_export.BatchExport, editorID *string) (string, error)
	// UpdateBatchExport updates an existing BatchExport entity.
	UpdateBatchExport(entity batch_export.BatchExport, editorID *string) (string, error)
	// SaveBatchExport saves a BatchExport entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveBatchExport(entity batch_export.BatchExport, editorID *string) (string, error)
	// DeleteBatchExport deletes a BatchExport entity by its parents' path and batchExportID.
	DeleteBatchExport(batchExportID string) error
	// GetExaminationTask retrieves a single ExaminationTask entity by its ID and examinationTaskID.
	GetExaminationTask(examinationTaskID string) (examination_task.ExaminationTask, error)
	// GetExaminationTaskOrNew retrieves a single ExaminationTask entity by its ID and examinationTaskID.
	GetExaminationTaskOrNew(examinationTaskID string) (examination_task.ExaminationTask, error)
	// GetExaminationTask retrieves a single ExaminationTask entity by its document path.
	GetExaminationTaskByPath(path string) (examination_task.ExaminationTask, error)
	// FindExaminationTask retrieves a ExaminationTask entity according to given queries.
	FindExaminationTask(queries []datastore.Query) (examination_task.ExaminationTask, error)
	// GetAllExaminationTasks retrieves all ExaminationTask entities.
	GetAllExaminationTasks() ([]examination_task.ExaminationTask, error)
	// GetAllExaminationTasksPaginated retrieves all ExaminationTask entities in a paginated manner.
	GetAllExaminationTasksPaginated(pagination datastore.Pagination) ([]examination_task.ExaminationTask, datastore.Pagination, error)
	// QueryExaminationTasks retrieves all ExaminationTask entities according to given queries.
	QueryExaminationTasks(queries []datastore.Query, pagination datastore.Pagination) ([]examination_task.ExaminationTask, error)
	// QueryExaminationTasksGroup retrieves all ExaminationTask entities according to given queries.
	QueryExaminationTasksGroup(queries []datastore.Query) ([]examination_task.ExaminationTask, error)
	// CreateExaminationTask creates a new ExaminationTask entity.
	CreateExaminationTask(entity examination_task.ExaminationTask, editorID *string) (string, error)
	// UpdateExaminationTask updates an existing ExaminationTask entity.
	UpdateExaminationTask(entity examination_task.ExaminationTask, editorID *string) (string, error)
	// SaveExaminationTask saves a ExaminationTask entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveExaminationTask(entity examination_task.ExaminationTask, editorID *string) (string, error)
	// DeleteExaminationTask deletes a ExaminationTask entity by its parents' path and examinationTaskID.
	DeleteExaminationTask(examinationTaskID string) error
	// GetAttribute retrieves a single Attribute entity by its ID and attributeID.
	GetAttribute(attributeID string) (attribute.Attribute, error)
	// GetAttributeOrNew retrieves a single Attribute entity by its ID and attributeID.
	GetAttributeOrNew(attributeID string) (attribute.Attribute, error)
	// GetAttribute retrieves a single Attribute entity by its document path.
	GetAttributeByPath(path string) (attribute.Attribute, error)
	// FindAttribute retrieves a Attribute entity according to given queries.
	FindAttribute(queries []datastore.Query) (attribute.Attribute, error)
	// GetAllAttributes retrieves all Attribute entities.
	GetAllAttributes() ([]attribute.Attribute, error)
	// GetAllAttributesPaginated retrieves all Attribute entities in a paginated manner.
	GetAllAttributesPaginated(pagination datastore.Pagination) ([]attribute.Attribute, datastore.Pagination, error)
	// QueryAttributes retrieves all Attribute entities according to given queries.
	QueryAttributes(queries []datastore.Query, pagination datastore.Pagination) ([]attribute.Attribute, error)
	// QueryAttributesGroup retrieves all Attribute entities according to given queries.
	QueryAttributesGroup(queries []datastore.Query) ([]attribute.Attribute, error)
	// CreateAttribute creates a new Attribute entity.
	CreateAttribute(entity attribute.Attribute, editorID *string) (string, error)
	// UpdateAttribute updates an existing Attribute entity.
	UpdateAttribute(entity attribute.Attribute, editorID *string) (string, error)
	// SaveAttribute saves a Attribute entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveAttribute(entity attribute.Attribute, editorID *string) (string, error)
	// DeleteAttribute deletes a Attribute entity by its parents' path and attributeID.
	DeleteAttribute(attributeID string) error
	// GetAttributeOptionSet retrieves a single AttributeOptionSet entity by its ID and attributeOptionSetID.
	GetAttributeOptionSet(attributeOptionSetID string) (attribute_option_set.AttributeOptionSet, error)
	// GetAttributeOptionSetOrNew retrieves a single AttributeOptionSet entity by its ID and attributeOptionSetID.
	GetAttributeOptionSetOrNew(attributeOptionSetID string) (attribute_option_set.AttributeOptionSet, error)
	// GetAttributeOptionSet retrieves a single AttributeOptionSet entity by its document path.
	GetAttributeOptionSetByPath(path string) (attribute_option_set.AttributeOptionSet, error)
	// FindAttributeOptionSet retrieves a AttributeOptionSet entity according to given queries.
	FindAttributeOptionSet(queries []datastore.Query) (attribute_option_set.AttributeOptionSet, error)
	// GetAllAttributeOptionSets retrieves all AttributeOptionSet entities.
	GetAllAttributeOptionSets() ([]attribute_option_set.AttributeOptionSet, error)
	// GetAllAttributeOptionSetsPaginated retrieves all AttributeOptionSet entities in a paginated manner.
	GetAllAttributeOptionSetsPaginated(pagination datastore.Pagination) ([]attribute_option_set.AttributeOptionSet, datastore.Pagination, error)
	// QueryAttributeOptionSets retrieves all AttributeOptionSet entities according to given queries.
	QueryAttributeOptionSets(queries []datastore.Query, pagination datastore.Pagination) ([]attribute_option_set.AttributeOptionSet, error)
	// QueryAttributeOptionSetsGroup retrieves all AttributeOptionSet entities according to given queries.
	QueryAttributeOptionSetsGroup(queries []datastore.Query) ([]attribute_option_set.AttributeOptionSet, error)
	// CreateAttributeOptionSet creates a new AttributeOptionSet entity.
	CreateAttributeOptionSet(entity attribute_option_set.AttributeOptionSet, editorID *string) (string, error)
	// UpdateAttributeOptionSet updates an existing AttributeOptionSet entity.
	UpdateAttributeOptionSet(entity attribute_option_set.AttributeOptionSet, editorID *string) (string, error)
	// SaveAttributeOptionSet saves a AttributeOptionSet entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveAttributeOptionSet(entity attribute_option_set.AttributeOptionSet, editorID *string) (string, error)
	// DeleteAttributeOptionSet deletes a AttributeOptionSet entity by its parents' path and attributeOptionSetID.
	DeleteAttributeOptionSet(attributeOptionSetID string) error
	// GetAttributeOption retrieves a single AttributeOption entity by its ID and attributeOptionID.
	GetAttributeOption(attributeOptionID string) (attribute_option.AttributeOption, error)
	// GetAttributeOptionOrNew retrieves a single AttributeOption entity by its ID and attributeOptionID.
	GetAttributeOptionOrNew(attributeOptionID string) (attribute_option.AttributeOption, error)
	// GetAttributeOption retrieves a single AttributeOption entity by its document path.
	GetAttributeOptionByPath(path string) (attribute_option.AttributeOption, error)
	// FindAttributeOption retrieves a AttributeOption entity according to given queries.
	FindAttributeOption(queries []datastore.Query) (attribute_option.AttributeOption, error)
	// GetAllAttributeOptions retrieves all AttributeOption entities.
	GetAllAttributeOptions() ([]attribute_option.AttributeOption, error)
	// GetAllAttributeOptionsPaginated retrieves all AttributeOption entities in a paginated manner.
	GetAllAttributeOptionsPaginated(pagination datastore.Pagination) ([]attribute_option.AttributeOption, datastore.Pagination, error)
	// QueryAttributeOptions retrieves all AttributeOption entities according to given queries.
	QueryAttributeOptions(queries []datastore.Query, pagination datastore.Pagination) ([]attribute_option.AttributeOption, error)
	// QueryAttributeOptionsGroup retrieves all AttributeOption entities according to given queries.
	QueryAttributeOptionsGroup(queries []datastore.Query) ([]attribute_option.AttributeOption, error)
	// CreateAttributeOption creates a new AttributeOption entity.
	CreateAttributeOption(entity attribute_option.AttributeOption, editorID *string) (string, error)
	// UpdateAttributeOption updates an existing AttributeOption entity.
	UpdateAttributeOption(entity attribute_option.AttributeOption, editorID *string) (string, error)
	// SaveAttributeOption saves a AttributeOption entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
	SaveAttributeOption(entity attribute_option.AttributeOption, editorID *string) (string, error)
	// DeleteAttributeOption deletes a AttributeOption entity by its parents' path and attributeOptionID.
	DeleteAttributeOption(attributeOptionID string) error
}


type Repository struct {}
func NewRepository() IRepository {
	return Repository{}
}
func (Repository) GetUserProfile(userProfileID string) (user_profile.UserProfile, error) {
	return GetUserProfile(userProfileID)
}
func (Repository) GetUserProfileOrNew(userProfileID string) (user_profile.UserProfile, error) {
	return GetUserProfileOrNew(userProfileID)
}
func (Repository) GetUserProfileByPath(path string) (user_profile.UserProfile, error) {
	return GetUserProfileByPath(path)
}
func (Repository) FindUserProfile(queries []datastore.Query) (user_profile.UserProfile, error) {
	return FindUserProfile(queries)
}
func (Repository) GetAllUserProfiles() ([]user_profile.UserProfile, error) {
	return GetAllUserProfiles()
}
func (Repository) GetAllUserProfilesPaginated(pagination datastore.Pagination) ([]user_profile.UserProfile, datastore.Pagination, error) {
	return GetAllUserProfilesPaginated(pagination)
}
func (Repository) QueryUserProfiles(queries []datastore.Query, pagination datastore.Pagination) ([]user_profile.UserProfile, error) {
	return QueryUserProfiles(queries, pagination)
}
func (Repository) QueryUserProfilesGroup(queries []datastore.Query) ([]user_profile.UserProfile, error) {
	return QueryUserProfilesGroup(queries)
}
func (Repository) CreateUserProfile(entity user_profile.UserProfile, editorID *string) (string, error) {
	return CreateUserProfile(entity, editorID)
}
func (Repository) UpdateUserProfile(entity user_profile.UserProfile, editorID *string) (string, error) {
	return UpdateUserProfile(entity, editorID)
}
func (Repository) SaveUserProfile(entity user_profile.UserProfile, editorID *string) (string, error) {
	return SaveUserProfile(entity, editorID)
}
func (Repository) DeleteUserProfile(userProfileID string) error {
	return DeleteUserProfile(userProfileID)
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
func (Repository) FindUser(queries []datastore.Query) (user.User, error) {
	return FindUser(queries)
}
func (Repository) GetAllUsers() ([]user.User, error) {
	return GetAllUsers()
}
func (Repository) GetAllUsersPaginated(pagination datastore.Pagination) ([]user.User, datastore.Pagination, error) {
	return GetAllUsersPaginated(pagination)
}
func (Repository) QueryUsers(queries []datastore.Query, pagination datastore.Pagination) ([]user.User, error) {
	return QueryUsers(queries, pagination)
}
func (Repository) QueryUsersGroup(queries []datastore.Query) ([]user.User, error) {
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
func (Repository) GetPortalEvent(portalEventID string) (portal_event.PortalEvent, error) {
	return GetPortalEvent(portalEventID)
}
func (Repository) GetPortalEventOrNew(portalEventID string) (portal_event.PortalEvent, error) {
	return GetPortalEventOrNew(portalEventID)
}
func (Repository) GetPortalEventByPath(path string) (portal_event.PortalEvent, error) {
	return GetPortalEventByPath(path)
}
func (Repository) FindPortalEvent(queries []datastore.Query) (portal_event.PortalEvent, error) {
	return FindPortalEvent(queries)
}
func (Repository) GetAllPortalEvents() ([]portal_event.PortalEvent, error) {
	return GetAllPortalEvents()
}
func (Repository) GetAllPortalEventsPaginated(pagination datastore.Pagination) ([]portal_event.PortalEvent, datastore.Pagination, error) {
	return GetAllPortalEventsPaginated(pagination)
}
func (Repository) QueryPortalEvents(queries []datastore.Query, pagination datastore.Pagination) ([]portal_event.PortalEvent, error) {
	return QueryPortalEvents(queries, pagination)
}
func (Repository) QueryPortalEventsGroup(queries []datastore.Query) ([]portal_event.PortalEvent, error) {
	return QueryPortalEventsGroup(queries)
}
func (Repository) CreatePortalEvent(entity portal_event.PortalEvent, editorID *string) (string, error) {
	return CreatePortalEvent(entity, editorID)
}
func (Repository) UpdatePortalEvent(entity portal_event.PortalEvent, editorID *string) (string, error) {
	return UpdatePortalEvent(entity, editorID)
}
func (Repository) SavePortalEvent(entity portal_event.PortalEvent, editorID *string) (string, error) {
	return SavePortalEvent(entity, editorID)
}
func (Repository) DeletePortalEvent(portalEventID string) error {
	return DeletePortalEvent(portalEventID)
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
func (Repository) FindMedia(queries []datastore.Query) (media.Media, error) {
	return FindMedia(queries)
}
func (Repository) GetAllMedia() ([]media.Media, error) {
	return GetAllMedia()
}
func (Repository) GetAllMediaPaginated(pagination datastore.Pagination) ([]media.Media, datastore.Pagination, error) {
	return GetAllMediaPaginated(pagination)
}
func (Repository) QueryMedia(queries []datastore.Query, pagination datastore.Pagination) ([]media.Media, error) {
	return QueryMedia(queries, pagination)
}
func (Repository) QueryMediaGroup(queries []datastore.Query) ([]media.Media, error) {
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
func (Repository) GetAirlineSettings(airlineSettingsID string) (airline_settings.AirlineSettings, error) {
	return GetAirlineSettings(airlineSettingsID)
}
func (Repository) GetAirlineSettingsOrNew(airlineSettingsID string) (airline_settings.AirlineSettings, error) {
	return GetAirlineSettingsOrNew(airlineSettingsID)
}
func (Repository) GetAirlineSettingsByPath(path string) (airline_settings.AirlineSettings, error) {
	return GetAirlineSettingsByPath(path)
}
func (Repository) FindAirlineSettings(queries []datastore.Query) (airline_settings.AirlineSettings, error) {
	return FindAirlineSettings(queries)
}
func (Repository) GetAllAirlineSettings() ([]airline_settings.AirlineSettings, error) {
	return GetAllAirlineSettings()
}
func (Repository) GetAllAirlineSettingsPaginated(pagination datastore.Pagination) ([]airline_settings.AirlineSettings, datastore.Pagination, error) {
	return GetAllAirlineSettingsPaginated(pagination)
}
func (Repository) QueryAirlineSettings(queries []datastore.Query, pagination datastore.Pagination) ([]airline_settings.AirlineSettings, error) {
	return QueryAirlineSettings(queries, pagination)
}
func (Repository) QueryAirlineSettingsGroup(queries []datastore.Query) ([]airline_settings.AirlineSettings, error) {
	return QueryAirlineSettingsGroup(queries)
}
func (Repository) CreateAirlineSettings(entity airline_settings.AirlineSettings, editorID *string) (string, error) {
	return CreateAirlineSettings(entity, editorID)
}
func (Repository) UpdateAirlineSettings(entity airline_settings.AirlineSettings, editorID *string) (string, error) {
	return UpdateAirlineSettings(entity, editorID)
}
func (Repository) SaveAirlineSettings(entity airline_settings.AirlineSettings, editorID *string) (string, error) {
	return SaveAirlineSettings(entity, editorID)
}
func (Repository) DeleteAirlineSettings(airlineSettingsID string) error {
	return DeleteAirlineSettings(airlineSettingsID)
}
func (Repository) GetDeliveryEntry(deliveryEntryID string) (delivery_entry.DeliveryEntry, error) {
	return GetDeliveryEntry(deliveryEntryID)
}
func (Repository) GetDeliveryEntryOrNew(deliveryEntryID string) (delivery_entry.DeliveryEntry, error) {
	return GetDeliveryEntryOrNew(deliveryEntryID)
}
func (Repository) GetDeliveryEntryByPath(path string) (delivery_entry.DeliveryEntry, error) {
	return GetDeliveryEntryByPath(path)
}
func (Repository) FindDeliveryEntry(queries []datastore.Query) (delivery_entry.DeliveryEntry, error) {
	return FindDeliveryEntry(queries)
}
func (Repository) GetAllDeliveryEntries() ([]delivery_entry.DeliveryEntry, error) {
	return GetAllDeliveryEntries()
}
func (Repository) GetAllDeliveryEntriesPaginated(pagination datastore.Pagination) ([]delivery_entry.DeliveryEntry, datastore.Pagination, error) {
	return GetAllDeliveryEntriesPaginated(pagination)
}
func (Repository) QueryDeliveryEntries(queries []datastore.Query, pagination datastore.Pagination) ([]delivery_entry.DeliveryEntry, error) {
	return QueryDeliveryEntries(queries, pagination)
}
func (Repository) QueryDeliveryEntriesGroup(queries []datastore.Query) ([]delivery_entry.DeliveryEntry, error) {
	return QueryDeliveryEntriesGroup(queries)
}
func (Repository) CreateDeliveryEntry(entity delivery_entry.DeliveryEntry, editorID *string) (string, error) {
	return CreateDeliveryEntry(entity, editorID)
}
func (Repository) UpdateDeliveryEntry(entity delivery_entry.DeliveryEntry, editorID *string) (string, error) {
	return UpdateDeliveryEntry(entity, editorID)
}
func (Repository) SaveDeliveryEntry(entity delivery_entry.DeliveryEntry, editorID *string) (string, error) {
	return SaveDeliveryEntry(entity, editorID)
}
func (Repository) DeleteDeliveryEntry(deliveryEntryID string) error {
	return DeleteDeliveryEntry(deliveryEntryID)
}
func (Repository) GetPrintingOrderSupplier(printingOrderSupplierID string) (printing_order_supplier.PrintingOrderSupplier, error) {
	return GetPrintingOrderSupplier(printingOrderSupplierID)
}
func (Repository) GetPrintingOrderSupplierOrNew(printingOrderSupplierID string) (printing_order_supplier.PrintingOrderSupplier, error) {
	return GetPrintingOrderSupplierOrNew(printingOrderSupplierID)
}
func (Repository) GetPrintingOrderSupplierByPath(path string) (printing_order_supplier.PrintingOrderSupplier, error) {
	return GetPrintingOrderSupplierByPath(path)
}
func (Repository) FindPrintingOrderSupplier(queries []datastore.Query) (printing_order_supplier.PrintingOrderSupplier, error) {
	return FindPrintingOrderSupplier(queries)
}
func (Repository) GetAllPrintingOrderSuppliers() ([]printing_order_supplier.PrintingOrderSupplier, error) {
	return GetAllPrintingOrderSuppliers()
}
func (Repository) GetAllPrintingOrderSuppliersPaginated(pagination datastore.Pagination) ([]printing_order_supplier.PrintingOrderSupplier, datastore.Pagination, error) {
	return GetAllPrintingOrderSuppliersPaginated(pagination)
}
func (Repository) QueryPrintingOrderSuppliers(queries []datastore.Query, pagination datastore.Pagination) ([]printing_order_supplier.PrintingOrderSupplier, error) {
	return QueryPrintingOrderSuppliers(queries, pagination)
}
func (Repository) QueryPrintingOrderSuppliersGroup(queries []datastore.Query) ([]printing_order_supplier.PrintingOrderSupplier, error) {
	return QueryPrintingOrderSuppliersGroup(queries)
}
func (Repository) CreatePrintingOrderSupplier(entity printing_order_supplier.PrintingOrderSupplier, editorID *string) (string, error) {
	return CreatePrintingOrderSupplier(entity, editorID)
}
func (Repository) UpdatePrintingOrderSupplier(entity printing_order_supplier.PrintingOrderSupplier, editorID *string) (string, error) {
	return UpdatePrintingOrderSupplier(entity, editorID)
}
func (Repository) SavePrintingOrderSupplier(entity printing_order_supplier.PrintingOrderSupplier, editorID *string) (string, error) {
	return SavePrintingOrderSupplier(entity, editorID)
}
func (Repository) DeletePrintingOrderSupplier(printingOrderSupplierID string) error {
	return DeletePrintingOrderSupplier(printingOrderSupplierID)
}
func (Repository) GetProduct(productID string) (product.Product, error) {
	return GetProduct(productID)
}
func (Repository) GetProductOrNew(productID string) (product.Product, error) {
	return GetProductOrNew(productID)
}
func (Repository) GetProductByPath(path string) (product.Product, error) {
	return GetProductByPath(path)
}
func (Repository) FindProduct(queries []datastore.Query) (product.Product, error) {
	return FindProduct(queries)
}
func (Repository) GetAllProducts() ([]product.Product, error) {
	return GetAllProducts()
}
func (Repository) GetAllProductsPaginated(pagination datastore.Pagination) ([]product.Product, datastore.Pagination, error) {
	return GetAllProductsPaginated(pagination)
}
func (Repository) QueryProducts(queries []datastore.Query, pagination datastore.Pagination) ([]product.Product, error) {
	return QueryProducts(queries, pagination)
}
func (Repository) QueryProductsGroup(queries []datastore.Query) ([]product.Product, error) {
	return QueryProductsGroup(queries)
}
func (Repository) CreateProduct(entity product.Product, editorID *string) (string, error) {
	return CreateProduct(entity, editorID)
}
func (Repository) UpdateProduct(entity product.Product, editorID *string) (string, error) {
	return UpdateProduct(entity, editorID)
}
func (Repository) SaveProduct(entity product.Product, editorID *string) (string, error) {
	return SaveProduct(entity, editorID)
}
func (Repository) DeleteProduct(productID string) error {
	return DeleteProduct(productID)
}
func (Repository) GetExaminationSettings(examinationSettingsID string) (examination_settings.ExaminationSettings, error) {
	return GetExaminationSettings(examinationSettingsID)
}
func (Repository) GetExaminationSettingsOrNew(examinationSettingsID string) (examination_settings.ExaminationSettings, error) {
	return GetExaminationSettingsOrNew(examinationSettingsID)
}
func (Repository) GetExaminationSettingsByPath(path string) (examination_settings.ExaminationSettings, error) {
	return GetExaminationSettingsByPath(path)
}
func (Repository) FindExaminationSettings(queries []datastore.Query) (examination_settings.ExaminationSettings, error) {
	return FindExaminationSettings(queries)
}
func (Repository) GetAllExaminationSettings() ([]examination_settings.ExaminationSettings, error) {
	return GetAllExaminationSettings()
}
func (Repository) GetAllExaminationSettingsPaginated(pagination datastore.Pagination) ([]examination_settings.ExaminationSettings, datastore.Pagination, error) {
	return GetAllExaminationSettingsPaginated(pagination)
}
func (Repository) QueryExaminationSettings(queries []datastore.Query, pagination datastore.Pagination) ([]examination_settings.ExaminationSettings, error) {
	return QueryExaminationSettings(queries, pagination)
}
func (Repository) QueryExaminationSettingsGroup(queries []datastore.Query) ([]examination_settings.ExaminationSettings, error) {
	return QueryExaminationSettingsGroup(queries)
}
func (Repository) CreateExaminationSettings(entity examination_settings.ExaminationSettings, editorID *string) (string, error) {
	return CreateExaminationSettings(entity, editorID)
}
func (Repository) UpdateExaminationSettings(entity examination_settings.ExaminationSettings, editorID *string) (string, error) {
	return UpdateExaminationSettings(entity, editorID)
}
func (Repository) SaveExaminationSettings(entity examination_settings.ExaminationSettings, editorID *string) (string, error) {
	return SaveExaminationSettings(entity, editorID)
}
func (Repository) DeleteExaminationSettings(examinationSettingsID string) error {
	return DeleteExaminationSettings(examinationSettingsID)
}
func (Repository) GetPriceConfiguration(priceConfigurationID string) (price_configuration.PriceConfiguration, error) {
	return GetPriceConfiguration(priceConfigurationID)
}
func (Repository) GetPriceConfigurationOrNew(priceConfigurationID string) (price_configuration.PriceConfiguration, error) {
	return GetPriceConfigurationOrNew(priceConfigurationID)
}
func (Repository) GetPriceConfigurationByPath(path string) (price_configuration.PriceConfiguration, error) {
	return GetPriceConfigurationByPath(path)
}
func (Repository) FindPriceConfiguration(queries []datastore.Query) (price_configuration.PriceConfiguration, error) {
	return FindPriceConfiguration(queries)
}
func (Repository) GetAllPriceConfigurations() ([]price_configuration.PriceConfiguration, error) {
	return GetAllPriceConfigurations()
}
func (Repository) GetAllPriceConfigurationsPaginated(pagination datastore.Pagination) ([]price_configuration.PriceConfiguration, datastore.Pagination, error) {
	return GetAllPriceConfigurationsPaginated(pagination)
}
func (Repository) QueryPriceConfigurations(queries []datastore.Query, pagination datastore.Pagination) ([]price_configuration.PriceConfiguration, error) {
	return QueryPriceConfigurations(queries, pagination)
}
func (Repository) QueryPriceConfigurationsGroup(queries []datastore.Query) ([]price_configuration.PriceConfiguration, error) {
	return QueryPriceConfigurationsGroup(queries)
}
func (Repository) CreatePriceConfiguration(entity price_configuration.PriceConfiguration, editorID *string) (string, error) {
	return CreatePriceConfiguration(entity, editorID)
}
func (Repository) UpdatePriceConfiguration(entity price_configuration.PriceConfiguration, editorID *string) (string, error) {
	return UpdatePriceConfiguration(entity, editorID)
}
func (Repository) SavePriceConfiguration(entity price_configuration.PriceConfiguration, editorID *string) (string, error) {
	return SavePriceConfiguration(entity, editorID)
}
func (Repository) DeletePriceConfiguration(priceConfigurationID string) error {
	return DeletePriceConfiguration(priceConfigurationID)
}
func (Repository) GetPrintingOrderItem(printingOrderItemID string) (printing_order_item.PrintingOrderItem, error) {
	return GetPrintingOrderItem(printingOrderItemID)
}
func (Repository) GetPrintingOrderItemOrNew(printingOrderItemID string) (printing_order_item.PrintingOrderItem, error) {
	return GetPrintingOrderItemOrNew(printingOrderItemID)
}
func (Repository) GetPrintingOrderItemByPath(path string) (printing_order_item.PrintingOrderItem, error) {
	return GetPrintingOrderItemByPath(path)
}
func (Repository) FindPrintingOrderItem(queries []datastore.Query) (printing_order_item.PrintingOrderItem, error) {
	return FindPrintingOrderItem(queries)
}
func (Repository) GetAllPrintingOrderItems() ([]printing_order_item.PrintingOrderItem, error) {
	return GetAllPrintingOrderItems()
}
func (Repository) GetAllPrintingOrderItemsPaginated(pagination datastore.Pagination) ([]printing_order_item.PrintingOrderItem, datastore.Pagination, error) {
	return GetAllPrintingOrderItemsPaginated(pagination)
}
func (Repository) QueryPrintingOrderItems(queries []datastore.Query, pagination datastore.Pagination) ([]printing_order_item.PrintingOrderItem, error) {
	return QueryPrintingOrderItems(queries, pagination)
}
func (Repository) QueryPrintingOrderItemsGroup(queries []datastore.Query) ([]printing_order_item.PrintingOrderItem, error) {
	return QueryPrintingOrderItemsGroup(queries)
}
func (Repository) CreatePrintingOrderItem(entity printing_order_item.PrintingOrderItem, editorID *string) (string, error) {
	return CreatePrintingOrderItem(entity, editorID)
}
func (Repository) UpdatePrintingOrderItem(entity printing_order_item.PrintingOrderItem, editorID *string) (string, error) {
	return UpdatePrintingOrderItem(entity, editorID)
}
func (Repository) SavePrintingOrderItem(entity printing_order_item.PrintingOrderItem, editorID *string) (string, error) {
	return SavePrintingOrderItem(entity, editorID)
}
func (Repository) DeletePrintingOrderItem(printingOrderItemID string) error {
	return DeletePrintingOrderItem(printingOrderItemID)
}
func (Repository) GetPackingSlipPackage(packingSlipPackageID string) (packing_slip_package.PackingSlipPackage, error) {
	return GetPackingSlipPackage(packingSlipPackageID)
}
func (Repository) GetPackingSlipPackageOrNew(packingSlipPackageID string) (packing_slip_package.PackingSlipPackage, error) {
	return GetPackingSlipPackageOrNew(packingSlipPackageID)
}
func (Repository) GetPackingSlipPackageByPath(path string) (packing_slip_package.PackingSlipPackage, error) {
	return GetPackingSlipPackageByPath(path)
}
func (Repository) FindPackingSlipPackage(queries []datastore.Query) (packing_slip_package.PackingSlipPackage, error) {
	return FindPackingSlipPackage(queries)
}
func (Repository) GetAllPackingSlipPackages() ([]packing_slip_package.PackingSlipPackage, error) {
	return GetAllPackingSlipPackages()
}
func (Repository) GetAllPackingSlipPackagesPaginated(pagination datastore.Pagination) ([]packing_slip_package.PackingSlipPackage, datastore.Pagination, error) {
	return GetAllPackingSlipPackagesPaginated(pagination)
}
func (Repository) QueryPackingSlipPackages(queries []datastore.Query, pagination datastore.Pagination) ([]packing_slip_package.PackingSlipPackage, error) {
	return QueryPackingSlipPackages(queries, pagination)
}
func (Repository) QueryPackingSlipPackagesGroup(queries []datastore.Query) ([]packing_slip_package.PackingSlipPackage, error) {
	return QueryPackingSlipPackagesGroup(queries)
}
func (Repository) CreatePackingSlipPackage(entity packing_slip_package.PackingSlipPackage, editorID *string) (string, error) {
	return CreatePackingSlipPackage(entity, editorID)
}
func (Repository) UpdatePackingSlipPackage(entity packing_slip_package.PackingSlipPackage, editorID *string) (string, error) {
	return UpdatePackingSlipPackage(entity, editorID)
}
func (Repository) SavePackingSlipPackage(entity packing_slip_package.PackingSlipPackage, editorID *string) (string, error) {
	return SavePackingSlipPackage(entity, editorID)
}
func (Repository) DeletePackingSlipPackage(packingSlipPackageID string) error {
	return DeletePackingSlipPackage(packingSlipPackageID)
}
func (Repository) GetCloudFunction(cloudFunctionID string) (cloud_function.CloudFunction, error) {
	return GetCloudFunction(cloudFunctionID)
}
func (Repository) GetCloudFunctionOrNew(cloudFunctionID string) (cloud_function.CloudFunction, error) {
	return GetCloudFunctionOrNew(cloudFunctionID)
}
func (Repository) GetCloudFunctionByPath(path string) (cloud_function.CloudFunction, error) {
	return GetCloudFunctionByPath(path)
}
func (Repository) FindCloudFunction(queries []datastore.Query) (cloud_function.CloudFunction, error) {
	return FindCloudFunction(queries)
}
func (Repository) GetAllCloudFunctions() ([]cloud_function.CloudFunction, error) {
	return GetAllCloudFunctions()
}
func (Repository) GetAllCloudFunctionsPaginated(pagination datastore.Pagination) ([]cloud_function.CloudFunction, datastore.Pagination, error) {
	return GetAllCloudFunctionsPaginated(pagination)
}
func (Repository) QueryCloudFunctions(queries []datastore.Query, pagination datastore.Pagination) ([]cloud_function.CloudFunction, error) {
	return QueryCloudFunctions(queries, pagination)
}
func (Repository) QueryCloudFunctionsGroup(queries []datastore.Query) ([]cloud_function.CloudFunction, error) {
	return QueryCloudFunctionsGroup(queries)
}
func (Repository) CreateCloudFunction(entity cloud_function.CloudFunction, editorID *string) (string, error) {
	return CreateCloudFunction(entity, editorID)
}
func (Repository) UpdateCloudFunction(entity cloud_function.CloudFunction, editorID *string) (string, error) {
	return UpdateCloudFunction(entity, editorID)
}
func (Repository) SaveCloudFunction(entity cloud_function.CloudFunction, editorID *string) (string, error) {
	return SaveCloudFunction(entity, editorID)
}
func (Repository) DeleteCloudFunction(cloudFunctionID string) error {
	return DeleteCloudFunction(cloudFunctionID)
}
func (Repository) GetApiKey(apiKeyID string) (api_key.ApiKey, error) {
	return GetApiKey(apiKeyID)
}
func (Repository) GetApiKeyOrNew(apiKeyID string) (api_key.ApiKey, error) {
	return GetApiKeyOrNew(apiKeyID)
}
func (Repository) GetApiKeyByPath(path string) (api_key.ApiKey, error) {
	return GetApiKeyByPath(path)
}
func (Repository) FindApiKey(queries []datastore.Query) (api_key.ApiKey, error) {
	return FindApiKey(queries)
}
func (Repository) GetAllApiKeys() ([]api_key.ApiKey, error) {
	return GetAllApiKeys()
}
func (Repository) GetAllApiKeysPaginated(pagination datastore.Pagination) ([]api_key.ApiKey, datastore.Pagination, error) {
	return GetAllApiKeysPaginated(pagination)
}
func (Repository) QueryApiKeys(queries []datastore.Query, pagination datastore.Pagination) ([]api_key.ApiKey, error) {
	return QueryApiKeys(queries, pagination)
}
func (Repository) QueryApiKeysGroup(queries []datastore.Query) ([]api_key.ApiKey, error) {
	return QueryApiKeysGroup(queries)
}
func (Repository) CreateApiKey(entity api_key.ApiKey, editorID *string) (string, error) {
	return CreateApiKey(entity, editorID)
}
func (Repository) UpdateApiKey(entity api_key.ApiKey, editorID *string) (string, error) {
	return UpdateApiKey(entity, editorID)
}
func (Repository) SaveApiKey(entity api_key.ApiKey, editorID *string) (string, error) {
	return SaveApiKey(entity, editorID)
}
func (Repository) DeleteApiKey(apiKeyID string) error {
	return DeleteApiKey(apiKeyID)
}
func (Repository) GetQuotation(quotationID string) (quotation.Quotation, error) {
	return GetQuotation(quotationID)
}
func (Repository) GetQuotationOrNew(quotationID string) (quotation.Quotation, error) {
	return GetQuotationOrNew(quotationID)
}
func (Repository) GetQuotationByPath(path string) (quotation.Quotation, error) {
	return GetQuotationByPath(path)
}
func (Repository) FindQuotation(queries []datastore.Query) (quotation.Quotation, error) {
	return FindQuotation(queries)
}
func (Repository) GetAllQuotations() ([]quotation.Quotation, error) {
	return GetAllQuotations()
}
func (Repository) GetAllQuotationsPaginated(pagination datastore.Pagination) ([]quotation.Quotation, datastore.Pagination, error) {
	return GetAllQuotationsPaginated(pagination)
}
func (Repository) QueryQuotations(queries []datastore.Query, pagination datastore.Pagination) ([]quotation.Quotation, error) {
	return QueryQuotations(queries, pagination)
}
func (Repository) QueryQuotationsGroup(queries []datastore.Query) ([]quotation.Quotation, error) {
	return QueryQuotationsGroup(queries)
}
func (Repository) CreateQuotation(entity quotation.Quotation, editorID *string) (string, error) {
	return CreateQuotation(entity, editorID)
}
func (Repository) UpdateQuotation(entity quotation.Quotation, editorID *string) (string, error) {
	return UpdateQuotation(entity, editorID)
}
func (Repository) SaveQuotation(entity quotation.Quotation, editorID *string) (string, error) {
	return SaveQuotation(entity, editorID)
}
func (Repository) DeleteQuotation(quotationID string) error {
	return DeleteQuotation(quotationID)
}
func (Repository) GetExaminationProperty(examinationPropertyID string) (examination_property.ExaminationProperty, error) {
	return GetExaminationProperty(examinationPropertyID)
}
func (Repository) GetExaminationPropertyOrNew(examinationPropertyID string) (examination_property.ExaminationProperty, error) {
	return GetExaminationPropertyOrNew(examinationPropertyID)
}
func (Repository) GetExaminationPropertyByPath(path string) (examination_property.ExaminationProperty, error) {
	return GetExaminationPropertyByPath(path)
}
func (Repository) FindExaminationProperty(queries []datastore.Query) (examination_property.ExaminationProperty, error) {
	return FindExaminationProperty(queries)
}
func (Repository) GetAllExaminationProperties() ([]examination_property.ExaminationProperty, error) {
	return GetAllExaminationProperties()
}
func (Repository) GetAllExaminationPropertiesPaginated(pagination datastore.Pagination) ([]examination_property.ExaminationProperty, datastore.Pagination, error) {
	return GetAllExaminationPropertiesPaginated(pagination)
}
func (Repository) QueryExaminationProperties(queries []datastore.Query, pagination datastore.Pagination) ([]examination_property.ExaminationProperty, error) {
	return QueryExaminationProperties(queries, pagination)
}
func (Repository) QueryExaminationPropertiesGroup(queries []datastore.Query) ([]examination_property.ExaminationProperty, error) {
	return QueryExaminationPropertiesGroup(queries)
}
func (Repository) CreateExaminationProperty(entity examination_property.ExaminationProperty, editorID *string) (string, error) {
	return CreateExaminationProperty(entity, editorID)
}
func (Repository) UpdateExaminationProperty(entity examination_property.ExaminationProperty, editorID *string) (string, error) {
	return UpdateExaminationProperty(entity, editorID)
}
func (Repository) SaveExaminationProperty(entity examination_property.ExaminationProperty, editorID *string) (string, error) {
	return SaveExaminationProperty(entity, editorID)
}
func (Repository) DeleteExaminationProperty(examinationPropertyID string) error {
	return DeleteExaminationProperty(examinationPropertyID)
}
func (Repository) GetAirlineOrderGrouping(airlineOrderGroupingID string) (airline_order_grouping.AirlineOrderGrouping, error) {
	return GetAirlineOrderGrouping(airlineOrderGroupingID)
}
func (Repository) GetAirlineOrderGroupingOrNew(airlineOrderGroupingID string) (airline_order_grouping.AirlineOrderGrouping, error) {
	return GetAirlineOrderGroupingOrNew(airlineOrderGroupingID)
}
func (Repository) GetAirlineOrderGroupingByPath(path string) (airline_order_grouping.AirlineOrderGrouping, error) {
	return GetAirlineOrderGroupingByPath(path)
}
func (Repository) FindAirlineOrderGrouping(queries []datastore.Query) (airline_order_grouping.AirlineOrderGrouping, error) {
	return FindAirlineOrderGrouping(queries)
}
func (Repository) GetAllAirlineOrderGroupings() ([]airline_order_grouping.AirlineOrderGrouping, error) {
	return GetAllAirlineOrderGroupings()
}
func (Repository) GetAllAirlineOrderGroupingsPaginated(pagination datastore.Pagination) ([]airline_order_grouping.AirlineOrderGrouping, datastore.Pagination, error) {
	return GetAllAirlineOrderGroupingsPaginated(pagination)
}
func (Repository) QueryAirlineOrderGroupings(queries []datastore.Query, pagination datastore.Pagination) ([]airline_order_grouping.AirlineOrderGrouping, error) {
	return QueryAirlineOrderGroupings(queries, pagination)
}
func (Repository) QueryAirlineOrderGroupingsGroup(queries []datastore.Query) ([]airline_order_grouping.AirlineOrderGrouping, error) {
	return QueryAirlineOrderGroupingsGroup(queries)
}
func (Repository) CreateAirlineOrderGrouping(entity airline_order_grouping.AirlineOrderGrouping, editorID *string) (string, error) {
	return CreateAirlineOrderGrouping(entity, editorID)
}
func (Repository) UpdateAirlineOrderGrouping(entity airline_order_grouping.AirlineOrderGrouping, editorID *string) (string, error) {
	return UpdateAirlineOrderGrouping(entity, editorID)
}
func (Repository) SaveAirlineOrderGrouping(entity airline_order_grouping.AirlineOrderGrouping, editorID *string) (string, error) {
	return SaveAirlineOrderGrouping(entity, editorID)
}
func (Repository) DeleteAirlineOrderGrouping(airlineOrderGroupingID string) error {
	return DeleteAirlineOrderGrouping(airlineOrderGroupingID)
}
func (Repository) GetPriceLayer(priceLayerID string) (price_layer.PriceLayer, error) {
	return GetPriceLayer(priceLayerID)
}
func (Repository) GetPriceLayerOrNew(priceLayerID string) (price_layer.PriceLayer, error) {
	return GetPriceLayerOrNew(priceLayerID)
}
func (Repository) GetPriceLayerByPath(path string) (price_layer.PriceLayer, error) {
	return GetPriceLayerByPath(path)
}
func (Repository) FindPriceLayer(queries []datastore.Query) (price_layer.PriceLayer, error) {
	return FindPriceLayer(queries)
}
func (Repository) GetAllPriceLayers() ([]price_layer.PriceLayer, error) {
	return GetAllPriceLayers()
}
func (Repository) GetAllPriceLayersPaginated(pagination datastore.Pagination) ([]price_layer.PriceLayer, datastore.Pagination, error) {
	return GetAllPriceLayersPaginated(pagination)
}
func (Repository) QueryPriceLayers(queries []datastore.Query, pagination datastore.Pagination) ([]price_layer.PriceLayer, error) {
	return QueryPriceLayers(queries, pagination)
}
func (Repository) QueryPriceLayersGroup(queries []datastore.Query) ([]price_layer.PriceLayer, error) {
	return QueryPriceLayersGroup(queries)
}
func (Repository) CreatePriceLayer(entity price_layer.PriceLayer, editorID *string) (string, error) {
	return CreatePriceLayer(entity, editorID)
}
func (Repository) UpdatePriceLayer(entity price_layer.PriceLayer, editorID *string) (string, error) {
	return UpdatePriceLayer(entity, editorID)
}
func (Repository) SavePriceLayer(entity price_layer.PriceLayer, editorID *string) (string, error) {
	return SavePriceLayer(entity, editorID)
}
func (Repository) DeletePriceLayer(priceLayerID string) error {
	return DeletePriceLayer(priceLayerID)
}
func (Repository) GetAirlineOrderGroup(airlineOrderGroupID string) (airline_order_group.AirlineOrderGroup, error) {
	return GetAirlineOrderGroup(airlineOrderGroupID)
}
func (Repository) GetAirlineOrderGroupOrNew(airlineOrderGroupID string) (airline_order_group.AirlineOrderGroup, error) {
	return GetAirlineOrderGroupOrNew(airlineOrderGroupID)
}
func (Repository) GetAirlineOrderGroupByPath(path string) (airline_order_group.AirlineOrderGroup, error) {
	return GetAirlineOrderGroupByPath(path)
}
func (Repository) FindAirlineOrderGroup(queries []datastore.Query) (airline_order_group.AirlineOrderGroup, error) {
	return FindAirlineOrderGroup(queries)
}
func (Repository) GetAllAirlineOrderGroups() ([]airline_order_group.AirlineOrderGroup, error) {
	return GetAllAirlineOrderGroups()
}
func (Repository) GetAllAirlineOrderGroupsPaginated(pagination datastore.Pagination) ([]airline_order_group.AirlineOrderGroup, datastore.Pagination, error) {
	return GetAllAirlineOrderGroupsPaginated(pagination)
}
func (Repository) QueryAirlineOrderGroups(queries []datastore.Query, pagination datastore.Pagination) ([]airline_order_group.AirlineOrderGroup, error) {
	return QueryAirlineOrderGroups(queries, pagination)
}
func (Repository) QueryAirlineOrderGroupsGroup(queries []datastore.Query) ([]airline_order_group.AirlineOrderGroup, error) {
	return QueryAirlineOrderGroupsGroup(queries)
}
func (Repository) CreateAirlineOrderGroup(entity airline_order_group.AirlineOrderGroup, editorID *string) (string, error) {
	return CreateAirlineOrderGroup(entity, editorID)
}
func (Repository) UpdateAirlineOrderGroup(entity airline_order_group.AirlineOrderGroup, editorID *string) (string, error) {
	return UpdateAirlineOrderGroup(entity, editorID)
}
func (Repository) SaveAirlineOrderGroup(entity airline_order_group.AirlineOrderGroup, editorID *string) (string, error) {
	return SaveAirlineOrderGroup(entity, editorID)
}
func (Repository) DeleteAirlineOrderGroup(airlineOrderGroupID string) error {
	return DeleteAirlineOrderGroup(airlineOrderGroupID)
}
func (Repository) GetAirlineOrder(airlineOrderID string) (airline_order.AirlineOrder, error) {
	return GetAirlineOrder(airlineOrderID)
}
func (Repository) GetAirlineOrderOrNew(airlineOrderID string) (airline_order.AirlineOrder, error) {
	return GetAirlineOrderOrNew(airlineOrderID)
}
func (Repository) GetAirlineOrderByPath(path string) (airline_order.AirlineOrder, error) {
	return GetAirlineOrderByPath(path)
}
func (Repository) FindAirlineOrder(queries []datastore.Query) (airline_order.AirlineOrder, error) {
	return FindAirlineOrder(queries)
}
func (Repository) GetAllAirlineOrders() ([]airline_order.AirlineOrder, error) {
	return GetAllAirlineOrders()
}
func (Repository) GetAllAirlineOrdersPaginated(pagination datastore.Pagination) ([]airline_order.AirlineOrder, datastore.Pagination, error) {
	return GetAllAirlineOrdersPaginated(pagination)
}
func (Repository) QueryAirlineOrders(queries []datastore.Query, pagination datastore.Pagination) ([]airline_order.AirlineOrder, error) {
	return QueryAirlineOrders(queries, pagination)
}
func (Repository) QueryAirlineOrdersGroup(queries []datastore.Query) ([]airline_order.AirlineOrder, error) {
	return QueryAirlineOrdersGroup(queries)
}
func (Repository) CreateAirlineOrder(entity airline_order.AirlineOrder, editorID *string) (string, error) {
	return CreateAirlineOrder(entity, editorID)
}
func (Repository) UpdateAirlineOrder(entity airline_order.AirlineOrder, editorID *string) (string, error) {
	return UpdateAirlineOrder(entity, editorID)
}
func (Repository) SaveAirlineOrder(entity airline_order.AirlineOrder, editorID *string) (string, error) {
	return SaveAirlineOrder(entity, editorID)
}
func (Repository) DeleteAirlineOrder(airlineOrderID string) error {
	return DeleteAirlineOrder(airlineOrderID)
}
func (Repository) GetAirlinePricing(airlinePricingID string) (airline_pricing.AirlinePricing, error) {
	return GetAirlinePricing(airlinePricingID)
}
func (Repository) GetAirlinePricingOrNew(airlinePricingID string) (airline_pricing.AirlinePricing, error) {
	return GetAirlinePricingOrNew(airlinePricingID)
}
func (Repository) GetAirlinePricingByPath(path string) (airline_pricing.AirlinePricing, error) {
	return GetAirlinePricingByPath(path)
}
func (Repository) FindAirlinePricing(queries []datastore.Query) (airline_pricing.AirlinePricing, error) {
	return FindAirlinePricing(queries)
}
func (Repository) GetAllAirlinePricings() ([]airline_pricing.AirlinePricing, error) {
	return GetAllAirlinePricings()
}
func (Repository) GetAllAirlinePricingsPaginated(pagination datastore.Pagination) ([]airline_pricing.AirlinePricing, datastore.Pagination, error) {
	return GetAllAirlinePricingsPaginated(pagination)
}
func (Repository) QueryAirlinePricings(queries []datastore.Query, pagination datastore.Pagination) ([]airline_pricing.AirlinePricing, error) {
	return QueryAirlinePricings(queries, pagination)
}
func (Repository) QueryAirlinePricingsGroup(queries []datastore.Query) ([]airline_pricing.AirlinePricing, error) {
	return QueryAirlinePricingsGroup(queries)
}
func (Repository) CreateAirlinePricing(entity airline_pricing.AirlinePricing, editorID *string) (string, error) {
	return CreateAirlinePricing(entity, editorID)
}
func (Repository) UpdateAirlinePricing(entity airline_pricing.AirlinePricing, editorID *string) (string, error) {
	return UpdateAirlinePricing(entity, editorID)
}
func (Repository) SaveAirlinePricing(entity airline_pricing.AirlinePricing, editorID *string) (string, error) {
	return SaveAirlinePricing(entity, editorID)
}
func (Repository) DeleteAirlinePricing(airlinePricingID string) error {
	return DeleteAirlinePricing(airlinePricingID)
}
func (Repository) GetPrintingOrder(printingOrderID string) (printing_order.PrintingOrder, error) {
	return GetPrintingOrder(printingOrderID)
}
func (Repository) GetPrintingOrderOrNew(printingOrderID string) (printing_order.PrintingOrder, error) {
	return GetPrintingOrderOrNew(printingOrderID)
}
func (Repository) GetPrintingOrderByPath(path string) (printing_order.PrintingOrder, error) {
	return GetPrintingOrderByPath(path)
}
func (Repository) FindPrintingOrder(queries []datastore.Query) (printing_order.PrintingOrder, error) {
	return FindPrintingOrder(queries)
}
func (Repository) GetAllPrintingOrders() ([]printing_order.PrintingOrder, error) {
	return GetAllPrintingOrders()
}
func (Repository) GetAllPrintingOrdersPaginated(pagination datastore.Pagination) ([]printing_order.PrintingOrder, datastore.Pagination, error) {
	return GetAllPrintingOrdersPaginated(pagination)
}
func (Repository) QueryPrintingOrders(queries []datastore.Query, pagination datastore.Pagination) ([]printing_order.PrintingOrder, error) {
	return QueryPrintingOrders(queries, pagination)
}
func (Repository) QueryPrintingOrdersGroup(queries []datastore.Query) ([]printing_order.PrintingOrder, error) {
	return QueryPrintingOrdersGroup(queries)
}
func (Repository) CreatePrintingOrder(entity printing_order.PrintingOrder, editorID *string) (string, error) {
	return CreatePrintingOrder(entity, editorID)
}
func (Repository) UpdatePrintingOrder(entity printing_order.PrintingOrder, editorID *string) (string, error) {
	return UpdatePrintingOrder(entity, editorID)
}
func (Repository) SavePrintingOrder(entity printing_order.PrintingOrder, editorID *string) (string, error) {
	return SavePrintingOrder(entity, editorID)
}
func (Repository) DeletePrintingOrder(printingOrderID string) error {
	return DeletePrintingOrder(printingOrderID)
}
func (Repository) GetPackingSlip(packingSlipID string) (packing_slip.PackingSlip, error) {
	return GetPackingSlip(packingSlipID)
}
func (Repository) GetPackingSlipOrNew(packingSlipID string) (packing_slip.PackingSlip, error) {
	return GetPackingSlipOrNew(packingSlipID)
}
func (Repository) GetPackingSlipByPath(path string) (packing_slip.PackingSlip, error) {
	return GetPackingSlipByPath(path)
}
func (Repository) FindPackingSlip(queries []datastore.Query) (packing_slip.PackingSlip, error) {
	return FindPackingSlip(queries)
}
func (Repository) GetAllPackingSlips() ([]packing_slip.PackingSlip, error) {
	return GetAllPackingSlips()
}
func (Repository) GetAllPackingSlipsPaginated(pagination datastore.Pagination) ([]packing_slip.PackingSlip, datastore.Pagination, error) {
	return GetAllPackingSlipsPaginated(pagination)
}
func (Repository) QueryPackingSlips(queries []datastore.Query, pagination datastore.Pagination) ([]packing_slip.PackingSlip, error) {
	return QueryPackingSlips(queries, pagination)
}
func (Repository) QueryPackingSlipsGroup(queries []datastore.Query) ([]packing_slip.PackingSlip, error) {
	return QueryPackingSlipsGroup(queries)
}
func (Repository) CreatePackingSlip(entity packing_slip.PackingSlip, editorID *string) (string, error) {
	return CreatePackingSlip(entity, editorID)
}
func (Repository) UpdatePackingSlip(entity packing_slip.PackingSlip, editorID *string) (string, error) {
	return UpdatePackingSlip(entity, editorID)
}
func (Repository) SavePackingSlip(entity packing_slip.PackingSlip, editorID *string) (string, error) {
	return SavePackingSlip(entity, editorID)
}
func (Repository) DeletePackingSlip(packingSlipID string) error {
	return DeletePackingSlip(packingSlipID)
}
func (Repository) GetEmail(emailID string) (email.Email, error) {
	return GetEmail(emailID)
}
func (Repository) GetEmailOrNew(emailID string) (email.Email, error) {
	return GetEmailOrNew(emailID)
}
func (Repository) GetEmailByPath(path string) (email.Email, error) {
	return GetEmailByPath(path)
}
func (Repository) FindEmail(queries []datastore.Query) (email.Email, error) {
	return FindEmail(queries)
}
func (Repository) GetAllEmails() ([]email.Email, error) {
	return GetAllEmails()
}
func (Repository) GetAllEmailsPaginated(pagination datastore.Pagination) ([]email.Email, datastore.Pagination, error) {
	return GetAllEmailsPaginated(pagination)
}
func (Repository) QueryEmails(queries []datastore.Query, pagination datastore.Pagination) ([]email.Email, error) {
	return QueryEmails(queries, pagination)
}
func (Repository) QueryEmailsGroup(queries []datastore.Query) ([]email.Email, error) {
	return QueryEmailsGroup(queries)
}
func (Repository) CreateEmail(entity email.Email, editorID *string) (string, error) {
	return CreateEmail(entity, editorID)
}
func (Repository) UpdateEmail(entity email.Email, editorID *string) (string, error) {
	return UpdateEmail(entity, editorID)
}
func (Repository) SaveEmail(entity email.Email, editorID *string) (string, error) {
	return SaveEmail(entity, editorID)
}
func (Repository) DeleteEmail(emailID string) error {
	return DeleteEmail(emailID)
}
func (Repository) GetBatchExport(batchExportID string) (batch_export.BatchExport, error) {
	return GetBatchExport(batchExportID)
}
func (Repository) GetBatchExportOrNew(batchExportID string) (batch_export.BatchExport, error) {
	return GetBatchExportOrNew(batchExportID)
}
func (Repository) GetBatchExportByPath(path string) (batch_export.BatchExport, error) {
	return GetBatchExportByPath(path)
}
func (Repository) FindBatchExport(queries []datastore.Query) (batch_export.BatchExport, error) {
	return FindBatchExport(queries)
}
func (Repository) GetAllBatchExports() ([]batch_export.BatchExport, error) {
	return GetAllBatchExports()
}
func (Repository) GetAllBatchExportsPaginated(pagination datastore.Pagination) ([]batch_export.BatchExport, datastore.Pagination, error) {
	return GetAllBatchExportsPaginated(pagination)
}
func (Repository) QueryBatchExports(queries []datastore.Query, pagination datastore.Pagination) ([]batch_export.BatchExport, error) {
	return QueryBatchExports(queries, pagination)
}
func (Repository) QueryBatchExportsGroup(queries []datastore.Query) ([]batch_export.BatchExport, error) {
	return QueryBatchExportsGroup(queries)
}
func (Repository) CreateBatchExport(entity batch_export.BatchExport, editorID *string) (string, error) {
	return CreateBatchExport(entity, editorID)
}
func (Repository) UpdateBatchExport(entity batch_export.BatchExport, editorID *string) (string, error) {
	return UpdateBatchExport(entity, editorID)
}
func (Repository) SaveBatchExport(entity batch_export.BatchExport, editorID *string) (string, error) {
	return SaveBatchExport(entity, editorID)
}
func (Repository) DeleteBatchExport(batchExportID string) error {
	return DeleteBatchExport(batchExportID)
}
func (Repository) GetExaminationTask(examinationTaskID string) (examination_task.ExaminationTask, error) {
	return GetExaminationTask(examinationTaskID)
}
func (Repository) GetExaminationTaskOrNew(examinationTaskID string) (examination_task.ExaminationTask, error) {
	return GetExaminationTaskOrNew(examinationTaskID)
}
func (Repository) GetExaminationTaskByPath(path string) (examination_task.ExaminationTask, error) {
	return GetExaminationTaskByPath(path)
}
func (Repository) FindExaminationTask(queries []datastore.Query) (examination_task.ExaminationTask, error) {
	return FindExaminationTask(queries)
}
func (Repository) GetAllExaminationTasks() ([]examination_task.ExaminationTask, error) {
	return GetAllExaminationTasks()
}
func (Repository) GetAllExaminationTasksPaginated(pagination datastore.Pagination) ([]examination_task.ExaminationTask, datastore.Pagination, error) {
	return GetAllExaminationTasksPaginated(pagination)
}
func (Repository) QueryExaminationTasks(queries []datastore.Query, pagination datastore.Pagination) ([]examination_task.ExaminationTask, error) {
	return QueryExaminationTasks(queries, pagination)
}
func (Repository) QueryExaminationTasksGroup(queries []datastore.Query) ([]examination_task.ExaminationTask, error) {
	return QueryExaminationTasksGroup(queries)
}
func (Repository) CreateExaminationTask(entity examination_task.ExaminationTask, editorID *string) (string, error) {
	return CreateExaminationTask(entity, editorID)
}
func (Repository) UpdateExaminationTask(entity examination_task.ExaminationTask, editorID *string) (string, error) {
	return UpdateExaminationTask(entity, editorID)
}
func (Repository) SaveExaminationTask(entity examination_task.ExaminationTask, editorID *string) (string, error) {
	return SaveExaminationTask(entity, editorID)
}
func (Repository) DeleteExaminationTask(examinationTaskID string) error {
	return DeleteExaminationTask(examinationTaskID)
}
func (Repository) GetAttribute(attributeID string) (attribute.Attribute, error) {
	return GetAttribute(attributeID)
}
func (Repository) GetAttributeOrNew(attributeID string) (attribute.Attribute, error) {
	return GetAttributeOrNew(attributeID)
}
func (Repository) GetAttributeByPath(path string) (attribute.Attribute, error) {
	return GetAttributeByPath(path)
}
func (Repository) FindAttribute(queries []datastore.Query) (attribute.Attribute, error) {
	return FindAttribute(queries)
}
func (Repository) GetAllAttributes() ([]attribute.Attribute, error) {
	return GetAllAttributes()
}
func (Repository) GetAllAttributesPaginated(pagination datastore.Pagination) ([]attribute.Attribute, datastore.Pagination, error) {
	return GetAllAttributesPaginated(pagination)
}
func (Repository) QueryAttributes(queries []datastore.Query, pagination datastore.Pagination) ([]attribute.Attribute, error) {
	return QueryAttributes(queries, pagination)
}
func (Repository) QueryAttributesGroup(queries []datastore.Query) ([]attribute.Attribute, error) {
	return QueryAttributesGroup(queries)
}
func (Repository) CreateAttribute(entity attribute.Attribute, editorID *string) (string, error) {
	return CreateAttribute(entity, editorID)
}
func (Repository) UpdateAttribute(entity attribute.Attribute, editorID *string) (string, error) {
	return UpdateAttribute(entity, editorID)
}
func (Repository) SaveAttribute(entity attribute.Attribute, editorID *string) (string, error) {
	return SaveAttribute(entity, editorID)
}
func (Repository) DeleteAttribute(attributeID string) error {
	return DeleteAttribute(attributeID)
}
func (Repository) GetAttributeOptionSet(attributeOptionSetID string) (attribute_option_set.AttributeOptionSet, error) {
	return GetAttributeOptionSet(attributeOptionSetID)
}
func (Repository) GetAttributeOptionSetOrNew(attributeOptionSetID string) (attribute_option_set.AttributeOptionSet, error) {
	return GetAttributeOptionSetOrNew(attributeOptionSetID)
}
func (Repository) GetAttributeOptionSetByPath(path string) (attribute_option_set.AttributeOptionSet, error) {
	return GetAttributeOptionSetByPath(path)
}
func (Repository) FindAttributeOptionSet(queries []datastore.Query) (attribute_option_set.AttributeOptionSet, error) {
	return FindAttributeOptionSet(queries)
}
func (Repository) GetAllAttributeOptionSets() ([]attribute_option_set.AttributeOptionSet, error) {
	return GetAllAttributeOptionSets()
}
func (Repository) GetAllAttributeOptionSetsPaginated(pagination datastore.Pagination) ([]attribute_option_set.AttributeOptionSet, datastore.Pagination, error) {
	return GetAllAttributeOptionSetsPaginated(pagination)
}
func (Repository) QueryAttributeOptionSets(queries []datastore.Query, pagination datastore.Pagination) ([]attribute_option_set.AttributeOptionSet, error) {
	return QueryAttributeOptionSets(queries, pagination)
}
func (Repository) QueryAttributeOptionSetsGroup(queries []datastore.Query) ([]attribute_option_set.AttributeOptionSet, error) {
	return QueryAttributeOptionSetsGroup(queries)
}
func (Repository) CreateAttributeOptionSet(entity attribute_option_set.AttributeOptionSet, editorID *string) (string, error) {
	return CreateAttributeOptionSet(entity, editorID)
}
func (Repository) UpdateAttributeOptionSet(entity attribute_option_set.AttributeOptionSet, editorID *string) (string, error) {
	return UpdateAttributeOptionSet(entity, editorID)
}
func (Repository) SaveAttributeOptionSet(entity attribute_option_set.AttributeOptionSet, editorID *string) (string, error) {
	return SaveAttributeOptionSet(entity, editorID)
}
func (Repository) DeleteAttributeOptionSet(attributeOptionSetID string) error {
	return DeleteAttributeOptionSet(attributeOptionSetID)
}
func (Repository) GetAttributeOption(attributeOptionID string) (attribute_option.AttributeOption, error) {
	return GetAttributeOption(attributeOptionID)
}
func (Repository) GetAttributeOptionOrNew(attributeOptionID string) (attribute_option.AttributeOption, error) {
	return GetAttributeOptionOrNew(attributeOptionID)
}
func (Repository) GetAttributeOptionByPath(path string) (attribute_option.AttributeOption, error) {
	return GetAttributeOptionByPath(path)
}
func (Repository) FindAttributeOption(queries []datastore.Query) (attribute_option.AttributeOption, error) {
	return FindAttributeOption(queries)
}
func (Repository) GetAllAttributeOptions() ([]attribute_option.AttributeOption, error) {
	return GetAllAttributeOptions()
}
func (Repository) GetAllAttributeOptionsPaginated(pagination datastore.Pagination) ([]attribute_option.AttributeOption, datastore.Pagination, error) {
	return GetAllAttributeOptionsPaginated(pagination)
}
func (Repository) QueryAttributeOptions(queries []datastore.Query, pagination datastore.Pagination) ([]attribute_option.AttributeOption, error) {
	return QueryAttributeOptions(queries, pagination)
}
func (Repository) QueryAttributeOptionsGroup(queries []datastore.Query) ([]attribute_option.AttributeOption, error) {
	return QueryAttributeOptionsGroup(queries)
}
func (Repository) CreateAttributeOption(entity attribute_option.AttributeOption, editorID *string) (string, error) {
	return CreateAttributeOption(entity, editorID)
}
func (Repository) UpdateAttributeOption(entity attribute_option.AttributeOption, editorID *string) (string, error) {
	return UpdateAttributeOption(entity, editorID)
}
func (Repository) SaveAttributeOption(entity attribute_option.AttributeOption, editorID *string) (string, error) {
	return SaveAttributeOption(entity, editorID)
}
func (Repository) DeleteAttributeOption(attributeOptionID string) error {
	return DeleteAttributeOption(attributeOptionID)
}