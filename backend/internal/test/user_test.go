package repository

import (
    "context"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/Shiluco/UniTimetable/backend/ent"
    "github.com/Shiluco/UniTimetable/backend/internal/domain/model"
    _ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *ent.Client {
    client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
    if err != nil {
        t.Fatalf("failed opening connection to sqlite: %v", err)
    }
    
    if err := client.Schema.Create(context.Background()); err != nil {
        t.Fatalf("failed creating schema resources: %v", err)
    }
    
    return client
}

func TestUserRepository_CreateUser(t *testing.T) {
    client := setupTestDB(t)
    defer client.Close()

    repo := NewUserRepository(client)
    ctx := context.Background()

    testUser := &model.User{
        Name:     "Test User",
        Email:    "test@example.com",
        Password: "hashedpassword",
    }

    // ユーザー作成のテスト
    createdUser, err := repo.CreateUser(ctx, testUser)
    assert.NoError(t, err)
    assert.NotNil(t, createdUser)
    assert.Equal(t, testUser.Name, createdUser.Name)
    assert.Equal(t, testUser.Email, createdUser.Email)
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
    client := setupTestDB(t)
    defer client.Close()

    repo := NewUserRepository(client)
    ctx := context.Background()

    // テストデータの作成
    testUser := &model.User{
        Name:     "Test User",
        Email:    "test@example.com",
        Password: "hashedpassword",
    }
    createdUser, _ := repo.CreateUser(ctx, testUser)

    // メールアドレスによるユーザー取得のテスト
    foundUser, err := repo.GetUserByEmail(ctx, testUser.Email)
    assert.NoError(t, err)
    assert.NotNil(t, foundUser)
    assert.Equal(t, createdUser.Email, foundUser.Email)

    // 存在しないメールアドレスのテスト
    _, err = repo.GetUserByEmail(ctx, "nonexistent@example.com")
    assert.Equal(t, model.ErrUserNotFound, err)
}

func TestUserRepository_SearchUsersByName(t *testing.T) {
    client := setupTestDB(t)
    defer client.Close()

    repo := NewUserRepository(client)
    ctx := context.Background()

    // テストデータの作成
    testUsers := []*model.User{
        {Name: "Test User 1", Email: "test1@example.com", Password: "pass1"},
        {Name: "Test User 2", Email: "test2@example.com", Password: "pass2"},
        {Name: "Another User", Email: "another@example.com", Password: "pass3"},
    }

    for _, u := range testUsers {
        _, err := repo.CreateUser(ctx, u)
        assert.NoError(t, err)
    }

    // 検索テスト
    users, total, err := repo.SearchUsersByName(ctx, "Test", 10, 0)
    assert.NoError(t, err)
    assert.Equal(t, 2, total)
    assert.Len(t, users, 2)

    // ページネーションのテスト
    users, total, err = repo.SearchUsersByName(ctx, "Test", 1, 0)
    assert.NoError(t, err)
    assert.Equal(t, 2, total)
    assert.Len(t, users, 1)
}

func TestUserRepository_SearchUsersByEmail(t *testing.T) {
    client := setupTestDB(t)
    defer client.Close()

    repo := NewUserRepository(client)
    ctx := context.Background()

    // テストデータの作成
    testUsers := []*model.User{
        {Name: "User 1", Email: "test1@example.com", Password: "pass1"},
        {Name: "User 2", Email: "test2@example.com", Password: "pass2"},
        {Name: "User 3", Email: "another@example.com", Password: "pass3"},
    }

    for _, u := range testUsers {
        _, err := repo.CreateUser(ctx, u)
        assert.NoError(t, err)
    }

    // 検索テスト
    users, total, err := repo.SearchUsersByEmail(ctx, "test", 10, 0)
    assert.NoError(t, err)
    assert.Equal(t, 2, total)
    assert.Len(t, users, 2)

    // ページネーションのテスト
    users, total, err = repo.SearchUsersByEmail(ctx, "test", 1, 0)
    assert.NoError(t, err)
    assert.Equal(t, 2, total)
    assert.Len(t, users, 1)
}

func TestUserRepository_GetUserWithPassword(t *testing.T) {
    client := setupTestDB(t)
    defer client.Close()

    repo := NewUserRepository(client)
    ctx := context.Background()

    // テストデータの作成
    testUser := &model.User{
        Name:     "Test User",
        Email:    "test@example.com",
        Password: "hashedpassword",
    }
    _, err := repo.CreateUser(ctx, testUser)
    assert.NoError(t, err)

    // パスワードを含むユーザー情報の取得テスト
    user, err := repo.GetUserWithPassword(ctx, testUser.Email)
    assert.NoError(t, err)
    assert.NotNil(t, user)
    assert.Equal(t, testUser.Password, user.Password)

    // 通常のユーザー取得ではパスワードが含まれないことを確認
    userWithoutPass, err := repo.GetUserByEmail(ctx, testUser.Email)
    assert.NoError(t, err)
    assert.NotNil(t, userWithoutPass)
    assert.Empty(t, userWithoutPass.Password)
}
