package user

import (
	"sample-tabungan2/internal/repository"

	"github.com/gocraft/work"
)

type UserService struct {
	repo      *repository.UserRepository
	publisher Publisher
}

type Publisher interface {
	Enqueue(jobName string, args map[string]interface{}) (*work.Job, error)
}

func NewUserService(repo *repository.UserRepository, publisher Publisher) *UserService {
	return &UserService{repo: repo, publisher: publisher}
}
