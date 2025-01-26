package handler

// import (
//     "encoding/json"
//     "fmt"
//     "io"
//     "mime/multipart"
//     "net/http"
//     "path/filepath"
//     "strings"

//     "github.com/gin-gonic/gin"
//     "github.com/Shiluco/UniTimetable/backend/ent"
//     "github.com/Shiluco/UniTimetable/backend/internal/api/middleware"
//     "github.com/Shiluco/UniTimetable/backend/internal/schedule"
// )

// type FileHandler struct {
//     client *ent.Client
// }

// func NewFileHandler(client *ent.Client) *FileHandler {
//     return &FileHandler{client: client}
// }

// // ScheduleData 時間割データの構造体
// type ScheduleData struct {
//     DayOfWeek int    `json:"day_of_week" binding:"required,min=1,max=8"`
//     TimeSlot  int    `json:"time_slot" binding:"required,min=1,max=8"`
//     Subject   string `json:"subject" binding:"required,max=100"`
//     Location  string `json:"location" binding:"required,max=100"`
// }

// // ProcessFile ファイル処理ハンドラー
// func (h *FileHandler) ProcessFile(c *gin.Context) {
//     // 現在のユーザーを取得
//     currentUser := middleware.GetCurrentUser(c)
//     if currentUser == nil {
//         c.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーが見つかりません"})
//         return
//     }

//     // ファイルを取得
//     file, err := c.FormFile("file")
//     if err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "ファイルが見つかりません"})
//         return
//     }

//     // ファイルサイズの制限（10MB）
//     if file.Size > 10<<20 {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "ファイルサイズは10MB以下にしてください"})
//         return
//     }

//     // ファイル拡張子の確認
//     ext := strings.ToLower(filepath.Ext(file.Filename))
//     if ext != ".html" {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "HTMLファイルのみ対応しています"})
//         return
//     }

//     // ファイルを開く
//     src, err := file.Open()
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "ファイルを開けません"})
//         return
//     }
//     defer src.Close()

//     // ファイルの内容を読み込む
//     content, err := processHTMLFile(src)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("ファイルの処理に失敗しました: %v", err)})
//         return
//     }

//     // 解析結果をScheduleData配列に変換
//     var scheduleResponse struct {
//         Schedules []ScheduleData `json:"schedules"`
//     }
//     if err := json.Unmarshal(content, &scheduleResponse); err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "時間割データの解析に失敗しました"})
//         return
//     }

//     // デバッグ用のログ
//     fmt.Printf("Parsed schedules: %+v\n", scheduleResponse.Schedules)

//     ctx := c.Request.Context()
//     var savedSchedules []*ent.Schedule

//     // トランザクションを開始
//     tx, err := h.client.Tx(ctx)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "トランザクションの開始に失敗しました"})
//         return
//     }

//     // 各時間割をDBに保存
//     for _, scheduleData := range scheduleResponse.Schedules {
//         schedule, err := tx.Schedule.Create().
//             SetUserID(currentUser.ID).
//             SetDayOfWeek(scheduleData.DayOfWeek).
//             SetTimeSlot(scheduleData.TimeSlot).
//             SetSubject(scheduleData.Subject).
//             SetLocation(scheduleData.Location).
//             Save(ctx)

//         if err != nil {
//             tx.Rollback()
//             c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("時間割の保存に失敗しました: %v", err)})
//             return
//         }

//         savedSchedules = append(savedSchedules, schedule)
//     }

//     // トランザクションをコミット
//     if err := tx.Commit(); err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "トランザクションのコミットに失敗しました"})
//         return
//     }

//     c.JSON(http.StatusOK, gin.H{
//         "message": "時間割を保存しました",
//         "schedules": savedSchedules,
//     })
// }

// // processHTMLFile HTMLファイルを処理する関数
// func processHTMLFile(file multipart.File) ([]byte, error) {
//     content, err := io.ReadAll(file)
//     if err != nil {
//         return nil, fmt.Errorf("ファイルの読み込みに失敗しました: %v", err)
//     }

//     scheduleData, err := schedule.GetSchedule(content)
//     if err != nil {
//         return nil, fmt.Errorf("時間割の解析に失敗しました: %v", err)
//     }

//     return scheduleData, nil
// } 