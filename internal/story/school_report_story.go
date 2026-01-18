package story

import (
	"context"

	"School/internal/domain"
	"School/internal/repository"
)

type SchoolReportStory struct {
	studentRepository *repository.StudentRepository
}

func NewSchoolReportStory(repo *repository.StudentRepository) *SchoolReportStory {
	return &SchoolReportStory{studentRepository: repo}
}

func (s *SchoolReportStory) BuildReport(ctx context.Context) (*domain.SchoolReport, error) {
	students, err := s.studentRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	report := &domain.SchoolReport{
		HeightGroups: []domain.HeightGroupReport{
			{Group: "<150", Count: 0},
			{Group: "150-160", Count: 0},
			{Group: "160-170", Count: 0},
			{Group: "170-180", Count: 0},
			{Group: ">=180", Count: 0},
		},
	}

	for _, st := range students {

		switch {
		case st.Height < 150:
			report.HeightGroups[0].Count++
		case st.Height < 160:
			report.HeightGroups[1].Count++
		case st.Height < 170:
			report.HeightGroups[2].Count++
		case st.Height < 180:
			report.HeightGroups[3].Count++
		default:
			report.HeightGroups[4].Count++
		}

		if st.Gender == "Male" {
			report.Males++
		} else {
			report.Females++
		}

		if st.Age >= 18 {
			report.Adults++
		}
	}

	return report, nil
}
