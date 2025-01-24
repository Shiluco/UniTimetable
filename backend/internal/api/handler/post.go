package handler

import (
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/Shiluco/UniTimetable/backend/ent"
    "github.com/Shiluco/UniTimetable/backend/ent/post"
    "github.com/Shiluco/UniTimetable/backend/internal/api/middleware"
)

type PostHandler struct {
    client *ent.Client
}

func NewPostHandler(client *ent.Client) *PostHandler {
    return &PostHandler{client: client}
}

// GetPosts 投稿取得ハンドラー（単一投稿または一覧）
func (h *PostHandler) GetPosts(c *gin.Context) {
    ctx := c.Request.Context()
    
    // IDパラメータの取得（単一投稿の場合）
    if id := c.Param("id"); id != "" {
        postID, err := strconv.Atoi(id)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
            return
        }

        // 単一の投稿を取得
        post, err := h.client.Post.Query().
            Where(post.ID(postID)).
            WithUser().           // 投稿者
            WithParent().        // 親投稿
            WithSchedule().      // 関連する時間割
            WithReplies().       // 返信
            Only(ctx)

        if err != nil {
            if ent.IsNotFound(err) {
                c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
                return
            }
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, post)
        return
    }

    // クエリパラメータの取得（一覧の場合）
    scheduleID, _ := strconv.Atoi(c.Query("schedule_id"))
    userID, _ := strconv.Atoi(c.Query("user_id"))
    parentID, _ := strconv.Atoi(c.Query("parent_id"))
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

    if page < 1 {
        page = 1
    }
    if pageSize < 1 {
        pageSize = 10
    }

    offset := (page - 1) * pageSize
    postQuery := h.client.Post.Query()

    // フィルター条件の適用
    if scheduleID > 0 {
        postQuery.Where(post.ScheduleIDEQ(scheduleID))
    }
    if userID > 0 {
        postQuery.Where(post.UserIDEQ(userID))
    }
    if parentID > 0 {
        // 特定の投稿への返信を取得
        postQuery.Where(post.ParentPostIDEQ(parentID))
    } else if parentID == 0 {
        // 親投稿のみを取得（返信を除外）
        postQuery.Where(post.ParentPostIDIsNil())
    }

    // 総件数の取得
    total, err := postQuery.Count(ctx)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // データが存在しない場合は空配列を返す
    if total == 0 {
        c.JSON(http.StatusOK, gin.H{
            "posts":       []interface{}{},
            "total":      0,
            "page":       page,
            "page_size":  pageSize,
            "total_pages": 0,
        })
        return
    }

    // ページネーションとソートの適用
    posts, err := postQuery.
        Limit(pageSize).
        Offset(offset).
        Order(ent.Desc(post.FieldCreatedAt)).
        WithUser().
        WithSchedule().
        WithParent().
        WithReplies().
        All(ctx)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "posts":       posts,
        "total":      total,
        "page":       page,
        "page_size":  pageSize,
        "total_pages": (total + pageSize - 1) / pageSize,
    })
}

// CreatePost 投稿を作成（通常の投稿または返信）
func (h *PostHandler) CreatePost(c *gin.Context) {
    var req struct {
        ParentPostID *int   `json:"parent_post_id"`
        Content      string `json:"content" binding:"required"`
        ScheduleID   *int   `json:"schedule_id"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 現在のユーザーを取得
    currentUser := middleware.GetCurrentUser(c)
    if currentUser == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
        return
    }

    // 投稿の作成
    create := h.client.Post.Create().
        SetContent(req.Content).
        SetUserID(currentUser.ID).
        SetCreatedAt(time.Now()).
        SetUpdatedAt(time.Now())

    // オプションフィールドの設定
    if req.ParentPostID != nil {
        create.SetParentPostID(*req.ParentPostID)
    }
    if req.ScheduleID != nil {
        create.SetScheduleID(*req.ScheduleID)
    }

    post, err := create.Save(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, post)
}

// UpdatePost 投稿を更新
func (h *PostHandler) UpdatePost(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }

    var req struct {
        Content string `json:"content" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 現在のユーザーを取得
    currentUser := middleware.GetCurrentUser(c)
    if currentUser == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
        return
    }

    // 投稿の所有者を確認
    p, err := h.client.Post.Query().
        Where(post.ID(id)).
        Only(c.Request.Context())
    if err != nil {
        if ent.IsNotFound(err) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 所有者でない場合は更新を許可しない
    if p.UserID != currentUser.ID {
        c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to update this post"})
        return
    }

    post, err := h.client.Post.UpdateOneID(id).
        SetContent(req.Content).
        SetUpdatedAt(time.Now()).
        Save(c.Request.Context())

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, post)
}

// DeletePost 投稿を削除
func (h *PostHandler) DeletePost(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }

    // 現在のユーザーを取得
    currentUser := middleware.GetCurrentUser(c)
    if currentUser == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
        return
    }

    // 投稿の所有者を確認
    p, err := h.client.Post.Query().
        Where(post.ID(id)).
        Only(c.Request.Context())
    if err != nil {
        if ent.IsNotFound(err) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 所有者でない場合は削除を許可しない
    if p.UserID != currentUser.ID {
        c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to delete this post"})
        return
    }

    err = h.client.Post.DeleteOne(p).Exec(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
