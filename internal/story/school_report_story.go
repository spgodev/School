package story

import (
	"School/internal/domain"
	"context"

	"fmt"
)

type SchoolReportStory struct {
	repo StudentGetter
}

func New(repo StudentGetter) *SchoolReportStory {
	return &SchoolReportStory{repo: repo}
}

const AdultAge = 18

func (s *SchoolReportStory) BuildReport(ctx context.Context) (*domain.SchoolReport, error) {
	students, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	report := &domain.SchoolReport{
		HeightGroups: map[domain.HeightGroupName]int{
			domain.Below150:     0,
			domain.From150To160: 0,
			domain.From160To170: 0,
			domain.From170To180: 0,
			domain.Higher180:    0,
		},
	}

	for _, st := range students {
		var group domain.HeightGroupName
		switch {
		case st.Height < 150:
			group = domain.Below150
		case st.Height < 160:
			group = domain.From150To160
		case st.Height < 170:
			group = domain.From160To170
		case st.Height < 180:
			group = domain.From170To180
		default:
			group = domain.Higher180
		}
		report.HeightGroups[group]++

		switch st.Gender {
		case domain.Male:
			report.Males++
		case domain.Female:
			report.Females++
		default:
			return nil, fmt.Errorf("unknown gender: %v", st.Gender)
		}

		if st.Age >= AdultAge {
			report.Adults++
		}
	}

	return report, nil
}
