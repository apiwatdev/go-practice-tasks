package ginx

import (
	"net/http"
	"strconv"
	"time"

	"github.com/apiwatdev/go-practice-tasks/tasks/01-crud-api/internal/adapters/http/gin/response"
	dom "github.com/apiwatdev/go-practice-tasks/tasks/01-crud-api/internal/domain/todo"
	uc "github.com/apiwatdev/go-practice-tasks/tasks/01-crud-api/internal/usecase/todo"
	"github.com/gin-gonic/gin"
)

type Handler struct{ svc uc.Service }

func NewHandler(svc uc.Service) *Handler {
	return &Handler{svc: svc}
}

type createReq struct {
	Title string `json:"title" binding:"required, min=1, max=200"`
}

type updateReq struct {
	Title  string `json:"title" binding:"required, min=1, max=200"`
	IsDone *bool  `json:"isDone" binding:"required"`
}

func (h *Handler) Register(r *gin.Engine) {
	v1 := r.Group("/v1")
	todo := v1.Group("/todos")

	todo.POST("", h.Create)
	todo.GET("", h.List)
	todo.GET("/:id", h.Get)
	todo.PUT("/:id", h.Update)
	todo.DELETE("/:id", h.Delete)

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"time":   time.Now().Format("2006-01-02 15:04:05"),
		})
	})

}

func (h *Handler) Create(c *gin.Context) {
	var req createReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid_body", err.Error())
		return
	}
	td, err := h.svc.Create(dom.Todo{Title: req.Title})
	if err != nil {
		response.BadRequest(c, "create_failed", err.Error())
		return
	}
	c.JSON(http.StatusCreated, toJSON(td))
}

func (h *Handler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	items, totalPage, total, err := h.svc.List(page, size)
	if err != nil {
		response.BadRequest(c, "list_failed", err.Error())
		return
	}
	out := make([]gin.H, 0, len(items))
	for _, v := range items {
		out = append(out, toJSON(v))
	}
	c.JSON(http.StatusOK, gin.H{
		"items":      out,
		"total":      total,
		"page":       page,
		"page_size":  size,
		"total_page": totalPage,
	})
}

func (h *Handler) Get(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	td, ok, err := h.svc.Get(id)
	if !ok || err == uc.ErrNotFound {
		response.NotFound(c, "todo_not_found", "todo not found")
		return
	}
	if err != nil {
		response.BadRequest(c, "get_failed", err.Error())
		return
	}
	c.JSON(http.StatusOK, toJSON(td))
}

func (h *Handler) Update(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var req updateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid_body", err.Error())
		return
	}
	td, err := h.svc.Update(id, req.Title, *req.IsDone)
	if err == uc.ErrNotFound {
		response.NotFound(c, "todo_not_found", "todo not found")
		return
	}
	if err != nil {
		response.BadRequest(c, "update_failed", err.Error())
		return
	}
	c.JSON(http.StatusOK, toJSON(td))
}

func (h *Handler) Delete(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	if ok, err := h.svc.Delete(id); !ok || err == uc.ErrNotFound {
		response.NotFound(c, "todo_not_found", "todo not found")
		return
	} else if err != nil {
		response.BadRequest(c, "delete_failed", err.Error())
		return
	}
	c.Status(http.StatusNoContent)
}

func parseID(c *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid_id", "id must be integer")
		return 0, false
	}
	return id, true
}
func toJSON(t dom.Todo) gin.H {
	return gin.H{"id": t.ID, "title": t.Title, "is_done": t.IsDone}
}
