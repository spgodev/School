package controller

import (
	"School/internal/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	repo   StudentRepo
	report ReportBuilder
}

func NewStudentController(repo StudentRepo, report ReportBuilder) *StudentController {
	return &StudentController{repo: repo, report: report}
}

type CreateStudentRequest struct {
	Name   string `json:"name" binding:"required"`
	Age    int    `json:"age" binding:"required"`
	Gender string `json:"gender" binding:"required"`
	Height int    `json:"height" binding:"required"`
}

type StudentResponse struct {
	ID     int64         `json:"id"`
	Name   string        `json:"name"`
	Age    int           `json:"age"`
	Gender domain.Gender `json:"gender"`
	Height int           `json:"height"`
}

func (h *StudentController) CreateStudent(c *gin.Context) {
	var req CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Age <= 0 || req.Height <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "age and height must be > 0"})
		return
	}
	if req.Gender != string(domain.Male) && req.Gender != string(domain.Female) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender must be 'Male' or 'Female'"})
		return
	}

	id, err := h.repo.Create(c.Request.Context(), domain.Student{
		Name:   req.Name,
		Age:    req.Age,
		Gender: domain.Gender(req.Gender),
		Height: req.Height,
	})
	if err != nil {
		log.Printf("CreateStudent DB error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *StudentController) ListStudents(c *gin.Context) {
	students, err := h.repo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}

	out := make([]StudentResponse, 0, len(students))
	for _, s := range students {
		out = append(out, StudentResponse{
			ID:     s.ID,
			Name:   s.Name,
			Age:    s.Age,
			Gender: s.Gender,
			Height: s.Height,
		})
	}

	c.JSON(http.StatusOK, out)
}

func (h *StudentController) GetReport(c *gin.Context) {
	report, err := h.report.BuildReport(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "report error"})
		return
	}
	c.JSON(http.StatusOK, report)
}
