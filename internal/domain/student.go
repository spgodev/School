package domain

type Student struct {
	ID     int64
	Name   string
	Age    int
	Gender Gender
	Height int
}

type Gender string

var (
	Male   Gender = "Male"
	Female Gender = "Female"
)

type HeightGroupReport struct {
	Group HeightGroupName
	Count int
}

type HeightGroupName string

const (
	Below150     HeightGroupName = "<150"
	From150To160 HeightGroupName = "150-160"
	From160To170 HeightGroupName = "160-170"
	From170To180 HeightGroupName = "170-180"
	Higher180    HeightGroupName = ">=180"
)

type SchoolReport struct {
	HeightGroups map[HeightGroupName]int
	Males        int
	Females      int
	Adults       int
}
