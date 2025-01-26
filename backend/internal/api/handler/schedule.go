package handler

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/Shiluco/UniTimetable/backend/ent"
    "github.com/Shiluco/UniTimetable/backend/ent/schedule"
    //"github.com/Shiluco/UniTimetable/backend/internal/api/middleware"
)

type ScheduleHandler struct {
    client *ent.Client
}

func NewScheduleHandler(client *ent.Client) *ScheduleHandler {
    return &ScheduleHandler{client: client}
}

// GetSchedule 時間割取得ハンドラー（単一または一覧）
func (h *ScheduleHandler) GetSchedule(c *gin.Context) {
    // IDパラメータの取得（単一スケジュールの場合）
    idParam := c.Param("id")
    
    if idParam != "" {
        // 単一のスケジュールを取得
        id, err := strconv.Atoi(idParam)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID"})
            return
        }

        schedule, err := h.client.Schedule.Query().
            Where(schedule.ID(id)).
            WithPost(). // 投稿情報を取得
            Only(c.Request.Context())

        if err != nil {
            if ent.IsNotFound(err) {
                c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
                return
            }
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // 単一スケジュール情報をJSON形式で返す
        c.JSON(http.StatusOK, schedule)
        return
    }

    // スケジュールの一覧を取得
    schedules, err := h.client.Schedule.Query().
        WithPost(). // 投稿情報を取得
        All(c.Request.Context())

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // スケジュール一覧をJSON形式で返す
    c.JSON(http.StatusOK, schedules)
}

// CreateSchedule 時間割を作成
// func (h *ScheduleHandler) CreateSchedule(c *gin.Context) {
//     var req struct {
//         DayOfWeek int   `json:"day_of_week" binding:"omitempty,min=1,max=7"`
//         TimeSlot  int   `json:"time_slot" binding:"omitempty,min=1,max=7"`
//         Subject   string `json:"subject" binding:"omitempty,max=100"`
//         Location  string `json:"location" binding:"omitempty,max=100"`
//     }

//     if err := c.ShouldBindJSON(&req); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     // 現在のユーザーを取得
//     currentUser := middleware.GetCurrentUser(c)
//     if currentUser == nil {
//         c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
//         return
//     }

//     schedule, err := h.client.Schedule.Create().
//         SetDayOfWeek(req.DayOfWeek).
//         SetTimeSlot(req.TimeSlot).
//         SetSubject(req.Subject).
//         SetLocation(req.Location).
//         SetPostID(currentUser.ID).
//         SetCreatedAt(time.Now()).
//         SetUpdatedAt(time.Now()).
//         Save(c.Request.Context())

//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     c.JSON(http.StatusCreated, schedule)
// }

// UpdateSchedule 時間割を更新
// func (h *ScheduleHandler) UpdateSchedule(c *gin.Context) {
//     id, err := strconv.Atoi(c.Param("id"))
//     if err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID"})
//         return
//     }

//     var req struct {
//         DayOfWeek int   `json:"day_of_week" binding:"omitempty,min=0,max=6"`
//         TimeSlot  int   `json:"time_slot" binding:"omitempty,min=1,max=7"`
//         Subject   string `json:"subject" binding:"omitempty,max=100"`
//         Location  string `json:"location" binding:"omitempty,max=100"`
//     }

//     if err := c.ShouldBindJSON(&req); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     // 現在のユーザーを取得
//     currentUser := middleware.GetCurrentUser(c)
//     if currentUser == nil {
//         c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
//         return
//     }

//     // 時間割の所有者を確認
//     s, err := h.client.Schedule.Query().
//         Where(schedule.ID(id)).
//         Only(c.Request.Context())
//     if err != nil {
//         if ent.IsNotFound(err) {
//             c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
//             return
//         }
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     // 所有者でない場合は更新を許可しない
//     if s.UserID != currentUser.ID {
//         c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to update this schedule"})
//         return
//     }

//     update := h.client.Schedule.UpdateOneID(id).
//         SetUpdatedAt(time.Now())

//     if req.DayOfWeek != 0 {
//         update.SetDayOfWeek(req.DayOfWeek)
//     }
//     if req.TimeSlot != 0 {
//         update.SetTimeSlot(req.TimeSlot)
//     }
//     if req.Subject != "" {
//         update.SetSubject(req.Subject)
//     }
//     if req.Location != "" {
//         update.SetLocation(req.Location)
//     }

//     schedule, err := update.Save(c.Request.Context())
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     c.JSON(http.StatusOK, schedule)
// }

// DeleteSchedule 時間割を削除
// func (h *ScheduleHandler) DeleteSchedule(c *gin.Context) {
//     id, err := strconv.Atoi(c.Param("id"))
//     if err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID"})
//         return
//     }

//     // 現在のユーザーを取得
//     currentUser := middleware.GetCurrentUser(c)
//     if currentUser == nil {
//         c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
//         return
//     }

//     // 時間割の所有者を確認
//     s, err := h.client.Schedule.Query().
//         Where(schedule.ID(id)).
//         Only(c.Request.Context())
//     if err != nil {
//         if ent.IsNotFound(err) {
//             c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
//             return
//         }
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     // 所有者でない場合は削除を許可しない
//     if s.UserID != currentUser.ID {
//         c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to delete this schedule"})
//         return
//     }

//     err = h.client.Schedule.DeleteOne(s).Exec(c.Request.Context())
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     c.JSON(http.StatusOK, gin.H{"message": "Schedule deleted successfully"})
// }
