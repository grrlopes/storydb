package helper

import "errors"

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
	ErrEnvFailed    = errors.New("You must set up env as\n export storydb=$(tty)")
)
