package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"homework1/internal/models"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetById(id uuid.UUID) (model.User, error)
	Create(user model.User) (model.User, error)
	Update(id uuid.UUID,user model.User) (model.User, error)
	Delete(id uuid.UUID) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}



func (r *userRepository) GetAll() ([]model.User, error) {
	var users []model.User
	if err := r.db.Preload("Product").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetById(id uuid.UUID) (model.User, error) {
	var user model.User
	if err := r.db.Preload("Product").First(&user, "id = ?", id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Create(user model.User) (model.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Update(id uuid.UUID, updateUser model.User) (model.User, error) {

	var existingUser model.User
    if err := r.db.First(&existingUser, "id = ?", id).Error; err != nil {
        return existingUser, err
    }

	
	 existingUser.FirstName = updateUser.FirstName
	 existingUser.LastName = updateUser.LastName

	if err := r.db.Save(&existingUser).Error; err != nil {
		return existingUser, err
	}
	return existingUser, nil
}

func (r *userRepository) Delete(id uuid.UUID) error {
	if err := r.db.Delete(&model.User{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
