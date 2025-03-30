package equipamentusecases

import (
	"gym-system/src/inventory/Users/domain/repository"
)

type GetApprovedUsersUseCase struct {
	Repo repository.IUserRepository
}

func NewGetApprovedUsersUseCase(repo repository.IUserRepository) *GetApprovedUsersUseCase {
	return &GetApprovedUsersUseCase{Repo: repo}
}

func (uc *GetApprovedUsersUseCase) Execute() ([]map[string]interface{}, error) {
	return uc.Repo.GetApprovedUsers()
}