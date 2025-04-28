package service

import (
	"async_user_service/app/dto"
	"async_user_service/app/model"
	repostiory "async_user_service/app/repository"
	"async_user_service/app/util"
	"sync"
	"time"
)

type UserService struct {
	userRepository repostiory.UserRepository
	cacheService   CacheService
}

var (
	userServiceInstance *UserService
	once                sync.Once
)

func NewUserService() *UserService {
	once.Do(func() {
		userServiceInstance = &UserService{
			userRepository: *repostiory.NewUserRepository(),
			cacheService:   *NewCacheService(),
		}
	})
	return userServiceInstance
}

func (us *UserService) GetAllUsers() ([]dto.UserDTO, error) {
	const cacheKey = "users:all"

	var userDTOs []dto.UserDTO
	err := us.cacheService.Get(cacheKey, &userDTOs)
	if err == nil {
		// Cache hit
		return userDTOs, nil
	}

	// Cache miss: fetch from DB
	userDTOs, err = us.getUsersFromDb()
	if err != nil {
		return nil, err
	}

	// Save into Redis with 5 minutes TTL
	_ = us.cacheService.Set(cacheKey, userDTOs, 5*time.Minute)

	return userDTOs, nil
}

func (us *UserService) CreateUser(userDTO dto.UserDTO) (dto.UserDTO, error) {
	user := model.User{
		FirstName: userDTO.FirstName,
		LastName:  userDTO.LastName,
	}

	createdUser, err := us.userRepository.CreateUser(user)
	if err != nil {
		return dto.UserDTO{}, err
	}

	return dto.UserDTO{
		ID:        createdUser.ID,
		FirstName: createdUser.FirstName,
		LastName:  createdUser.LastName,
	}, nil
}

func (s *UserService) getUsersFromDb() ([]dto.UserDTO, error) {
	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	userDTOs := util.Map(users, func(u model.User) dto.UserDTO {
		return dto.UserDTO{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
		}
	})
	return userDTOs, nil
}
