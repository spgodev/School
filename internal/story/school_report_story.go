package story

import (
	"context"

	"School/internal/domain"
)

type SchoolReportStory struct {
	repo domain.StudentGetter
}

const AdultAge = 18

func New(repo domain.StudentGetter) *SchoolReportStory {
	return &SchoolReportStory{repo: repo}
}

func (s *SchoolReportStory) BuildReport(ctx context.Context) (*domain.SchoolReport, error) {
	students, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	report := &domain.SchoolReport{
		HeightGroups: []domain.HeightGroupReport{
			{Group: domain.Below150, Count: 0},
			{Group: domain.From150To160, Count: 0},
			{Group: domain.From160To170, Count: 0},
			{Group: domain.From170To180, Count: 0},
			{Group: domain.Higher180, Count: 0},
		},
	}

	counts := map[domain.HeightGroupName]int{
		domain.Below150:     0,
		domain.From150To160: 0,
		domain.From160To170: 0,
		domain.From170To180: 0,
		domain.Higher180:    0,
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
		counts[group]++

		if st.Gender == domain.Male {
			report.Males++
		} else {
			report.Females++
		}

		if st.Age >= AdultAge {
			report.Adults++
		}
	}

	report.HeightGroups = make([]domain.HeightGroupReport, 0, len(counts))
	for group, count := range counts {
		report.HeightGroups = append(report.HeightGroups, domain.HeightGroupReport{
			Group: group,
			Count: count,
		})
	}

	return report, nil
}
