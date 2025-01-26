package handler

import (
    "encoding/json"
    "fmt"
    "io"
    "mime/multipart"
    "net/http"
    "path/filepath"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/Shiluco/UniTimetable/backend/ent"
    "github.com/Shiluco/UniTimetable/backend/internal/api/middleware"
    "github.com/Shiluco/UniTimetable/backend/internal/schedule"
)


// ScheduleData 時間割データの構造体
type ScheduleData struct {
    DayOfWeek int    `json:"day_of_week" binding:"required,min=1,max=8"`
    TimeSlot  int    `json:"time_slot" binding:"required,min=1,max=8"`
    Subject   string `json:"subject" binding:"required,max=100"`
    Location  string `json:"location" binding:"required,max=100"`
}

// ProcessFile ファイル処理ハンドラー
func ProcessFile(file multipart.File) ([]byte, error) {

    // ファイルサイズの制限（10MB）
    if file.Size > 10<<20 {
		return nil, fmt.Errorf("ファイルサイズは10MB以下にしてください")
    }

    // ファイル拡張子の確認
    ext := strings.ToLower(filepath.Ext(file.Filename))
    if ext != ".html" {
        return nil, fmt.Errorf("HTMLファイルのみ対応しています")
    }

    // ファイルを開く
    src, err := file.Open()
    if err != nil {
        return nil, fmt.Errorf("ファイルを開けません")
    }
    defer src.Close()

    // ファイルの内容を読み込む
    content, err := processHTMLFile(src)
    if err != nil {
        return nil, fmt.Errorf("ファイルの処理に失敗しました: %v", err)
    }

    // 解析結果をScheduleData配列に変換
    var scheduleResponse struct {
        Schedules []ScheduleData `json:"schedules"`
    }
    if err := json.Unmarshal(content, &scheduleResponse); err != nil {
        return nil, fmt.Errorf("時間割データの解析に失敗しました")
    }

    // デバッグ用のログ
    fmt.Printf("Parsed schedules: %+v\n", scheduleResponse.Schedules)

    return scheduleResponse.Schedules, nil
}

// processHTMLFile HTMLファイルを処理する関数
func processHTMLFile(file multipart.File) ([]byte, error) {
    content, err := io.ReadAll(file)
    if err != nil {
        return nil, fmt.Errorf("ファイルの読み込みに失敗しました: %v", err)
    }

    scheduleData, err := schedule.GetSchedule(content)
    if err != nil {
        return nil, fmt.Errorf("時間割の解析に失敗しました: %v", err)
    }

    return scheduleData, nil
} 