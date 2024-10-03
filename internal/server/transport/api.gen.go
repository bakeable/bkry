package transport

import (
	"github.com/bakeable/bkry/internal/server/transport/middleware"
    author_routes "github.com/bakeable/bkry/internal/server/transport/routes/author"
    category_routes "github.com/bakeable/bkry/internal/server/transport/routes/category"
    game_modus_routes "github.com/bakeable/bkry/internal/server/transport/routes/game_modus"
    media_routes "github.com/bakeable/bkry/internal/server/transport/routes/media"
    product_package_routes "github.com/bakeable/bkry/internal/server/transport/routes/product_package"
    purchase_routes "github.com/bakeable/bkry/internal/server/transport/routes/purchase"
    question_bundle_routes "github.com/bakeable/bkry/internal/server/transport/routes/question_bundle"
    question_context_routes "github.com/bakeable/bkry/internal/server/transport/routes/question_context"
    tag_routes "github.com/bakeable/bkry/internal/server/transport/routes/tag"
    user_routes "github.com/bakeable/bkry/internal/server/transport/routes/user"
    category_localisation_routes "github.com/bakeable/bkry/internal/server/transport/routes/category_localisation"
    game_modus_localisation_routes "github.com/bakeable/bkry/internal/server/transport/routes/game_modus_localisation"
    product_package_localisation_routes "github.com/bakeable/bkry/internal/server/transport/routes/product_package_localisation"
    question_bundle_localisation_routes "github.com/bakeable/bkry/internal/server/transport/routes/question_bundle_localisation"
    question_context_localisation_routes "github.com/bakeable/bkry/internal/server/transport/routes/question_context_localisation"
    tag_localisation_routes "github.com/bakeable/bkry/internal/server/transport/routes/tag_localisation"

	"github.com/gin-gonic/gin"
)

// InitRoutes initializes all routes for different models
func InitRoutes(r *gin.RouterGroup) {
	r.Use(gin.RecoveryWithWriter(gin.DefaultErrorWriter))
	r.Use(middleware.RecoveryWithWriter())


	// Initialize routes for author
	author_routes.InitRoutes(r)

	// Initialize routes for category
	category_routes.InitRoutes(r)

	// Initialize routes for game_modus
	game_modus_routes.InitRoutes(r)

	// Initialize routes for media
	media_routes.InitRoutes(r)

	// Initialize routes for product_package
	product_package_routes.InitRoutes(r)

	// Initialize routes for purchase
	purchase_routes.InitRoutes(r)

	// Initialize routes for question_bundle
	question_bundle_routes.InitRoutes(r)

	// Initialize routes for question_context
	question_context_routes.InitRoutes(r)

	// Initialize routes for tag
	tag_routes.InitRoutes(r)

	// Initialize routes for user
	user_routes.InitRoutes(r)

	// Initialize routes for category_localisation
	category_localisation_routes.InitRoutes(r)

	// Initialize routes for game_modus_localisation
	game_modus_localisation_routes.InitRoutes(r)

	// Initialize routes for product_package_localisation
	product_package_localisation_routes.InitRoutes(r)

	// Initialize routes for question_bundle_localisation
	question_bundle_localisation_routes.InitRoutes(r)

	// Initialize routes for question_context_localisation
	question_context_localisation_routes.InitRoutes(r)

	// Initialize routes for tag_localisation
	tag_localisation_routes.InitRoutes(r)


	// Initialize custom routes
	InitCustomRoutes(r)
}
