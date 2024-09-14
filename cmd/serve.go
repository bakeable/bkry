package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bakeable/bkry/api"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	"github.com/bakeable/bkry/third_party/gcloud/firebase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the application",
	Long:  `This command starts the web server and serves the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get path to this script
		path, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}
		godotenv.Load(path + "/.env")

		// Initialize Firebase
		firebase.Init()

		// Initialize datastore
		datastore.Init(os.Getenv("GOOGLE_CLOUD_PROJECT"))

		// Initialize Gin
		if os.Getenv("GO_ENV") == "production" {
			// Set Gin to production mode
			gin.SetMode(gin.ReleaseMode)
		}

		// Create Gin router
		r := gin.Default()

		// Allow cross origin requests for development purposes
		config := cors.DefaultConfig()

		// Get allowed origins
		allowedOrigins := os.Getenv("GO_ALLOWED_ORIGINS")
		config.AllowOrigins = []string{"http://localhost:5173"}
		if allowedOrigins != "" {
			config.AllowOrigins = append(config.AllowOrigins, strings.Split(os.Getenv("GO_ALLOWED_ORIGINS"), ",")...)
		}
		log.Println("Allowed origins: ", config.AllowOrigins)

		// Allow credentials
		config.AllowCredentials = true

		// Allow headers
		config.AllowHeaders = []string{"access-control-allow-origin", "content-type", "authorization"}

		// Use cors
		r.Use(cors.New(config))

		// Setup route group for the API
		apiGroup := r.Group("/api")
		{
			// Ping test
			apiGroup.GET("/ping", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "pong",
				})
			})

			// Initialize other routes
			api.InitRoutes(apiGroup)
		}

		r.Run(":8080")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
