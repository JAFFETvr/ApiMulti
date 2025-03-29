package entities

type user struct {
	id         int32
	username   string
	password   string
	macAddress string
	role       string
}

func NewUser(username string, password string, macAddress string, role string) *user {
	return &user{username: username, password: password, macAddress: macAddress, role: role}
}

type registrationRequest struct {
	id       int32
	username string
	password string
	status   string
}

func NewRegistrationRequest(username string, password string, status string) *registrationRequest {
	return &registrationRequest{username: username, password: password, status: status}
}
