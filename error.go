package atm

import (
	"github.com/twitchtv/twirp"
)

func IsErrorNotFound(err error) bool {
	if e, ok := err.(twirp.Error); ok {
		return e.Code() == twirp.NotFound
	}

	return false
}
