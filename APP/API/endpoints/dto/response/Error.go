package response

type Error struct {
	StatusCode int    `json:"Status_code"`
	Msg        string `json:"msg"`
	Erro       string `json:"erro"`
}

func NewError(statusCode int, msg, erro string) *Error {
	return &Error{
		StatusCode: statusCode,
		Msg:        msg,
		Erro:       erro,
	}
}
