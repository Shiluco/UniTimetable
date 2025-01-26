package handler

import (
	"net/http"
	"time"

	"github.com/Shiluco/UniTimetable/backend/ent"
	"github.com/Shiluco/UniTimetable/backend/ent/user"
	"github.com/Shiluco/UniTimetable/backend/internal/auth"
	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusCreated, gin.H{
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
	token, refreshToken, err := auth.GenerateTokens(u.ID, u.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Login successful",
		"data": gin.H{
			"accessToken":  token,
			"refreshToken": refreshToken, // リフレッシュトークンの生成ロジックが必要
			"token_type":   "Bearer",
			"expires_in":   3600, // トークンの有効期限設定
			"user": gin.H{
				"id":    u.ID,
				"name":  u.Name,
				"email": u.Email,
			},
		},
	})
}
