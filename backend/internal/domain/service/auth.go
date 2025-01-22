package service

import (
	"context"
	"fmt"

	"github.com/Shiluco/UniTimetable/backend/internal/domain/model"
	"github.com/Shiluco/UniTimetable/backend/internal/repository"
	"github.com/Shiluco/UniTimetable/backend/internal/util"
)

type AuthService interface {
	Login(ctx context.Context, req model.LoginRequest) (string, *model.UserResponse, error)
	Register(ctx context.Context, req model.RegisterRequest) (*model.UserResponse, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

// Login ユーザー認証とトークン生成
func (s *authService) Login(ctx context.Context, req model.LoginRequest) (string, *model.UserResponse, error) {
	// メールアドレスでユーザーを検索
	user, err := s.userRepo.GetUserWithPassword(ctx, req.Email)
	if err != nil {
		return "", nil, fmt.Errorf("authentication failed")
	}

	// パスワードを検証
	if err := util.CheckPassword(req.Password, user.Password); err != nil {
		return "", nil, fmt.Errorf("authentication failed")
	}

	// JWTトークンを生成
	token, err := util.GenerateJWT(user.Email)
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate token")
	}

	return token, user.ToResponse(), nil
}

// Register 新規ユーザー登録
func (s *authService) Register(ctx context.Context, req model.RegisterRequest) (*model.UserResponse, error) {
	// パスワードをハッシュ化
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	// ユーザーを作成
	createdUser, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return createdUser.ToResponse(), nil
}
