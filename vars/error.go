package vars

import "fmt"

type ErrorResp struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func (e *ErrorResp) Error() string {
	return fmt.Sprintf("(message=%s) (type: %s)", e.Message, e.Type)
}
