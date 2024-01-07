package controller

import (
	"time"

	"citycodes/store"

	"github.com/go-fuego/fuego"
)

// Ressource is the struct that holds useful sources of informations available for the controllers.
func NewRessource(queries store.Queries) Ressource {
	return Ressource{
		Queries: queries,
	}
}

// Ressource is the struct that holds useful sources of informations available for the controllers.
type Ressource struct {
	Queries     store.Queries          // Database queries
	UserQueries store.Queries          // Database queries from another store
	ExternalAPI interface{}            // External API
	Cache       map[string]interface{} // Some cache
	Now         func() time.Time       // Function to get the current time. Mocked in tests.
	Security    fuego.Security            // Security configuration
}

func (rs Ressource) Routes(s *fuego.Server) {
	fuego.Get(s, "/caches/", rs.getAllCaches)
	fuego.Get(s, "/caches/:id", rs.getCacheById).WithQueryParam("id", "The cache id.")
	fuego.Post(s, "/caches/new", rs.newCache)
}
