package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Shiluco/UniTimetable/backend/internal/domain/model"
    "github.com/Shiluco/UniTimetable/backend/internal/domain/service"
)

type AuthHandler struct {
    authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
    return &AuthHandler{
        authService: authService,
    }
}

// Login ログインハンドラー
func (h *AuthHandler) Login(c *gin.Context) {
    var req model.LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // ユーザー認証とトークン生成
    token, user, err := h.authService.Login(c.Request.Context(), req)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "token": token,
        "user":  user,
    })
}

// Register 新規ユーザー登録ハンドラー
func (h *AuthHandler) Register(c *gin.Context) {
    var req model.RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.authService.Register(c.Request.Context(), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, user)
}
