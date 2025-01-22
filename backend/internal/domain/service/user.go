package service

import (
    "context"
    "fmt"

    "github.com/Shiluco/UniTimetable/backend/internal/domain/model"
    "github.com/Shiluco/UniTimetable/backend/internal/repository"
    "github.com/Shiluco/UniTimetable/backend/internal/util"
)

type UserService interface {
    GetUser(ctx context.Context, id int) (*model.UserResponse, error)
    ListUsers(ctx context.Context) ([]*model.UserResponse, error)
    UpdateUser(ctx context.Context, id int, req model.UpdateUserRequest) (*model.UserResponse, error)
    DeleteUser(ctx context.Context, id int) error
    GetUserByEmail(ctx context.Context, email string) (*model.UserResponse, error)
    GetUserByName(ctx context.Context, name string) (*model.UserResponse, error)
    SearchUsersByName(ctx context.Context, query string, page, pageSize int) (*model.SearchResponse, error)
    SearchUsersByEmail(ctx context.Context, query string, page, pageSize int) (*model.SearchResponse, error)
}

type userService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo: repo}
}

func (s *userService) GetUser(ctx context.Context, id int) (*model.UserResponse, error) {
    user, err := s.repo.GetUserByID(ctx, id)
    if err != nil {
        return nil, fmt.Errorf("failed to get user: %w", err)
    }
    return user.ToResponse(), nil
}

func (s *userService) ListUsers(ctx context.Context) ([]*model.UserResponse, error) {
    users, err := s.repo.GetAllUsers(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to list users: %w", err)
    }

    responses := make([]*model.UserResponse, len(users))
    for i, user := range users {
        responses[i] = user.ToResponse()
    }
    return responses, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int, req model.UpdateUserRequest) (*model.UserResponse, error) {
    user := &model.User{
        ID:    id,
        Name:  req.Name,
        Email: req.Email,
    }

    if req.Password != "" {
        hashedPassword, err := util.HashPassword(req.Password)
        if err != nil {
            return nil, fmt.Errorf("failed to hash password: %w", err)
        }
        user.Password = hashedPassword
    }

    updatedUser, err := s.repo.UpdateUser(ctx, user)
    if err != nil {
        return nil, fmt.Errorf("failed to update user: %w", err)
    }
    return updatedUser.ToResponse(), nil
}

func (s *userService) DeleteUser(ctx context.Context, id int) error {
    if err := s.repo.DeleteUser(ctx, id); err != nil {
        return fmt.Errorf("failed to delete user: %w", err)
    }
    return nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*model.UserResponse, error) {
    user, err := s.repo.GetUserByEmail(ctx, email)
    if err != nil {
        return nil, fmt.Errorf("failed to get user by email: %w", err)
    }
    return user.ToResponse(), nil
}

func (s *userService) GetUserByName(ctx context.Context, name string) (*model.UserResponse, error) {
    user, err := s.repo.GetUserByName(ctx, name)
    if err != nil {
        return nil, fmt.Errorf("failed to get user by name: %w", err)
    }
    return user.ToResponse(), nil
}

func (s *userService) SearchUsersByName(ctx context.Context, query string, page, pageSize int) (*model.SearchResponse, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 {
        pageSize = 10
    }

    offset := (page - 1) * pageSize
    users, total, err := s.repo.SearchUsersByName(ctx, query, pageSize, offset)
    if err != nil {
        return nil, fmt.Errorf("failed to search users by name: %w", err)
    }

    responses := make([]*model.UserResponse, len(users))
    for i, user := range users {
        responses[i] = user.ToResponse()
    }

    return &model.SearchResponse{
        Users:      responses,
        Total:      total,
        Page:       page,
        PageSize:   pageSize,
        TotalPages: (total + pageSize - 1) / pageSize,
    }, nil
}

func (s *userService) SearchUsersByEmail(ctx context.Context, query string, page, pageSize int) (*model.SearchResponse, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 {
        pageSize = 10
    }

    offset := (page - 1) * pageSize
    users, total, err := s.repo.SearchUsersByEmail(ctx, query, pageSize, offset)
    if err != nil {
        return nil, fmt.Errorf("failed to search users by email: %w", err)
    }

    responses := make([]*model.UserResponse, len(users))
    for i, user := range users {
        responses[i] = user.ToResponse()
    }

    return &model.SearchResponse{
        Users:      responses,
        Total:      total,
        Page:       page,
        PageSize:   pageSize,
        TotalPages: (total + pageSize - 1) / pageSize,
    }, nil
}