package shapeshift

type ErrorMsg interface {
	ErrorMsg() string
	isOk() bool
}

func (e Error) ErrorMsg() string {
	return e.Message
}

func (e Error) isOk() bool {
	if e.Message == "" {
		return true
	}
	return false
}

type Error struct {
	Message string `json:"error,omitempty"`
}
