package interfaces

type UserResonseCreate struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Message  string `json:"message"`
}

type MessageUserDeleted struct {
	Message string `json:"message"`
}

type UserEdit struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}
