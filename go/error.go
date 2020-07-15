package atm

import (
	"github.com/twitchtv/twirp"
)

func IsErrorNotFound(err error) bool {
	return IsErrorCode(err, twirp.NotFound)
}

func IsErrorUnauthenticated(err error) bool {
	return IsErrorCode(err, twirp.Unauthenticated)
}

func IsErrorCode(err error, code twirp.ErrorCode) bool {
	if e, ok := err.(twirp.Error); ok {
		return e.Code() == code
	}

	return false
}
