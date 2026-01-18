package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

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
	
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	if err := app.RunMigrations(db); err != nil {
		log.Fatal(err)
	}

	studentRepo := repository.NewStudentRepository(db)
	reportStory := story.NewSchoolReportStory(studentRepo)

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
