package controllers

import (
	"citycodes/services"
	"time"

	"github.com/go-fuego/fuego"
)

// Ressource is the struct that holds useful sources of informations available for the controllers.
type Ressource struct {
	SecretCacheRepository services.SecretCacheRepository
	ExternalAPI           interface{}            // External API
	Cache                 map[string]interface{} // Some cache
	Now                   func() time.Time       // Function to get the current time. Mocked in tests.
	Security              fuego.Security         // Security configuration
}

func (rs Ressource) Routes(s *fuego.Server) {
	SecretCacheControllerRessource{
		SecretCacheService: services.SecretCacheServiceRessource{
			SecretCacheRepository: rs.SecretCacheRepository,
		},
	}.Routes(s)
}
