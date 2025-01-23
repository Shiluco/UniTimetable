package handler

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/Shiluco/UniTimetable/backend/internal/domain/model"
    "github.com/Shiluco/UniTimetable/backend/internal/domain/service"
)

type UserHandler struct {
    service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
    return &UserHandler{
        service: service,
    }
}

// GetUsers ユーザー一覧を取得
func (h *UserHandler) GetUsers(c *gin.Context) {
    users, err := h.service.ListUsers(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

// GetUser 特定のユーザーを取得
func (h *UserHandler) GetUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    user, err := h.service.GetUser(c.Request.Context(), id)
    if err != nil {
        if err == model.ErrUserNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

// CreateUser 新規ユーザーを作成
// func (h *UserHandler) CreateUser(c *gin.Context) {
//     var req model.CreateUserRequest
//     if err := c.ShouldBindJSON(&req); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     user, err := h.service.CreateUser(c.Request.Context(), req)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     c.JSON(http.StatusCreated, user)
// }

// UpdateUser ユーザー情報を更新
func (h *UserHandler) UpdateUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    var req model.UpdateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.service.UpdateUser(c.Request.Context(), id, req)
    if err != nil {
        if err == model.ErrUserNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

// DeleteUser ユーザーを削除
func (h *UserHandler) DeleteUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    err = h.service.DeleteUser(c.Request.Context(), id)
    if err != nil {
        if err == model.ErrUserNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// GetUserByName ユーザー名による検索
func (h *UserHandler) GetUserByName(c *gin.Context) {
    name := c.Query("name")
    if name == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
        return
    }

    user, err := h.service.GetUserByName(c.Request.Context(), name)
    if err != nil {
        if err == model.ErrUserNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

// SearchUsersByName ユーザー名による部分一致検索
func (h *UserHandler) SearchUsersByName(c *gin.Context) {
    query := c.Query("q")
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

    result, err := h.service.SearchUsersByName(c.Request.Context(), query, page, pageSize)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, result)
}

// SearchUsersByEmail メールアドレスによる部分一致検索
func (h *UserHandler) SearchUsersByEmail(c *gin.Context) {
    query := c.Query("q")
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

    result, err := h.service.SearchUsersByEmail(c.Request.Context(), query, page, pageSize)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, result)
}

// GetUserByEmail メールアドレスでユーザーを検索
func (h *UserHandler) GetUserByEmail(c *gin.Context) {
    email := c.Query("email")
    if email == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
        return
    }

    user, err := h.service.GetUserByEmail(c.Request.Context(), email)
    if err != nil {
        if err == model.ErrUserNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

// SearchUsers 統合された検索ハンドラー
func (h *UserHandler) SearchUsers(c *gin.Context) {
    searchType := c.Query("type")
    query := c.Query("q")
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

    var result *model.SearchResponse
    var err error

    switch searchType {
    case "email":
        result, err = h.service.SearchUsersByEmail(c.Request.Context(), query, page, pageSize)
    case "name":
        result, err = h.service.SearchUsersByName(c.Request.Context(), query, page, pageSize)
    default:
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid search type"})
        return
    }

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, result)
}
