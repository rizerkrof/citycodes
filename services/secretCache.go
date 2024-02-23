package services

import (
	"citycodes/interfaces"
	"citycodes/store"
	"context"
	"io"
	"log/slog"
	"os"

	"path/filepath"

	"github.com/go-fuego/fuego"
)

type SecretCacheRepository interface {
	CreateSecretCache(ctx context.Context, arg store.CreateSecretCacheParams) (store.SecretCache, error)
	GetSecretCache(ctx context.Context, id string) (store.SecretCache, error)
	GetSecretCaches(ctx context.Context) ([]store.SecretCache, error)
	PatchSecretCacheImageUrl(ctx context.Context, arg store.PatchSecretCacheImageUrlParams) (store.SecretCache, error)
}

type SecretCacheServiceRessource struct {
	SecretCacheRepository SecretCacheRepository
}

func (rs SecretCacheServiceRessource) GetAllSecretCaches(c context.Context) ([]store.SecretCache, error) {
	secretCache, err := rs.SecretCacheRepository.GetSecretCaches(c)
	if err != nil {
		return nil, err
	}

	return secretCache, nil
}

func (rs SecretCacheServiceRessource) GetSecretCacheById(c fuego.ContextNoBody) (*store.SecretCache, error) {
	id := c.Request().PathValue("id")

	secretCache, err := rs.SecretCacheRepository.GetSecretCache(c.Context(), id)
	if err != nil {
		return nil, err
	}

	return &secretCache, nil
}

func (rs SecretCacheServiceRessource) CreateSecretCache(c *fuego.ContextWithBody[interfaces.CreateSecretCache]) (*store.SecretCache, error) {
	body, err := c.Body()
	if err != nil {
		return nil, err
	}

	payload := store.CreateSecretCacheParams{
		ID:       generateID(),
		Name:     body.Name,
		ImageUrl: body.ImageUrl,
	}

	createdSecretCache, err := rs.SecretCacheRepository.CreateSecretCache(c.Context(), payload)
	if err != nil {
		return nil, err
	}

	return &createdSecretCache, nil
}

func (rs SecretCacheServiceRessource) PostSecretCacheImage(c *fuego.ContextNoBody) (*store.SecretCache, error) {
	slog.Info("1")
	err := c.Request().ParseMultipartForm(10 >> 20)
	file, _, err := c.Request().FormFile("image")
	if err != nil {
		slog.Error("laaa")
		return nil, err
	}
	defer file.Close()

	secretCodeId := c.Request().PathValue("id")

	secretCodeDirectoryPath := filepath.Join(".", filepath.Dir("static/images/secretCacheCodes/"+secretCodeId+"/"))

	err = os.MkdirAll(secretCodeDirectoryPath, os.ModePerm)
	if err != nil {
		return nil, err
	}

	tempFile, err := os.CreateTemp(secretCodeDirectoryPath, "*.png")
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	tempFile.Write(fileBytes)

	secretCodeUpdateImageUrlPayload := store.PatchSecretCacheImageUrlParams{
		ID:       secretCodeId,
		ImageUrl: tempFile.Name(),
	}

	postedSecretCode, err := rs.SecretCacheRepository.PatchSecretCacheImageUrl(c.Context(), secretCodeUpdateImageUrlPayload)
	if err != nil {
		return nil, err
	}

	return &postedSecretCode, nil
}
