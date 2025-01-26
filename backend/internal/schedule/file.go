package schedule

import (
    "fmt"
    "mime/multipart"
    "path/filepath"
    "strings"
    "encoding/json"
    "io"
)

// ScheduleData 時間割データの構造体
type ScheduleData struct {
    DayOfWeek int    `json:"day_of_week" binding:"required,min=1,max=8"`
    TimeSlot  int    `json:"time_slot" binding:"required,min=1,max=8"`
    Subject   string `json:"subject" binding:"required,max=100"`
    Location  string `json:"location" binding:"required,max=100"`
}

// ProcessFile ファイル処理ハンドラー
func ProcessFile(file multipart.File, fileHeader *multipart.FileHeader) ([]ScheduleData, error) {

    // ファイルサイズの制限（10MB）
    if fileHeader.Size > 10<<20 {
        return nil, fmt.Errorf("ファイルサイズは10MB以下にしてください")
    }

    // ファイル拡張子の確認
    ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
    if ext != ".html" {
        return nil, fmt.Errorf("HTMLファイルのみ対応しています")
    }

    // multipart.Fileから[]byteに変換
    content, err := io.ReadAll(file)
    if err != nil {
        return nil, fmt.Errorf("ファイルの読み込みに失敗しました: %v", err)
    }

    // HTMLコンテンツを処理して時間割データを取得
    scheduleData, err := GetSchedule(content)
    if err != nil {
        return nil, fmt.Errorf("ファイルの処理に失敗しました: %v", err)
    }

    // 解析結果をScheduleData配列に変換
    var scheduleResponse struct {
        Schedules []ScheduleData `json:"schedules"`
    }
    if err := json.Unmarshal(scheduleData, &scheduleResponse); err != nil {
        return nil, fmt.Errorf("時間割データの解析に失敗しました")
    }

    // デバッグ用のログ
    fmt.Printf("Parsed schedules: %+v\n", scheduleResponse.Schedules)

    return scheduleResponse.Schedules, nil
}
