package views

import (
	"citycodes/services"

	"github.com/go-fuego/fuego"
)

type Ressource struct {
	SecretCacheRepository services.SecretCacheRepository
}

func (rs Ressource) Routes(s *fuego.Server) {
	secretCacheRessource{
		SecretCacheService: services.SecretCacheServiceRessource{
			SecretCacheRepository: rs.SecretCacheRepository,
		},
	}.Routes(s)
}
