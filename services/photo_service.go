package services

import (
	"errors"

	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/app/request"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/models"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/repository"
)

type PhotoService interface {
	Create(inputID int, inputData request.PhotoInput) (models.Photo, error)
	GetByUserID(userID int) (models.Photo, error)
	GetByID(ID int) (models.Photo, error)
	Update(inputID int, inputData request.PhotoUpdate) (models.Photo, error)
	Delete(ID int) error
}

func NewPhotoService(photoRepository repository.PhotoRepository) PhotoService {
	return &PhotoServiceImpl{photoRepository: photoRepository}
}

type PhotoServiceImpl struct {
	photoRepository repository.PhotoRepository
}

func (s *PhotoServiceImpl) Create(inputID int, inputData request.PhotoInput) (models.Photo, error) {

	photo := models.Photo{}
	photo.Title = inputData.Title
	photo.Caption = inputData.Caption
	photo.PhotoURL = inputData.PhotoURL
	photo.UserID = inputID

	created, err := s.photoRepository.Create(photo)
	if err != nil {
		return created, err
	}
	return created, nil
}

func (s *PhotoServiceImpl) GetByUserID(userID int) (models.Photo, error) {
	photo, err := s.photoRepository.FindByUserID(userID)
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (s *PhotoServiceImpl) GetByID(ID int) (models.Photo, error) {
	photo, err := s.photoRepository.FindByID(ID)
	if err != nil {
		return photo, err
	}

	if photo.ID == 0 {
		return photo, errors.New("Photo not found")
	}
	return photo, nil
}

func (s *PhotoServiceImpl) Update(inputID int, inputData request.PhotoUpdate) (models.Photo, error) {

	photo, err := s.photoRepository.FindByUserID(inputID)
	if err != nil {
		return photo, err
	}

	photo.Title = inputData.Title
	photo.Caption = inputData.Caption
	photo.PhotoURL = inputData.PhotoURL

	updatedPhoto, err := s.photoRepository.Update(photo)
	if err != nil {
		return updatedPhoto, err
	}

	return updatedPhoto, nil

}

func (s *PhotoServiceImpl) Delete(ID int) error {
	err := s.photoRepository.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
