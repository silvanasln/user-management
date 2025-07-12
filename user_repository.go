package repository

import (
    "errors"
    "github.com/silvanasln/user-management/models"
)

type UserRepository interface {
    Create(user models.User) (models.User, error)
    GetByID(id int) (models.User, error)
    GetAll() ([]models.User, error)
    Update(id int, user models.User) (models.User, error)
    Delete(id int) error
}

type InMemoryUserRepository struct {
    users  []models.User
    nextID int
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
    return &InMemoryUserRepository{
        users:  make([]models.User, 0),
        nextID: 1,
    }
}

func (r *InMemoryUserRepository) Create(user models.User) (models.User, error) {
    for _, u := range r.users {
        if u.Email == user.Email {
            return models.User{}, errors.New("ایمیل قبلاً ثبت شده است")
        }
    }
    user.ID = r.nextID
    r.nextID++
    r.users = append(r.users, user)
    return user, nil
}

func (r *InMemoryUserRepository) GetByID(id int) (models.User, error) {
    for _, user := range r.users {
        if user.ID == id {
            return user, nil
        }
    }
    return models.User{}, errors.New("کاربر یافت نشد")
}

func (r *InMemoryUserRepository) GetAll() ([]models.User, error) {
    return r.users, nil
}

func (r *InMemoryUserRepository) Update(id int, updatedUser models.User) (models.User, error) {
    for i, user := range r.users {
        if user.ID == id {
            r.users[i].Name = updatedUser.Name
            r.users[i].Email = updatedUser.Email
            return r.users[i], nil
        }
    }
    return models.User{}, errors.New("کاربر یافت نشد")
}

func (r *InMemoryUserRepository) Delete(id int) error {
    for i, user := range r.users {
        if user.ID == id {
            r.users = append(r.users[:i], r.users[i+1:]...)
            return nil
        }
    }
    return errors.New("کاربر یافت نشد")
}
