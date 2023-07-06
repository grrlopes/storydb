package repositories

import "os"

type IFileParsedRepository interface {
	All() *os.File
}
