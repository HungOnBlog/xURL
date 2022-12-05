package dto

type CreateLinkDTO struct {
	OriginalLink string `json:"originalLink" validate:"required,url"`
}
