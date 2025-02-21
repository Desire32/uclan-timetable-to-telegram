package timetable

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"time"
)

type TimeService struct {
}

func (t *TimeService) TimetableAuth(ctx context.Context) error {
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://timetable.cyprus.uclan.ac.uk/"),
		chromedp.Sleep(3*time.Second),                
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}
	return chromedp.Run(ctx,
		chromedp.Sleep(3*time.Second),
		chromedp.WaitVisible(`input[name="ctl00$ContentPlaceHolder1$txtEmail"]`, chromedp.ByQuery),
		chromedp.SendKeys(`input[name="ctl00$ContentPlaceHolder1$txtEmail"]`, os.Getenv("EMAIL"), chromedp.ByQuery),
		chromedp.WaitVisible(`input[name="ctl00$ContentPlaceHolder1$txtPassword"]`, chromedp.ByQuery),
		chromedp.SendKeys(`input[name="ctl00$ContentPlaceHolder1$txtPassword"]`, os.Getenv("TIMETABLE_PASS"), chromedp.ByQuery),
		chromedp.Click(`input[name="ctl00$ContentPlaceHolder1$btnLogin"]`, chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
	)
}

func (t *TimeService) TimetableRetrieve(ctx context.Context) string {
	var result string
	err := chromedp.Run(ctx,
		chromedp.Click(`a[href="/myTimetable.aspx"]`, chromedp.NodeVisible),
		chromedp.Sleep(2*time.Second),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.WaitVisible(".fc-content-skeleton tbody tr td:not(.fc-axis)", chromedp.ByQueryAll),
		chromedp.Sleep(2*time.Second), // Дополнительная пауза для надежности
		chromedp.Evaluate(`
			const days = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];
			let schedule = [];
			const dayCells = document.querySelectorAll(".fc-content-skeleton tbody tr td:not(.fc-axis)");
			console.log("dayCells length:", dayCells.length);

			if (dayCells.length === 0) {
				schedule = days.map(day => ({ day, modules: [{ time: "Free", name: "Выходной" }] }));
			} else if (dayCells.length === 14) {
				for (let dayIndex = 0; dayIndex < 7; dayIndex++) {
					let daySchedule = { day: days[dayIndex], modules: [] };
					const cell = dayCells[dayIndex + 7];
					const events = cell.querySelectorAll(".fc-event-container .fc-content");
					if (events.length > 0) {
						events.forEach(event => {
							let module = {
								time: event.querySelector(".fc-time")?.textContent.trim() || "No time",
								name: event.querySelector(".fc-title")?.textContent.trim() || "No title"
							};
							daySchedule.modules.push(module);
						});
					} else {
						daySchedule.modules.push({ time: "Free", name: "Выходной" });
					}
					schedule.push(daySchedule);
				}
			} else {
				console.error("Unexpected number of cells:", dayCells.length);
			}
			JSON.stringify(schedule);
		`, &result),
	)

	if err != nil {
		log.Fatal("Error retrieving data:", err)
	}

	return result
}
