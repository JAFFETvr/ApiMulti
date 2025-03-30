package entities

type user struct {
	id         int32
	username   string
	password_hash   string
	macAddress string
	role       string
}

func NewUser(username string, password_hash string, macAddress string, role string) *user {
	return &user{username: username, password_hash: password_hash, macAddress: macAddress, role: role}
}

type registrationRequest struct {
	id       int32
	username string
	password_hash string
	status   string
}

func NewRegistrationRequest(username string,password_hash string, status string) *registrationRequest {
	return &registrationRequest{username: username, password_hash: password_hash, status: status}
}
