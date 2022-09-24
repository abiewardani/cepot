package manager

import (
	"sync"

	"gitlab.com/abiewardani/cepot/internal/system"
	"gitlab.com/abiewardani/cepot/internal/usecase"
)

type UsecaseManager interface {
	SystemUser() usecase.SystemUser
}

type usecaseManager struct {
	sys  system.System
	repo RepositoryManager
}

// NewUsecaseManager ...
func NewUsecaseManager(sys *system.System) UsecaseManager {
	var ucManager usecaseManager
	var once sync.Once

	once.Do(func() {
		ucManager = usecaseManager{
			sys:  *sys,
			repo: NewRepositoryManager(sys),
		}
	})

	return &ucManager
}

func (c *usecaseManager) SystemUser() usecase.SystemUser {
	var systemUserUc usecase.SystemUser
	var once sync.Once

	once.Do(func() {
		systemUserUc = usecase.NewSystemUserUc(c.sys, c.repo.SystemUser(), c.repo.ActiveStorage())
	})

	return systemUserUc
}
