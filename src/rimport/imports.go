package rimport

import (
	"tgsms/internal/repository"
	"tgsms/internal/repository/mysql"
)

type Repository struct {
	IncomingMessages repository.IncomingMessages
}

func NewRepositoryImports() *Repository {
	return &Repository{
		IncomingMessages: mysql.NewIncomingMessages(),
	}
}
