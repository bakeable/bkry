package airline_pricing_controllers

import (
	repo "github.com/bakeable/bkry/data/repository/entities"

	"github.com/gin-gonic/gin"
)

func getRepository() repo.IRepository {
	return repo.NewRepository()
}

// Create interface for the controllers
type IControllers interface {
	Add(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	GetAllPaginated(c *gin.Context)
	Query(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// Create the controllers struct
type Controllers struct {}
func NewControllers() IControllers {
	return Controllers{}
}

// Wrap the controllers
func (Controllers) Add(c *gin.Context) {
	Add(c)
}
func (Controllers) Get(c *gin.Context) {
	Get(c)
}
func (Controllers) GetAll(c *gin.Context) {
	GetAll(c)
}
func (Controllers) GetAllPaginated(c *gin.Context) {
	GetAllPaginated(c)
}
func (Controllers) Query(c *gin.Context) {
	Query(c)
}
func (Controllers) Update(c *gin.Context) {
	Update(c)
}
func (Controllers) Delete(c *gin.Context) {
	Delete(c)
}


