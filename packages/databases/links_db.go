package databases

import "hungon.space/xurl/app/models"

type LinkRepository interface {
	CreateLink(link *models.Link) (*models.Link, error)
	GetLinkById(id string) (*models.Link, error)
	GetLinkByLinkId(linkId string) (*models.Link, error)
	GetLinksByApiKey(apiKey string) ([]*models.Link, error)
	DeleteLinkById(id string) error
	DeleteLinkByLinkId(linkId string) error
	LastId() (uint, error)
}

type linkRepository struct {
}

func NewLinkRepository() LinkRepository {
	return &linkRepository{}
}

func (r *linkRepository) CreateLink(link *models.Link) (*models.Link, error) {
	error := GetDB().Create(&link).Error
	if error != nil {
		return nil, error
	}

	return link, nil
}

func (r *linkRepository) GetLinkById(id string) (*models.Link, error) {
	var link models.Link
	error := GetDB().Where("id = ?", id).First(&link).Error
	if error != nil {
		return nil, error
	}

	return &link, nil
}

func (r *linkRepository) GetLinkByLinkId(linkId string) (*models.Link, error) {
	var link models.Link
	error := GetDB().Where("link_id = ?", linkId).First(&link).Error
	if error != nil {
		return nil, error
	}

	return &link, nil
}

func (r *linkRepository) GetLinksByApiKey(apiKey string) ([]*models.Link, error) {
	var links []*models.Link
	error := GetDB().Where("api_key = ?", apiKey).Find(&links).Error
	if error != nil {
		return nil, error
	}

	return links, nil
}

func (r *linkRepository) DeleteLinkById(id string) error {
	error := GetDB().Where("id = ?", id).Delete(&models.Link{})

	if error != nil {
		return error.Error
	}

	return nil
}

func (r *linkRepository) DeleteLinkByLinkId(linkId string) error {
	error := GetDB().Where("link_id = ?", linkId).Delete(&models.Link{})

	if error != nil {
		return error.Error
	}

	return nil
}

func (r *linkRepository) LastId() (uint, error) {
	var link models.Link
	error := GetDB().Last(&link).Error
	if error != nil {
		return 0, error
	}

	return link.Id, nil
}
