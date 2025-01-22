package repository

import (
	"context"
	//"fmt"

	"github.com/Shiluco/UniTimetable/backend/ent"
	"github.com/Shiluco/UniTimetable/backend/ent/user"
	"github.com/Shiluco/UniTimetable/backend/internal/domain/model"
)

// UserRepository データベース操作のインターフェース
type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id int) error
	//ResetSequence(ctx context.Context) error
}

// userRepository リポジトリの実装
type userRepository struct {
	client *ent.Client
}

// NewUserRepository リポジトリのインスタンスを作成
func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepository{
		client: client,
	}
}

// entのUserをドメインモデルに変換
func toModel(u *ent.User) *model.User {
	return &model.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// CreateUser 新規ユーザーを作成
func (r *userRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	u, err := r.client.User.
		Create().
		SetName(user.Name).
		SetEmail(user.Email).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return toModel(u), nil
}

// GetUserByID IDによるユーザー取得
func (r *userRepository) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	u, err := r.client.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.ErrUserNotFound
		}
		return nil, err
	}
	return toModel(u), nil
}

// GetAllUsers 全ユーザーを取得
func (r *userRepository) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	users, err := r.client.User.
		Query().
		Order(ent.Asc(user.FieldID)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	
	result := make([]*model.User, len(users))
	for i, u := range users {
		result[i] = toModel(u)
	}
	return result, nil
}

// UpdateUser ユーザー情報を更新
func (r *userRepository) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	u, err := r.client.User.
		UpdateOneID(user.ID).
		SetName(user.Name).
		SetEmail(user.Email).
		Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.ErrUserNotFound
		}
		return nil, err
	}
	return toModel(u), nil
}

// DeleteUser ユーザーを削除
func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	err := r.client.User.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return model.ErrUserNotFound
		}
		return err
	}
	return nil
}

// ResetSequence IDシーケンスをリセット
// func (r *userRepository) ResetSequence(ctx context.Context) error {
// 	// テーブルが空の場合のみシーケンスをリセット
// 	count, err := r.client.User.Query().Count(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	if count == 0 {
// 		// QueryContextを使用してSQLを実行
// 		_, err := r.client.QueryContext(ctx, "ALTER SEQUENCE users_id_seq RESTART WITH 1")
// 		if err != nil {
// 			return fmt.Errorf("failed to reset sequence: %w", err)
// 		}
// 	}
// 	return nil
// } 