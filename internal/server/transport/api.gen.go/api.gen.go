package api

import (
	"github.com/bakeable/bkry/api/middleware"
    user_profile_routes "github.com/bakeable/bkry/api/user_profile/routes"
    user_routes "github.com/bakeable/bkry/api/user/routes"
    portal_event_routes "github.com/bakeable/bkry/api/portal_event/routes"
    media_routes "github.com/bakeable/bkry/api/media/routes"
    airline_settings_routes "github.com/bakeable/bkry/api/airline_settings/routes"
    delivery_entry_routes "github.com/bakeable/bkry/api/delivery_entry/routes"
    printing_order_supplier_routes "github.com/bakeable/bkry/api/printing_order_supplier/routes"
    product_routes "github.com/bakeable/bkry/api/product/routes"
    examination_settings_routes "github.com/bakeable/bkry/api/examination_settings/routes"
    price_configuration_routes "github.com/bakeable/bkry/api/price_configuration/routes"
    printing_order_item_routes "github.com/bakeable/bkry/api/printing_order_item/routes"
    packing_slip_package_routes "github.com/bakeable/bkry/api/packing_slip_package/routes"
    cloud_function_routes "github.com/bakeable/bkry/api/cloud_function/routes"
    api_key_routes "github.com/bakeable/bkry/api/api_key/routes"
    quotation_routes "github.com/bakeable/bkry/api/quotation/routes"
    examination_property_routes "github.com/bakeable/bkry/api/examination_property/routes"
    airline_order_grouping_routes "github.com/bakeable/bkry/api/airline_order_grouping/routes"
    price_layer_routes "github.com/bakeable/bkry/api/price_layer/routes"
    airline_order_group_routes "github.com/bakeable/bkry/api/airline_order_group/routes"
    airline_order_routes "github.com/bakeable/bkry/api/airline_order/routes"
    airline_pricing_routes "github.com/bakeable/bkry/api/airline_pricing/routes"
    printing_order_routes "github.com/bakeable/bkry/api/printing_order/routes"
    packing_slip_routes "github.com/bakeable/bkry/api/packing_slip/routes"
    email_routes "github.com/bakeable/bkry/api/email/routes"
    batch_export_routes "github.com/bakeable/bkry/api/batch_export/routes"
    examination_task_routes "github.com/bakeable/bkry/api/examination_task/routes"
    attribute_routes "github.com/bakeable/bkry/api/attribute/routes"
    attribute_option_set_routes "github.com/bakeable/bkry/api/attribute_option_set/routes"
    attribute_option_routes "github.com/bakeable/bkry/api/attribute_option/routes"

	"github.com/gin-gonic/gin"
)

// InitRoutes initializes all routes for different models
func InitRoutes(r *gin.RouterGroup) {
	r.Use(gin.RecoveryWithWriter(gin.DefaultErrorWriter))
	r.Use(middleware.RecoveryWithWriter())


	// Initialize routes for user_profile
	user_profile_routes.InitRoutes(r)

	// Initialize routes for user
	user_routes.InitRoutes(r)

	// Initialize routes for portal_event
	portal_event_routes.InitRoutes(r)

	// Initialize routes for media
	media_routes.InitRoutes(r)

	// Initialize routes for airline_settings
	airline_settings_routes.InitRoutes(r)

	// Initialize routes for delivery_entry
	delivery_entry_routes.InitRoutes(r)

	// Initialize routes for printing_order_supplier
	printing_order_supplier_routes.InitRoutes(r)

	// Initialize routes for product
	product_routes.InitRoutes(r)

	// Initialize routes for examination_settings
	examination_settings_routes.InitRoutes(r)

	// Initialize routes for price_configuration
	price_configuration_routes.InitRoutes(r)

	// Initialize routes for printing_order_item
	printing_order_item_routes.InitRoutes(r)

	// Initialize routes for packing_slip_package
	packing_slip_package_routes.InitRoutes(r)

	// Initialize routes for cloud_function
	cloud_function_routes.InitRoutes(r)

	// Initialize routes for api_key
	api_key_routes.InitRoutes(r)

	// Initialize routes for quotation
	quotation_routes.InitRoutes(r)

	// Initialize routes for examination_property
	examination_property_routes.InitRoutes(r)

	// Initialize routes for airline_order_grouping
	airline_order_grouping_routes.InitRoutes(r)

	// Initialize routes for price_layer
	price_layer_routes.InitRoutes(r)

	// Initialize routes for airline_order_group
	airline_order_group_routes.InitRoutes(r)

	// Initialize routes for airline_order
	airline_order_routes.InitRoutes(r)

	// Initialize routes for airline_pricing
	airline_pricing_routes.InitRoutes(r)

	// Initialize routes for printing_order
	printing_order_routes.InitRoutes(r)

	// Initialize routes for packing_slip
	packing_slip_routes.InitRoutes(r)

	// Initialize routes for email
	email_routes.InitRoutes(r)

	// Initialize routes for batch_export
	batch_export_routes.InitRoutes(r)

	// Initialize routes for examination_task
	examination_task_routes.InitRoutes(r)

	// Initialize routes for attribute
	attribute_routes.InitRoutes(r)

	// Initialize routes for attribute_option_set
	attribute_option_set_routes.InitRoutes(r)

	// Initialize routes for attribute_option
	attribute_option_routes.InitRoutes(r)


	// Initialize custom routes
	InitCustomRoutes(r)
}
