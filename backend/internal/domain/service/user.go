package service

import (
    "context"
    "fmt"

    "github.com/Shiluco/UniTimetable/backend/internal/domain/model"
    "github.com/Shiluco/UniTimetable/backend/internal/repository"
)

type UserService interface {
    CreateUser(ctx context.Context, req model.CreateUserRequest) (*model.UserResponse, error)
    GetUser(ctx context.Context, id int) (*model.UserResponse, error)
    ListUsers(ctx context.Context) ([]*model.UserResponse, error)
    UpdateUser(ctx context.Context, id int, req model.UpdateUserRequest) (*model.UserResponse, error)
    DeleteUser(ctx context.Context, id int) error
}

type userService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req model.CreateUserRequest) (*model.UserResponse, error) {
    // if err := s.repo.ResetSequence(ctx); err != nil {
    //     return nil, fmt.Errorf("failed to reset sequence: %w", err)
    // }

    user := &model.User{
        Name:  req.Name,
        Email: req.Email,
    }

    createdUser, err := s.repo.CreateUser(ctx, user)
    if err != nil {
        return nil, fmt.Errorf("failed to create user: %w", err)
    }

    return createdUser.ToResponse(), nil
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