package data

// Цельное расписание
type Schedules struct {
	Day     string    `bson:"day" json:"day"`
	Modules []Modules `bson:"modules" json:"modules"`
}

type Modules struct {
	Time string `bson:"time" json:"time"`
	Name string `bson:"name" json:"name"`
}

type Badges struct {
	Number            string `json:"number"`
	ModuleName        string `json:"module_name"`
	Semester          string `json:"semester"`
	TotalSessions     string `json:"total_sessions"`
	AttendanceAmount  string `json:"attendances_amount"`
	AttendanceOverall string `json:"attendance_overall"`
	AbsensesOverall   string `json:"absences_overall"`
}

var Badge []Badges
var Schedule []Schedules
var Module []Modules
var BadgeAmount string
