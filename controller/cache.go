package controller

import (
	"citycodes/store"
	"github.com/go-fuego/fuego"
)

func (rs Ressource) getAllCaches(c fuego.Ctx[any]) ([]store.Cache, error) {
	caches, err:= rs.Queries.GetCaches(c.Context())
	if err != nil {
		return nil, err
	}

	return caches, nil
}

func (rs Ressource) getCacheById(c fuego.Ctx[any]) (store.Cache, error) {
	cache, err := rs.Queries.GetCache(c.Context(), "id")
	if err != nil {
		return store.Cache{}, err
	}

	return cache, nil
}

type CreateCache struct {
	Name string `json:"name" validate:"required"`
}

func (rs Ressource) newCache(c fuego.Ctx[CreateCache]) (store.Cache, error) {
	body, err := c.Body()
	if err != nil {
		return store.Cache{}, err
	}

	payload:= store.CreateCacheParams{
		ID: generateID(),
		Name: body.Name,
	}

	cache, err := rs.Queries.CreateCache(c.Context(), payload)
	if err != nil {
		return store.Cache{}, err
	}

	return cache, nil
}
