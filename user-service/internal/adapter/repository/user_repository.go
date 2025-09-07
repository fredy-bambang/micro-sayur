package repository

import (
	"context"
	"errors"
	"time"
	"user-service/internal/core/domain/entity"
	"user-service/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
	CreateUserAccount(ctx context.Context, req entity.UserEntity) error
	UpdateUserVerified(ctx context.Context, userID int64) (*entity.UserEntity, error)
	UpdatePasswordByID(ctx context.Context, req entity.UserEntity) error
}

type userRepository struct {
	db *gorm.DB
}

// UpdatePasswordByID implement UserRepositoryInterface
func (u *userRepository) UpdatePasswordByID(ctx context.Context, req entity.UserEntity) error {
	modelUser := model.User{}

	if err := u.db.Where("id = ?", req.ID).First(&modelUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("404")
			log.Errorf("[UserRepository-9] UpdatePasswordByID: %v", err)
			return err
		}
		log.Errorf("[UserRepository-10] UpdatePasswordByID: %v", err)
		return err
	}
	modelUser.Password = req.Password
	if err := u.db.Save(&modelUser).Error; err != nil {
		log.Errorf("[UserRepository-11] UpdatePasswordByID: %v", err)
		return err
	}

	return nil
}

func (u *userRepository) UpdateUserVerified(ctx context.Context, userID int64) (*entity.UserEntity, error) {
	modelUser := model.User{}

	if err := u.db.Where("id = ?", userID).Preload("Roles").First(&modelUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("404")
			log.Errorf("[UserRepository-6] UpdateUserVerified: %v", err)
			return nil, err
		}
		log.Errorf("[UserRepository-7] UpdateUserVerified: %v", err)
		return nil, err
	}

	modelUser.IsVerified = true
	if err := u.db.Save(&modelUser).Error; err != nil {
		log.Errorf("[UserRepository-8] UpdateUserVerified: %v", err)
		return nil, err
	}

	return &entity.UserEntity{
		ID:         userID,
		Name:       modelUser.Name,
		Email:      modelUser.Email,
		RoleName:   modelUser.Roles[0].Name,
		Address:    modelUser.Address,
		Lat:        modelUser.Lat,
		Lng:        modelUser.Lng,
		Phone:      modelUser.Phone,
		Photo:      modelUser.Photo,
		IsVerified: modelUser.IsVerified,
	}, nil
}

// GetUserByEmail implement UserRepositoryInterface
func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	modelUser := model.User{}

	// Preload Roles itu bisa cek di model.User.Roles
	if err := u.db.Where("email = ? AND is_verified = ?", email, true).
		Preload("Roles").First(&modelUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("404")
			log.Infof("[UserRepository-1] GetUserByEmail: User not found")
			return nil, err
		}
		log.Errorf("[UserRepository-2] GetUserByEmail: %v", err)
		return nil, err
	}

	return &entity.UserEntity{
		ID:         modelUser.ID,
		Name:       modelUser.Name,
		Email:      modelUser.Email,
		Password:   modelUser.Password,
		RoleName:   modelUser.Roles[0].Name,
		Address:    modelUser.Address,
		Lat:        modelUser.Lat,
		Lng:        modelUser.Lng,
		Phone:      modelUser.Phone,
		Photo:      modelUser.Photo,
		IsVerified: modelUser.IsVerified,
	}, nil
}

func (u *userRepository) CreateUserAccount(ctx context.Context, req entity.UserEntity) error {
	modelRole := model.Role{}
	err := u.db.Where("name = ?", "Customer").First(&modelRole).Error
	if err != nil {
		log.Errorf("[UserRepository-5] CreateUserAccount: %v", err)
		return err
	}

	modelUser := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Roles:    []model.Role{modelRole},
	}

	if err := u.db.Create(&modelUser).Error; err != nil {
		log.Errorf("[UserRepository-3] CreateUserAccount: %v", err)
		return err
	}

	modelVerify := model.VerificationToken{
		UserID:    modelUser.ID,
		Token:     req.Token,
		TokenType: "email_verification",
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}

	if err := u.db.Create(&modelVerify).Error; err != nil {
		log.Errorf("[UserRepository-4] CreateUserAccount: %v", err)
		return err
	}

	return nil
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}
