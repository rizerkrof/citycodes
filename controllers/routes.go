package controllers

import (
	"time"

	"github.com/go-fuego/fuego"
)

// Ressource is the struct that holds useful sources of informations available for the controllers.
type Ressource struct {
	SecretCacheRepository SecretCacheRepository
	ExternalAPI           interface{}            // External API
	Cache                 map[string]interface{} // Some cache
	Now                   func() time.Time       // Function to get the current time. Mocked in tests.
	Security              fuego.Security         // Security configuration
}

func (rs Ressource) Routes(s *fuego.Server) {
	secretCacheRessource{
		SecretCacheRepository: rs.SecretCacheRepository,
	}.Routes(s)
}
