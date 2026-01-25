package mock_repositories

import (
	"gold-dashboard-be/internal/mocks"
	"gold-dashboard-be/internal/repositories"

	"github.com/golang/mock/gomock"
)

var (
	UserRepository    *mocks.MockUserRepository
)

func NewMockRepository(ctrl *gomock.Controller) *repositories.Repository {
	UserRepository = mocks.NewMockUserRepository(ctrl)

	return &repositories.Repository{
		User: UserRepository,
	}
}