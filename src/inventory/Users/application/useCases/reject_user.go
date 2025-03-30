package equipamentusecases

import "gym-system/src/inventory/Users/domain/repository"

type RejectUserUseCase struct {
	Repo repository.IUserRepository
}

func NewRejectUserUseCase(repo repository.IUserRepository) *RejectUserUseCase {
	return &RejectUserUseCase{Repo: repo}
}

func (uc *RejectUserUseCase) Execute(id int) error {
	return uc.Repo.RejectUser(id)
}