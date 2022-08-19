package services

import (
	"bitbucket.org/devsu/go-template/repositories"
	"bitbucket.org/devsu/go-template/services/dto"
)

type SampleService interface {
	GetHelloWorld() string
	GetUserByID(id int) (dto.UserDto, error)
}

type SampleServiceImpl struct {
	sampleRepository repositories.SampleRepository
}

func NewSampleServiceImpl() *SampleServiceImpl {
	sampleRepository := repositories.NewSampleDummyRepository()
	return &SampleServiceImpl{
		sampleRepository: sampleRepository,
	}
}

func (s *SampleServiceImpl) GetHelloWorld() string {
	return "Hello World"
}

func (s *SampleServiceImpl) GetUserByID(id int) (dto.UserDto, error) {
	user, err := s.sampleRepository.GetUserByID(id)

	if err != nil {
		return dto.UserDto{}, err
	}

	userDto := dto.UserDto{
		Name: user.Name,
		Age:  user.Age,
	}

	return userDto, nil
}
