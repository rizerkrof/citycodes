package interfaces

type CreateSecretCache struct {
	Name     string `json:"name" validate:"required"`
	ImageUrl string `json:"imageUrl" validate:"required"`
}
