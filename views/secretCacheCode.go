package views

import (
	"citycodes/services"
	"citycodes/templates/pages"
	"fmt"

	"github.com/go-fuego/fuego"
)

type secretCacheRessource struct {
	SecretCacheService services.SecretCacheServiceRessource
}

func (rs secretCacheRessource) Routes(s *fuego.Server) {
	secretCacheRoutesGroupe := fuego.Group(s, "/secret-caches")
	fuego.Get(secretCacheRoutesGroupe, "/{id}", rs.SecretCachePage)
	fuego.Post(secretCacheRoutesGroupe, "/{id}", rs.PostSecretCacheImageUrl)
}

func (rs secretCacheRessource) SecretCachePage(c fuego.ContextNoBody) (fuego.Templ, error) {
	secretCache, err := rs.SecretCacheService.GetSecretCacheById(c)
	if err != nil {
		return nil, fmt.Errorf("Error. Getting secret cache failed.")
	}

	return pages.SecretCachePage(*secretCache), nil
}

func (rs secretCacheRessource) PostSecretCacheImageUrl(c *fuego.ContextNoBody) (any, error) {
	secretCache, err := rs.SecretCacheService.PostSecretCacheImage(c)
	if err != nil {
		return nil, fmt.Errorf("Error. Uploading secret cache image failed.")
	}

	return c.Redirect(301, "/secret-caches/"+secretCache.ID)
}
