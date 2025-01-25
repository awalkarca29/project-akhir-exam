package repository

import (
	"project-akhir-exam/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByID(ID int) (entity.User, error)
	FindByRoleID(RoleID int) (entity.User, error)
	// Update(user entity.User) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByID(ID int) (entity.User, error) {
	var user entity.User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByRoleID(RoleID int) (entity.User, error) {
	var user entity.User

	err := r.db.Where("role_id = ?", RoleID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// func (r *userRepository) Update(user entity.User) (entity.User, error) {
// 	err := r.db.Save(&user).Error
// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }
