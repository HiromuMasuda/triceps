package usecase

import "app/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (i *UserInteractor) Add(u domain.User) (err error) {
	_, err := i.UserRepository.Store(u)
	return
}

func (i *UserInteractor) Users() (user domain.Users, err error) {
	user, err = i.UserRepository.FindAll()
	return
}

func (i *UserInteractor) UserById(identifier int) (user domain.User, err error) {
	user, err = i.UserRepository.FindById(identifier)
	return
}
