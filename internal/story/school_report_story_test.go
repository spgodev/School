package story

import (
	"context"
	"errors"
	"testing"

	"School/internal/domain"
	"School/internal/repository"
)

func newStory(students []domain.Student, repoErr error) *SchoolReportStory {
	repo := repository.FakeRepo{
		Students: students,
		Err:      repoErr,
	}
	return New(repo)
}

func TestBuildReport_OK(t *testing.T) {
	students := []domain.Student{
		{ID: 1, Name: "A", Age: 17, Gender: domain.Male, Height: 149},
		{ID: 2, Name: "B", Age: 18, Gender: domain.Female, Height: 150},
		{ID: 3, Name: "C", Age: 30, Gender: domain.Male, Height: 160},
		{ID: 4, Name: "D", Age: 18, Gender: domain.Female, Height: 179},
		{ID: 5, Name: "E", Age: 10, Gender: domain.Male, Height: 180},
	}

	s := newStory(students, nil)

	report, err := s.BuildReport(context.Background())

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if report == nil {
		t.Fatalf("expected report, got nil")
	}

	if report.Males != 3 {
		t.Fatalf("expected Males=3, got %d", report.Males)
	}
	if report.Females != 2 {
		t.Fatalf("expected Females=2, got %d", report.Females)
	}

	if report.Adults != 3 {
		t.Fatalf("expected Adults=3, got %d", report.Adults)
	}

	want := map[domain.HeightGroupName]int{
		domain.Below150:     1, // 149
		domain.From150To160: 1, // 150
		domain.From160To170: 1, // 160
		domain.From170To180: 1, // 179
		domain.Higher180:    1, // 180
	}

	if report.HeightGroups == nil {
		t.Fatalf("expected HeightGroups map, got nil")
	}

	for group, exp := range want {
		got := report.HeightGroups[group]
		if got != exp {
			t.Fatalf("group %s: expected %d, got %d", group, exp, got)
		}
	}
}

func TestBuildReport_RepoError(t *testing.T) {
	repoErr := errors.New("db down")
	s := newStory(nil, repoErr)

	report, err := s.BuildReport(context.Background())

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !errors.Is(err, repoErr) {
		t.Fatalf("expected repoErr, got %v", err)
	}
	if report != nil {
		t.Fatalf("expected nil report, got %+v", report)
	}
}

func TestBuildReport_UnknownGender(t *testing.T) {
	students := []domain.Student{
		{ID: 1, Name: "X", Age: 20, Gender: "Other", Height: 170},
	}

	s := newStory(students, nil)

	report, err := s.BuildReport(context.Background())

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if report != nil {
		t.Fatalf("expected nil report, got %+v", report)
	}
}

func TestBuildReport_EmptyStudents(t *testing.T) {
	s := newStory(nil, nil)

	report, err := s.BuildReport(context.Background())
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if report == nil {
		t.Fatalf("expected report, got nil")
	}

	if report.Males != 0 || report.Females != 0 || report.Adults != 0 {
		t.Fatalf("expected all counters 0, got males=%d females=%d adults=%d",
			report.Males, report.Females, report.Adults)
	}

	if report.HeightGroups == nil {
		t.Fatalf("expected HeightGroups map, got nil")
	}

	groups := []domain.HeightGroupName{
		domain.Below150,
		domain.From150To160,
		domain.From160To170,
		domain.From170To180,
		domain.Higher180,
	}

	for _, g := range groups {
		if report.HeightGroups[g] != 0 {
			t.Fatalf("expected group %s = 0, got %d", g, report.HeightGroups[g])
		}
	}
}
