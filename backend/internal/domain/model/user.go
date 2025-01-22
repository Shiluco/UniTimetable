package model

import (
    "errors"
    "time"
)

// User ドメインモデル
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserRequest ユーザー作成リクエスト
type CreateUserRequest struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
}

// UpdateUserRequest ユーザー更新リクエスト
type UpdateUserRequest struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
}

// UserResponse ユーザーレスポンス
type UserResponse struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// ToResponse ドメインモデルからレスポンスモデルへの変換
func (u *User) ToResponse() *UserResponse {
    return &UserResponse{
        ID:        u.ID,
        Name:      u.Name,
        Email:     u.Email,
        CreatedAt: u.CreatedAt,
        UpdatedAt: u.UpdatedAt,
    }
}

var (
    ErrUserNotFound = errors.New("user not found")
)
