package repository

import (
	"github.com/febrielven/go_crud_posgres/module/model"
)

//ProfileRepository
type ProfileRepository interface {
	Save(*model.Profile) error

	Update(int, *model.Profile) error
	Delete(int) error
	FindById(int) (*model.Profile, error)
	FindAll() (model.Profiles, error)
}
