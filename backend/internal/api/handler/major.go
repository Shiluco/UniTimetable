package handler

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/Shiluco/UniTimetable/backend/ent"
    "github.com/Shiluco/UniTimetable/backend/ent/major"
)

type MajorHandler struct {
    client *ent.Client
}

func NewMajorHandler(client *ent.Client) *MajorHandler {
    return &MajorHandler{client: client}
}

// GetMajors 学科一覧を取得
func (h *MajorHandler) GetMajors(c *gin.Context) {
    ctx := c.Request.Context()
    
    // 学部IDによるフィルタリング（オプション）
    departmentID := c.Query("department_id")
    
    query := h.client.Major.Query().
        Order(ent.Asc(major.FieldName))

    // 学部IDが指定されている場合はフィルタリング
    if departmentID != "" {
        deptID, err := strconv.Atoi(departmentID)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
            return
        }
        query.Where(major.DepartmentID(deptID))
    }

    // WithDepartmentを使用して学部情報も取得
    majors, err := query.
        WithDepartment().
        All(ctx)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, majors)
}

// GetMajor 特定の学科を取得
func (h *MajorHandler) GetMajor(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid major ID"})
        return
    }

    major, err := h.client.Major.Query().
        Where(major.ID(id)).
        WithDepartment().  // 学部情報も含める
        Only(c.Request.Context())

    if err != nil {
        if ent.IsNotFound(err) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Major not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, major)
}

// GetMajorsByDepartment 特定の学部に属する学科一覧を取得
func (h *MajorHandler) GetMajorsByDepartment(c *gin.Context) {
    departmentID, err := strconv.Atoi(c.Param("department_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
        return
    }

    majors, err := h.client.Major.Query().
        Where(major.DepartmentID(departmentID)).
        Order(ent.Asc(major.FieldName)).
        WithDepartment().
        All(c.Request.Context())

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, majors)
}
