package data

type Schedules struct {
	Day     string    `json:"day"`
	Modules []Modules `json:"modules"`
}

type Modules struct {
	Time string `json:"time"`
	Name string `json:"name"`
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
