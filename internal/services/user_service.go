package services


import (
	"github.com/google/uuid"
	"homework1/internal/models"
	"homework1/internal/repository"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserById(id uuid.UUID) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	UpdateUser(id uuid.UUID,user model.User) (model.User, error)
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s *userService) GetUserById(id uuid.UUID) (model.User, error) {
	return s.repo.GetById(id)
}

func (s *userService) CreateUser(user model.User) (model.User, error) {
	user.ID = uuid.New()
	return s.repo.Create(user)
}

func (s *userService) UpdateUser(id uuid.UUID, user model.User) (model.User, error) {
	return s.repo.Update(id,user)
}

func (s *userService) DeleteUser(id uuid.UUID) error {
	return s.repo.Delete(id)
}
