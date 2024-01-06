package helper

import "errors"

var (
	ErrDuplicate     = errors.New("record already exists")
	ErrNotExists     = errors.New("row not exists")
	ErrUpdateFailed  = errors.New("update failed")
	ErrDeleteFailed  = errors.New("delete failed")
	ErrEnvFailed     = errors.New("You must set up env as\n export storydb=$(tty)")
	ErrEnvHISTFailed = errors.New("You must set up env as\n export HISTFILE=$HOME/.bash_history \n or" +
		"\n export HISTFILE=$HOME/.zsh_history ")
	ErrAddFavorite = errors.New("That cmd is already added to favorite!..")
	OKAddFavorite  = "That cmd has been added to Favorite!!!"
)
