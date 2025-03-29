package repository

type IUserRepository interface {
	SaveRequest(username string, password string) error
	GetPendingRequests() ([]map[string]interface{}, error)
	ApproveUser(id int, macAddress string) error
	RejectUser(id int) error
}
