package handler

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/Shiluco/UniTimetable/backend/ent"
    "github.com/Shiluco/UniTimetable/backend/ent/user"
    "github.com/Shiluco/UniTimetable/backend/internal/auth"
)

type AuthHandler struct {
    client *ent.Client
}

func NewAuthHandler(client *ent.Client) *AuthHandler {
    return &AuthHandler{client: client}
}

// Register ユーザー登録
func (h *AuthHandler) Register(c *gin.Context) {
    var req struct {
        Name     string `json:"name" binding:"required"`
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=8"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // メールアドレスの重複チェック
    exists, err := h.client.User.Query().
        Where(user.Email(req.Email)).
        Exist(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if exists {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
        return
    }

    // パスワードのハッシュ化
    hashedPassword, err := auth.HashPassword(req.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    // ユーザーの作成
    u, err := h.client.User.Create().
        SetName(req.Name).
        SetEmail(req.Email).
        SetPassword(hashedPassword).
        SetCreatedAt(time.Now()).
        SetUpdatedAt(time.Now()).
        Save(c.Request.Context())

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // JWTトークンの生成
    token, err := auth.GenerateJWT(u.ID, u.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "token": token,
        "user": gin.H{
            "id":    u.ID,
            "name":  u.Name,
            "email": u.Email,
        },
    })
}

// Login ログイン
func (h *AuthHandler) Login(c *gin.Context) {
    var req struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // ユーザーの検索
    u, err := h.client.User.Query().
        Where(user.Email(req.Email)).
        Only(c.Request.Context())
    if err != nil {
        if ent.IsNotFound(err) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // パスワードの検証
    if err := auth.CheckPassword(req.Password, u.Password); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    // JWTトークンの生成
    token, err := auth.GenerateJWT(u.ID, u.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "token": token,
        "user": gin.H{
            "id":    u.ID,
            "name":  u.Name,
            "email": u.Email,
        },
    })
}
