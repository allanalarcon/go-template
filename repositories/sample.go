package repositories

import "bitbucket.org/devsu/go-template/models"

type SampleRepository interface {
	GetUserByID(id int) (models.User, error)
}

type SampleDummyRepository struct{}

func NewSampleDummyRepository() *SampleDummyRepository {
	repository := &SampleDummyRepository{}
	return repository
}

func (repo *SampleDummyRepository) GetUserByID(id int) (models.User, error) {
	return models.User{
		ID:   10,
		Name: "Jack Sparrow",
		Age:  32,
	}, nil
}
