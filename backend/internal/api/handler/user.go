package handler

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/Shiluco/UniTimetable/backend/ent"
    "github.com/Shiluco/UniTimetable/backend/ent/user"
)

type UserHandler struct {
    client *ent.Client
}

func NewUserHandler(client *ent.Client) *UserHandler {
    return &UserHandler{client: client}
}

// GetUsers ユーザー取得ハンドラー（一覧・個別・検索に対応）
func (h *UserHandler) GetUsers(c *gin.Context) {
    ctx := c.Request.Context()
    
    // IDパラメータの取得（個別取得用）
    userID := c.Param("id")

    // 個別取得の場合
    if userID != "" {
        id, err := strconv.Atoi(userID)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
            return
        }

        user, err := h.client.User.Query().
            Where(user.ID(id)).
            WithDepartment().
            WithMajor().
            WithPosts().
            WithSchedules().
            Only(ctx)

        if err != nil {
            if ent.IsNotFound(err) {
                c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
                return
            }
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "status":  "success",
            "message": "Users fetched successfully",
            "data": gin.H{
                "user": user,
            }
        })
        return
    }

    // クエリパラメータの取得
    searchType := c.Query("type")    // 検索タイプ（email or name）
    query := c.Query("q")           // 検索クエリ
    total := c.Query("total") == "true" // 全件取得フラグ

    userQuery := h.client.User.Query()

    // 検索条件の適用
    if searchType != "" && query != "" {
        switch searchType {
        case "email":
            userQuery.Where(user.EmailContains(query))
        case "name":
            userQuery.Where(user.NameContains(query))
        default:
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid search type"})
            return
        }
    }

    // 全件取得の場合
    if total {
        users, err := userQuery.
            Order(ent.Asc(user.FieldName)).
            All(ctx)

        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "status":  "success",
            "message": "Users fetched successfully",
            "data": gin.H{
                "users": users,
            }
        })
        return
    }

    // ページネーション付き取得の場合
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

    if page < 1 {
        page = 1
    }
    if pageSize < 1 {
        pageSize = 10
    }

    offset := (page - 1) * pageSize

    // 総件数の取得
    total_count, err := userQuery.Count(ctx)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // ユーザー一覧の取得
    users, err := userQuery.
        Limit(pageSize).
        Offset(offset).
        Order(ent.Asc(user.FieldName)).
        All(ctx)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Users fetched successfully",
		"data": gin.H{
            "users": users,
            "total": total_count,
            "page": page,
            "page_size": pageSize,
            "total_pages": (total_count + pageSize - 1) / pageSize,
        }
    })
}

// UpdateUser ユーザー情報を更新
func (h *UserHandler) UpdateUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    var req struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.client.User.UpdateOneID(id).
        SetName(req.Name).
        SetEmail(req.Email).
        Save(c.Request.Context())

    if err != nil {
        if ent.IsNotFound(err) {
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

    err = h.client.User.DeleteOneID(id).Exec(c.Request.Context())
    if err != nil {
        if ent.IsNotFound(err) {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}


