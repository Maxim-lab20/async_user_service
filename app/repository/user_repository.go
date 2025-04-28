package repostiory

import (
	"async_user_service/app/config"
	"async_user_service/app/model"
	"sync"
)

type UserRepository struct{}

var (
	UserRepositoryInstance *UserRepository
	once                   sync.Once
)

func NewUserRepository() *UserRepository {
	once.Do(func() {
		UserRepositoryInstance = &UserRepository{}
	})
	return UserRepositoryInstance
}

func (s *UserRepository) GetAllUsers() ([]model.User, error) {
	db := config.GetDBConnection()
	var users []model.User
	result := db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r *UserRepository) CreateUser(user model.User) (model.User, error) {
	db := config.GetDBConnection()

	result := db.Create(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}
