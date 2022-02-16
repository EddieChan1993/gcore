package gmongo

import "errors"

var (
	ErrRepeatInit       = errors.New("repeat init")
	ErrNilConfig        = errors.New("config is nil")
	ErrNotInit          = errors.New("client is no init")
	ErrNilOption        = errors.New("option is nil")
	ErrNameSpaceInvalid = errors.New("name_space is invalid")
)
