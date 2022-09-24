package usecase

import (
	"context"

	"gitlab.com/abiewardani/cepot/internal/domain"
	"gitlab.com/abiewardani/cepot/internal/repository"
	"gitlab.com/abiewardani/cepot/internal/system"
	serializer "gitlab.com/abiewardani/cepot/internal/usecase/serializers"
	"gitlab.com/abiewardani/cepot/pkg/shared"
)

type SystemUser interface {
	Show(ctx context.Context, id string) (sysUser *serializer.SystemUserShortAttributesSerializer, err error)
}

// systemUserCtx ..
type systemUserCtx struct {
	sys                  system.System
	systemUserRepository repository.SystemUserRepository
	activeStorageRepo    repository.ActiveStorageRepository
}

func NewSystemUserUc(sys system.System, systemUserRepo repository.SystemUserRepository, activeStorageRepository repository.ActiveStorageRepository) SystemUser {
	return &systemUserCtx{
		sys:                  sys,
		systemUserRepository: systemUserRepo,
		activeStorageRepo:    activeStorageRepository,
	}
}

func (c *systemUserCtx) Show(ctx context.Context, id string) (sysUser *serializer.SystemUserShortAttributesSerializer, err error) {
	systemUser, err := c.systemUserRepository.FindOne(ctx, &domain.SystemUserParams{
		ID: id,
	})

	if err != nil {
		return nil, shared.NewMultiStringBadRequestError(shared.HTTPErrorBadRequest, "An error occurred while getting system user")
	}

	if systemUser == nil {
		return nil, shared.NewMultiStringBadRequestError(shared.HTTPErrorDataNotFound, "Data not found")
	}

	return serializer.NewSystemUserShortAttributesSerializer(systemUser), nil
}
