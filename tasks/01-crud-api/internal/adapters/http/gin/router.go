package ginx

import (
	"fmt"
	"sort"

	"github.com/apiwatdev/go-practice-tasks/tasks/01-crud-api/internal/adapters/http/gin/middleware"
	domrepo "github.com/apiwatdev/go-practice-tasks/tasks/01-crud-api/internal/adapters/repo/memory"
	uc "github.com/apiwatdev/go-practice-tasks/tasks/01-crud-api/internal/usecase/todo"
	"github.com/gin-gonic/gin"
)

func BuildRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.RequestID())

	// wiring: repo -> usecase -> handler
	repo := domrepo.NewTodoRepo()
	svc := uc.NewService(repo)
	h := NewHandler(svc)
	h.Register(r)

	logRoutes(r)

	return r
}

func logRoutes(r *gin.Engine) {
	rts := r.Routes()
	sort.Slice(rts, func(i, j int) bool {
		if rts[i].Method == rts[j].Method {
			return rts[i].Path < rts[j].Path
		}
		return rts[i].Method < rts[j].Method
	})

	fmt.Println("=== Registered Routes ===")
	for _, rt := range rts {
		fmt.Printf("%-6s  %-40s\n", rt.Method, rt.Path)
	}
	fmt.Println("=========================")
}
