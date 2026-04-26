package service

import (
	"context"
	"errors"

	"github.com/user/go-gin-gorm-starter/internal/domain"
	"github.com/user/go-gin-gorm-starter/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
)

type UserService interface {
	CreateUser(ctx context.Context, name, email, password string) (*domain.User, error)
	GetUserByID(ctx context.Context, id uint) (*domain.User, error)
	ListUsers(ctx context.Context, page, pageSize int) ([]domain.User, int64, error)
	UpdateUser(ctx context.Context, id uint, name, email string) (*domain.User, error)
	DeleteUser(ctx context.Context, id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, name, email, password string) (*domain.User, error) {
	// Check if email already exists
	existingUser, _ := s.repo.FindByEmail(ctx, email)
	if existingUser != nil {
		return nil, ErrEmailAlreadyExists
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (s *userService) ListUsers(ctx context.Context, page, pageSize int) ([]domain.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	return s.repo.FindAll(ctx, offset, pageSize)
}

func (s *userService) UpdateUser(ctx context.Context, id uint, name, email string) (*domain.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrUserNotFound
	}

	if email != "" && email != user.Email {
		existingUser, _ := s.repo.FindByEmail(ctx, email)
		if existingUser != nil {
			return nil, ErrEmailAlreadyExists
		}
		user.Email = email
	}

	if name != "" {
		user.Name = name
	}

	err = s.repo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return ErrUserNotFound
	}
	return s.repo.Delete(ctx, id)
}
