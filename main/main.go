package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"School/internal/app"
	"School/internal/controller"
	"School/internal/repository"
	"School/internal/story"
)

func main() {
	setupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:zXc12026@localhost:5434/students?sslmode=disable"
	}

	pool, err := pgxpool.New(setupCtx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	if err := pool.Ping(setupCtx); err != nil {
		log.Fatal(err)
	}

	if err := app.RunMigrations(dsn); err != nil {
		log.Fatal(err)
	}

	studentRepo := repository.NewStudentRepository(pool)
	reportStory := story.New(studentRepo)

	studentController := controller.NewStudentController(studentRepo, reportStory)

	r := gin.Default()
	r.POST("/students", studentController.CreateStudent)
	r.GET("/students", studentController.ListStudents)
	r.GET("/students/report", studentController.GetReport)

	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	log.Fatal(r.Run(addr))
}
