package schedule

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetSchedule(htmlPath string) ([]byte, error) {
	// HTMLファイルを開く
	file, err := os.Open(htmlPath)
	if err != nil {
		return nil, fmt.Errorf("ファイルを開けません: %v", err)
	}
	defer file.Close()

	// goqueryでHTMLを解析
	doc, err := goquery.NewDocumentFromReader(file)
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
	var schedule []map[string]string

	// 行データを解析
	doc.Find(".schedule-table tbody tr").Each(func(rowIdx int, tr *goquery.Selection) {
		// 最初の行（ヘッダー）はスキップ
		if rowIdx == 0 {
			return
		}

		// 時間帯（例: "1・2"）を取得
		timeSlot := strings.TrimSpace(tr.Find("th").First().Text())

		// 各曜日列を解析
		tr.Find("td").Each(func(colIdx int, td *goquery.Selection) {
			if colIdx < len(weekdays) { // 対応する曜日が存在する場合のみ処理
				day := weekdays[colIdx]
				td.Find("li").Each(func(liIdx int, li *goquery.Selection) { // 修正箇所
					course := make(map[string]string)
					course["曜日"] = day
					course["時間"] = timeSlot
					course["講義名"] = strings.TrimSpace(li.Find("h4").Text())
					course["担当"] = strings.TrimSpace(li.Find("p").First().Text())
					course["単位"] = strings.TrimSpace(li.Find("p").Eq(1).Text())
					course["教室"] = strings.TrimSpace(li.Find("p span").Text())
					schedule = append(schedule, course)
				})
			}
		})
	})

	// JSONに変換
	jsonData, err := json.MarshalIndent(schedule, "", "  ")
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
