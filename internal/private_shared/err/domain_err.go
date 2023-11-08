package err

import "errors"

var (
	Unauthorized = errors.New("unauthorized")
	NotFound     = errors.New("not_found")
)
