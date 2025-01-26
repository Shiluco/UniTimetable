package univData

import (
    "context"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"

    "github.com/Shiluco/UniTimetable/backend/ent"
    // "github.com/Shiluco/UniTimetable/backend/ent/department"
    // "github.com/Shiluco/UniTimetable/backend/ent/major"
)

// Department は学部の情報を表す構造体です。
type Department struct {
    DepartmentID int      `json:"department_id"`
    Name         string   `json:"name"`
    Majors       []Major  `json:"majors"`
}

// Major は学科の情報を表す構造体です。
type Major struct {
    MajorID      int    `json:"major_id"`
    DepartmentID int    `json:"department_id"`
    Name         string `json:"name"`
}

// University は大学の情報を表す構造体です。
type University struct {
    Departments []Department `json:"departments"`
}

// SaveUniversityData は JSON データをデータベースに保存する関数です。
func SaveUniversityData(ctx context.Context, client *ent.Client) error {
	jsonFilePath := "./univ.json"
    jsonFile, err := os.Open(jsonFilePath)
    if err != nil {
        return fmt.Errorf("failed to open JSON file: %w", err)
    }
    defer jsonFile.Close()

    byteValue, err := ioutil.ReadAll(jsonFile)
    if err != nil {
        return fmt.Errorf("failed to read JSON file: %w", err)
    }

    var university University
    if err := json.Unmarshal(byteValue, &university); err != nil {
        return fmt.Errorf("failed to unmarshal JSON: %w", err)
    }

    // データベースに保存
    for _, dept := range university.Departments {
        d, err := client.Department.Create().
            SetID(dept.DepartmentID).
            SetName(dept.Name).
            Save(ctx)
        if err != nil {
            return fmt.Errorf("failed to create department: %w", err)
        }

        for _, major := range dept.Majors {
            _, err := client.Major.Create().
                SetID(major.MajorID).
                SetDepartmentID(d.ID).
                SetName(major.Name).
                Save(ctx)
            if err != nil {
                return fmt.Errorf("failed to create major: %w", err)
            }
        }
    }

    return nil
}