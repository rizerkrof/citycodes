package controllers

import (
	"citycodes/store"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"path/filepath"

	"github.com/go-fuego/fuego"
)

type secretCacheRessource struct {
	SecretCacheRepository SecretCacheRepository
}

type SecretCacheRepository interface {
	CreateSecretCache(ctx context.Context, arg store.CreateSecretCacheParams) (store.SecretCache, error)
	GetSecretCache(ctx context.Context, id string) (store.SecretCache, error)
	GetSecretCaches(ctx context.Context) ([]store.SecretCache, error)
	PatchSecretCacheImageUrl(ctx context.Context, arg store.PatchSecretCacheImageUrlParams) (store.SecretCache, error)
}

func (rs secretCacheRessource) Routes(s *fuego.Server) {
	secretCacheRoutesGroupe := fuego.Group(s, "/secret-caches")
	fuego.Get(secretCacheRoutesGroupe, "/", rs.getAllSecretCaches)
	fuego.Get(secretCacheRoutesGroupe, "/unique", rs.getSecretCacheById).WithQueryParam("id", "The secret cache id.")
	fuego.Post(secretCacheRoutesGroupe, "/new", rs.newSecretCache)
	fuego.Post(secretCacheRoutesGroupe, "/image/new", rs.addSecretCacheImage).WithQueryParam("id", "The secret cache id.")
}

func (rs secretCacheRessource) getAllSecretCaches(c fuego.Ctx[any]) ([]store.SecretCache, error) {
	caches, err := rs.SecretCacheRepository.GetSecretCaches(c.Context())
	if err != nil {
		return nil, err
	}

	slog.Info("Neww cache ------------------")

	return caches, nil
}

func (rs secretCacheRessource) getSecretCacheById(c fuego.Ctx[any]) (store.SecretCache, error) {
	id := c.Request().URL.Query().Get("id")
	cache, err := rs.SecretCacheRepository.GetSecretCache(c.Context(), id)
	if err != nil {
		return store.SecretCache{}, err
	}

	return cache, nil
}

type CreateSecretCache struct {
	Name     string `json:"name" validate:"required"`
	ImageUrl string `json:"imageUrl" validate:"required"`
}

func (rs secretCacheRessource) newSecretCache(c fuego.Ctx[CreateSecretCache]) (store.SecretCache, error) {
	slog.Info("Neww cache ------------------")

	body, err := c.Body()
	if err != nil {
		return store.SecretCache{}, err
	}

	payload := store.CreateSecretCacheParams{
		ID:       generateID(),
		Name:     body.Name,
		ImageUrl: body.ImageUrl,
	}

	cache, err := rs.SecretCacheRepository.CreateSecretCache(c.Context(), payload)
	if err != nil {
		return store.SecretCache{}, err
	}

	return cache, nil
}

func (rs secretCacheRessource) addSecretCacheImage(c fuego.Ctx[CreateSecretCache]) (any, error) {
	slog.Info("image")
	slog.Info(c.Request().URL.Query().Get("id"))
	slog.Info("done")
	// Parse the multipart form in the request
	err := c.Request().ParseMultipartForm(10 >> 20) // limit your maxMemory here

	file, handler, err := c.Request().FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return store.SecretCache{}, err
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	secretCodeId := c.Request().URL.Query().Get("id")
	secretCodeDirectoryPath := filepath.Join(".", filepath.Dir("static/images/secretCacheCodes/"+secretCodeId+"/"))
	slog.Info(secretCodeDirectoryPath)
	err = os.MkdirAll(secretCodeDirectoryPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	// a particular naming pattern
	tempFile, err := os.CreateTemp(secretCodeDirectoryPath, "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	secretCodeUpdateImageUrlPayload := store.PatchSecretCacheImageUrlParams{
		ID:       secretCodeId,
		ImageUrl: tempFile.Name(),
	}

	slog.Info(secretCodeUpdateImageUrlPayload.ID)

	secretCache, err := rs.SecretCacheRepository.PatchSecretCacheImageUrl(c.Context(), secretCodeUpdateImageUrlPayload)

	if err != nil {
		return store.SecretCache{}, err
	}

	return c.Redirect(301, "/secretCacheCode?id="+secretCache.ID)
}
