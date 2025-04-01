package repository

type IUserRepository interface {
	SaveRequest(username string, passwordHash string) error
	GetPendingRequests() ([]map[string]interface{}, error)
	ApproveUser(id int, macAddress string) error
	RejectUser(id int) error
	GetApprovedUsers() ([]map[string]interface{}, error) 
}

//