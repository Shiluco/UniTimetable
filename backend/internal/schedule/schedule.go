package schedule

import (
	"fmt"
	"bytes"
	//"os"
	"strings"

	"encoding/json"
	"github.com/PuerkitoBio/goquery"
)
var timeSlotMap = map[string]int{
	"1・2": 1,
	"3・4": 2,
	"5・6": 3,
	"7・8": 4,
	"9・10": 5,
	"11・12": 6,
	"13・14": 7,
}
type ScheduleBody struct {
	ID int `json:"id"`
	PostID int `json:"post_id"`
	DayOfWeek int `json:"day_of_week"`
	TimeSlot int `json:"time_slot"`
	Subject string `json:"subject"`
	Location string `json:"location"`
}

func GetSchedule(content []byte) ([]byte, error) {
	// HTMLファイルを開く
	// file, err := os.Open(htmlPath)
	// if err != nil {
	// 	return nil, fmt.Errorf("ファイルを開けません: %v", err)
	// }
	// defer file.Close()

	// goqueryでHTMLを解析
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
	if err != nil {
		return nil, fmt.Errorf("HTMLを解析できません: %v", err)
	}

	// 曜日と列インデックスの対応を取得
	var weekdays []string
	doc.Find("#weeksRow th").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if text != "" {
			weekdays = append(weekdays, text) // 曜日を追加
		}
	})

	// 時間割データを格納するスライス
	var schedule []map[string]interface{}

	// 行データを解析
	doc.Find(".schedule-table tbody tr").Each(func(rowIdx int, tr *goquery.Selection) {
		// 最初の行（ヘッダー）はスキップ
		if rowIdx == 0 {
			return
		}

		// 時間帯（例: "1・2"）を取得
		timeSlot := strings.TrimSpace(tr.Find("th").First().Text())
		timeSlotInt, ok := timeSlotMap[timeSlot];
		if !ok {
			return
		}
		// 各曜日列を解析
		tr.Find("td").Each(func(colIdx int, td *goquery.Selection) {
			if colIdx < len(weekdays) { // 対応する曜日が存在する場合のみ処理
				day := colIdx + 1
				td.Find("li").Each(func(liIdx int, li *goquery.Selection) { // 修正箇所
					course := make(map[string]interface{})
					course["day_of_week"] = day
					course["time_slot"] = timeSlotInt
					course["subject"] = strings.TrimSpace(li.Find("h4").Text())
					course["location"] = strings.TrimSpace(li.Find("p span").Text())
					schedule = append(schedule, course)
				})
			}
		})
	})
	schedules := map[string]interface{}{
		"schedules": schedule,
	}
	// JSONに変換
	jsonData, err := json.MarshalIndent(schedules, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("JSONへの変換に失敗しました: %v", err)
	}

	// // JSONをファイルに保存
	// err = os.WriteFile("時間割2.json", jsonData, 0644)
	// if err != nil {
	// 	log.Fatalf("JSONファイルの保存に失敗しました: %v", err)
	// }

	// fmt.Println("時間割をJSONに変換しました。'時間割.json'に保存しました。")

	return jsonData, nil
}
