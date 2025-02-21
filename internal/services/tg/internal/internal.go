package internal

import (
	"fmt"
	tg "github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"time"
	"timetable/internal/services/timetable/data"
)

func GetScheduleVariations(schedules []data.Schedules, bot *tg.Bot, query tg.CallbackQuery, scheduleType string) *tg.SendMessageParams {
	currentDay := time.Now().Weekday().String()
	var scheduleText string
	chatID := query.Message.GetChat().ID
	switch scheduleType {
	case "full_schedule":
		for _, schedule := range schedules {
			scheduleText += fmt.Sprintf("Day: %s\n", schedule.Day)
			for _, module := range schedule.Modules {
				scheduleText += fmt.Sprintf("  %s: %s\n", module.Time, module.Name)
			}
		}
	case "today_schedule":
		for _, schedule := range schedules {
			if schedule.Day == currentDay {
				scheduleText += fmt.Sprintf("Day: %s\n", schedule.Day)
				for _, module := range schedule.Modules {
					scheduleText += fmt.Sprintf("  %s: %s\n", module.Time, module.Name)
				}
				break
			}
		}
	default:
		scheduleText = "Unknown selection."
	}
	msg := tu.Message(tu.ID(chatID), scheduleText)
	_, _ = bot.SendMessage(msg)
	return msg
}

func GetBadgesInfo(badges []data.Badges, bot *tg.Bot, query tg.CallbackQuery, badgesType string) *tg.SendMessageParams {
	var badgesText string
	chatID := query.Message.GetChat().ID
	switch badgesType {
	case "badge_info":
		badgesText += data.BadgeAmount
	case "modules_percentages":
		for _, badge := range badges {
			badgesText += fmt.Sprintf("Module: %s\nSemester: %s\nTotal Sessions: %s\nAttendance Amount: %s\nAttendance Overall: %s\nAbsences Overall: %s\n\n",
				badge.ModuleName, badge.Semester, badge.TotalSessions, badge.AttendanceAmount, badge.AttendanceOverall, badge.AbsensesOverall)
		}
	default:
		badgesText = "Unknown selection."
	}
	msg := tu.Message(tu.ID(chatID), badgesText)
	_, _ = bot.SendMessage(msg)
	return msg
}
