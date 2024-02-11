package views

import (
	"citycodes/store"
	"citycodes/templates/pages"
	"context"
	"fmt"

	"github.com/go-fuego/fuego"
)

type secretCacheRessource struct {
	SecretCacheRepository SecretCacheRepository
}

type SecretCacheRepository interface {
	CreateSecretCache(ctx context.Context, arg store.CreateSecretCacheParams) (store.SecretCache, error)
	GetSecretCache(ctx context.Context, id string) (store.SecretCache, error)
	GetSecretCaches(ctx context.Context) ([]store.SecretCache, error)
}

func (rs Ressource) SecretCachePage(c fuego.Ctx[any]) (fuego.Templ, error) {
	id := c.QueryParam("id")

	secretCache, err := rs.SecretCacheRepository.GetSecretCache(c.Context(), id)

	if err != nil {
		return nil, fmt.Errorf("error getting secret cache %s: %w", id, err)
	}

	return pages.SecretCachePage(secretCache), nil
}
