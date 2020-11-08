package infrastructure

import (
	"io"
	"os"
	"os/user"
)

type User struct {
	UserName string
	Email    string
}

type Repository interface {
	GetCurrentBranch() (string, error)
	GetUser() (User, error)
	GetLastTag() (string, error)
}

type FileSystem interface {
	Write(path, content string) error
	Read(path string) (string, error)
	Reader(path string) (io.Reader, error)
	Exist(path string) bool
	Delete(path string) error
	Chmod(path string, mode os.FileMode) error
	Chown(name string, user *user.User) error
}

type Shell interface {
	Exec(commands []string, env *map[string]string, output bool) (string, int, error)
}
