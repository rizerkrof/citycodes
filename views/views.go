package views

import (
	"github.com/go-fuego/fuego"
)

type Ressource struct {
	SecretCacheRepository SecretCacheRepository
}

func (rs Ressource) Routes(s *fuego.Server) {
	fuego.Get(s, "/secretCacheCode", rs.SecretCachePage)
}
