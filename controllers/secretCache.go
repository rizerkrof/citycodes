package controllers

import (
	"citycodes/interfaces"
	"citycodes/services"
	"citycodes/store"

	"github.com/go-fuego/fuego"
)

type SecretCacheControllerRessource struct {
	SecretCacheService services.SecretCacheServiceRessource
}

func (rs SecretCacheControllerRessource) Routes(s *fuego.Server) {
	secretCacheRoutesGroupe := fuego.Group(s, "/secret-caches")
	fuego.Get(secretCacheRoutesGroupe, "/", rs.getAllSecretCaches)
	fuego.Post(secretCacheRoutesGroupe, "/", rs.createSecretCache)
	fuego.Get(secretCacheRoutesGroupe, "/{id}", rs.getSecretCacheById)
	fuego.Post(secretCacheRoutesGroupe, "/{id}", rs.postSecretCacheImage)
}

func (rs SecretCacheControllerRessource) getAllSecretCaches(c fuego.ContextNoBody) ([]store.SecretCache, error) {
	secretCaches, err := rs.SecretCacheService.GetAllSecretCaches(c.Context())
	if err != nil {
		return nil, err
	}

	return secretCaches, nil
}

func (rs SecretCacheControllerRessource) getSecretCacheById(c fuego.ContextNoBody) (*store.SecretCache, error) {
	secretCache, err := rs.SecretCacheService.GetSecretCacheById(c)
	if err != nil {
		return nil, err
	}

	return secretCache, nil
}

func (rs SecretCacheControllerRessource) createSecretCache(c *fuego.ContextWithBody[interfaces.CreateSecretCache]) (*store.SecretCache, error) {
	createdSecretCache, err := rs.SecretCacheService.CreateSecretCache(c)
	if err != nil {
		return nil, err
	}

	return createdSecretCache, nil
}

func (rs SecretCacheControllerRessource) postSecretCacheImage(c *fuego.ContextNoBody) (*store.SecretCache, error) {
	postedSecretCode, err := rs.SecretCacheService.PostSecretCacheImage(c)
	if err != nil {
		return nil, err
	}

	return postedSecretCode, nil
}
