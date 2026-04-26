package repository

import (
	"context"
	"github.com/user/go-gin-gorm-starter/internal/domain"
	"gorm.io/gorm"
)

// UserRepository defines the interface for user database operations
type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByID(ctx context.Context, id uint) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindAll(ctx context.Context, offset, limit int) ([]domain.User, int64, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id uint) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindAll(ctx context.Context, offset, limit int) ([]domain.User, int64, error) {
	var users []domain.User
	var total int64

	db := r.db.WithContext(ctx).Model(&domain.User{})

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.User{}, id).Error
}
