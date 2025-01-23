package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Shiluco/UniTimetable/backend/ent"
	"github.com/Shiluco/UniTimetable/backend/ent/department"
)

type DepartmentHandler struct {
	client *ent.Client
}

func NewDepartmentHandler(client *ent.Client) *DepartmentHandler {
	return &DepartmentHandler{client: client}
}

// GetDepartments 学部一覧を取得
func (h *DepartmentHandler) GetDepartments(c *gin.Context) {
	departments, err := h.client.Department.Query().
		Order(ent.Asc(department.FieldName)).
		All(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, departments)
}

// GetDepartment 特定の学部を取得
func (h *DepartmentHandler) GetDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}

	dept, err := h.client.Department.Query().
		Where(department.ID(id)).
		Only(c.Request.Context())

	if err != nil {
		if ent.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dept)
}
