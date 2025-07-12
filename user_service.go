package service

import (
"github.com/silvanasln/user-management/models"
   "github.com/silvanasln/user-management/repository"
)

type UserService interface {
    CreateUser(name, email string) (models.User, error)
    GetUserByID(id int) (models.User, error)
    GetAllUsers() ([]models.User, error)
    UpdateUser(id int, name, email string) (models.User, error)
    DeleteUser(id int) error
}

type userService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo: repo}
}

func (s *userService) CreateUser(name, email string) (models.User, error) {
    user := models.User{Name: name, Email: email}
    return s.repo.Create(user)
}

func (s *userService) GetUserByID(id int) (models.User, error) {
    return s.repo.GetByID(id) 
}

func (s *userService) GetAllUsers() ([]models.User, error) {
    return s.repo.GetAll() 
}

func (s *userService) UpdateUser(id int, name, email string) (models.User, error) {
    user := models.User{ID: id, Name: name, Email: email}
    return s.repo.Update(id, user) 
}

func (s *userService) DeleteUser(id int) error {
    return s.repo.Delete(id)
}
