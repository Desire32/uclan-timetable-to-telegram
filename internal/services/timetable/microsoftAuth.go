package timetable

import (
	"bufio"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"os"
	"regexp"
	"time"
)

type AuthService struct{}

func (a *AuthService) MicrosoftLogin(ctx context.Context) error {
	var loginURL string
	fmt.Println("Переход на страницу расписания...")
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://timetable.cyprus.uclan.ac.uk/"),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),             
		chromedp.Location(&loginURL),                              
	)
	if err != nil {
		return fmt.Errorf("ошибка при переходе на сайт: %w", err)
	}
	re := regexp.MustCompile(`https://login\.microsoftonline\.com/(.+)`)
	if !re.MatchString(loginURL) {
		return fmt.Errorf("не удалось переадресоваться на страницу входа, текущий URL: %s", loginURL)
	}
	fmt.Println("Переадресация на:", loginURL)

	err = chromedp.Run(ctx,
		// EMAIL
		chromedp.WaitVisible(`input[name="loginfmt"]`, chromedp.ByQuery),
		chromedp.SendKeys(`input[name="loginfmt"]`, os.Getenv("EMAIL"), chromedp.ByQuery),
		chromedp.Click(`input[id="idSIButton9"]`, chromedp.ByQuery),
		chromedp.Sleep(3*time.Second),

		// PASSWORD
		chromedp.WaitVisible(`input[name="passwd"]`, chromedp.ByQuery),
		chromedp.SendKeys(`input[name="passwd"]`, os.Getenv("PASSWORD"), chromedp.ByQuery),
		chromedp.Click(`input[id="idSIButton9"]`, chromedp.ByQuery),

		// SMS
		chromedp.Sleep(3*time.Second),
		chromedp.WaitVisible(`div[data-value="OneWaySMS"]`, chromedp.ByQuery),
		chromedp.Sleep(3*time.Second),
		chromedp.Click(`div[data-value="OneWaySMS"]`, chromedp.ByQuery),
		chromedp.Sleep(1*time.Second),
		chromedp.WaitVisible(`#idTxtBx_SAOTCC_OTC`, chromedp.ByID),
	)
	if err != nil {
		return fmt.Errorf("ошибка при вводе email, пароля или выборе SMS: %w", err)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите SMS-код: ")
	smsCode, _ := reader.ReadString('\n')

	err = chromedp.Run(ctx,
		chromedp.Sleep(1*time.Second),
		chromedp.SendKeys(`#idTxtBx_SAOTCC_OTC`, smsCode, chromedp.ByID),
		chromedp.Click(`#idSubmit_SAOTCC_Continue`, chromedp.ByID),
	)
	if err != nil {
		return fmt.Errorf("ошибка при вводе SMS-кода: %w", err)
	}
	err = chromedp.Run(ctx,
		chromedp.Click(`input[id="idSubmit_SAOTCC_Continue"]`, chromedp.ByQuery),
		chromedp.Sleep(3*time.Second),
	)
	if err != nil {
		return fmt.Errorf("ошибка при подтверждении SMS-кода: %w", err)
	}
	fmt.Println("Ждем переадресации...")

	return nil
}
