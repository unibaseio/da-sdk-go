package error

import "github.com/unibaseio/da-sdk-go/lib/log"

var logger = log.Logger("lerror")

type APIError struct {
	Type    string
	Message string
}

func ToAPIError(typ string, err error) APIError {
	logger.Debugf("%s %s", typ, err)
	return APIError{
		Type:    typ,
		Message: err.Error(),
	}
}

func (e APIError) Error() string {
	return e.Type + ":" + e.Message
}
