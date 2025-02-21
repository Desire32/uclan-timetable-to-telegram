package internal

import (
	tg "github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func GeneralMarkUp() *tg.InlineKeyboardMarkup {
	general := &tg.InlineKeyboardMarkup{
		InlineKeyboard: [][]tg.InlineKeyboardButton{
			{
				tu.InlineKeyboardButton("Расписание").WithCallbackData("show_schedule"),
				tu.InlineKeyboardButton("Личное дело").WithCallbackData("show_badges"),
			},
		},
	}
	return general
}
func ScheduleMarkUp() *tg.InlineKeyboardMarkup {
	schedule := &tg.InlineKeyboardMarkup{
		InlineKeyboard: [][]tg.InlineKeyboardButton{
			{
				tu.InlineKeyboardButton("На сегодня").WithCallbackData("daySchedule"),
				tu.InlineKeyboardButton("Общее").WithCallbackData("fullSchedule"),
			},
		},
	}
	return schedule
}

func BadgesMarkUp() *tg.InlineKeyboardMarkup {
	badges := &tg.InlineKeyboardMarkup{
		InlineKeyboard: [][]tg.InlineKeyboardButton{
			{
				tu.InlineKeyboardButton("Успеваемость").WithCallbackData("badges"),
				tu.InlineKeyboardButton("Проценты модулей").WithCallbackData("modules"),
			},
		},
	}
	return badges
}
