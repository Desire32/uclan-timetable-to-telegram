package timetable

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

var (
	percentages string
	modules     string
)

type BadgeService struct {
}

func (b *BadgeService) BadgeRetrieve(ctx context.Context) string {
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://timetable.cyprus.uclan.ac.uk/STUMyBadges.aspx"),
		chromedp.Sleep(3*time.Second),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.WaitVisible(`#circleP`, chromedp.ByID),
		chromedp.Evaluate(`
			const circleElement = document.getElementById('circleP');
			const strongElement = circleElement.querySelector('strong');
			const strongValue = strongElement.textContent || strongElement.innerText;
JSON.stringify(strongValue);
		`, &percentages))

	if err != nil {
		log.Fatal("Error retrieving data:", err)
	}
	return percentages
}
func (b *BadgeService) ModulesRetrieve(ctx context.Context) string {
	err := chromedp.Run(ctx,
		chromedp.Sleep(3*time.Second),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.Evaluate(`
const tbodyElement = document.getElementById('divTable');
const rows = tbodyElement.querySelectorAll('tr');
const tableData = [];
rows.forEach(row => {
    const cells = row.querySelectorAll('td');
    const rowData = {
        number: cells[0]?.textContent.trim() || "No number",
        module_name: cells[1]?.textContent.trim() || "No module name",
        semester: cells[2]?.textContent.trim() || "No semester",
        total_sessions: cells[3]?.textContent.trim() || "No total sessions",
        attendances_amount: cells[4]?.textContent.trim() || "No attendancy",
        attendance_overall: cells[5]?.querySelector('small')?.textContent.trim() || "No attendance status",
        absences_overall: cells[6]?.querySelector('small')?.textContent.trim() || "No absences status"
    };
    tableData.push(rowData);
});
JSON.stringify(tableData);
		`, &modules),
	)
	if err != nil {
		log.Fatal("Error retrieving data:", err)
	}
	return modules
}
