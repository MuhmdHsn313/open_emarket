package jwt

type NotValidToken struct {
	Message string
}

func (receiver NotValidToken) Error() string {
	return receiver.Message
}
