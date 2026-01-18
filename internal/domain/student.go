package domain

type Student struct {
	ID     int64
	Name   string
	Age    int
	Gender string
	Height int
}

type HeightGroupReport struct {
	Group string
	Count int
}

type SchoolReport struct {
	HeightGroups []HeightGroupReport
	Males        int
	Females      int
	Adults       int
}
