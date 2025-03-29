package equipamentusecases


import "gym-system/src/inventory/Users/domain/repository"

type ApproveUserUseCase struct {
	Repo repository.IUserRepository
}

func NewApproveUserUseCase(repo repository.IUserRepository) *ApproveUserUseCase {
	return &ApproveUserUseCase{Repo: repo}
}

func (uc *ApproveUserUseCase) Execute(id int, macAddress string) error {
	return uc.Repo.ApproveUser(id, macAddress)
}
