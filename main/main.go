package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"School/internal/app"
	"School/internal/repository"
	"School/internal/story"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:zXc12026@localhost:5434/students?sslmode=disable"
	}

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	if err := app.RunMigrations(dsn); err != nil {
		log.Fatal(err)
	}

	studentRepo := repository.NewStudentRepository(pool)
	reportStory := story.New(studentRepo)

	report, err := reportStory.BuildReport(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("SCHOOL REPORT")
	for _, g := range report.HeightGroups {
		log.Printf("  %-7s : %d\n", g.Group, g.Count)
	}
	log.Printf("  Males   : %d\n", report.Males)
	log.Printf("  Females : %d\n", report.Females)
	log.Printf("  Adults  : %d\n", report.Adults)
}
