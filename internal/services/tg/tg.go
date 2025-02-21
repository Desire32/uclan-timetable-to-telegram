package tg

import (
	"fmt"
	tg "github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"os"
	"timetable/internal/services/tg/internal"
	"timetable/internal/services/timetable"
	"timetable/internal/services/timetable/data"
)

type TelegramBot struct {
	Bot        *tg.Bot
	Handler    *th.BotHandler
	Timetable  *timetable.TimeService
	BadgesPage *timetable.BadgeService
}

type ServiceTg struct{}

func (t *ServiceTg) TgConnection() error {
	bot, err := tg.NewBot(os.Getenv("BOT_TOKEN"), tg.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// LONGPOLLING ПРОЩЕ
	updates, _ := bot.UpdatesViaLongPolling(nil)
	botHandler, _ := th.NewBotHandler(bot, updates)

	defer bot.StopLongPolling()
	defer botHandler.Stop()
	t.TgHandlers(botHandler)
	botHandler.Start()

	return nil
}

func (t *ServiceTg) TgHandlers(botHandler *th.BotHandler) {
	botHandler.Handle(func(bot *tg.Bot, updates tg.Update) {
		chatID := updates.Message.Chat.ID
		msg := tu.Message(tu.ID(chatID), "Hello there! Please select: ").WithReplyMarkup(internal.GeneralMarkUp())
		_, _ = bot.SendMessage(msg)
	}, th.CommandEqual("menu"))

	botHandler.HandleCallbackQuery(func(bot *tg.Bot, query tg.CallbackQuery) {
		chatID := query.Message.GetChat().ID
		switch query.Data {
		// schedule
		case "show_schedule":
			msg := tu.Message(tu.ID(chatID), "Расписание")
			msg.ReplyMarkup = internal.ScheduleMarkUp()
			_, _ = bot.SendMessage(msg)
		// badges
		case "show_badges":
			msg := tu.Message(tu.ID(chatID), "Личное дело")
			msg.ReplyMarkup = internal.BadgesMarkUp()
			_, _ = bot.SendMessage(msg)
		// schedule cases
		case "daySchedule":
			internal.GetScheduleVariations(data.Schedule, bot, query, "today_schedule")
		case "fullSchedule":
			internal.GetScheduleVariations(data.Schedule, bot, query, "full_schedule")
		// badges cases
		case "badges":
			internal.GetBadgesInfo(data.Badge, bot, query, "badge_info")
		case "modules":
			internal.GetBadgesInfo(data.Badge, bot, query, "modules_percentages")
		}
	})
}
